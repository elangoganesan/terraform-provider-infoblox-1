package infoblox

import (
	"fmt"
	"github.com/CARFAX/skyinfoblox/api/common/v261/model"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strconv"
	"testing"
)

func TestResourceMXRecord(t *testing.T) {
	mailExchanger := fmt.Sprintf("mx-%s.bskyb.com", strconv.Itoa(acctest.RandIntRange(0, 100)))
	resourceName := "infoblox_mx_record.mxtest1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.RecordMXObj, "mail_exchanger", mailExchanger)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMXRecordCreateTemplate(mailExchanger),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceMXRecordExists("mail_exchanger", mailExchanger),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mailExchanger),
					resource.TestCheckResourceAttr(resourceName, "name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "preference", "120"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			{
				Config: testAccResourceMXRecordUpdateTemplate(mailExchanger),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceMXRecordExists("mail_exchanger", mailExchanger),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mailExchanger),
					resource.TestCheckResourceAttr(resourceName, "name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment edited for a disabled zone"),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "4000"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
		},
	})

}

func testAccResourceMXRecordExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.RecordMXObj, key, value)
	}
}

func testAccResourceMXRecordCreateTemplate(mailExchanger string) string {
	return fmt.Sprintf(`
	resource "infoblox_mx_record" "mxtest1" {
        	name = "example.com"
        	comment = "this is a comment"
        	disable = false
        	mail_exchanger = "%s"
        	preference = 120
        	ttl = 3600
        	use_ttl = true
        	view = "default"
	}
	`, mailExchanger)

}

func testAccResourceMXRecordUpdateTemplate(mailExchanger string) string {
	return fmt.Sprintf(`
	resource "infoblox_mx_record" "mxtest1" {
        	name = "example.com"
        	comment = "this is a comment edited for a disabled zone"
        	disable = true
        	mail_exchanger = "%s"
        	preference = 10
        	ttl = 4000
        	use_ttl = true
        	view = "default"
	}
	`, mailExchanger)

}
