package slack

import (
	"context"
	"os"

	"github.com/slack-go/slack"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &slackProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &slackProvider{}
}

// slackProvider is the provider implementation.
type slackProvider struct{}

// slackProviderModel maps provider schema data to a Go type.
type slackProviderModel struct {
	Token types.String `tfsdk:"token"`
}

// Metadata returns the provider type name.
func (p *slackProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "slack"
}

// Schema defines the provider-level schema for configuration data.
func (p *slackProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				Required:    true,
				Description: "A valid token for the Slack API. May also be provided via SLACK_TOKEN environment variable.",
			},
		},
		Blocks:      map[string]schema.Block{},
		Description: "Interface with the Slack API.",
	}
}

// Configure prepares a Slack API client for data sources and resources.
//
//gocyclo:ignore
func (p *slackProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Slack client")

	// Retrieve provider data from configuration
	var config slackProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown Slack API Token",
			"The provider cannot create the Slack API client as there is an unknown configuration value for the Slack API token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SLACK_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	token := os.Getenv("SLACK_TOKEN")

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}
	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if token == "" {
		resp.Diagnostics.AddAttributeWarning(
			path.Root("token"),
			"Missing Slack API Token",
			"The provider cannot create the Slack API client as there is a missing or empty value for the Slack API token. "+
				"Set the token value in the configuration or use the SLACK_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Mask sensitive value in logs
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "slack_token")

	tflog.Debug(ctx, "Creating Slack client")

	// Enable debugging in the Slack client, if it is enabled for Terraform
	debugProvider := os.Getenv("TF_LOG")
	debug := false
	if (debugProvider == "debug") || (debugProvider == "trace") {
		debug = true
	}

	// Instantiate the client that we will use to talk to the Slack server
	api := slack.New(token, slack.OptionDebug(debug))
	// Test that we have some basic connectivity
	params := slack.NewListReactionsParameters()
	params.Count = int(1)
	_, _, err := api.ListReactions(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Slack API Client",
			"An unexpected error occurred when creating the Slack API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Slack Client Error: "+err.Error(),
		)
		return
	}

	// Make the Slack client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = api
	resp.ResourceData = api

	tflog.Info(ctx, "Configured Slack client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *slackProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUserDataSource,
	}
	//return nil
}

// Resources defines the resources implemented in the provider.
func (p *slackProvider) Resources(_ context.Context) []func() resource.Resource {
	//	return []func() resource.Resource{
	//		NewSlackResource,
	//	}
	return nil
}
