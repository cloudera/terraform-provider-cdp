package datalake

import "github.com/hashicorp/terraform-plugin-framework/types"

type awsDatalakeResourceModel struct {
	ID types.String `tfsdk:"id"`

	InstanceProfile types.String `tfsdk:"instance_profile"`

	StorageBucketLocation types.String `tfsdk:"storage_bucket_location"`

	CertificateExpirationState types.String `tfsdk:"certificate_expiration_state"`

	CloudStorageBaseLocation types.String `tfsdk:"cloud_storage_base_location"`

	CloudbreakVersion types.String `tfsdk:"cloudbreak_version"`

	ClouderaManager types.Object `tfsdk:"cloudera_manager"`

	CreationDate types.String `tfsdk:"creation_date"`

	CredentialCrn types.String `tfsdk:"credential_crn"`

	Crn types.String `tfsdk:"crn"`

	CustomInstanceGroups []*awsDatalakeInstanceGroup `tfsdk:"custom_instance_groups"`

	DatabaseAvailabilityType types.String `tfsdk:"database_availability_type"`

	DatabaseEngineVersion types.String `tfsdk:"database_engine_version"`

	DatalakeName types.String `tfsdk:"datalake_name"`

	DatalakeTemplate types.String `tfsdk:"datalake_template"`

	EnableRangerRaz types.Bool `tfsdk:"enable_ranger_raz"`

	Endpoints types.List `tfsdk:"endpoints"`

	EnvironmentCrn types.String `tfsdk:"environment_crn"`

	EnvironmentName types.String `tfsdk:"environment_name"`

	Image *awsDatalakeImage `tfsdk:"image"`

	InstanceGroups types.List `tfsdk:"instance_groups"`

	LoadBalancers types.List `tfsdk:"load_balancers"`

	ProductVersions types.List `tfsdk:"product_versions"`

	Region types.String `tfsdk:"region"`

	JavaVersion types.Int64 `tfsdk:"java_version"`

	MultiAz types.Bool `tfsdk:"multi_az"`

	Recipes []*instanceGroupRecipe `tfsdk:"recipes"`

	Runtime types.String `tfsdk:"runtime"`

	Scale types.String `tfsdk:"scale"`

	SpotMaxPrice types.Float64 `tfsdk:"spot_max_price"`

	SpotPercentage types.Int64 `tfsdk:"spot_percentage"`

	Status types.String `tfsdk:"status"`

	StatusReason types.String `tfsdk:"status_reason"`

	Tags types.Map `tfsdk:"tags"`
}

type awsDatalakeInstanceGroup struct {
	InstanceType types.String `tfsdk:"instance_type"`

	Name types.String `tfsdk:"name"`
}
type awsDatalakeImage struct {
	CatalogName types.String `tfsdk:"catalog_name"`

	ID types.String `tfsdk:"id"`
}

type instanceGroupRecipe struct {
	InstanceGroupName types.String `tfsdk:"instance_group_name"`

	RecipeNames types.List `tfsdk:"recipe_names"`
}

type instanceGroup struct {
	Instances types.List `tfsdk:"instances"`

	Name types.String `tfsdk:"name"`
}

type instance struct {
	AmbariServer types.Bool `tfsdk:"ambari_server"`

	DiscoveryFQDN types.String `tfsdk:"discovery_fqdn"`

	ID types.String `tfsdk:"id"`

	InstanceGroup types.String `tfsdk:"instance_group"`

	InstanceStatus types.String `tfsdk:"instance_status"`

	InstanceTypeVal types.String `tfsdk:"instance_type_val"`

	LifeCycle types.String `tfsdk:"life_cycle"`

	MountedVolumes types.List `tfsdk:"mounted_volumes"`

	PrivateIP types.String `tfsdk:"private_ip"`

	PublicIP types.String `tfsdk:"public_ip"`

	SSHPort types.Int64 `tfsdk:"ssh_port"`

	State types.String `tfsdk:"state"`

	StatusReason types.String `tfsdk:"status_reason"`
}

type mountedVolume struct {
	Device types.String `tfsdk:"device"`

	VolumeID types.String `tfsdk:"volume_id"`

	VolumeSize types.String `tfsdk:"volume_size"`

	VolumeType types.String `tfsdk:"volume_type"`
}

type clouderaManagerDetails struct {
	ClouderaManagerRepositoryURL types.String `tfsdk:"cloudera_manager_repository_url"`

	ClouderaManagerServerURL types.String `tfsdk:"cloudera_manager_server_url"`

	Version types.String `tfsdk:"version"`
}

type endpoint struct {
	DisplayName types.String `tfsdk:"display_name"`

	KnoxService types.String `tfsdk:"knox_service"`

	Mode types.String `tfsdk:"mode"`

	Open types.Bool `tfsdk:"open"`

	ServiceName types.String `tfsdk:"service_name"`

	ServiceURL types.String `tfsdk:"service_url"`
}

type loadBalancer struct {
	ResourceID types.String `tfsdk:"resource_id"`

	CloudDNS types.String `tfsdk:"cloud_dns"`

	Fqdn types.String `tfsdk:"fqdn"`

	IP types.String `tfsdk:"ip"`

	LoadBalancerType types.String `tfsdk:"load_balancer_type"`

	Targets types.List `tfsdk:"targets"`
}

type targetGroup struct {
	ListenerID types.String `tfsdk:"listener_id"`

	TargetGroupID types.String `tfsdk:"target_group_id"`

	Port types.Int64 `tfsdk:"port"`

	TargetInstances types.List `tfsdk:"target_instances"`
}

type productVersion struct {
	Name types.String `tfsdk:"name"`

	Version types.String `tfsdk:"version"`
}
