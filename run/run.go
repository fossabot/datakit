package run

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	teleagent "github.com/influxdata/telegraf/agent"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/models"
)

type Agent struct {
	Config     *config.Config
	outputsMgr *outputsMgr
}

func NewAgent(config *config.Config) (*Agent, error) {
	a := &Agent{
		Config: config,
		outputsMgr: &outputsMgr{
			outputChannels: make(map[string]*outputChannel),
		},
	}
	return a, nil
}

func (a *Agent) Run(ctx context.Context) error {

	if ctx.Err() != nil {
		return ctx.Err()
	}

	select {
	case <-ctx.Done():
		return context.Canceled
	default:
	}

	log.Printf("Loading outputs")
	if err := a.outputsMgr.LoadOutputs(a.Config); err != nil {
		return err
	}

	if err := a.outputsMgr.ConnectOutputs(ctx); err != nil {
		return err
	}

	var outputsNames []string
	for _, oc := range a.outputsMgr.outputChannels {
		for _, op := range oc.outputs {
			outputsNames = append(outputsNames, op.Config.Name)
		}
	}
	log.Printf("avariable outputs: %s", strings.Join(outputsNames, ","))

	startTime := time.Now()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Printf("Starting inputs")
		err := a.runInputs(ctx, startTime)
		if err != nil && err != context.Canceled {
			log.Printf("E! Error running inputs: %v", err)
		}

		log.Printf("Inputs done")
	}()

	go func() {
		select {
		case <-ctx.Done():
			a.stopInputs()
			a.outputsMgr.Stop()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.outputsMgr.Start(startTime,
			a.Config.MainCfg.FlushInterval.Duration,
			a.Config.MainCfg.FlushJitter.Duration,
			a.Config.MainCfg.RoundInterval)

		if err != nil && err != context.Canceled {
			log.Printf("E! Error starting outputs: %v", err)
		}

	}()

	wg.Wait()
	a.outputsMgr.Close()

	log.Printf("datakit stopped successfully")
	return nil
}

func (a *Agent) runIntervalInput(
	ctx context.Context,
	input *models.RunningInput,
	startTime time.Time,
	dst chan<- telegraf.Metric,
	wg *sync.WaitGroup,
) error {

	if _, ok := input.Input.(telegraf.ServiceInput); ok {
		return nil
	}
	interval := a.Config.MainCfg.Interval.Duration
	jitter := time.Duration(0) //a.Config.Agent.CollectionJitter.Duration

	// Overwrite agent interval if this plugin has its own.
	if input.Config.Interval != 0 {
		interval = input.Config.Interval
	}

	acc := teleagent.NewAccumulator(input, dst)
	acc.SetPrecision(time.Nanosecond)

	wg.Add(1)
	go func(input *models.RunningInput) {
		defer wg.Done()

		if a.Config.MainCfg.RoundInterval {
			err := internal.SleepContext(ctx, internal.AlignDuration(startTime, interval))
			if err != nil {
				return
			}
		}

		a.gatherOnInterval(ctx, acc, input, interval, jitter)
	}(input)

	return nil
}

func (a *Agent) runServiceInput(ctx context.Context, input *models.RunningInput, dst chan<- telegraf.Metric) error {

	if si, ok := input.Input.(telegraf.ServiceInput); ok {

		// Service input plugins are not subject to timestamp rounding.
		// This only applies to the accumulator passed to Start(), the
		// Gather() accumulator does apply rounding according to the
		// precision agent setting.
		log.Printf("D! starting service input: %s", input.Config.Name)
		acc := teleagent.NewAccumulator(input, dst)
		acc.SetPrecision(time.Nanosecond)

		err := si.Start(acc)
		if err != nil {
			log.Printf("E! Service for [%s] failed to start: %v", input.LogName(), err)
			return err
		}
	}

	return nil
}

func (a *Agent) runInputs(ctx context.Context, startTime time.Time) error {

	var wg sync.WaitGroup

	for _, input := range a.Config.Inputs {

		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}

		err := input.Init()
		if err != nil {
			return fmt.Errorf("could not initialize input %s: %v", input.LogName(), err)
		}

		dst := a.outputsMgr.findMetricChannel(input.Config.Name)
		if dst == nil {
			continue
		}
		if _, ok := input.Input.(telegraf.ServiceInput); ok {
			if err := a.runServiceInput(ctx, input, dst.ch); err != nil {
				return err
			}
		} else {
			if err := a.runIntervalInput(ctx, input, startTime, dst.ch, &wg); err != nil {
				return err
			}
		}
	}

	wg.Wait()

	return nil
}

func (a *Agent) stopInputs() {
	for _, input := range a.Config.Inputs {
		if _, ok := input.Input.(telegraf.ServiceInput); !ok {
			continue
		}
		log.Printf("D! stopping service input: %s", input.Config.Name)
		if si, ok := input.Input.(telegraf.ServiceInput); ok {
			si.Stop()
		}
	}
}

// panicRecover displays an error if an input panics.
func panicRecover(input *models.RunningInput) {
	if err := recover(); err != nil {
		trace := make([]byte, 2048)
		runtime.Stack(trace, true)
		log.Printf("E! FATAL: [%s] panicked: %s, Stack:\n%s",
			input.LogName(), err, trace)
		log.Println("E! PLEASE REPORT THIS PANIC ON GITHUB with " +
			"stack trace, configuration, and OS information: " +
			"https://github.com/influxdata/telegraf/issues/new/choose")
	}
}

func (a *Agent) gatherOnInterval(
	ctx context.Context,
	acc telegraf.Accumulator,
	input *models.RunningInput,
	interval time.Duration,
	jitter time.Duration,
) {
	defer panicRecover(input)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		err := internal.SleepContext(ctx, internal.RandomDuration(jitter))
		if err != nil {
			return
		}

		err = a.gatherOnce(acc, input, interval)
		if err != nil {
			acc.AddError(err)
		}

		select {
		case <-ticker.C:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (a *Agent) gatherOnce(
	acc telegraf.Accumulator,
	input *models.RunningInput,
	timeout time.Duration,
) error {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	done := make(chan error)
	go func() {
		done <- input.Gather(acc)
	}()

	for {
		select {
		case err := <-done:
			return err
		case <-ticker.C:
			log.Printf("W! [%s] did not complete within its interval",
				input.LogName())
		}
	}
}
