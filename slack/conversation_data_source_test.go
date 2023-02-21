package slack

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccConversationDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(`
data "slack_conversation" "test" {
	id = "%s"
}
`, slackTestConversationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify placeholder id attribute
					resource.TestCheckResourceAttrSet("data.slack_conversation.test", "id"),
				),
			},
		},
	})
}
