package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bluemountaincyber/falco_postprocessor/internal/processors"
)

func TestRetrieveIMDSQueryHost(t *testing.T) {
	data := "R0VUIC9tZXRhZGF0YS9pbnN0YW5jZT9hcGktdmVyc2lvbj0yMDE4LTAyLTAxIEhUVFAvMS4xDQpIb3N0OiAxNjkuMjU0LjE2OS4yNTQNCkE="
	hostName, err := processors.RetrieveMetadataAccessPath(data)
	assert.Nil(t, err)
	assert.Equal(t, "/metadata/instance?api-version=2018-02-01", hostName)
}

func TestRetrieveIMDSQueryHostError(t *testing.T) {
	data := "R0VUIC9tZXRhZGF0YS9pbnN0YW5jZT9hcGktdmVyc2lvbj0yMDE4LTAyLTAxIEhUVFAvMS4xDQpIb3N0OiAxNjkuMjU0LjE2OS4yNTQNCkE"
	hostName, err := processors.RetrieveMetadataAccessPath(data)
	assert.NotNil(t, err)
	assert.Empty(t, hostName)
}

func TestRetrieveIMDSQueryHostBlank(t *testing.T) {
	data := ""
	hostName, err := processors.RetrieveMetadataAccessPath(data)
	assert.NotNil(t, err)
	assert.Empty(t, hostName)
}
