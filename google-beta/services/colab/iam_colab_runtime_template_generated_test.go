// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package colab_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccColabRuntimeTemplateIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccColabRuntimeTemplateIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_colab_runtime_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-colab-runtime-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccColabRuntimeTemplateIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_colab_runtime_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-colab-runtime-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccColabRuntimeTemplateIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccColabRuntimeTemplateIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_colab_runtime_template_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-colab-runtime-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccColabRuntimeTemplateIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccColabRuntimeTemplateIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_colab_runtime_template_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_colab_runtime_template_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-colab-runtime-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccColabRuntimeTemplateIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_colab_runtime_template_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-colab-runtime-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccColabRuntimeTemplateIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "runtime-template" {
  name = "tf-test-colab-runtime-template%{random_suffix}"
  display_name = "Runtime template basic"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

resource "google_colab_runtime_template_iam_member" "foo" {
  project = google_colab_runtime_template.runtime-template.project
  location = google_colab_runtime_template.runtime-template.location
  runtime_template = google_colab_runtime_template.runtime-template.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccColabRuntimeTemplateIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "runtime-template" {
  name = "tf-test-colab-runtime-template%{random_suffix}"
  display_name = "Runtime template basic"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_colab_runtime_template_iam_policy" "foo" {
  project = google_colab_runtime_template.runtime-template.project
  location = google_colab_runtime_template.runtime-template.location
  runtime_template = google_colab_runtime_template.runtime-template.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_colab_runtime_template_iam_policy" "foo" {
  project = google_colab_runtime_template.runtime-template.project
  location = google_colab_runtime_template.runtime-template.location
  runtime_template = google_colab_runtime_template.runtime-template.name
  depends_on = [
    google_colab_runtime_template_iam_policy.foo
  ]
}
`, context)
}

func testAccColabRuntimeTemplateIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "runtime-template" {
  name = "tf-test-colab-runtime-template%{random_suffix}"
  display_name = "Runtime template basic"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

data "google_iam_policy" "foo" {
}

resource "google_colab_runtime_template_iam_policy" "foo" {
  project = google_colab_runtime_template.runtime-template.project
  location = google_colab_runtime_template.runtime-template.location
  runtime_template = google_colab_runtime_template.runtime-template.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccColabRuntimeTemplateIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "runtime-template" {
  name = "tf-test-colab-runtime-template%{random_suffix}"
  display_name = "Runtime template basic"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

resource "google_colab_runtime_template_iam_binding" "foo" {
  project = google_colab_runtime_template.runtime-template.project
  location = google_colab_runtime_template.runtime-template.location
  runtime_template = google_colab_runtime_template.runtime-template.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccColabRuntimeTemplateIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "runtime-template" {
  name = "tf-test-colab-runtime-template%{random_suffix}"
  display_name = "Runtime template basic"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

resource "google_colab_runtime_template_iam_binding" "foo" {
  project = google_colab_runtime_template.runtime-template.project
  location = google_colab_runtime_template.runtime-template.location
  runtime_template = google_colab_runtime_template.runtime-template.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
