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

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func ConvertFreeIpaInstances(freeIpaInstances []*environmentsmodels.FreeIpaInstance) *[]FreeIpaInstance {
	var instances []FreeIpaInstance
	if freeIpaInstances != nil || len(freeIpaInstances) > 0 {
		instances = make([]FreeIpaInstance, 0)
		for _, instance := range freeIpaInstances {
			var attachedVolumes []*AttachedVolumeDetail
			if instance.AttachedVolumes != nil {
				attachedVolumes = make([]*AttachedVolumeDetail, 0)
				for _, volume := range instance.AttachedVolumes {
					attachedVolumes = append(attachedVolumes, &AttachedVolumeDetail{
						Count:      utils.ConvertInt32ToTypesInt64(volume.Count),
						Size:       utils.ConvertInt32ToTypesInt64(volume.Size),
						VolumeType: types.StringValue(volume.VolumeType),
					})
				}
			}
			instances = append(instances, FreeIpaInstance{
				AttachedVolumes:      attachedVolumes,
				AvailabilityZone:     types.StringValue(instance.AvailabilityZone),
				DiscoveryFQDN:        types.StringValue(instance.DiscoveryFQDN),
				InstanceGroup:        types.StringValue(instance.InstanceGroup),
				InstanceID:           types.StringValue(instance.InstanceID),
				InstanceStatus:       types.StringValue(instance.InstanceStatus),
				InstanceStatusReason: types.StringValue(instance.InstanceStatusReason),
				InstanceType:         types.StringValue(instance.InstanceType),
				InstanceVMType:       types.StringValue(instance.InstanceVMType),
				LifeCycle:            types.StringValue(instance.LifeCycle),
				PrivateIP:            types.StringValue(instance.PrivateIP),
				PublicIP:             types.StringValue(instance.PublicIP),
				SSHPort:              utils.ConvertInt32ToTypesInt64(instance.SSHPort),
				SubnetID:             types.StringValue(instance.SubnetID),
			})
		}
	}
	return &instances
}
