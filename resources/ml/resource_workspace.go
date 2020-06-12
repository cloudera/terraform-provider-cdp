package ml

import (
	"fmt"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	mlclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/client/operations"
	mlmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
	"time"
)

func ResourceWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkspaceCreate,
		Read:   resourceWorkspaceRead,
		Delete: resourceWorkspaceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"workspace_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"use_public_load_balancer": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"disable_tls": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"kubernetes": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_group": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"autoscaling": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"min_instances": &schema.Schema{
													Type:     schema.TypeInt,
													Required: true,
													ForceNew: true,
												},
												"max_instances": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceWorkspaceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Ml

	workspaceName := d.Get("workspace_name").(string)
	environmentName := d.Get("environment_name").(string)
	usePublicLoadBalancer := d.Get("use_public_load_balancer").(bool)
	disableTls := d.Get("disable_tls").(bool)

	kubernetes, kubernetesErr := utils.GetMapFromSingleItemList(d, "kubernetes")
	if kubernetesErr != nil {
		return kubernetesErr
	}

	instanceGroups := []*mlmodels.InstanceGroup{}
	instanceGroupSet := kubernetes["instance_group"].(*schema.Set)
	for _, instanceGroupObject := range instanceGroupSet.List() {
		instanceGroup := instanceGroupObject.(map[string]interface{})

		instanceType := instanceGroup["instance_type"].(string)

		autoscaling, autoscalingErr := utils.GetMapFromSingleItemListInMap(instanceGroup, "autoscaling")
		if autoscalingErr != nil {
			return autoscalingErr
		}

		minInstances := int32(autoscaling["min_instances"].(int))
		maxInstances := int32(autoscaling["max_instances"].(int))

		instanceGroups = append(instanceGroups, &mlmodels.InstanceGroup{
			InstanceType: &instanceType,
			Autoscaling: &mlmodels.Autoscaling{
				MinInstances: &minInstances,
				MaxInstances: &maxInstances,
			},
			IngressRules: []string{},
		})
	}

	params := operations.NewCreateWorkspaceParams()
	params.WithInput(&mlmodels.CreateWorkspaceRequest{
		WorkspaceName:         &workspaceName,
		EnvironmentName:       &environmentName,
		UsePublicLoadBalancer: usePublicLoadBalancer,
		DisableTLS:            disableTls,
		ProvisionK8sRequest: &mlmodels.ProvisionK8sRequest{
			EnvironmentName: &environmentName,
			InstanceGroups:  instanceGroups,
			Tags:            []*mlmodels.ProvisionTag{},
		},
		LoadBalancerIPWhitelists: []string{},
	})
	_, err := client.Operations.CreateWorkspace(params)
	if err != nil {
		return err
	}

	d.SetId(makeId(environmentName, workspaceName))

	if err := waitForWorkspaceToBeAvailable(environmentName, workspaceName, d.Timeout(schema.TimeoutCreate), client); err != nil {
		return err
	}

	return resourceWorkspaceRead(d, m)
}

func makeId(environmentName string, workspaceName string) string {
	return fmt.Sprintf("%s::%s", environmentName, workspaceName)
}

func getId(d *schema.ResourceData) (string, string) {
	parts := strings.SplitN(d.Id(), "::", 2)
	return parts[0], parts[1]
}

func describeWorkspace(environmentName string, workspaceName string, client *mlclient.Ml) (*mlmodels.WorkspaceSummary, error) {
	params := operations.NewListWorkspacesParams()
	params.WithInput(struct{}{})
	resp, err := client.Operations.ListWorkspaces(params)
	if err != nil {
		return nil, err
	}
	for _, workspace := range resp.GetPayload().Workspaces {
		if *workspace.EnvironmentName == environmentName && *workspace.InstanceName == workspaceName {
			return workspace, nil
		}
	}
	return nil, nil
}

func waitForWorkspaceToBeAvailable(environmentName string, workspaceName string, timeout time.Duration, client *mlclient.Ml) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"provision:started"},
		Target:       []string{"installation:finished"},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh:      makePollWorkspaceStatus(environmentName, workspaceName, client),
	}
	_, err := stateConf.WaitForState()

	return err
}

func resourceWorkspaceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Ml

	environmentName, workspaceName := getId(d)

	workspace, err := describeWorkspace(environmentName, workspaceName, client)
	if err != nil {
		return err
	}
	if workspace == nil {
		d.SetId("") // deleted
		return nil
	}

	d.Set("crn", workspace.Crn)

	return nil
}

func resourceWorkspaceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*cdp.Client).Ml

	environmentName, workspaceName := getId(d)
	removeStorage := true
	force := false

	params := operations.NewDeleteWorkspaceParams()
	params.WithInput(&mlmodels.DeleteWorkspaceRequest{
		EnvironmentName: environmentName,
		WorkspaceName:   workspaceName,
		RemoveStorage:   &removeStorage,
		Force:           &force,
	})
	_, err := client.Operations.DeleteWorkspace(params)
	if err != nil {
		return err
	}

	if err := waitForWorkspaceToBeDeleted(environmentName, workspaceName, d.Timeout(schema.TimeoutDelete), client); err != nil {
		return err
	}

	return nil
}

func waitForWorkspaceToBeDeleted(environmentName string, workspaceName string, timeout time.Duration, client *mlclient.Ml) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"deprovision:started"},
		Target:       []string{},
		Delay:        5 * time.Second,
		Timeout:      timeout,
		PollInterval: 10 * time.Second,
		Refresh:      makePollWorkspaceStatus(environmentName, workspaceName, client),
	}
	_, err := stateConf.WaitForState()

	return err
}

func makePollWorkspaceStatus(environmentName string, workspaceName string, client *mlclient.Ml) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("About to describe workspace")
		workspace, err := describeWorkspace(environmentName, workspaceName, client)
		if err != nil {
			log.Printf("Error describing workspace: %s", err)
			return nil, "", err
		}
		if workspace == nil {
			log.Printf("Workspace not found")
			return nil, "", nil
		}
		log.Printf("Described workspace: %s", *workspace.InstanceStatus)
		return workspace, *workspace.InstanceStatus, nil
	}
}
