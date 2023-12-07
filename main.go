package main

import (
	"fmt"
	"github.com/CodeFishStudio/bitrise-step-infisical-secrets/infisical"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type stepInput struct {
	Token  string
	Env    string
	ApiUrl string
	Path   string
}

func main() {

	//Get ENVs for step
	infisicalEnvs := stepInput{
		Token:  os.Getenv("infisical_token"),
		Env:    os.Getenv("infisical_env"),
		ApiUrl: os.Getenv("infisical_url"),
		Path:   os.Getenv("infisical_path"),
	}

	serviceTokenParts := strings.SplitN(infisicalEnvs.Token, ".", 4)
	if len(serviceTokenParts) < 4 {
		log.Fatalln("Invalid token provided")
	}

	//Configure API Client
	api, err := infisical.API(infisical.Config{
		ApiUrl:       infisicalEnvs.ApiUrl,
		ServiceToken: infisicalEnvs.Token,
		Environment:  infisicalEnvs.Env,
		HttpClient:   nil,
	})

	serviceTokenDetails, err := api.GetServiceToken()
	if err != nil {
		log.Fatalln(err)
	}

	encryptedSecrets, err := api.GetSecrets(infisical.GetSecretRequest{
		WorkspaceId: serviceTokenDetails.Workspace,
		Path:        infisicalEnvs.Path,
	})
	if err != nil {
		log.Fatalln(err)
	}

	decodedEncryptionDetails, err := infisical.GetBase64DecodedSymmetricEncryptionDetails(serviceTokenParts[3],
		serviceTokenDetails.EncryptedKey, serviceTokenDetails.Iv, serviceTokenDetails.Tag)
	if err != nil {
		log.Fatalln(err)
	}

	plainTextWorkspaceKey, err := infisical.DecryptSymmetric([]byte(serviceTokenParts[3]),
		decodedEncryptionDetails.Cipher, decodedEncryptionDetails.Tag, decodedEncryptionDetails.IV)
	if err != nil {
		log.Fatalln(err)
	}

	plainTextSecrets, err := infisical.GetPlainTextSecrets(plainTextWorkspaceKey, encryptedSecrets.Secrets)
	if err != nil {
		log.Fatalln(err)
	}

	for _, secret := range plainTextSecrets {
		// --- Step Outputs: Export Environment Variables for other Steps:
		// You can export Environment Variables for other Steps with
		//  envman, which is automatically installed by `bitrise setup`.
		// A very simple example:
		cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", secret.Key, "--value",
			secret.Value).CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
			os.Exit(1)
		}
	}

	secretCountStr := strconv.Itoa(len(plainTextSecrets))

	fmt.Printf("\033[1;32m%s\033[0m", "Injecting "+secretCountStr+
		" Infisical secrets into your application process\n")
	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
