package okta

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-okta/sdk"
)

func resourcePageError() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePageErrorSet,
		ReadContext:   resourcePageErrorRead,
		UpdateContext: resourcePageErrorSet,
		DeleteContext: resourcePageErrorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"brand_id": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: stringAtLeast(20),
			},
			"page_content": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "HTML template of the error page",
			},
		},
	}
}

func resourcePageErrorSet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	brandID := d.Get("brand_id").(string)
	obj := sdk.ErrorPageCustomization{}
	if val, ok := d.GetOk("page_content"); ok {
		obj.PageContent = val.(string)
	}
	_, err := getSupplementFromMetadata(m).SetPageError(ctx, brandID, &obj)
	if err != nil {
		return diag.Errorf("failed to set %s page template: %v", "error", err)
	}
	d.SetId(brandID)
	return nil
}

func resourcePageErrorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()
	resp, _, err := getSupplementFromMetadata(m).GetPageError(ctx, id)
	if err != nil {
		return diag.Errorf("failed to get %s page template: %v", "error", err)
	}
	_ = d.Set("page_content", resp.PageContent)
	if err != nil {
		panic(err)
	}
	d.SetId(id)
	return nil
}

func resourcePageErrorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()
	_, err := getSupplementFromMetadata(m).DeletePageError(ctx, id)
	if err != nil {
		return diag.Errorf("failed to get %s page template: %v", "error", err)
	}
	return nil
}
