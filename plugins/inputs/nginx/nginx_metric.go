// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package nginx

import (
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

type NginxMeasurement struct {
	name     string
	tags     map[string]string
	fields   map[string]interface{}
	ts       time.Time
	election bool
}

func (m *NginxMeasurement) LineProto() (*point.Point, error) {
	return point.NewPoint(m.name, m.tags, m.fields, point.MOptElectionV2(m.election))
}

//nolint:lll
func (m *NginxMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: nginx,
		Fields: map[string]interface{}{
			"load_timestamp":      newOtherFieldInfo(inputs.Int, inputs.Gauge, inputs.TimestampMS, "Loaded process time in milliseconds, when exist by open vts"),
			"connection_active":   newCountFieldInfo("The current number of active client connections"),
			"connection_reading":  newCountFieldInfo("The total number of reading client connections"),
			"connection_writing":  newCountFieldInfo("The total number of writing client connections"),
			"connection_waiting":  newCountFieldInfo("The total number of waiting client connections"),
			"connection_handled":  newCountFieldInfo("The total number of handled client connections"),
			"connection_requests": newCountFieldInfo("The total number of requests client connections"),
		},
		Tags: map[string]interface{}{
			"nginx_server":  inputs.NewTagInfo("nginx server host"),
			"nginx_port":    inputs.NewTagInfo("nginx server port"),
			"host":          inputs.NewTagInfo("host mame which installed nginx,use vts exist"),
			"nginx_version": inputs.NewTagInfo("nginx version,use vts exist"),
		},
	}
}
