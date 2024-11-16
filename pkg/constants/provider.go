package constants

import (
	"errors"
)

type ProviderType string

const (
	AWS   ProviderType = "aws"
	AZURE ProviderType = "azure"
	GCP   ProviderType = "gcp"
)

func IsValidProviderType(providerType string) error {
	switch providerType {
	case string(AWS):
		return nil
	case string(GCP):
		return nil
	case string(AZURE):
		return nil
	default:
		return errors.New("Invalid provider type: " + providerType)
	}
}
