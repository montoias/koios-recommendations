package types

// EnvironmentType define the allowed environments
type EnvironmentType string

const (
	// EnvironmentTest represents the test environment
	EnvironmentTest EnvironmentType = "test"

	// EnvironmentDevelopment represents the development environment
	EnvironmentDevelopment EnvironmentType = "dev"

	// EnvironmentStaging represents the staging environment
	EnvironmentStaging EnvironmentType = "staging"

	// EnvironmentProduction represents the production environment
	EnvironmentProduction EnvironmentType = "prod"
)

// String return the string representation of the EnvironmentType
func (environmentType EnvironmentType) String() string {
	return string(environmentType)
}
