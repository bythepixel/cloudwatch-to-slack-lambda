package slack

import (
	"fmt"
	btp "github.com/bythepixel/cloudwatch-to-slack-lambda/aws"
	"os"
)

type Message struct {
	Text        string              `json:"text"`
	Attachments []MessageAttachment `json:"attachments"`
}

type MessageAttachment struct {
	Fallback   string            `json:"fallback"`
	Color      string            `json:"color"`
	AuthorName string            `json:"author_name"`
	Title      string            `json:"title"`
	TitleLink  string            `json:"title_link"`
	Text       string            `json:"text"`
	Fields     []AttachmentField `json:"fields"`
	Footer     string            `json:"footer"`
	FooterIcon string            `json:"footer_icon"`
	Timestamp  int64             `json:"ts"`
}

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func TranslateStateChangeToMessage(change btp.IdentifiedEC2StateChange) Message {
	text := fmt.Sprintf("*%s* in %s changed state: %s", change.Name, change.Region, change.State)

	return Message{
		Text: text,
		Attachments: []MessageAttachment{
			{
				Fallback: text,
				Color: "#ff9900",
				AuthorName: "CloudWatch Event",
				Title: change.Name,
				TitleLink: fmt.Sprintf("https://console.aws.amazon.com/ec2/v2/home?region=%s#Instances:sort=instanceId", change.Region),
				Fields: []AttachmentField{
					{
						Title: "State",
						Value: change.State,
						Short: true,
					},
					{
						Title: "Type",
						Value: change.InstanceType,
						Short: true,
					},
				},
				Footer: "Amazon AWS",
				FooterIcon: os.Getenv("SLACK_MESSAGE_FOOTER_ICON"),
				Timestamp: change.Time.Unix(),
			},
		},
	}
}
