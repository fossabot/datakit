// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package opentelemetry http method

package opentelemetry

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/trace"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/workerpool"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry/collector"
	collectormetricpb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	collectortracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	pbContentType   = "application/x-protobuf"
	jsonContentType = "application/json"
)

// handler collector.
type otlpHTTPCollector struct {
	storage         *collector.SpansStorage
	Enable          bool              `toml:"enable"`
	HTTPStatusOK    int               `toml:"http_status_ok"`
	ExpectedHeaders map[string]string // 用于检测是否包含特定的 header
}

type parameters struct {
	urlPath string
	media   string
	buf     []byte
	storage *collector.SpansStorage
}

// apiOtlpCollector :trace.
func (o *otlpHTTPCollector) apiOtlpTrace(resp http.ResponseWriter, req *http.Request) {
	if o.storage == nil {
		log.Error("storage is nil")
		resp.WriteHeader(http.StatusInternalServerError)

		return
	}

	if !o.checkHeaders(req) {
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	response := collectortracepb.ExportTraceServiceResponse{}
	rawResponse, err := proto.Marshal(&response)
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusInternalServerError)

		return
	}

	readbodycost := time.Now()
	media, encode, buf, err := itrace.ParseTracerRequest(req)
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	param := &parameters{
		urlPath: req.URL.Path,
		media:   media,
		buf:     buf,
		storage: o.storage,
	}

	log.Debugf("### path: %s, Content-Type: %s, Encode-Type: %s, body-size: %dkb, read-body-cost: %dms",
		req.URL.Path, media, encode, len(buf)>>10, time.Since(readbodycost)/time.Millisecond)

	if wpool == nil {
		if err = parseOtelTrace(param); err != nil {
			log.Error(err.Error())
			resp.WriteHeader(http.StatusBadRequest)

			return
		}
	} else {
		job, err := workerpool.NewJob(workerpool.WithInput(param),
			workerpool.WithProcess(parseOtelTraceAdapter),
			workerpool.WithProcessCallback(func(input, output interface{}, cost time.Duration) {
				log.Debugf("### job status: input: %v, output: %v, cost: %dms", input, output, cost/time.Millisecond)
			}),
		)
		if err != nil {
			log.Error(err.Error())
			resp.WriteHeader(http.StatusBadRequest)

			return
		}

		if err = wpool.MoreJob(job); err != nil {
			log.Error(err.Error())
			resp.WriteHeader(http.StatusTooManyRequests)

			return
		}
	}

	writeReply(resp, rawResponse, o.HTTPStatusOK, param.media, nil)
}

func parseOtelTraceAdapter(input interface{}) (output interface{}) {
	param, ok := input.(*parameters)
	if !ok {
		return errors.New("type assertion failed")
	}

	return parseOtelTrace(param)
}

func parseOtelTrace(param *parameters) error {
	request, err := unmarshalTraceRequest(param.buf, param.media)
	if err != nil {
		return err
	}

	if len(request.ResourceSpans) != 0 && param.storage != nil {
		param.storage.AddSpans(request.ResourceSpans)
	}

	return nil
}

func (o *otlpHTTPCollector) apiOtlpMetric(resp http.ResponseWriter, req *http.Request) {
	if o.storage == nil {
		log.Error("storage is nil")
		resp.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := collectormetricpb.ExportMetricsServiceResponse{}
	rawResponse, err := proto.Marshal(&response)
	if err != nil {
		log.Errorf("proto marshal error=%v", err)
		resp.WriteHeader(http.StatusInternalServerError)

		return
	}

	media, _, buf, err := itrace.ParseTracerRequest(req)
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	request, err := unmarshalMetricsRequest(buf, media)
	if err != nil {
		log.Errorf("unmarshalMetricsRequest err=%v", err)
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	writeReply(resp, rawResponse, o.HTTPStatusOK, media, nil)

	orms := o.storage.ToDatakitMetric(request.ResourceMetrics)
	o.storage.AddMetric(orms)
}

func (o *otlpHTTPCollector) checkHeaders(req *http.Request) bool {
	for k, v := range o.ExpectedHeaders {
		got := req.Header.Get(k)
		if got != v {
			return false
		}
	}

	return true
}

func unmarshalTraceRequest(rawRequest []byte, contentType string) (*collectortracepb.ExportTraceServiceRequest, error) {
	request := &collectortracepb.ExportTraceServiceRequest{}
	var err error
	switch contentType {
	case pbContentType:
		err = proto.Unmarshal(rawRequest, request)
	case jsonContentType:
		err = protojson.Unmarshal(rawRequest, request)
	default:
		err = fmt.Errorf("invalid content-type: %s, only application/x-protobuf and application/json is supported", contentType)
	}

	return request, err
}

func unmarshalMetricsRequest(rawRequest []byte, contentType string) (*collectormetricpb.ExportMetricsServiceRequest, error) {
	request := &collectormetricpb.ExportMetricsServiceRequest{}
	var err error
	switch contentType {
	case pbContentType:
		err = proto.Unmarshal(rawRequest, request)
	case jsonContentType:
		err = protojson.Unmarshal(rawRequest, request)
	default:
		err = fmt.Errorf("invalid content-type: %s, only application/x-protobuf and application/json is supported", contentType)
	}

	return request, err
}

func writeReply(resp http.ResponseWriter, rawResponse []byte, status int, ct string, h map[string]string) {
	contentType := "application/x-protobuf"
	if ct != "" {
		contentType = ct
	}
	resp.Header().Set("Content-Type", contentType)
	for k, v := range h {
		resp.Header().Add(k, v)
	}
	resp.WriteHeader(status)
	_, _ = resp.Write(rawResponse)
}
