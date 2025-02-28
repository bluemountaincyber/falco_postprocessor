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
		return err
	}
	client, err := azlogs.NewClient("https://"+endpoint, credential, nil)
	if err != nil {
		return err
	}
	var outputMap map[string]interface{}
	if err := json.Unmarshal(output, &outputMap); err != nil {
		return err
	}
	var events []FalcoEvent
	outputFieldsJSON, err := json.Marshal(outputMap["output_fields"])
	if err != nil {
		return err
	}
	events = append(events, FalcoEvent{
		EventTime:    outputMap["time"].(string),
		Hostname:     outputMap["hostname"].(string),
		Rule:         outputMap["rule"].(string),
		OutputFields: string(outputFieldsJSON),
	})
	logs, err := json.Marshal(events)
	fmt.Println(string(logs))
	if err != nil {
		return err
	}
	_, err = client.Upload(context.TODO(), ruleId, streamName, logs, nil)
	if err != nil {
		return err
	}
	return nil
}
