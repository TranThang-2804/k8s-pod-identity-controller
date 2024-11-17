package provider

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/log"
)

type AWSProviderClient struct {
	awsRoleArn string
	context    context.Context
}

func NewAwsProviderClient(awsRoleArn string, context context.Context) (*AWSProviderClient, error) {
	return &AWSProviderClient{
		awsRoleArn: awsRoleArn,
		context:    context,
	}, nil
}

func (c *AWSProviderClient) AssumeRole() error {
	logger := log.FromContext(c.context)
  logger.Info("Assuming role successfully")

	return nil
}
