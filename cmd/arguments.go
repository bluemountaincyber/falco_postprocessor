package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "falco_postprocessor",
	Short: "A post-processor for Falco events",
	Long: `falco_postprocessor is a tool that processes Falco events and writes them to a log file.
	example: cat falco.json | falco_postprocessor
	`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err)
		}
		logfile, err := cmd.Flags().GetString("logfile")
		if err != nil {
			fmt.Println(err)
		}
		awslogsGroup, err := cmd.Flags().GetString("awslogs-group")
		if err != nil {
			fmt.Println(err)
		}
		awslogsStream, err := cmd.Flags().GetString("awslogs-stream")
		if err != nil {
			fmt.Println(err)
		}
		awslogsRegion, err := cmd.Flags().GetString("awslogs-region")
		if err != nil {
			fmt.Println(err)
		}
		azurelogsDCRStreamUrl, err := cmd.Flags().GetString("azurelogs-dcr-stream-url")
		if err != nil {
			fmt.Println(err)
		}
		if output == "json" && logfile == "" {
			fmt.Println("Error: logfile is required when output is json")
			os.Exit(1)
		}
		if output == "awslogs" && (awslogsGroup == "" || awslogsStream == "" || awslogsRegion == "") {
			fmt.Println("Error: awslogs-group, awslogs-stream, and awslogs-region are required when output is awslogs")
			os.Exit(1)
		}
		if output == "azurelogs" && azurelogsDCRStreamUrl == "" {
			fmt.Println("Error: azurelogs-dcr-stream-url is required when output is azurelogs")
			os.Exit(1)
		}
	},
}

// init initializes the command flags.
func init() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "falcopostprocessor"
	}

	rootCmd.Flags().StringP("logfile", "l", "", "The log file to write to")
	rootCmd.Flags().StringP("output", "o", "stdout", "The output format")
	rootCmd.Flags().StringP("awslogs-group", "g", "FalcoEvents", "The AWS CloudWatch Logs group")
	rootCmd.Flags().StringP("awslogs-stream", "s", hostname, "The AWS CloudWatch Logs stream")
	rootCmd.Flags().StringP("awslogs-region", "r", "us-east-1", "The AWS region")
	rootCmd.Flags().StringP("azurelogs-dcr-stream-url", "d", "", "The Azure Data Collection Stream URL")
}
