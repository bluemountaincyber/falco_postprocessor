package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bluemountaincyber/falco_postprocessor/internal/processors"
)

func TestRetrieveDNSQueryHost(t *testing.T) {
	data := "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H"
	hostName, err := processors.RetrieveDNSQueryHost(data)
	assert.Nil(t, err)
	assert.Equal(t, "google.com", hostName)
}

func TestRetrieveDNSQueryHostBadBase64(t *testing.T) {
	data := "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H======"
	hostName, err := processors.RetrieveDNSQueryHost(data)
	assert.NotNil(t, err)
	assert.Equal(t, "", hostName)
}

func TestProcessData(t *testing.T) {
	data := &processors.FalcoEvent{
		Time:     "2021-08-25T16:45:00Z",
		HostName: "host1",
		Rule:     "DNS Query Logging",
		OutputFields: map[string]interface{}{
			"evt.arg.data": "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H",
			"evt.time":     "2021-08-25T16:45:00Z",
		},
	}
	err := processors.ProcessData(data)
	assert.Nil(t, err)
	assert.Equal(t, "google.com", data.OutputFields["dns_query"])
	_, ok := data.OutputFields["evt.time"]
	assert.False(t, ok)
}

func TestProcessDataMissingEvtArgData(t *testing.T) {
	data := &processors.FalcoEvent{
		Time:     "2021-08-25T16:45:00Z",
		HostName: "host1",
		Rule:     "DNS Query Logging",
		OutputFields: map[string]interface{}{
			"evt.arg.data": "",
			"evt.time":     "2021-08-25T16:45:00Z",
		},
	}
	err := processors.ProcessData(data)
	assert.NotNil(t, err)
}
