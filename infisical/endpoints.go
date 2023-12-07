package infisical

import "fmt"

type GetSecretRequest struct {
	WorkspaceId string
	Path        string
}

const UserAgent = "Bitrise"

func (api APIConf) GetServiceToken() (ServiceTokenResponse, error) {
	var serviceTokenResponse ServiceTokenResponse

	response, err := api.config.HttpClient.
		R().
		SetResult(&serviceTokenResponse).
		SetHeader("User-Agent", UserAgent).
		Get("/api/v2/service-token")

	if err != nil {
		return ServiceTokenResponse{}, fmt.Errorf("get service token error response before fetching token [err=%s]", err)
	}

	if response.IsError() {
		return ServiceTokenResponse{}, fmt.Errorf("get service token response error [err=%s]", err)
	}

	return serviceTokenResponse, nil
}

func (api APIConf) GetSecrets(params GetSecretRequest) (SecretsResponse, error) {
	var secrets SecretsResponse

	req := api.config.HttpClient.R().
		SetResult(&secrets).
		SetHeader("User-Agent", UserAgent).
		SetQueryParam("environment", api.config.Environment).
		SetQueryParam("workspaceId", params.WorkspaceId)

	if params.Path != "" {
		req.SetQueryParam("secretPath", params.Path)
	}

	response, err := req.Get("/api/v3/secrets")

	if err != nil {
		return SecretsResponse{}, fmt.Errorf("get secret error response before fetching secrets [err=%s]", err)
	}

	if response.IsError() {
		return SecretsResponse{}, fmt.Errorf("get secrets response error [err=%s]", err)
	}

	return secrets, nil

}
