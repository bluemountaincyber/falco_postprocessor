package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bluemountaincyber/falco_postprocessor/internal/outputs"
	"github.com/bluemountaincyber/falco_postprocessor/internal/processors"
)

// Execute runs the root command
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if rootCmd.Flag("help").Changed {
		os.Exit(0)
	}

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

	switch outputType := rootCmd.Flag("output").Value.String(); outputType {
	case "json":
		logfile, err := rootCmd.Flags().GetString("logfile")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting logfile: %v\n", err)
			os.Exit(1)
		}

		outputs.WriteToFile(logfile, output)

	case "awslogs":
		awslogsGroup, err := rootCmd.Flags().GetString("awslogs-group")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting awslogs-group: %v\n", err)
			os.Exit(1)
		}

		awslogsStream, err := rootCmd.Flags().GetString("awslogs-stream")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting awslogs-stream: %v\n", err)
			os.Exit(1)
		}

		awslogsRegion, err := rootCmd.Flags().GetString("awslogs-region")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting awslogs-region: %v\n", err)
			os.Exit(1)
		}

		err = outputs.WriteToCloudWatch(output, awslogsGroup, awslogsStream, awslogsRegion)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to CloudWatch: %v\n", err)
			os.Exit(1)
		}

	case "azurelogs":
		azurelogsDCRStreamUrl, err := rootCmd.Flags().GetString("azurelogs-dcr-stream-url")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting azurelogs-dcr-stream-url: %v\n", err)
			os.Exit(1)
		}

		err = outputs.WriteToMonitor(output, azurelogsDCRStreamUrl)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to Azure Monitor: %v\n", err)
			os.Exit(1)
		}

	default:
		outputs.WriteToStdOut(output)
	}
}
