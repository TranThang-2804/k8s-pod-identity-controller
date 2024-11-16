package constants

type ProviderType string

const (
	AWS   ProviderType = "aws"
	AZURE ProviderType = "azure"
	GCP   ProviderType = "gcp"
)

func IsValidProviderType(providerType string) bool {
	switch providerType {
	case string(AWS), string(AZURE), string(GCP):
		return true
	default:
		return false
	}
}
