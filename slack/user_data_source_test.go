package slack

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing - FIXME: remove hardcoded user id (maybe, create, read, and then delete?)
			{
				Config: providerConfig + fmt.Sprintf(`
data "slack_user" "test" {
	id = "%s"
}
`, slackTestUserID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify placeholder id attribute
					resource.TestCheckResourceAttrSet("data.slack_user.test", "id"),
				),
			},
		},
	})
}
