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
	"github.com/stretchr/testify/assert"
	"testing"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func TestConvertFreeIpaInstancesWhenResultShouldBeEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input []*environmentsmodels.FreeIpaInstance
	}{
		{name: "test when input slice is nil", input: nil},
		{name: "test when input slice is empty", input: make([]*environmentsmodels.FreeIpaInstance, 0)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ConvertFreeIpaInstances(test.input)
			assert.NotNil(t, result)
			assert.Equal(t, 0, len(*result))
		})
	}
}

func TestConvertFreeIpaInstancesForRootBasicComponents(t *testing.T) {
	testElement := createTestFreeIpaInstanceOfEnvironmentmodels()
	input := make([]*environmentsmodels.FreeIpaInstance, 0)
	input = append(input, testElement)

	result := ConvertFreeIpaInstances(input)

	assert.Equal(t, testElement.PublicIP, (*result)[0].PublicIP.ValueString())
	assert.Equal(t, testElement.SubnetID, (*result)[0].SubnetID.ValueString())
	assert.Equal(t, testElement.LifeCycle, (*result)[0].LifeCycle.ValueString())
	assert.Equal(t, testElement.PrivateIP, (*result)[0].PrivateIP.ValueString())
	assert.Equal(t, testElement.InstanceID, (*result)[0].InstanceID.ValueString())
	assert.Equal(t, testElement.InstanceType, (*result)[0].InstanceType.ValueString())
	assert.Equal(t, testElement.DiscoveryFQDN, (*result)[0].DiscoveryFQDN.ValueString())
	assert.Equal(t, testElement.InstanceGroup, (*result)[0].InstanceGroup.ValueString())
	assert.Equal(t, testElement.InstanceStatus, (*result)[0].InstanceStatus.ValueString())
	assert.Equal(t, testElement.InstanceVMType, (*result)[0].InstanceVMType.ValueString())
	assert.Equal(t, testElement.InstanceStatusReason, (*result)[0].InstanceStatusReason.ValueString())
	assert.Equal(t, func(i int32) int64 { return int64(i) }(testElement.SSHPort), (*result)[0].SSHPort.ValueInt64())
}

func TestConvertFreeIpaInstancesForAttachedVolumesSize(t *testing.T) {
	testElement := createTestFreeIpaInstanceOfEnvironmentmodels()
	testElement.AttachedVolumes = append(testElement.AttachedVolumes, &environmentsmodels.AttachedVolumeDetail{
		Count:      4321,
		Size:       4321,
		VolumeType: "someOtherVolumeType",
	})
	input := make([]*environmentsmodels.FreeIpaInstance, 0)
	input = append(input, testElement)

	result := ConvertFreeIpaInstances(input)

	assert.Equal(t, len(testElement.AttachedVolumes), len((*result)[0].AttachedVolumes))
}

func TestConvertFreeIpaInstancesForAttachedVolumesContent(t *testing.T) {
	testElement := createTestFreeIpaInstanceOfEnvironmentmodels()
	input := make([]*environmentsmodels.FreeIpaInstance, 0)
	input = append(input, testElement)

	result := ConvertFreeIpaInstances(input)

	assert.Equal(t, testElement.AttachedVolumes[0].VolumeType, (*result)[0].AttachedVolumes[0].VolumeType.ValueString())
	assert.Equal(t, utils.ConvertInt32ToTypesInt64(testElement.AttachedVolumes[0].Count), (*result)[0].AttachedVolumes[0].Count)
	assert.Equal(t, utils.ConvertInt32ToTypesInt64(testElement.AttachedVolumes[0].Size), (*result)[0].AttachedVolumes[0].Size)
}

func TestConvertFreeIpaInstancesIfThereAreMultipleInputsThenMultipleOutputShoutReturn(t *testing.T) {
	input := make([]*environmentsmodels.FreeIpaInstance, 0)
	input = append(input, createTestFreeIpaInstanceOfEnvironmentmodels())
	input = append(input, createTestFreeIpaInstanceOfEnvironmentmodels())

	assert.Equal(t, len(input), len(*ConvertFreeIpaInstances(input)))
}

func createTestFreeIpaInstanceOfEnvironmentmodels() *environmentsmodels.FreeIpaInstance {
	return &environmentsmodels.FreeIpaInstance{
		AttachedVolumes: func(slc []*environmentsmodels.AttachedVolumeDetail) []*environmentsmodels.AttachedVolumeDetail {
			vol := &environmentsmodels.AttachedVolumeDetail{
				Count:      1234,
				Size:       1234,
				VolumeType: "someVolumeType",
			}
			slc = append(slc, vol)
			return slc
		}(make([]*environmentsmodels.AttachedVolumeDetail, 0)),
		AvailabilityZone:     "someAvailabilityZone",
		DiscoveryFQDN:        "someDiscoveryFQDN",
		InstanceGroup:        "someInstanceGroup",
		InstanceID:           "someInstanceID",
		InstanceStatus:       "someInstanceStatus",
		InstanceStatusReason: "someInstanceStatusReason",
		InstanceType:         "someInstanceType",
		InstanceVMType:       "someInstanceVMType",
		LifeCycle:            "someLifeCycle",
		PrivateIP:            "somePrivateIP",
		PublicIP:             "somePublicIP",
		SSHPort:              1234,
		SubnetID:             "someSubnetID",
	}
}
