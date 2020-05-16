package aliyunlog

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/models"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	consumerLibrary "github.com/aliyun/aliyun-log-go-sdk/consumer"
)

type AliyunLog struct {
	Consumer []*ConsumerInstance

	runningInstances []*runningInstance

	ctx       context.Context
	cancelFun context.CancelFunc

	accumulator telegraf.Accumulator

	logger *models.Logger
}

type runningInstance struct {
	cfg *ConsumerInstance

	agent *AliyunLog

	logger *models.Logger

	runningProjects []*runningProject
}

type runningProject struct {
	inst *runningInstance
	cfg  *LogProject

	logger *models.Logger

	runningStores []*runningStore
}

type runningStore struct {
	proj       *runningProject
	cfg        *LogStoreCfg
	metricName string

	fieldsInfo map[string]string

	logger *models.Logger
}

func (_ *AliyunLog) SampleConfig() string {
	return aliyunlogConfigSample
}

func (_ *AliyunLog) Description() string {
	return "Collect logs from aliyun SLS"
}

func (_ *AliyunLog) Gather(telegraf.Accumulator) error {
	return nil
}

func (al *AliyunLog) Start(acc telegraf.Accumulator) error {

	al.logger = &models.Logger{
		Name: `aliyunlog`,
	}

	if len(al.Consumer) == 0 {
		al.logger.Warnf("no configuration found")
		return nil
	}

	al.logger.Info("starting...")

	al.accumulator = acc

	for _, instCfg := range al.Consumer {
		r := &runningInstance{
			cfg:    instCfg,
			agent:  al,
			logger: al.logger,
		}
		al.runningInstances = append(al.runningInstances, r)

		go r.run(al.ctx)
	}

	return nil
}

func (al *AliyunLog) Stop() {
	al.cancelFun()
}

func (r *runningInstance) run(ctx context.Context) error {

	for _, c := range r.cfg.Projects {

		p := &runningProject{
			cfg:    c,
			inst:   r,
			logger: r.logger,
		}
		r.runningProjects = append(r.runningProjects, p)

		go p.run(ctx)
	}

	return nil
}

func (r *runningProject) run(ctx context.Context) error {

	for _, c := range r.cfg.Stores {

		s := &runningStore{
			cfg:    c,
			proj:   r,
			logger: r.logger,
		}
		s.metricName = c.MetricName
		if s.metricName == "" {
			s.metricName = `aliyunlog_` + c.Name
		}
		r.runningStores = append(r.runningStores, s)

		go s.run(ctx)
	}

	return nil
}

func (r *runningStore) run(ctx context.Context) error {

	r.fieldsInfo = map[string]string{}

	for _, fitem := range r.cfg.Fields {
		parts := strings.Split(fitem, ":")
		if len(parts) != 2 {
			r.logger.Warnf("invalid field type specification")
			continue
		}
		fieldType := parts[0]
		fieldNames := strings.Split(parts[1], ",")
		for _, f := range fieldNames {
			r.fieldsInfo[f] = fieldType
		}
	}

	option := consumerLibrary.LogHubConfig{
		Endpoint:          r.proj.inst.cfg.Endpoint,
		AccessKeyID:       r.proj.inst.cfg.AccessKeyID,
		AccessKeySecret:   r.proj.inst.cfg.AccessKeySecret,
		Project:           r.proj.cfg.Name,
		Logstore:          r.cfg.Name,
		ConsumerGroupName: r.cfg.ConsumerGroupName,
		ConsumerName:      r.cfg.ConsumerName,
		// This options is used for initialization, will be ignored once consumer group is created and each shard has been started to be consumed.
		// Could be "begin", "end", "specific time format in time stamp", it's log receiving time.
		CursorPosition: consumerLibrary.BEGIN_CURSOR,
	}

	consumerWorker := consumerLibrary.InitConsumerWorker(option, r.logProcess)
	consumerWorker.Start()

	select {
	case <-ctx.Done():
		consumerWorker.StopAndWait()
	}

	r.logger.Infof("%s done", r.cfg.Name)

	return nil

}

func (r *runningStore) checkAsTag(key string) bool {
	for _, k := range r.cfg.Tags {
		if k == key {
			return true
		}
	}
	return false
}

func (r *runningStore) checkFieldType(field string) string {
	if ftype, ok := r.fieldsInfo[field]; ok {
		return ftype
	}
	return "string"
}

func (r *runningStore) logProcess(shardId int, logGroupList *sls.LogGroupList) string {

	r.logger.Debugf("shardId:%d, grouplist:%s", shardId, logGroupList.String())
	for _, lg := range logGroupList.LogGroups {

		for _, l := range lg.GetLogs() {

			fields := map[string]interface{}{}

			tags := map[string]string{}
			tags["store"] = r.cfg.Name
			tags["project"] = r.proj.cfg.Name
			tags["__topic__"] = lg.GetTopic()

			for _, lt := range lg.GetLogTags() {
				k := lt.GetKey()
				if k == "" || lt.GetValue() == "" {
					continue
				}
				if r.checkAsTag(k) {
					tags[k] = lt.GetValue()
				} else {
					fields[k] = lt.GetValue()
				}
			}

			if lg.GetSource() != "" {
				if r.checkAsTag("__source__") {
					tags["__source__"] = lg.GetSource()
				} else {
					fields["__source__"] = lg.GetSource()
				}
			}

			// if lg.GetCategory() != "" {
			// 	tags["__category__"] = lg.GetCategory()
			// }

			for _, lc := range l.Contents {
				k := lc.GetKey()
				if k != "" {
					if r.checkAsTag(k) {
						tags[k] = lc.GetValue()
					} else {
						strval := lc.GetValue()
						fieldType := r.checkFieldType(k)
						if fieldType != "string" {
							switch fieldType {
							case "int":
								nval, err := strconv.ParseInt(strval, 10, 64)
								if err != nil {
									r.logger.Warnf("you specify '%s' as int, but fail to convert '%s' to int", k, strval)
								} else {
									fields[k] = nval
								}
							case "float":
								fval, err := strconv.ParseFloat(strval, 64)
								if err != nil {
									r.logger.Warnf("you specify '%s' as float, but fail to convert '%s' to float", k, strval)
								} else {
									fields[k] = fval
								}
							}
						} else {
							fields[k] = strval
						}
					}
				}
			}

			tm := time.Unix(int64(l.GetTime()), 0)
			m, err := metric.New(r.metricName, tags, fields, tm)
			if err == nil {
				if r.proj.inst.agent.accumulator != nil {
					r.proj.inst.agent.accumulator.AddMetric(m)
				} else {
					fmt.Printf("%s", internal.Metric2InfluxLine(m))
				}
			} else {
				r.logger.Warnf("fail to generate metric, %s", err)
			}
		}
	}
	return ""
}

func NewAgent() *AliyunLog {
	ac := &AliyunLog{}
	ac.ctx, ac.cancelFun = context.WithCancel(context.Background())
	return ac
}

func init() {
	inputs.Add("aliyunlog", func() telegraf.Input {
		return NewAgent()
	})
}
