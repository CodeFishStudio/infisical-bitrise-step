package infisical

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Config struct {
	ApiUrl       string
	ServiceToken string
	Environment  string
	HttpClient   *resty.Client
}

type APIConf struct {
	config Config
}

func API(config Config) (*APIConf, error) {
	if config.HttpClient == nil {
		config.HttpClient = resty.New()
		config.HttpClient.SetBaseURL(config.ApiUrl)
	}

	if config.ServiceToken != "" {
		config.HttpClient.SetAuthToken(config.ServiceToken)
	} else {
		return nil, fmt.Errorf("API key is required for access to the infisical API")
	}

	config.HttpClient.SetHeader("Accept", "application/json")

	return &APIConf{config}, nil
}
