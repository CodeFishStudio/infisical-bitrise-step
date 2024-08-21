package helpers

import (
	"fmt"
)

type StepInput struct {
	Client       string
	ClientSecret string
	ProjectId    string
	Env          string
	ApiUrl       string // Optional
	Path         string // Optional
}

func ValidateConfig(config StepInput) error {
	// Check required fields
	if config.Client == "" {
		return fmt.Errorf("client is required")
	}
	if config.ClientSecret == "" {
		return fmt.Errorf("ClientSecret is required")
	}
	if config.ProjectId == "" {
		return fmt.Errorf("ProjectId is required")
	}
	if config.Env == "" {
		return fmt.Errorf("env is required")
	}
	// Optional fields don't need explicit checks unless they must meet certain criteria
	return nil
}
