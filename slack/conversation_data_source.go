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
	_ datasource.DataSource              = &conversationDataSource{}
	_ datasource.DataSourceWithConfigure = &conversationDataSource{}
)

// NewConversationDataSource is a helper function to simplify the provider implementation.
func NewConversationDataSource() datasource.DataSource {
	return &conversationDataSource{}
}

// conversationDataSource is the data source implementation.
type conversationDataSource struct {
	client *slack.Client
}

// conversationDataSourceModel maps the data source schema data.
type conversationDataSourceModel struct {
	Conversation conversationModel `tfsdk:"channels"`
}

type conversationModel struct {
	ID                 types.String  `tfsdk:"id"`
	Created            types.String  `tfsdk:"created"`
	Creator            types.String  `tfsdk:"creator"`
	IsArchived         types.Bool    `tfsdk:"is_archived"`
	IsChannel          types.Bool    `tfsdk:"is_channel"`
	IsExtShared        types.Bool    `tfsdk:"is_ext_shared"`
	IsGeneral          types.Bool    `tfsdk:"is_general"`
	IsGroup            types.Bool    `tfsdk:"is_group"`
	IsIM               types.Bool    `tfsdk:"is_im"`
	IsMember           types.Bool    `tfsdk:"is_member"`
	IsOpen             types.Bool    `tfsdk:"is_open"`
	IsOrgShared        types.Bool    `tfsdk:"is_org_shared"`
	IsPendingExtShared types.Bool    `tfsdk:"is_pending_ext_shared"`
	IsPrivate          types.Bool    `tfsdk:"is_private"`
	IsShared           types.Bool    `tfsdk:"is_shared"`
	LastRead           types.String  `tfsdk:"last_read"`
	Latest             *latestModel  `tfsdk:"latest"`
	Locale             types.String  `tfsdk:"locale"`
	Name               types.String  `tfsdk:"name"`
	NameNormalized     types.String  `tfsdk:"name_normalized"`
	NumMembers         types.Int64   `tfsdk:"num_members"`
	Priority           types.Int64   `tfsdk:"priority"`
	Purpose            *purposeModel `tfsdk:"purpose"`
	Topic              *topicModel   `tfsdk:"topic"`
	Unlinked           types.Int64   `tfsdk:"unlinked"`
	UnreadCount        types.Int64   `tfsdk:"unread_count"`
	UnreadCountDisplay types.Int64   `tfsdk:"unread_count_display"`
	User               types.String  `tfsdk:"user"`
}

type topicModel struct {
	Creator types.String `tfsdk:"creator"`
	LastSet types.Int64  `tfsdk:"last_set"`
	Value   types.String `tfsdk:"value"`
}

type purposeModel struct {
	Creator types.String `tfsdk:"creator"`
	LastSet types.Int64  `tfsdk:"last_set"`
	Value   types.String `tfsdk:"value"`
}

type latestModel struct {
	Type types.String `tfsdk:"type"`
	User types.String `tfsdk:"user"`
	Text types.String `tfsdk:"text"`
	TS   types.String `tfsdk:"ts"`
}

// Configure adds the provider configured client to the data source.
func (d *conversationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*slack.Client)

}

// Metadata returns the data source type name.
func (d *conversationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_conversation"
}

// Schema defines the schema for the data source.
func (d *conversationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetch a conversation.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this conversation.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "A Unix timestamp indicating when the conversation was created.",
				Computed:    true,
			},
			"creator": schema.StringAttribute{
				Description: "The ID for the user that created the conversation.",
				Computed:    true,
			},
			"is_archived": schema.BoolAttribute{
				Description: "Indicates whether the conversation is archived.",
				Computed:    true,
			},
			"is_channel": schema.BoolAttribute{
				Description: "Indicates whether the conversation is a channel.",
				Computed:    true,
			},
			"is_ext_shared": schema.BoolAttribute{
				Description: "Indicates whether the conversation is externally shared.",
				Computed:    true,
			},
			"is_general": schema.BoolAttribute{
				Description: "Indicates whether the conversation is general.",
				Computed:    true,
			},
			"is_group": schema.BoolAttribute{
				Description: "Indicates whether the conversation is a group.",
				Computed:    true,
			},
			"is_im": schema.BoolAttribute{
				Description: "Indicates whether the conversation is an IM.",
				Computed:    true,
			},
			"is_member": schema.BoolAttribute{
				Description: "Indicates whether the conversation is a member.",
				Computed:    true,
			},
			"is_open": schema.BoolAttribute{
				Description: "Indicates whether the conversation is open.",
				Computed:    true,
			},
			"is_org_shared": schema.BoolAttribute{
				Description: "Indicates whether the conversation is org shared.",
				Computed:    true,
			},
			"is_pending_ext_shared": schema.BoolAttribute{
				Description: "Indicates whether the conversation is a pending external share.",
				Computed:    true,
			},
			"is_private": schema.BoolAttribute{
				Description: "Indicates whether the conversation is private.",
				Computed:    true,
			},
			"is_shared": schema.BoolAttribute{
				Description: "Indicates whether the conversation is shared.",
				Computed:    true,
			},
			"last_read": schema.StringAttribute{
				Description: "The last time the conversation was read.",
				Computed:    true,
			},
			"locale": schema.StringAttribute{
				Description: "The locale set for the conversation.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the conversation.",
				Computed:    true,
			},
			"name_normalized": schema.StringAttribute{
				Description: "The name field, but with any non-Latin characters filtered out.",
				Computed:    true,
			},
			"num_members": schema.Int64Attribute{
				Description: "The number of members in the conversation.",
				Computed:    true,
			},
			"priority": schema.Int64Attribute{
				Description: "The conversation's priority value.",
				Computed:    true,
			},
			"unlinked": schema.Int64Attribute{
				Description: "Conversation unlinked value.",
				Computed:    true,
			},
			"unread_count": schema.Int64Attribute{
				Description: "The count of unread messages in the conversation.",
				Computed:    true,
			},
			"unread_count_display": schema.Int64Attribute{
				Description: "The unread count display value.",
				Computed:    true,
			},
			"user": schema.StringAttribute{
				Description: "The ID for the user who started the conversation.",
				Computed:    true,
			},
			"purpose": schema.SingleNestedAttribute{
				Description: "The conversation's purpose object.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"creator": schema.StringAttribute{
						Description: "The ID for the creator of the purpose.",
						Computed:    true,
					},
					"last_set": schema.Int64Attribute{
						Description: "A Unix timestamp indicating when the purpose was last set.",
						Computed:    true,
					},
					"value": schema.StringAttribute{
						Description: "The conversation's purpose.",
						Computed:    true,
					},
				},
			},
			"topic": schema.SingleNestedAttribute{
				Description: "The conversation's topic object.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"creator": schema.StringAttribute{
						Description: "The ID for the creator of the topic.",
						Computed:    true,
					},
					"last_set": schema.Int64Attribute{
						Description: "A Unix timestamp indicating when the topic was last set.",
						Computed:    true,
					},
					"value": schema.StringAttribute{
						Description: "The conversation's topic.",
						Computed:    true,
					},
				},
			},
			"latest": schema.SingleNestedAttribute{
				Description: "The latest post in the conversation.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: "The type of post.",
						Computed:    true,
					},
					"user": schema.StringAttribute{
						Description: "The ID of the user that made the post.",
						Computed:    true,
					},
					"text": schema.StringAttribute{
						Description: "The text of the post.",
						Computed:    true,
					},
					"ts": schema.StringAttribute{
						Description: "A Unix timestamp for the post.",
						Computed:    true,
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *conversationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read conversation data source")
	var state conversationModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	conversation := slack.GetConversationInfoInput{
		ChannelID:         state.ID.ValueString(),
		IncludeLocale:     true,
		IncludeNumMembers: true,
	}

	conversationResponse, err := d.client.GetConversationInfo(&conversation)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Conversation",
			err.Error(),
		)
		return
	}

	latestData := latestModel{
		Type: types.StringValue(""),
		User: types.StringValue(""),
		Text: types.StringValue(""),
		TS:   types.StringValue(""),
	}

	if conversationResponse.Latest != nil {
		latestData = latestModel{
			Type: types.StringValue(conversationResponse.Latest.Type),
			User: types.StringValue(conversationResponse.Latest.User),
			Text: types.StringValue(conversationResponse.Latest.Text),
			TS:   types.StringValue(conversationResponse.Topic.LastSet.Time().Format("Mon Jan 2 15:04:05 MST 2006")),
		}
	}

	topicData := topicModel{
		Creator: types.StringValue(conversationResponse.Topic.Creator),
		LastSet: types.Int64Value(conversationResponse.Topic.LastSet.Time().Unix()),
		Value:   types.StringValue(conversationResponse.Topic.Value),
	}

	purposeData := purposeModel{
		Creator: types.StringValue(conversationResponse.Purpose.Creator),
		LastSet: types.Int64Value(conversationResponse.Purpose.LastSet.Time().Unix()),
		Value:   types.StringValue(conversationResponse.Purpose.Value),
	}

	// Map response body to model
	state = conversationModel{
		Created:            types.StringValue(conversationResponse.Created.Time().Format("Mon Jan 2 15:04:05 MST 2006")),
		Creator:            types.StringValue(conversationResponse.Creator),
		ID:                 types.StringValue(conversationResponse.ID),
		IsArchived:         types.BoolValue(conversationResponse.IsArchived),
		IsChannel:          types.BoolValue(conversationResponse.IsChannel),
		IsExtShared:        types.BoolValue(conversationResponse.IsExtShared),
		IsGeneral:          types.BoolValue(conversationResponse.IsGeneral),
		IsGroup:            types.BoolValue(conversationResponse.IsGroup),
		IsIM:               types.BoolValue(conversationResponse.IsIM),
		IsMember:           types.BoolValue(conversationResponse.IsMember),
		IsOpen:             types.BoolValue(conversationResponse.IsOpen),
		IsOrgShared:        types.BoolValue(conversationResponse.IsOrgShared),
		IsPendingExtShared: types.BoolValue(conversationResponse.IsPendingExtShared),
		IsPrivate:          types.BoolValue(conversationResponse.IsPrivate),
		IsShared:           types.BoolValue(conversationResponse.IsShared),
		LastRead:           types.StringValue(conversationResponse.LastRead),
		Latest:             &latestData,
		Locale:             types.StringValue(conversationResponse.Locale),
		Name:               types.StringValue(conversationResponse.Name),
		NameNormalized:     types.StringValue(conversationResponse.NameNormalized),
		NumMembers:         types.Int64Value(int64(conversationResponse.NumMembers)),
		Priority:           types.Int64Value(int64(conversationResponse.Priority)),
		Purpose:            &purposeData,
		Topic:              &topicData,
		Unlinked:           types.Int64Value(int64(conversationResponse.Unlinked)),
		UnreadCount:        types.Int64Value(int64(conversationResponse.UnreadCount)),
		UnreadCountDisplay: types.Int64Value(int64(conversationResponse.UnreadCountDisplay)),
		User:               types.StringValue(conversationResponse.User),
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Debug(ctx, "Read conversation data source", map[string]any{"success": true})
}
