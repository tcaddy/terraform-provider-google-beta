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

package iap

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceIapClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceIapClientCreate,
		Read:   resourceIapClientRead,
		Delete: resourceIapClientDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIapClientImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"brand": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Identifier of the brand to which this client
is attached to. The format is
'projects/{project_number}/brands/{brand_id}'.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Human-friendly name given to the OAuth client.`,
			},
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Unique identifier of the OAuth client.`,
			},
			"secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Client secret of the OAuth client.`,
				Sensitive:   true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIapClientCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandIapClientDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}{{brand}}/identityAwareProxyClients")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Client: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IapClient409Operation},
	})
	if err != nil {
		return fmt.Errorf("Error creating Client: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{brand}}/identityAwareProxyClients/{{client_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	brand := d.Get("brand")
	clientId := flattenIapClientClientId(res["name"], d, config)

	if err := d.Set("client_id", clientId); err != nil {
		return fmt.Errorf("Error setting client_id: %s", err)
	}
	d.SetId(fmt.Sprintf("%s/identityAwareProxyClients/%s", brand, clientId))

	log.Printf("[DEBUG] Finished creating Client %q: %#v", d.Id(), res)

	return resourceIapClientRead(d, meta)
}

func resourceIapClientRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}{{brand}}/identityAwareProxyClients/{{client_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IapClient409Operation},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IapClient %q", d.Id()))
	}

	if err := d.Set("secret", flattenIapClientSecret(res["secret"], d, config)); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}
	if err := d.Set("display_name", flattenIapClientDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}
	if err := d.Set("client_id", flattenIapClientClientId(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}

	return nil
}

func resourceIapClientDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{IapBasePath}}{{brand}}/identityAwareProxyClients/{{client_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Client %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IapClient409Operation},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Client")
	}

	log.Printf("[DEBUG] Finished deleting Client %q: %#v", d.Id(), res)
	return nil
}

func resourceIapClientImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<brand>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("brand").(string), "/")
	if len(nameParts) != 6 {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s",
			d.Get("brand").(string),
			"projects/{{project_number}}/brands/{{brand_id}}/identityAwareProxyClients/{{client_id}}",
		)
	}

	if err := d.Set("brand", fmt.Sprintf("projects/%s/brands/%s", nameParts[1], nameParts[3])); err != nil {
		return nil, fmt.Errorf("Error setting brand: %s", err)
	}
	if err := d.Set("client_id", nameParts[5]); err != nil {
		return nil, fmt.Errorf("Error setting client_id: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenIapClientSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIapClientDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIapClientClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func expandIapClientDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
