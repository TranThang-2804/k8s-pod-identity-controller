package provider

type AWSProviderClient struct{}

func NewAwsProviderClient() (*AWSProviderClient, error) {
	return &AWSProviderClient{}, nil
}

func (c *AWSProviderClient) AssumeRole() error {
	return nil
}
