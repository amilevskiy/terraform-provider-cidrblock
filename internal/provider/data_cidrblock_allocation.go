package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/amilevskiy/cidrblock"
)

func dataSourceCidrBlockAllocation() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get the list of CIDRs to create subnets.",

		ReadContext: dataSourceCidrBlockAllocationRead,

		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Description:  "The IPv4 address and network mask in CIDR notation.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},

			"exclude_cidr_blocks": {
				Description: "The list of the IPv4 address and network mask in CIDR notation to exclude from allocation.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsCIDR,
				},
			},

			"prefix_lengths": {
				Description: "The list of allocating CIDR prefix lengths.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"cidr_blocks": {
				Description: "The list of the allocated IPv4 CIDRs.",
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

func dataSourceCidrBlockAllocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error

	cb := cidrblock.New()

	if v, ok := d.GetOk("cidr_block"); ok {
		if err = cb.SetCidrBlock(v.(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	if v, ok := d.GetOk("exclude_cidr_blocks"); ok && v.(*schema.Set).Len() > 0 {
		s := make([]string, 0)
		for _, vv := range v.(*schema.Set).List() {
			s = append(s, vv.(string))
		}
		cb.SetExcludeCidrBlocks(s)
	}

	prefixes := make([]int, 0)
	if v, ok := d.GetOk("prefix_lengths"); ok {
		for _, vv := range v.([]interface{}) {
			if 0 <= vv.(int) && vv.(int) <= 32 {
				prefixes = append(prefixes, vv.(int))
			}
		}
	}

	cidrBlocks, err := cb.AllocCidrBlocks(prefixes...)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cb.CidrBlock())
	d.Set("cidr_blocks", cidrBlocks)

	return nil
}
