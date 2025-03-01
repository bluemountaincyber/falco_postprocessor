package outputs

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/ingestion/azlogs"
)

type FalcoEvent struct {
	EventTime    string `json:"EventTime"`
	Hostname     string `json:"Hostname"`
	Rule         string `json:"Rule"`
	OutputFields string `json:"OutputFields"`
}

func WriteToMonitor(output []byte, stream_url string) error {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	endpoint := strings.Split(stream_url, "/")[2]
	ruleId := strings.Split(stream_url, "/")[4]
	streamName := strings.Split(strings.Split(stream_url, "/")[6], "?")[0]

	if err != nil {
		return fmt.Errorf("failed to get Azure credential: %v", err)
	}

	client, err := azlogs.NewClient("https://"+endpoint, credential, nil)

	if err != nil {
		return fmt.Errorf("failed to create Azure Monitor client: %v", err)
	}

	var outputMap map[string]interface{}

	if err := json.Unmarshal(output, &outputMap); err != nil {
		return fmt.Errorf("failed to unmarshal output: %v", err)
	}

	var events []FalcoEvent

	outputFieldsJSON, err := json.Marshal(outputMap["output_fields"])

	if err != nil {
		return fmt.Errorf("failed to marshal output fields: %v", err)
	}

	events = append(events, FalcoEvent{
		EventTime:    outputMap["time"].(string),
		Hostname:     outputMap["hostname"].(string),
		Rule:         outputMap["rule"].(string),
		OutputFields: string(outputFieldsJSON),
	})

	logs, err := json.Marshal(events)

	if err != nil {
		return fmt.Errorf("failed to marshal logs: %v", err)
	}

	_, err = client.Upload(context.TODO(), ruleId, streamName, logs, nil)

	if err != nil {
		return fmt.Errorf("failed to upload logs: %v", err)
	}

	return nil
}
