package run

import (
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var (
	l *logger.Logger
)

type Agent struct {
}

func NewAgent() (*Agent, error) {
	a := &Agent{}
	return a, nil
}

func (a *Agent) Run() error {

	l = logger.SLogger("run")

	io.Start()

	if err := a.runInputs(); err != nil {
		l.Error("error running inputs: %v", err)
	}

	// wait all plugin start
	time.Sleep(time.Second * 3)
	datakit.WG.Add(1)
	go func() {
		defer datakit.WG.Done()
		io.HTTPServer()
		l.Info("HTTPServer goroutine exit")
	}()

	return nil
}

func (a *Agent) runInputs() error {

	for name, ips := range config.Cfg.Inputs {

		for _, input := range ips {

			switch input.(type) {

			case inputs.Input:
				l.Infof("starting input %s ...", name)
				datakit.WG.Add(1)
				go func(i inputs.Input, name string) {
					defer datakit.WG.Done()
					i.Run()
					l.Infof("input %s exited", name)
				}(input, name)

			default:
				l.Warn("ignore input %s", name)
			}
		}

	}

	return nil
}
