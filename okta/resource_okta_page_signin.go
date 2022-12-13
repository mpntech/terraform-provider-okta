package okta

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-okta/sdk"
	"unicode"
)

func resourcePageSignIn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePageSignInSet,
		ReadContext:   resourcePageSignInRead,
		UpdateContext: resourcePageSignInSet,
		DeleteContext: resourcePageSignInDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"brand_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"page_content": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTML template of the sign-in page",
			},
			"widget_customizations": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signin_label":                                  {Type: schema.TypeString, Optional: true},
						"username_label":                                {Type: schema.TypeString, Optional: true},
						"username_info_tip":                             {Type: schema.TypeString, Optional: true},
						"password_label":                                {Type: schema.TypeString, Optional: true},
						"password_info_tip":                             {Type: schema.TypeString, Optional: true},
						"forgot_password_label":                         {Type: schema.TypeString, Optional: true},
						"forgot_password_url":                           {Type: schema.TypeString, Optional: true},
						"unlock_account_label":                          {Type: schema.TypeString, Optional: true},
						"unlock_account_url":                            {Type: schema.TypeString, Optional: true},
						"help_label":                                    {Type: schema.TypeString, Optional: true},
						"help_url":                                      {Type: schema.TypeString, Optional: true},
						"custom_link1_label":                            {Type: schema.TypeString, Optional: true},
						"custom_link1_url":                              {Type: schema.TypeString, Optional: true},
						"custom_link2_label":                            {Type: schema.TypeString, Optional: true},
						"custom_link2_url":                              {Type: schema.TypeString, Optional: true},
						"authenticator_page_custom_link_label":          {Type: schema.TypeString, Optional: true},
						"authenticator_page_custom_link_url":            {Type: schema.TypeString, Optional: true},
						"classic_recovery_flow_email_or_username_label": {Type: schema.TypeString, Optional: true},
						"show_password_visibility_toggle":               {Type: schema.TypeBool, Optional: true, Default: true},
						"show_user_identifier":                          {Type: schema.TypeBool, Optional: true, Default: true},
					},
				},
			},
			"widget_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Version of sign-in widget",
			},
		},
	}
}

func snakeToCamelMap(m map[string]interface{}) map[string]interface{} {
	rekeyed := make(map[string]interface{})
	for k, v := range m {
		newRunes := []rune{}
		runes := []rune(k)
		match := false
		for _, r := range runes {
			if r == rune('_') {
				match = true
				continue
			}
			if match {
				r = unicode.ToUpper(r)
				match = false
			}
			newRunes = append(newRunes, r)
		}
		rekeyed[string(newRunes)] = v
	}
	return rekeyed
}

func camelToSnakeMap(m map[string]interface{}) map[string]interface{} {
	rekeyed := make(map[string]interface{})
	for k, v := range m {
		newRunes := []rune{}
		runes := []rune(k)
		for _, r := range runes {
			if unicode.IsUpper(r) {
				newRunes = append(newRunes, '_', unicode.ToLower(r))
				continue
			}
			newRunes = append(newRunes, r)
		}
		rekeyed[string(newRunes)] = v
	}
	return rekeyed
}

func resourcePageSignInSet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	brandID := d.Get("brand_id").(string)
	obj := sdk.SignInPageCustomization{}
	if val, ok := d.GetOk("page_content"); ok {
		obj.PageContent = val.(string)
	}
	if val, ok := d.GetOk("widget_version"); ok {
		obj.WidgetVersion = val.(string)
	}
	if val, ok := d.GetOk("widget_customizations"); ok {
		m := val.([]interface{})[0].(map[string]interface{})
		marsh, _ := json.Marshal(snakeToCamelMap(m))
		json.Unmarshal(marsh, &obj.WidgetCustomizations)
	}
	_, err := getSupplementFromMetadata(m).SetPageSignIn(ctx, brandID, &obj)
	if err != nil {
		return diag.Errorf("failed to set %s page template: %v", "sign-in", err)
	}
	d.SetId(brandID)
	return nil
}

func resourcePageSignInRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()
	resp, _, err := getSupplementFromMetadata(m).GetPageSignIn(ctx, id)
	if err != nil {
		return diag.Errorf("failed to get %s page template: %v", "sign-in", err)
	}
	_ = d.Set("page_content", resp.PageContent)
	_ = d.Set("widget_version", resp.WidgetVersion)
	err = d.Set("widget_customizations", []interface{}{flattenWidgetCustomizations(&resp.WidgetCustomizations)})
	if err != nil {
		panic(err)
	}
	d.SetId(id)
	return nil
}

func flattenWidgetCustomizations(wc *sdk.SignInPageWidgetCustomizations) interface{} {
	marsh, err := json.Marshal(wc)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(marsh, &m); err != nil {
		panic(err)
	}
	return camelToSnakeMap(m)
}

func resourcePageSignInDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()
	_, err := getSupplementFromMetadata(m).DeletePageSignIn(ctx, id)
	if err != nil {
		return diag.Errorf("failed to get %s page template: %v", "sign-in", err)
	}
	return nil
}
