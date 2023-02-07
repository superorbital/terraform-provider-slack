terraform {
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

# Create a new Slack user
resource "slack_user" "test" {
  real_name = "Brighid Quinn"
}

# Read in a existing Slack user
data "slack_user" "example" {
 id = "U99ZZ9USZ9Z00"
}
