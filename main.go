package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	btp "github.com/bythepixel/cloudwatch-to-slack-lambda/aws"
	"github.com/bythepixel/cloudwatch-to-slack-lambda/slack"
	"log"
	"os"
)

func handler(_ context.Context, event events.CloudWatchEvent) error {
	var detail btp.EC2StateChangeEventDetail
	detail.Time = event.Time
	detail.Region = event.Region

	if err := json.Unmarshal(event.Detail, &detail); err != nil {
		return err
	}

	instance, err := btp.GetInstanceFromStateChangeDetail(detail)
	if err != nil {
		return err
	}

	change := btp.TranslateInstanceToStateChange(detail, instance)

	message := slack.TranslateStateChangeToMessage(change)

	if _, err := slack.NewHttpClient().Send(os.Getenv("SLACK_WEBHOOK"), message); err != nil {
		return err
	}

	log.Printf("Notification has been sent: %v", change)

	return nil
}

func main() {
	lambda.Start(handler)
}
