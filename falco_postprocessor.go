package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var LOGFILE string = "/var/log/falco.json"

type FalcoEvent struct {
	Time         string                 `json:"time"`
	HostName     string                 `json:"hostname"`
	Rule         string                 `json:"rule"`
	OutputFields map[string]interface{} `json:"output_fields"`
}

func retrieveDNSQueryHost(data string) (string, error) {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	var domain []byte
	byteCounter := 13
	if len(payload) < byteCounter {
		return "", fmt.Errorf("invalid payload")
	}
	wordLen := int(payload[12])

	for {
		domain = append(domain, payload[byteCounter:byteCounter+wordLen]...)
		if payload[byteCounter+wordLen] == 0 {
			break
		} else {
			domain = append(domain, '.')
			byteCounter += wordLen + 1
			wordLen = int(payload[byteCounter-1])
		}
	}

	return string(domain), nil
}

func processData(data *FalcoEvent) error {
	if data.Rule == "DNS Query Logging" {
		hostName, err := retrieveDNSQueryHost(data.OutputFields["evt.arg.data"].(string))
		if err != nil {
			return err
		}
		data.OutputFields["dns_query"] = hostName
	}
	delete(data.OutputFields, "evt.time")
	return nil
}

func main() {
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	var event FalcoEvent
	if err := json.Unmarshal(inputData, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	if err := processData(&event); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing data: %v\n", err)
		os.Exit(1)
	}

	output, err := json.Marshal(event)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening log file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(string(output) + "\n"); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to log file: %v\n", err)
		os.Exit(1)
	}
}
