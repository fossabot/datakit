package jvm

import (
	"time"

	"github.com/influxdata/telegraf/plugins/common/tls"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var (
	inputName = "jvm"
	l         *logger.Logger
)

const (
	defaultInterval = "60s"
)

type Input struct {
	URLs            string `toml:"urls"`
	Username        string
	Password        string
	ResponseTimeout time.Duration `toml:"response_timeout"`
	Interval        string
	MetricName      string `toml:"metric_name"`

	tls.ClientConfig

	client *Client

	collectCache []inputs.Measurement
}

type JvmMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
	ts     time.Time
}

func (j *JvmMeasurement) LineProto() (*io.Point, error) {
	return io.MakePoint(j.name, j.tags, j.fields, j.ts)
}

func (j *JvmMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: inputName,
		Fields: map[string]*inputs.FieldInfo{
			"heap_memory_init":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The initial Java heap memory allocated."},
			"heap_memory_committed": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The total Java heap memory committed to be used."},
			"heap_memory_max":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The maximum Java heap memory available."},
			"heap_memory":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The total Java heap memory used."},

			"non_heap_memory_init":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The initial Java non-heap memory allocated."},
			"non_heap_memory_committed": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The total Java non-heap memory committed to be used."},
			"non_heap_memory_max":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The maximum Java non-heap memory available."},
			"non_heap_memory":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.SizeByte, Desc: "The total Java non-heap memory used."},

			"thread_count":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Count, Unit: inputs.UnknownUnit, Desc: "The number of live threads."},
			"minor_collection_count": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Count, Unit: inputs.UnknownUnit, Desc: "The number of minor garbage collections that have occurred."},
			"minor_collection_time":  &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "The approximate minor garbage collection time elapsed."},
			"major_collection_count": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Count, Unit: inputs.UnknownUnit, Desc: "The number of major garbage collections that have occurred."},
			"major_collection_time":  &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.DurationMS, Desc: "The approximate major garbage collection time elapsed."},
		},
	}
}

func (i *Input) Run() {
	l = logger.DefaultSLogger(inputName)
	l.Infof("%s input started...", inputName)

	if i.Interval == "" {
		i.Interval = defaultInterval
	}

	if i.MetricName == "" {
		i.MetricName = inputName
	}

	i.gather()
}

func (i *Input) Catalog() string      { return inputName }
func (i *Input) SampleConfig() string { return javaConfSample }
func (i *Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{
		&JvmMeasurement{},
	}
}

func init() {
	initConvertDict()

	inputs.Add(inputName, func() inputs.Input {
		return &Input{}
	})
}
