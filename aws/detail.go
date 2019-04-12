package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

type EC2StateChangeEventDetail struct {
	InstanceID string `json:"instance-id"`
	State      string `json:"state"`
	Region     string
	Time       time.Time
}

type IdentifiedEC2StateChange struct {
	Name         string
	State        string
	InstanceType string
	Region       string
	Time         time.Time
}

func GetInstanceFromStateChangeDetail(detail EC2StateChangeEventDetail) (*ec2.Instance, error) {
	var instance *ec2.Instance

	s, err := session.NewSession()
	if err != nil {
		return instance, err
	}

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			&detail.InstanceID,
		},
	}

	result, err := ec2.New(s).DescribeInstances(input)
	if err != nil {
		return instance, err
	}

	for _, res := range result.Reservations {
		for _, inst := range res.Instances {
			instance = inst
		}
	}

	return instance, nil
}

func TranslateInstanceToStateChange(detail EC2StateChangeEventDetail, instance *ec2.Instance) IdentifiedEC2StateChange {
	name := *instance.InstanceId

	for _, tag := range instance.Tags {
		if *tag.Key == "Name" {
			name = *tag.Value
			break
		}
	}

	change := IdentifiedEC2StateChange{
		Name: name,
		State: detail.State,
		InstanceType: *instance.InstanceType,
		Region: detail.Region,
		Time: detail.Time,
	}

	return change
}


