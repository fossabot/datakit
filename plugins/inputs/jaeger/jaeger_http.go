package jaeger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/uber/jaeger-client-go/thrift"
	"github.com/uber/jaeger-client-go/thrift-gen/jaeger"
	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/io/trace"
)

func handleJaegerTrace(resp http.ResponseWriter, req *http.Request) {
	log.Debugf("%s: listen on path: %s", inputName, req.URL.Path)

	// contentType, body, err := itrace.ParseTracingRequest(req)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	resp.WriteHeader(http.StatusBadRequest)

	// 	return
	// }

	// if contentType != "application/x-thrift" {
	// 	log.Errorf("Jeager unsupported Content-Type: %s", contentType)
	// 	resp.WriteHeader(http.StatusBadRequest)

	// 	return
	// }

	buf := thrift.NewTMemoryBuffer()
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	var (
		transport = thrift.NewTBinaryProtocolConf(buf, &thrift.TConfiguration{})
		batch     = &jaeger.Batch{}
	)
	if err = batch.Read(context.TODO(), transport); err != nil {
		log.Error(err.Error())
		resp.WriteHeader(http.StatusBadRequest)

		return
	}

	if dktrace := batchToDkTrace(batch); len(dktrace) == 0 {
		log.Warn("empty datakit trace")
	} else {
		afterGatherRun.Run(inputName, dktrace, false)
	}
}

func batchToDkTrace(batch *jaeger.Batch) itrace.DatakitTrace {
	buf, err := json.Marshal(batch)
	if err != nil {
		log.Debug(err.Error())
	} else {
		log.Debug(string(buf))
	}

	var (
		project, version, env = getExpandInfo(batch)
		dktrace               itrace.DatakitTrace
		spanIDs, parentIDs    = getSpanIDsAndParentIDs(batch.Spans)
	)
	for _, span := range batch.Spans {
		if span == nil {
			continue
		}

		dkspan := &itrace.DatakitSpan{
			TraceID:   itrace.GetTraceStringID(span.TraceIdHigh, span.TraceIdLow),
			ParentID:  fmt.Sprintf("%d", span.ParentSpanId),
			SpanID:    fmt.Sprintf("%d", span.SpanId),
			Service:   batch.Process.ServiceName,
			Resource:  span.OperationName,
			Operation: span.OperationName,
			Source:    inputName,
			SpanType:  itrace.FindSpanTypeInt(span.SpanId, span.ParentSpanId, spanIDs, parentIDs),
			Env:       env,
			Project:   project,
			Start:     span.StartTime * int64(time.Microsecond),
			Duration:  span.Duration * int64(time.Microsecond),
			Version:   version,
		}

		dkspan.Status = itrace.STATUS_OK
		for _, tag := range span.Tags {
			if tag.Key == "error" {
				dkspan.Status = itrace.STATUS_ERR
				break
			}
		}

		sourceTags := make(map[string]string)
		for _, tag := range span.Tags {
			sourceTags[tag.Key] = tag.String()
		}
		dkspan.Tags = itrace.MergeInToCustomerTags(customerKeys, tags, sourceTags)

		if dkspan.ParentID == "0" && defSampler != nil {
			dkspan.Priority = defSampler.Priority
			dkspan.SamplingRateGlobal = defSampler.SamplingRateGlobal
		}

		if buf, err := json.Marshal(span); err != nil {
			log.Warn(err.Error())
		} else {
			dkspan.Content = string(buf)
		}

		dktrace = append(dktrace, dkspan)
	}

	return dktrace
}

func getSpanIDsAndParentIDs(trace []*jaeger.Span) (map[int64]bool, map[int64]bool) {
	var (
		spanIDs   = make(map[int64]bool)
		parentIDs = make(map[int64]bool)
	)
	for _, span := range trace {
		if span == nil {
			continue
		}
		spanIDs[span.SpanId] = true
		if span.ParentSpanId != 0 {
			parentIDs[span.ParentSpanId] = true
		}
	}

	return spanIDs, parentIDs
}

func getExpandInfo(batch *jaeger.Batch) (project, version, env string) {
	if batch.Process == nil {
		return
	}

	for _, tag := range batch.Process.Tags {
		if tag == nil {
			continue
		}

		switch tag.Key {
		case itrace.PROJECT:
			project = fmt.Sprintf("%v", getValueString(tag))
		case itrace.VERSION:
			version = fmt.Sprintf("%v", getValueString(tag))
		case itrace.ENV:
			env = fmt.Sprintf("%v", getValueString(tag))
		}
	}

	return
}

func getValueString(tag *jaeger.Tag) interface{} {
	switch tag.VType {
	case jaeger.TagType_STRING:
		return *(tag.VStr)
	case jaeger.TagType_DOUBLE:
		return *(tag.VDouble)
	case jaeger.TagType_BOOL:
		return *(tag.VBool)
	case jaeger.TagType_LONG:
		return *(tag.VLong)
	case jaeger.TagType_BINARY:
		return tag.VBinary
	default:
		return nil
	}
}
