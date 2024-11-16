package provider

type ProviderClient interface {
  AssumeRole() error
}
