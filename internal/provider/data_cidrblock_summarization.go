package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/PatrickCronin/routesum/pkg/routesum"
)

func dataSourceCidrBlockSummarization() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get the list of summarized CIDRs.",

		ReadContext: dataSourceCidrBlockSummarizationRead,

		Schema: map[string]*schema.Schema{
			"cidr_blocks": {
				Description: "The list of the IPv4 address and network mask in CIDR notation.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsCIDR,
				},
			},

			"summarized_cidr_blocks": {
				Description: "The list of the summarized IPv4 address and network mask in CIDR notation.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsCIDR,
				},
			},
		},
	}
}

func dataSourceCidrBlockSummarizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rtSum := routesum.NewRouteSum()
	if v, ok := d.GetOk("cidr_blocks"); ok && v.(*schema.Set).Len() > 0 {
		for _, vv := range v.(*schema.Set).List() {
			if err := rtSum.InsertFromString(vv.(string)); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	summarized := rtSum.SummaryStrings()

	d.SetId(strings.Join(summarized, ","))
	d.Set("summarized_cidr_blocks", summarized)

	return nil
}
