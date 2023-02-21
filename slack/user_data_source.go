package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

// NewUserDataSource is a helper function to simplify the provider implementation.
func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

// userDataSource is the data source implementation.
type userDataSource struct {
	client *slack.Client
}

// userDataSourceModel maps the data source schema data.
type userDataSourceModel struct {
	User userModel `tfsdk:"user"`
}

type userModel struct {
	ID                types.String         `tfsdk:"id"`
	Color             types.String         `tfsdk:"color"`
	Deleted           types.Bool           `tfsdk:"deleted"`
	EnterpriseUser    *enterpriseUserModel `tfsdk:"enterprise_user"`
	IsAdmin           types.Bool           `tfsdk:"is_admin"`
	IsAppUser         types.Bool           `tfsdk:"is_app_user"`
	IsBot             types.Bool           `tfsdk:"is_bot"`
	IsStranger        types.Bool           `tfsdk:"is_stranger"`
	IsOwner           types.Bool           `tfsdk:"is_owner"`
	IsPrimaryOwner    types.Bool           `tfsdk:"is_primary_owner"`
	IsRestricted      types.Bool           `tfsdk:"is_restricted"`
	IsUltraRestricted types.Bool           `tfsdk:"is_ultra_restricted"`
	Name              types.String         `tfsdk:"name"`
	Profile           *userProfileModel    `tfsdk:"profile"`
	RealName          types.String         `tfsdk:"real_name"`
	TeamID            types.String         `tfsdk:"team_id"`
	TZ                types.String         `tfsdk:"tz"`
	TZLabel           types.String         `tfsdk:"tz_label"`
	TZOffset          types.Int64          `tfsdk:"tz_offset"`
	Updated           types.String         `tfsdk:"updated"`
}

type userProfileModel struct {
	DisplayName           types.String `tfsdk:"display_name"`
	DisplayNameNormalized types.String `tfsdk:"display_name_normalized"`
	//Fields types.Map `tfsdk:"fields"`
	FirstName          types.String `tfsdk:"first_name"`
	Image192           types.String `tfsdk:"image_192"`
	Image24            types.String `tfsdk:"image_24"`
	Image32            types.String `tfsdk:"image_32"`
	Image48            types.String `tfsdk:"image_48"`
	Image512           types.String `tfsdk:"image_512"`
	Image72            types.String `tfsdk:"image_72"`
	ImageOriginal      types.String `tfsdk:"image_original"`
	LastName           types.String `tfsdk:"last_name"`
	Phone              types.String `tfsdk:"phone"`
	RealName           types.String `tfsdk:"real_name"`
	RealNameNormalized types.String `tfsdk:"real_name_normalized"`
	StatusEmoji        types.String `tfsdk:"status_emoji"`
	StatusExpiration   types.Int64  `tfsdk:"status_expiration"`
	StatusText         types.String `tfsdk:"status_text"`
	Team               types.String `tfsdk:"team"`
	Title              types.String `tfsdk:"title"`
}

type enterpriseUserModel struct {
	EnterpriseID   types.String `tfsdk:"enterprise_id"`
	EnterpriseName types.String `tfsdk:"enterprise_name"`
	ID             types.String `tfsdk:"id"`
	IsAdmin        types.Bool   `tfsdk:"is_admin"`
	IsOwner        types.Bool   `tfsdk:"is_owner"`
	Teams          types.List   `tfsdk:"teams"`
}

// Configure adds the provider configured client to the data source.
func (d *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*slack.Client)

}

// Metadata returns the data source type name.
func (d *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema defines the schema for the data source.
func (d *userDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetch a user.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this workspace user.",
				Required:    true,
			},
			"real_name": schema.StringAttribute{
				Description: "The user's first and last name",
				Computed:    true,
			},
			"color": schema.StringAttribute{
				Description: "Used in some clients to display a special username color.",
				Computed:    true,
			},
			"deleted": schema.BoolAttribute{
				Description: "This user has been deactivated when the value of this field is true.",
				Computed:    true,
			},
			"is_admin": schema.BoolAttribute{
				Description: "Indicates whether the user is an Admin of the current workspace.",
				Computed:    true,
			},
			"is_app_user": schema.BoolAttribute{
				Description: "Indicates whether the user is an authorized user of the calling app.",
				Computed:    true,
			},
			"is_bot": schema.BoolAttribute{
				Description: "Indicates whether the user is a bot user.",
				Computed:    true,
			},
			"is_stranger": schema.BoolAttribute{
				Description: "If true, this user belongs to a different workspace than the one associated with your app's token, and isn't in any shared channels visible to your app.",
				Computed:    true,
			},
			"is_owner": schema.BoolAttribute{
				Description: "Indicates whether the user is an Owner of the current workspace.",
				Computed:    true,
			},
			"is_primary_owner": schema.BoolAttribute{
				Description: "Indicates whether the user is the Primary Owner of the current workspace.",
				Computed:    true,
			},
			"is_restricted": schema.BoolAttribute{
				Description: "Indicates whether or not the user is a guest user.",
				Computed:    true,
			},
			"is_ultra_restricted": schema.BoolAttribute{
				Description: "Indicates whether or not the user is a single-channel guest.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Deprecated. It once indicated the preferred username for a user.",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "Identifier for this workspace user's team.",
				Computed:    true,
			},
			"tz": schema.StringAttribute{
				Description: "A human-readable string for the geographic timezone-related region this user has specified in their account.",
				Computed:    true,
			},
			"tz_label": schema.StringAttribute{
				Description: "Describes the commonly used name of the timezone defined in tz.",
				Computed:    true,
			},
			"tz_offset": schema.Int64Attribute{
				Description: "Indicates the number of seconds to offset UTC by for this user's timezone.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "A Unix timestamp indicating when the user object was last updated.",
				Computed:    true,
			},
			"profile": schema.SingleNestedAttribute{
				Description: "The profile object contains the default fields of a user's workspace profile.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"title": schema.StringAttribute{
						Description: "The user's title.",
						Computed:    true,
					},
					"phone": schema.StringAttribute{
						Description: "The user's phone number, in any format.",
						Computed:    true,
					},
					"real_name": schema.StringAttribute{
						Description: "The user's first and last name.",
						Computed:    true,
					},
					"real_name_normalized": schema.StringAttribute{
						Description: "The real_name field, but with any non-Latin characters filtered out.",
						Computed:    true,
					},
					"display_name": schema.StringAttribute{
						Description: "The display name the user has chosen to identify themselves by in their workspace profile.",
						Computed:    true,
					},
					"display_name_normalized": schema.StringAttribute{
						Description: "The display_name field, but with any non-Latin characters filtered out.",
						Computed:    true,
					},
					"status_text": schema.StringAttribute{
						Description: "The displayed text of up to 100 characters.",
						Computed:    true,
					},
					"status_emoji": schema.StringAttribute{
						Description: "The displayed emoji that is enabled for the Slack team, such as :train:.",
						Computed:    true,
					},
					"status_expiration": schema.Int64Attribute{
						Description: "The Unix Timestamp of when the status will expire.",
						Computed:    true,
					},
					"image_original": schema.StringAttribute{
						Description: "Contains the URL for the original square ratio, web-viewable images (GIFs, JPEGs, or PNGs) that represent a user's profile picture.",
						Computed:    true,
					},
					"first_name": schema.StringAttribute{
						Description: "The user's first name.",
						Computed:    true,
					},
					"last_name": schema.StringAttribute{
						Description: "The user's last name.",
						Computed:    true,
					},
					"image_24": schema.StringAttribute{
						Description: "Contains the URL for the 24-pixel square ratio, web-viewable images (GIFs, JPEGs, or PNGs) that represent a user's profile picture.",
						Computed:    true,
					},
					"image_32": schema.StringAttribute{
						Description: "Contains the URL for the 32-pixel square ratio, web-viewable images (GIFs, JPEGs, or PNGs) that represent a user's profile picture.",
						Computed:    true,
					},
					"image_48": schema.StringAttribute{
						Description: "Contains the URL for the 48-pixel square ratio, web-viewable images (GIFs, JPEGs, or PNGs) that represent a user's profile picture.",
						Computed:    true,
					},
					"image_72": schema.StringAttribute{
						Description: "Contains the URL for the 72-pixel square ratio, web-viewable images (GIFs, JPEGs, or PNGs) that represent a user's profile picture.",
						Computed:    true,
					},
					"image_192": schema.StringAttribute{
						Description: "Contains the URL for the 192-pixel square ratio, web-viewable images (GIFs, JPEGs, or PNGs) that represent a user's profile picture.",
						Computed:    true,
					},
					"image_512": schema.StringAttribute{
						Description: "Contains the URL for the 512-pixel square ratio, web-viewable images (GIFs, JPEGs, or PNGs) that represent a user's profile picture.",
						Computed:    true,
					},
					"team": schema.StringAttribute{
						Description: "The user's team ID.",
						Computed:    true,
					},
				},
			},
			"enterprise_user": schema.SingleNestedAttribute{
				Description: "An object containing info related to an Enterprise Grid user.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"enterprise_id": schema.StringAttribute{
						Description: "A unique ID for the Enterprise Grid organization this user belongs to.",
						Computed:    true,
					},
					"enterprise_name": schema.StringAttribute{
						Description: "A display name for the Enterprise Grid organization.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "This user's ID - some Grid users have a kind of dual identity — a local, workspace-centric user ID as well as a Grid-wise user ID, called the Enterprise user ID.",
						Computed:    true,
					},
					"is_admin": schema.BoolAttribute{
						Description: "Indicates whether the user is an Admin of the Enterprise Grid organization.",
						Computed:    true,
					},
					"is_owner": schema.BoolAttribute{
						Description: "Indicates whether the user is an Owner of the Enterprise Grid organization.",
						Computed:    true,
					},
					"teams": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "An array of workspace IDs that are in the Enterprise Grid organization.",
						Computed:    true,
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read user data source")
	var state userModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	userResponse, err := d.client.GetUserInfo(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read User",
			err.Error(),
		)
		return
	}

	userProfileData := userProfileModel{
		DisplayName:           types.StringValue(userResponse.Profile.DisplayName),
		DisplayNameNormalized: types.StringValue(userResponse.Profile.DisplayNameNormalized),
		FirstName:             types.StringValue(userResponse.Profile.FirstName),
		Image192:              types.StringValue(userResponse.Profile.Image192),
		Image24:               types.StringValue(userResponse.Profile.Image24),
		Image32:               types.StringValue(userResponse.Profile.Image32),
		Image48:               types.StringValue(userResponse.Profile.Image48),
		Image512:              types.StringValue(userResponse.Profile.Image512),
		Image72:               types.StringValue(userResponse.Profile.Image72),
		ImageOriginal:         types.StringValue(userResponse.Profile.ImageOriginal),
		LastName:              types.StringValue(userResponse.Profile.LastName),
		Phone:                 types.StringValue(userResponse.Profile.Phone),
		RealName:              types.StringValue(userResponse.Profile.RealName),
		RealNameNormalized:    types.StringValue(userResponse.Profile.RealNameNormalized),
		StatusEmoji:           types.StringValue(userResponse.Profile.StatusEmoji),
		StatusExpiration:      types.Int64Value(int64(userResponse.Profile.StatusExpiration)),
		StatusText:            types.StringValue(userResponse.Profile.StatusText),
		Team:                  types.StringValue(userResponse.Profile.Team),
		Title:                 types.StringValue(userResponse.Profile.Title),
	}

	enterpriseTeams, diags := types.ListValueFrom(ctx, types.StringType, userResponse.Enterprise.Teams)
	resp.Diagnostics.Append(diags...)

	enterpriseUserProfileData := enterpriseUserModel{
		EnterpriseID:   types.StringValue(userResponse.Enterprise.EnterpriseID),
		EnterpriseName: types.StringValue(userResponse.Enterprise.EnterpriseName),
		ID:             types.StringValue(userResponse.Enterprise.ID),
		IsAdmin:        types.BoolValue(userResponse.Enterprise.IsAdmin),
		IsOwner:        types.BoolValue(userResponse.Enterprise.IsOwner),
		Teams:          enterpriseTeams,
	}

	// Map response body to model
	state = userModel{
		Color:             types.StringValue(userResponse.Color),
		Deleted:           types.BoolValue(userResponse.Deleted),
		EnterpriseUser:    &enterpriseUserProfileData,
		ID:                types.StringValue(userResponse.ID),
		IsAdmin:           types.BoolValue(userResponse.IsAdmin),
		IsAppUser:         types.BoolValue(userResponse.IsAppUser),
		IsBot:             types.BoolValue(userResponse.IsBot),
		IsOwner:           types.BoolValue(userResponse.IsOwner),
		IsPrimaryOwner:    types.BoolValue(userResponse.IsPrimaryOwner),
		IsRestricted:      types.BoolValue(userResponse.IsRestricted),
		IsStranger:        types.BoolValue(userResponse.IsStranger),
		IsUltraRestricted: types.BoolValue(userResponse.IsUltraRestricted),
		Name:              types.StringValue(userResponse.Name),
		Profile:           &userProfileData,
		RealName:          types.StringValue(userResponse.RealName),
		TeamID:            types.StringValue(userResponse.TeamID),
		TZ:                types.StringValue(userResponse.TZ),
		TZLabel:           types.StringValue(userResponse.TZLabel),
		TZOffset:          types.Int64Value(int64(userResponse.TZOffset)),
		Updated:           types.StringValue(userResponse.Updated.Time().Format("Mon Jan 2 15:04:05 MST 2006")),
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Debug(ctx, "Read user data source", map[string]any{"success": true})
}
