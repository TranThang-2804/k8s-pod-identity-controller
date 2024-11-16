package provider

type AWSProviderClient struct{
  awsRoleArn string
}

func NewAwsProviderClient(awsRoleArn string) (*AWSProviderClient, error) {
	return &AWSProviderClient{
    awsRoleArn: awsRoleArn,
  }, nil
}

func (c *AWSProviderClient) AssumeRole() error {
	return nil
}
