// Copyright 2023 Cloudera. All Rights Reserved.
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

type FreeIpaInstance struct {
	AttachedVolumes      []*AttachedVolumeDetail `tfsdk:"attachedVolumes"`
	AvailabilityZone     types.String            `tfsdk:"availabilityZone"`
	DiscoveryFQDN        types.String            `tfsdk:"discoveryFQDN"`
	InstanceGroup        types.String            `tfsdk:"instanceGroup"`
	InstanceID           types.String            `tfsdk:"instanceId"`
	InstanceStatus       types.String            `tfsdk:"instanceStatus"`
	InstanceStatusReason types.String            `tfsdk:"instanceStatusReason"`
	InstanceType         types.String            `tfsdk:"instanceType"`
	InstanceVMType       types.String            `tfsdk:"instanceVmType"`
	LifeCycle            types.String            `tfsdk:"lifeCycle"`
	PrivateIP            types.String            `tfsdk:"privateIP"`
	PublicIP             types.String            `tfsdk:"publicIP"`
	SSHPort              types.Int64             `tfsdk:"sshPort"`
	SubnetID             types.String            `tfsdk:"subnetId"`
}

type AttachedVolumeDetail struct {
	Count      types.Int64  `tfsdk:"count"`
	Size       types.Int64  `tfsdk:"size"`
	VolumeType types.String `tfsdk:"volumeType"`
}
