package slack

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/joho/godotenv"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the Slack client is properly configured.
	// You should use the TF_VAR_slack_token environment variable
	// to pass in the token.
	providerConfig = `terraform {
  required_providers {
    slack = {
      source = "superorbital/slack"
    }
  }
}

variable slack_token {
  # You can set this in the environment via TF_VAR_slack_token
  type = string
  description = "A valid Slack API Token"
}

# Configure the connection details for Slack
provider "slack" {
  token = var.slack_token
}
`
	testEnvVarRoot = `terraform_provider_slack_test_`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"slack": providerserver.NewProtocol6WithError(New()),
	}

	// These test variables should be set in a .env file at the git repo root.
	slackTestUserID         = ""
	slackTestConversationID = ""
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	// Read in valid values for our acceptance tests
	if os.Getenv("TF_ACC") == "1" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading required .env file for acceptance tests.")
		}
		slackTestConversationID = os.Getenv(testEnvVarRoot + "conversation_id")
		slackTestUserID = os.Getenv(testEnvVarRoot + "user_id")
	}
}
