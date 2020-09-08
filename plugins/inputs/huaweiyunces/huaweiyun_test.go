package huaweiyunces

import (
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/influxdata/toml"
)

var (
	ak = ``
	sk = ``

	endPoint  = `ces.cn-east-3.myhuaweicloud.com`
	projectID = `0838fedce480f3982f39c0150293ac02`

	testInput = false
)

func TestGetMetric(t *testing.T) {

	//https://support.huaweicloud.com/api-ces/ces_03_0033.html

	cli := newHWClient(ak, sk, endPoint, projectID)
	dims := []*Dimension{
		{
			Name:  "instance_id",
			Value: "b5d7b7a3-681d-4c08-8e32-f14b640b3e12",
		},
	}
	resp, err := cli.getMetric("SYS.ECS", "cpu_util", "min", 300, time.Now().Add(-5*time.Minute).Unix()*1000, time.Now().Unix()*1000, dims)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v", resp)
}

func TestBatchMetrics(t *testing.T) {

	cli := newHWClient(ak, sk, endPoint, projectID)

	dims := []*Dimension{
		{
			Name:  "instance_id",
			Value: "b5d7b7a3-681d-4c08-8e32-f14b640b3e12",
		},
	}

	items := []*metricItem{
		{
			Namespace:  "SYS.ECS",
			MetricName: "cpu_util",
			Dimensions: dims,
		},
		{
			Namespace:  "SYS.ECS",
			MetricName: "disk_write_bytes_rate",
			Dimensions: dims,
		},
	}

	b := &batchReq{
		Period:  "300",
		Filter:  "min",
		From:    time.Now().Add(-1*time.Hour).Unix() * 1000,
		To:      time.Now().Unix() * 1000,
		Metrics: items,
	}

	resp, err := cli.batchMetrics(b)
	if err == nil {
		result := parseBatchResponse(resp, b.Filter)
		if result != nil {
			for _, item := range result.results {
				log.Printf("%s", item)
			}
		}
	}

}

func TestInput(t *testing.T) {

	testInput = true

	data, err := ioutil.ReadFile("test.conf")
	if err != nil {
		t.Error(err)
	}
	ag := newAgent()
	if err = toml.Unmarshal(data, &ag); err != nil {
		t.Error(err)
	}
	ag.Run()
}
