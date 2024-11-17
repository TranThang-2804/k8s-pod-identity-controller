package provider

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
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

	// Create a new session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		logger.Error(err, "Error creating session:", err)
		return err
	}

	// Create a new STS client
	svc := sts.New(sess)

	// Assume the role
	roleSessionName := "MySession"
	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(c.awsRoleArn),
		RoleSessionName: aws.String(roleSessionName),
	}

	result, err := svc.AssumeRole(input)
	if err != nil {
		logger.Error(err, "Error assuming role:", err)
    return err
	}

	// Print the session token
	logger.Info("Session Token:", *result.Credentials.SessionToken)
	logger.Info("Access Key ID:", *result.Credentials.AccessKeyId)
	logger.Info("Secret Access Key:", *result.Credentials.SecretAccessKey)
	logger.Info("Expiration:", *result.Credentials.Expiration)

	logger.Info("Assuming role successfully")

	return nil
}
