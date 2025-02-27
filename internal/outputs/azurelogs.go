package outputs

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/ingestion/azlogs"
)

type FalcoEvent struct {
	EventTime    time.Time `json:"time"`
	Hostname     string    `json:"hostname"`
	Rule         string    `json:"rule"`
	OutputFields string    `json:"output_fields"`
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
	var data []FalcoEvent
	logs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = client.Upload(context.TODO(), ruleId, streamName, logs, nil)
	if err != nil {
		return err
	}
	return nil
}
