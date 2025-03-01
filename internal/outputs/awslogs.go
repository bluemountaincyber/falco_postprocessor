package outputs

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// WriteToCloudWatch writes the output to AWS CloudWatch
//
// The input to this function is a byte slice representing the output, the log group, the log stream, and the region.
//
// An expected usage might be:
//
//	WriteToCloudWatch(output, group, stream, region)
func WriteToCloudWatch(output []byte, group string, stream string, region string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return fmt.Errorf("error creating session: %v", err)
	}

	svc := cloudwatchlogs.New(sess)
	err = PutCloudWatchEvent(svc, output, group, stream)

	for err != nil {
		if err.Error() == "ResourceNotFoundException: The specified log group does not exist." {
			_, err = svc.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
				LogGroupName: aws.String(group),
			})

			if err != nil {
				return fmt.Errorf("error creating log group: %v", err)
			}
		}

		if err.Error() == "ResourceNotFoundException: The specified log stream does not exist." {
			_, err = svc.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
				LogGroupName:  aws.String(group),
				LogStreamName: aws.String(stream),
			})

			if err != nil {
				return fmt.Errorf("error creating log stream: %v", err)
			}
		}

		err = PutCloudWatchEvent(svc, output, group, stream)
	}

	return nil
}

// PutCloudWatchEvent puts the output to AWS CloudWatch
//
// The input to this function is a pointer to the CloudWatchLogs service, a byte slice representing the output, the log group, and the log stream.
//
// An expected usage might be:
//
//	PutCloudWatchEvent(svc, output, group, stream)
func PutCloudWatchEvent(svc *cloudwatchlogs.CloudWatchLogs, output []byte, group string, stream string) error {
	_, err := svc.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(string(output)),
				Timestamp: aws.Int64(aws.TimeUnixMilli(time.Now())),
			},
		},
		LogGroupName:  aws.String(group),
		LogStreamName: aws.String(stream),
	})

	return fmt.Errorf("error putting log event: %v", err)
}
