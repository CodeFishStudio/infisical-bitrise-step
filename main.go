package main

import (
	"fmt"
	"github.com/CodeFishStudio/bitrise-step-infisical-secrets/helpers"
	infisical "github.com/infisical/go-sdk"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	//Get ENVs for step
	infisicalEnvs := helpers.StepInput{
		Client:       os.Getenv("infisical_client"),
		ClientSecret: os.Getenv("infisical_client_secret"),
		ProjectId:    os.Getenv("infisical_project_id"),
		Env:          os.Getenv("infisical_env"),
		ApiUrl:       os.Getenv("infisical_url"),
		Path:         os.Getenv("infisical_path"),
	}

	err := helpers.ValidateConfig(infisicalEnvs)
	if err != nil {
		log.Fatalln(err)
	}

	//Set config, Cloud API is default site URL.
	var config = infisical.Config{
		UserAgent: "Bitrise-Step-infisical-secrets",
	}

	if infisicalEnvs.ApiUrl != "" {
		config.SiteUrl = infisicalEnvs.ApiUrl
	}

	client := infisical.NewInfisicalClient(config)

	_, err = client.Auth().UniversalAuthLogin(infisicalEnvs.Client, infisicalEnvs.ClientSecret)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	secretsList, err := client.Secrets().List(infisical.ListSecretsOptions{
		ProjectID:          infisicalEnvs.ProjectId,
		Environment:        infisicalEnvs.Env,
		SecretPath:         infisicalEnvs.Path,
		AttachToProcessEnv: false,
	})
	if err != nil {
		log.Fatalf("Secret Error: %v", err)
	}

	for _, secret := range secretsList {
		// --- Step Outputs: Export Environment Variables for other Steps:
		// You can export Environment Variables for other Steps with
		//  envman, which is automatically installed by `bitrise setup`.
		// A very simple example:
		cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", secret.SecretKey, "--value",
			secret.SecretValue).CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
			os.Exit(1)
		}

		if os.Getenv("BITRISE_DEBUG") == "true" {
			fmt.Printf("%s=%s\n", secret.SecretKey, secret.SecretValue)
		}
	}

	secretCountStr := strconv.Itoa(len(secretsList))

	fmt.Printf("\033[1;32m%s\033[0m", "Injecting "+secretCountStr+
		" Infisical secrets into your application process\n")

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
