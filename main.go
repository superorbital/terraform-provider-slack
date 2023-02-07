package main

import (
	"context"

	"github.com/superorbital/terraform-provider-slack/slack"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name slack

func main() {
	providerserver.Serve(context.Background(), slack.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/superorbital/slack",
	})
}
