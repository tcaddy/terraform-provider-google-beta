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

func TestAccOSConfigPatchDeployment_osConfigPatchDeploymentBasicExample(t *testing.T) {
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
		CheckDestroy: testAccCheckOSConfigPatchDeploymentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigPatchDeployment_osConfigPatchDeploymentBasicExample(context),
			},
			{
				ResourceName:            "google_os_config_patch_deployment.patch",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"patch_deployment_id"},
			},
		},
	})
}

func testAccOSConfigPatchDeployment_osConfigPatchDeploymentBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_os_config_patch_deployment" "patch" {
  patch_deployment_id = "tf-test-patch-deploy-inst%{random_suffix}"

  instance_filter {
    all = true
  }

  one_time_schedule {
    execute_time = "2020-10-10T10:10:10.045123456Z"
  }
}
`, context)
}

func TestAccOSConfigPatchDeployment_osConfigPatchDeploymentInstanceExample(t *testing.T) {
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
		CheckDestroy: testAccCheckOSConfigPatchDeploymentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigPatchDeployment_osConfigPatchDeploymentInstanceExample(context),
			},
			{
				ResourceName:            "google_os_config_patch_deployment.patch",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"patch_deployment_id"},
			},
		},
	})
}

func testAccOSConfigPatchDeployment_osConfigPatchDeploymentInstanceExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_instance" "foobar" {
  name           = "tf-test-patch-deploy-inst%{random_suffix}"
  machine_type   = "n1-standard-1"
  zone           = "us-central1-a"
  can_ip_forward = false
  tags           = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }
}

resource "google_os_config_patch_deployment" "patch" {
  patch_deployment_id = "tf-test-patch-deploy-inst%{random_suffix}"

  instance_filter {
    instances = [google_compute_instance.foobar.id]
  }

  patch_config {
    yum {
      security = true
      minimal = true
      excludes = ["bash"]
    }
  }

  recurring_schedule {
    time_zone {
      id = "America/New_York"
    }

    time_of_day {
      hours = 0
      minutes = 30
      seconds = 30
      nanos = 20
    }

    monthly {
      month_day = 1
    }
  }
}
`, context)
}

func TestAccOSConfigPatchDeployment_osConfigPatchDeploymentFullExample(t *testing.T) {
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
		CheckDestroy: testAccCheckOSConfigPatchDeploymentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigPatchDeployment_osConfigPatchDeploymentFullExample(context),
			},
			{
				ResourceName:            "google_os_config_patch_deployment.patch",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"patch_deployment_id"},
			},
		},
	})
}

func testAccOSConfigPatchDeployment_osConfigPatchDeploymentFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_os_config_patch_deployment" "patch" {
  patch_deployment_id = "tf-test-patch-deploy-inst%{random_suffix}"

  instance_filter {
    group_labels {
      labels = {
        env = "dev",
        app = "web"
      }
    }

    instance_name_prefixes = ["test-"]

    zones = ["us-central1-a", "us-central-1c"]
  }

  patch_config {
    reboot_config = "ALWAYS"

    apt {
      type = "DIST"
      excludes = ["python"]
    }

    yum {
      security = true
      minimal = true
      excludes = ["bash"]
    }

    goo {
      enabled = true
    }

    zypper {
      categories = ["security"]
    }

    windows_update {
      classifications = ["CRITICAL", "SECURITY", "UPDATE"]
    }

    pre_step {
      linux_exec_step_config {
        allowed_success_codes = [0,3]
        local_path = "/tmp/pre_patch_script.sh"
      }

      windows_exec_step_config {
        interpreter = "SHELL"
        allowed_success_codes = [0,2]
        local_path  = "C:\\Users\\user\\pre-patch-script.cmd"
      }
    }

    post_step {
      linux_exec_step_config {
        gcs_object {
          bucket = "my-patch-scripts"
          generation_number = "1523477886880" 
          object = "linux/post_patch_script"
        }
      }

      windows_exec_step_config {
        interpreter = "POWERSHELL"
        gcs_object {
          bucket = "my-patch-scripts"
          generation_number = "135920493447"
          object = "windows/post_patch_script.ps1"
        }
      }
    }
  }

  duration = "10s"

  recurring_schedule {
    time_zone {
      id = "America/New_York"
    }

    time_of_day {
      hours = 0
      minutes = 30
      seconds = 30
      nanos = 20
    }

    monthly {
      week_day_of_month {
        week_ordinal = -1
        day_of_week  = "TUESDAY"
      }
    }
  }

  rollout {
    mode = "ZONE_BY_ZONE"
    disruption_budget {
      fixed = 1
    }
  }
}
`, context)
}

func testAccCheckOSConfigPatchDeploymentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_os_config_patch_deployment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{OSConfigBasePath}}{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("OSConfigPatchDeployment still exists at %s", url)
			}
		}

		return nil
	}
}
