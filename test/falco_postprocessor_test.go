package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveDNSQueryHost(t *testing.T) {
	data := "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H"
	hostName, err := retrieveDNSQueryHost(data)
	assert.Nil(t, err)
	assert.Equal(t, "google.com", hostName)
}

func TestRetrieveDNSQueryHostBadBase64(t *testing.T) {
	data := "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H======"
	hostName, err := retrieveDNSQueryHost(data)
	assert.NotNil(t, err)
	assert.Equal(t, "", hostName)
}

func TestProcessData(t *testing.T) {
	data := &FalcoEvent{
		Time:     "2021-08-25T16:45:00Z",
		HostName: "host1",
		Rule:     "DNS Query Logging",
		OutputFields: map[string]interface{}{
			"evt.arg.data": "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H",
			"evt.time":     "2021-08-25T16:45:00Z",
		},
	}
	err := processData(data)
	assert.Nil(t, err)
	assert.Equal(t, "google.com", data.OutputFields["dns_query"])
	_, ok := data.OutputFields["evt.time"]
	assert.False(t, ok)
}

func TestProcessDataMissingEvtArgData(t *testing.T) {
	data := &FalcoEvent{
		Time:     "2021-08-25T16:45:00Z",
		HostName: "host1",
		Rule:     "DNS Query Logging",
		OutputFields: map[string]interface{}{
			"evt.arg.data": "",
			"evt.time":     "2021-08-25T16:45:00Z",
		},
	}
	err := processData(data)
	assert.NotNil(t, err)
}

func TestMain(t *testing.T) {
	input := `{"hostname": "dev-vm", "output_fields": {"evt.arg.data": "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H", "evt.time": 1740244043454771209, "proc.name": "dig", "dns_query": "google.com"}, "priority": "Informational", "rule": "DNS Query Logging", "source": "syscall", "tags": ["network"], "time": "2025-02-22T17:07:23.454771209Z"}`
	tmpfile, err := os.CreateTemp("", "falco_postprocessor_test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(input)); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = tmpfile
	LOGFILE = "/tmp/testing.json"
	main()
	file, err := os.Open(LOGFILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, stat.Size())
	testOutput, err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	testOutputStr := string(buf[:testOutput])

	expectedOutput := "{\"time\":\"2025-02-22T17:07:23.454771209Z\",\"hostname\":\"dev-vm\",\"rule\":\"DNS Query Logging\",\"output_fields\":{\"dns_query\":\"google.com\",\"evt.arg.data\":\"rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+H\",\"proc.name\":\"dig\"}}\n"
	assert.Equal(t, expectedOutput, testOutputStr)

	os.Remove(LOGFILE)
}
