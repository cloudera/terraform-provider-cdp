package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGroupRead,

		Schema: map[string]*schema.Schema{
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sync_membership_on_user_login": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGroupRead(d *schema.ResourceData, m interface{}) error {
	return sharedGroupRead(d, m, d.Get("group_name").(string))
}
