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

func TestRetrieveDNSQueryHostError(t *testing.T) {
	data := "rh8BIAABAAAAAAABBmdvb2dsZQNjb20AAAEAAQAAKQTQAAAAAAAMAAoACEN4aks2Lk+"
	hostName, err := processors.RetrieveDNSQueryHost(data)
	assert.NotNil(t, err)
	assert.Empty(t, hostName)
}

func TestRetrieveDNSQueryHostBlank(t *testing.T) {
	data := ""
	hostName, err := processors.RetrieveDNSQueryHost(data)
	assert.NotNil(t, err)
	assert.Empty(t, hostName)
}
