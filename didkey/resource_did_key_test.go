package didkey

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccHashicupsOrderBasic(t *testing.T) {
	keeperKey := "1234"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHashicupsOrderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckHashicupsOrderConfigBasic(keeperKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHashicupsOrderExists("didkey.new"),
				),
			},
		},
	})
}

func testAccCheckHashicupsOrderDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "didkey" {
			continue
		}

		primaryID := rs.Primary.ID
		if primaryID == "" {
			return fmt.Errorf("primaryID is an empty string")
		}
	}

	return nil
}

func testAccCheckHashicupsOrderConfigBasic(keeperKey string) string {
	return fmt.Sprintf(`
	resource "didkey" "new" {
		keepers {
			"key" = %s
		}
	}
	`, keeperKey)
}

func testAccCheckHashicupsOrderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No primaryID set")
		}

		return nil
	}
}
