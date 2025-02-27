# Falco Post Processor

## Overview

This is a post processor for Falco. It takes the output of Falco, parses it, enriches it with additional information, and writes its output to a chosen location.

## Installation

To install the post processor, simply extract the current release into a directory of your choice.

## Configuration

There are sample Falco configuration and rules files in the [falco](./falco) directory. You can use these as a starting point for your own configuration.

## Supported Outputs

The post processor supports the following output (`-o` flag) formats:

- `stdout`: The output is written to the standard output. This is the default behavior.
- `json`: The output is written to a JSON file. The file name is specified with the `-f` flag.
- `none`: The output is discarded.
- `awslogs`: The output is written to AWS CloudWatch Logs. The log group and stream names are specified with the `-g` and `-s` flags, respectively.
- `azurelogs`: The output is written to Azure Monitor Logs. The log workspace ID and key are specified with the `-w` and `-k` flags, respectively.

## Usage

When referencing this post processor in the Falco configuration, you can use the following logic:

```yaml
program_output:
  enabled: true 
  keep_alive: false
  program: "/path/to/falco_postprocessor -o json -f /path/to/falco_logs.json"
```
