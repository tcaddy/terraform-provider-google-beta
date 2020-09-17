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
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDialogflowEntityType() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowEntityTypeCreate,
		Read:   resourceDialogflowEntityTypeRead,
		Update: resourceDialogflowEntityTypeUpdate,
		Delete: resourceDialogflowEntityTypeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowEntityTypeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of this entity type to be displayed on the console.`,
			},
			"kind": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"KIND_MAP", "KIND_LIST", "KIND_REGEXP"}, false),
				Description: `Indicates the kind of entity type.
* KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
* KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
types can contain references to other entity types (with or without aliases).
* KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"]`,
			},
			"enable_fuzzy_extraction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Enables fuzzy entity extraction during classification.`,
			},
			"entities": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The collection of entity entries associated with the entity type.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"synonyms": {
							Type:     schema.TypeList,
							Required: true,
							Description: `A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym
could be green onions.
For KIND_LIST entity types:
* This collection must contain exactly one synonym equal to value.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The primary value associated with this entity entry. For example, if the entity type is vegetable, the value
could be scallions.
For KIND_MAP entity types:
* A reference value to be used in place of synonyms.
For KIND_LIST entity types:
* A string that can contain references to other entity types (with or without aliases).`,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique identifier of the entity type. 
Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDialogflowEntityTypeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowEntityTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	kindProp, err := expandDialogflowEntityTypeKind(d.Get("kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kind"); !isEmptyValue(reflect.ValueOf(kindProp)) && (ok || !reflect.DeepEqual(v, kindProp)) {
		obj["kind"] = kindProp
	}
	enableFuzzyExtractionProp, err := expandDialogflowEntityTypeEnableFuzzyExtraction(d.Get("enable_fuzzy_extraction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_fuzzy_extraction"); !isEmptyValue(reflect.ValueOf(enableFuzzyExtractionProp)) && (ok || !reflect.DeepEqual(v, enableFuzzyExtractionProp)) {
		obj["enableFuzzyExtraction"] = enableFuzzyExtractionProp
	}
	entitiesProp, err := expandDialogflowEntityTypeEntities(d.Get("entities"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entities"); !isEmptyValue(reflect.ValueOf(entitiesProp)) && (ok || !reflect.DeepEqual(v, entitiesProp)) {
		obj["entities"] = entitiesProp
	}

	url, err := replaceVars(d, config, "{{DialogflowBasePath}}projects/{{project}}/agent/entityTypes/")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EntityType: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EntityType: %s", err)
	}
	if err := d.Set("name", flattenDialogflowEntityTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EntityType %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceDialogflowEntityTypeRead(d, meta)
}

func resourceDialogflowEntityTypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{DialogflowBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DialogflowEntityType %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}

	if err := d.Set("name", flattenDialogflowEntityTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowEntityTypeDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("kind", flattenDialogflowEntityTypeKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("enable_fuzzy_extraction", flattenDialogflowEntityTypeEnableFuzzyExtraction(res["enableFuzzyExtraction"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("entities", flattenDialogflowEntityTypeEntities(res["entities"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}

	return nil
}

func resourceDialogflowEntityTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowEntityTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	kindProp, err := expandDialogflowEntityTypeKind(d.Get("kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kind"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, kindProp)) {
		obj["kind"] = kindProp
	}
	enableFuzzyExtractionProp, err := expandDialogflowEntityTypeEnableFuzzyExtraction(d.Get("enable_fuzzy_extraction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_fuzzy_extraction"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableFuzzyExtractionProp)) {
		obj["enableFuzzyExtraction"] = enableFuzzyExtractionProp
	}
	entitiesProp, err := expandDialogflowEntityTypeEntities(d.Get("entities"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entities"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entitiesProp)) {
		obj["entities"] = entitiesProp
	}

	url, err := replaceVars(d, config, "{{DialogflowBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EntityType %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating EntityType %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating EntityType %q: %#v", d.Id(), res)
	}

	return resourceDialogflowEntityTypeRead(d, meta)
}

func resourceDialogflowEntityTypeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{DialogflowBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EntityType %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EntityType")
	}

	log.Printf("[DEBUG] Finished deleting EntityType %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowEntityTypeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) < 2 {
		return nil, fmt.Errorf(
			"Could not split project from name: %s",
			d.Get("name"),
		)
	}

	d.Set("project", stringParts[1])
	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowEntityTypeName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeEnableFuzzyExtraction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeEntities(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"value":    flattenDialogflowEntityTypeEntitiesValue(original["value"], d, config),
			"synonyms": flattenDialogflowEntityTypeEntitiesSynonyms(original["synonyms"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowEntityTypeEntitiesValue(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowEntityTypeEntitiesSynonyms(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandDialogflowEntityTypeDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeEnableFuzzyExtraction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeEntities(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedValue, err := expandDialogflowEntityTypeEntitiesValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedSynonyms, err := expandDialogflowEntityTypeEntitiesSynonyms(original["synonyms"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSynonyms); val.IsValid() && !isEmptyValue(val) {
			transformed["synonyms"] = transformedSynonyms
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowEntityTypeEntitiesValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeEntitiesSynonyms(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
