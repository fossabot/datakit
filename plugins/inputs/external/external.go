// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package external wraps all external command to collect various metrics
package external

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

const (
	configSample = `
[[inputs.external]]

	# 外部采集器名称
	name = 'some-external-inputs'  # required

	# 是否以后台方式运行外部采集器
	daemon = false

	# 如果以非 daemon 方式运行外部采集器，则以该间隔多次运行外部采集器
	#interval = '10s'

	# 运行外部采集器所需的环境变量
	#envs = ['LD_LIBRARY_PATH=/path/to/lib:$LD_LIBRARY_PATH',]

	# 外部采集器可执行程序路径(尽可能写绝对路径)
	cmd = "python" # required

	# 如果该外部采集器参与选举，则开启该选项
	# 注意，如果参与选举，则必须以 daemon 形式运行（即 daemon 自动为 true）
	election = false
	args = []

[[inputs.external.tags]]
	# tag1 = "val1"
	# tag2 = "val2"
	`
)

var (
	inputName = "external"
	l         = logger.DefaultSLogger(inputName)
)

type ExternalInput struct {
	Name     string            `toml:"name"`
	Daemon   bool              `toml:"daemon"`
	Election bool              `toml:"election"`
	Interval string            `toml:"interval"`
	Envs     []string          `toml:"envs"`
	Cmd      string            `toml:"cmd"`
	Args     []string          `toml:"args"`
	Tags     map[string]string `toml:"tags"`

	cmd      *exec.Cmd     `toml:"-"`
	duration time.Duration `toml:"-"`

	semStop        *cliutils.Sem // start stop signal
	semStopProcess *cliutils.Sem

	daemonStarted bool

	pauseCh chan bool
	pause   bool
}

func NewExternalInput() *ExternalInput {
	return &ExternalInput{
		semStop:        cliutils.NewSem(),
		semStopProcess: cliutils.NewSem(),
		pauseCh:        make(chan bool, inputs.ElectionPauseChannelLength),
	}
}

func (*ExternalInput) Catalog() string {
	return "external"
}

func (*ExternalInput) SampleConfig() string {
	return configSample
}

func (ex *ExternalInput) precheck() error {
	ex.duration = time.Second * 10
	if ex.Interval != "" {
		du, err := time.ParseDuration(ex.Interval)
		if err != nil {
			l.Errorf("parse external input %s interval failed: %s", ex.Name, err.Error())
			return err
		}

		ex.duration = du
	}

	// TODO: check ex.Cmd is ok

	return nil
}

func (ex *ExternalInput) start() error {
	ex.cmd = exec.Command(ex.Cmd, ex.Args...) //nolint:gosec
	if ex.Envs != nil {
		ex.cmd.Env = ex.Envs
	}

	if err := ex.cmd.Start(); err != nil {
		l.Errorf("start external input %s failed: %s", ex.Name, err.Error())
		return err
	}

	return nil
}

func (ex *ExternalInput) Run() {
	l = logger.SLogger(inputName)

	l.Infof("starting external input %s...", ex.Name)

	tagsStr := ""
	arr := []string{}
	for tagKey, tagVal := range ex.Tags {
		arr = append(arr, fmt.Sprintf("%s=%s", tagKey, tagVal))
	}
	if len(arr) > 0 {
		tagsStr = strings.Join(arr, ";")
	}

	if tagsStr != "" {
		ex.Args = append(ex.Args, []string{"--tags", tagsStr}...)
	}

	for {
		if err := ex.precheck(); err != nil {
			time.Sleep(time.Second)
			continue
		}
		break
	}

	tick := time.NewTicker(ex.duration)
	defer tick.Stop()

	for {
		if ex.Election && ex.pause {
			l.Debugf("%s not leader, skipped", ex.Name)
		} else {
			if ex.Daemon {
				ex.daemonRun()
			} else {
				// run as new process
				l.Debugf("non-daemon starting %s cmd %s %s, envs: %+#v", ex.Name, ex.Cmd, strings.Join(ex.Args, " "), ex.Envs)
				_ = ex.start() //nolint:errcheck
			}
		}

		select {
		case <-datakit.Exit.Wait():
			l.Infof("external input %s exiting", ex.Name)
			ex.semStopProcess.Close()
			return

		case <-ex.semStop.Wait():
			l.Infof("external input %s stopped", ex.Name)
			ex.semStopProcess.Close()
			return

		case ex.pause = <-ex.pauseCh:
			if ex.pause && ex.Election {
				l.Infof("%s paused", ex.Name)
				if ex.Daemon && ex.daemonStarted { // stop the daemon running process
					ex.semStopProcess.Close() // trigger the daemon process exit

					ex.daemonStarted = false
					ex.semStopProcess = cliutils.NewSem() // reopen the sem
				}
			}

		case <-tick.C:
		}
	}
}

func (ex *ExternalInput) daemonRun() {
	if ex.daemonStarted {
		return
	}

	// start failed, retry
	for {
		l.Debugf("daemon starting %s cmd %s %s, envs: %+#v", ex.Name, ex.Cmd, strings.Join(ex.Args, " "), ex.Envs)
		if err := ex.start(); err != nil {
			time.Sleep(time.Second)
			continue
		}
		ex.daemonStarted = true
		break
	}

	go func() {
		if err := datakit.MonitProc(ex.cmd.Process, ex.Name, ex.semStopProcess); err != nil { // blocking here...
			l.Errorf("datakit.MonitProc: %s", err.Error())
		}
	}()
}

func (ex *ExternalInput) Pause() error {
	tick := time.NewTicker(inputs.ElectionPauseTimeout)
	defer tick.Stop()
	select {
	case ex.pauseCh <- true:
		return nil
	case <-tick.C:
		return fmt.Errorf("pause %s failed", ex.Name)
	}
}

func (ex *ExternalInput) Resume() error {
	tick := time.NewTicker(inputs.ElectionResumeTimeout)
	defer tick.Stop()
	select {
	case ex.pauseCh <- false:
		return nil
	case <-tick.C:
		return fmt.Errorf("resume %s failed", ex.Name)
	}
}

func (ex *ExternalInput) Terminate() {
	if ex.semStop != nil {
		ex.semStop.Close()
	}
}

func init() { //nolint:gochecknoinits
	inputs.Add(inputName, func() inputs.Input {
		return NewExternalInput()
	})
}
