package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bluemountaincyber/falco_postprocessor/internal/processors"
)

type FalcoEvent struct {
	Time         string                 `json:"time"`
	HostName     string                 `json:"hostname"`
	Rule         string                 `json:"rule"`
	OutputFields map[string]interface{} `json:"output_fields"`
}

func TestProcessDataDNS(t *testing.T) {
	data := processors.FalcoEvent{
		Time:     "2021-08-10T15:00:00Z",
		HostName: "falco-1234",
		Rule:     "DNS Query Logging",
		OutputFields: map[string]interface{}{
			"evt.arg.data": "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H",
		},
	}
	err := processors.ProcessData(&data)
	assert.Nil(t, err)
	assert.Equal(t, "google.com", data.OutputFields["dns_query"])
}

func TestProcessDataIMDS(t *testing.T) {
	data := processors.FalcoEvent{
		Time:     "2021-08-10T15:00:00Z",
		HostName: "falco-1234",
		Rule:     "Metadata Access",
		OutputFields: map[string]interface{}{
			"evt.arg.data": "R0VUIC9tZXRhZGF0YS9pbnN0YW5jZT9hcGktdmVyc2lvbj0yMDE4LTAyLTAxIEhUVFAvMS4xDQpIb3N0OiAxNjkuMjU0LjE2OS4yNTQNCkE=",
		},
	}
	err := processors.ProcessData(&data)
	assert.Nil(t, err)
	assert.Equal(t, "/metadata/instance?api-version=2018-02-01", data.OutputFields["metadata_path"])
}

func TestProcessDataOther(t *testing.T) {
	data := processors.FalcoEvent{
		Time:     "2021-08-10T15:00:00Z",
		HostName: "falco-1234",
		Rule:     "Some Other Rule",
		OutputFields: map[string]interface{}{
			"evt.arg.data": "some data",
		},
	}
	err := processors.ProcessData(&data)
	assert.Nil(t, err)
	assert.NotContains(t, data.OutputFields, "dns_query")
	assert.NotContains(t, data.OutputFields, "metadata_path")
}
