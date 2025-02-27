package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bluemountaincyber/falco_postprocessor/internal/processors"
)

var LOGFILE string = "/var/log/falco.json"

func main() {
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	var event processors.FalcoEvent
	if err := json.Unmarshal(inputData, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	if err := processors.ProcessData(&event); err != nil {
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
