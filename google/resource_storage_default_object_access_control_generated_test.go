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
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccStorageDefaultObjectAccessControl_storageDefaultObjectAccessControlPublicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageDefaultObjectAccessControlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageDefaultObjectAccessControl_storageDefaultObjectAccessControlPublicExample(acctest.RandString(10)),
			},
			{
				ResourceName:            "google_storage_default_object_access_control.public_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bucket"},
			},
		},
	})
}

func testAccStorageDefaultObjectAccessControl_storageDefaultObjectAccessControlPublicExample(val string) string {
	return fmt.Sprintf(`
resource "google_storage_default_object_access_control" "public_rule" {
  bucket = "${google_storage_bucket.bucket.name}"
  role   = "READER"
  entity = "allUsers"
}

resource "google_storage_bucket" "bucket" {
	name = "static-content-bucket-%s"
}
`, val,
	)
}

func testAccCheckStorageDefaultObjectAccessControlDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_storage_default_object_access_control" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/storage/v1/b/{{bucket}}/defaultObjectAcl/{{entity}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("StorageDefaultObjectAccessControl still exists at %s", url)
		}
	}

	return nil
}
