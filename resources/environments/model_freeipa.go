// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import "github.com/hashicorp/terraform-plugin-framework/types"

type FreeIpaDetails struct {
	Catalog types.String `tfsdk:"catalog"`

	ImageID types.String `tfsdk:"image_id"`

	Os types.String `tfsdk:"os"`

	InstanceCountByGroup types.Int32 `tfsdk:"instance_count_by_group"`

	InstanceType types.String `tfsdk:"instance_type"`

	Instances types.Set `tfsdk:"instances"`

	MultiAz types.Bool `tfsdk:"multi_az"`

	Recipes types.Set `tfsdk:"recipes"`

	Architecture types.String `tfsdk:"architecture"`
}

type FreeIpaInstance struct {
	AvailabilityZone types.String `tfsdk:"availability_zone"`

	DiscoveryFQDN types.String `tfsdk:"discovery_fqdn"`

	InstanceGroup types.String `tfsdk:"instance_group"`

	InstanceID types.String `tfsdk:"instance_id"`

	InstanceStatus types.String `tfsdk:"instance_status"`

	InstanceStatusReason types.String `tfsdk:"instance_status_reason"`

	InstanceType types.String `tfsdk:"instance_type"`

	InstanceVMType types.String `tfsdk:"instance_vm_type"`

	LifeCycle types.String `tfsdk:"life_cycle"`

	PrivateIP types.String `tfsdk:"private_ip"`

	PublicIP types.String `tfsdk:"public_ip"`

	SSHPort types.Int64 `tfsdk:"ssh_port"`

	SubnetID types.String `tfsdk:"subnet_id"`
}
