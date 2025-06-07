package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmtypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func GetEC2Instances(region, profile string) ([]Instance, error) {
	ctx := context.TODO()
	var cfgOpts []func(*config.LoadOptions) error
	cfgOpts = append(cfgOpts, config.WithSharedConfigProfile(profile))
	if region != "" {
		cfgOpts = append(cfgOpts, config.WithRegion(region))
	}
	cfg, err := config.LoadDefaultConfig(ctx, cfgOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}
	ec2Client := ec2.NewFromConfig(cfg)
	ssmClient := ssm.NewFromConfig(cfg)

	desc, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		Filters: []ec2types.Filter{{
			Name:   aws.String("instance-state-name"),
			Values: []string{"running"},
		}},
	})
	if err != nil {
		return nil, err
	}

	instances := []Instance{}
	for _, reservation := range desc.Reservations {
		for _, inst := range reservation.Instances {
			id := aws.ToString(inst.InstanceId)
			ssmInfo, err := ssmClient.DescribeInstanceInformation(ctx, &ssm.DescribeInstanceInformationInput{
				Filters: []ssmtypes.InstanceInformationStringFilter{{
					Key:    aws.String("InstanceIds"),
					Values: []string{id},
				}},
			})
			if err != nil {
				fmt.Printf("Error describing instance information for instance %s: %v\n", id, err)
				continue
			}
			if len(ssmInfo.InstanceInformationList) == 0 {
				fmt.Printf("No instance information found for instance %s\n", id)
				continue
			}
			name := id
			for _, tag := range inst.Tags {
				if aws.ToString(tag.Key) == "Name" {
					name = aws.ToString(tag.Value)
				}
			}
			instances = append(instances, Instance{ID: id, Name: name})
		}
	}
	return instances, nil
}
