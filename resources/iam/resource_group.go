package iam

import (
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	iammodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sync_membership_on_user_login": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Iam

	groupName := d.Get("group_name").(string)
	syncMembershipOnUserLogin := d.Get("sync_membership_on_user_login").(bool)

	params := operations.NewCreateGroupParams()
	params.WithInput(&iammodels.CreateGroupRequest{
		GroupName:                 &groupName,
		SyncMembershipOnUserLogin: &syncMembershipOnUserLogin,
	})
	_, err := client.Operations.CreateGroup(params)
	if err != nil {
		return err
	}

	d.SetId(groupName)

	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	return sharedGroupRead(d, m, d.Id())
}

func sharedGroupRead(d *schema.ResourceData, m interface{}, groupName string) error {
	client := m.(*cdp.Client).Iam

	params := operations.NewListGroupsParams()
	params.WithInput(&iammodels.ListGroupsRequest{GroupNames: []string{groupName}})
	resp, err := client.Operations.ListGroups(params)
	if err != nil {
		return err
	}
	groups := resp.GetPayload().Groups
	if len(groups) == 0 || *groups[0].GroupName != groupName {
		d.SetId("") // deleted
		return nil
	}
	g := groups[0]

	d.SetId(*g.GroupName)
	d.Set("group_name", *g.GroupName)
	d.Set("sync_membership_on_user_login", g.SyncMembershipOnUserLogin)
	d.Set("crn", g.Crn)

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Iam

	groupName := d.Id()

	if d.HasChange("sync_membership_on_user_login") {
		syncMembershipOnUserLogin := d.Get("sync_membership_on_user_login").(bool)

		params := operations.NewUpdateGroupParams()
		params.WithInput(&iammodels.UpdateGroupRequest{
			GroupName:                 &groupName,
			SyncMembershipOnUserLogin: syncMembershipOnUserLogin,
		})
		_, err := client.Operations.UpdateGroup(params)
		if err != nil {
			return err
		}
	}

	return resourceGroupRead(d, m)
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Iam

	groupName := d.Id()
	params := operations.NewDeleteGroupParams()
	params.WithInput(&iammodels.DeleteGroupRequest{GroupName: &groupName})
	_, err := client.Operations.DeleteGroup(params)
	if err != nil {
		return err
	}

	return nil
}
