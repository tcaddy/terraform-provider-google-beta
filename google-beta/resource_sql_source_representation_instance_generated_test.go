// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckSQLSourceRepresentationInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstanceBasicExample(context),
			},
			{
				ResourceName:      "google_sql_source_representation_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSQLSourceRepresentationInstance_sqlSourceRepresentationInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_source_representation_instance" "instance" {
  name             = "tf-test-my-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "MYSQL_5_7"
  host             = "10.20.30.40"
  port             = 3306
}
`, context)
}

func testAccCheckSQLSourceRepresentationInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_sql_source_representation_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{SQLBasePath}}projects/{{project}}/instances/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("SQLSourceRepresentationInstance still exists at %s", url)
			}
		}

		return nil
	}
}
