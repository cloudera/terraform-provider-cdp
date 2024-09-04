// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package aws

import (
	"context"
	models2 "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type DwClusterModelTestSuite struct {
	suite.Suite
	rm *resourceModel
}

func TestDwModelClusterTestSuite(t *testing.T) {
	suite.Run(t, new(DwClusterModelTestSuite))
}

func (s *DwClusterModelTestSuite) SetupSuite() {
	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawClusterResource(),
			Schema: testDwClusterSchema,
		},
	}
	rm := &resourceModel{}
	req.Plan.Get(context.Background(), &rm)
	s.rm = rm
}

func (s *DwClusterModelTestSuite) TestConvertToCreateAwsClusterRequest() {
	awsCluster := s.rm.convertToCreateAwsClusterRequest()
	s.Equal("crn", *awsCluster.EnvironmentCrn)
	s.Equal(true, awsCluster.UseOverlayNetwork)
	s.Equal([]string{"cidr-1", "cidr-2", "cidr-3"}, awsCluster.WhitelistK8sClusterAccessIPCIDRs)
	s.Equal([]string{"cidr-4", "cidr-5", "cidr-6"}, awsCluster.WhitelistWorkloadAccessIPCIDRs)
	s.Equal(true, awsCluster.UsePrivateLoadBalancer)
	s.Equal(false, awsCluster.UsePublicWorkerNode)
	s.Equal([]string{"subnet-1", "subnet-2", "subnet-3"}, awsCluster.WorkerSubnetIds)
	s.Equal([]string{"subnet-4", "subnet-5", "subnet-6"}, awsCluster.LbSubnetIds)
	s.Equal("", awsCluster.NodeRoleCDWManagedPolicyArn)
	s.Equal(int32(0), *awsCluster.DatabaseBackupRetentionPeriod)
	s.Equal("", awsCluster.CustomSubdomain)
	s.Equal(models2.CustomRegistryOptions{RegistryType: "", RepositoryURL: ""}, *awsCluster.CustomRegistryOptions)
	s.Equal(false, *awsCluster.EnableSpotInstances)
	s.Equal("", awsCluster.CustomAmiID)
	s.Equal([]string{}, awsCluster.ComputeInstanceTypes)
	s.Equal([]string{}, awsCluster.AdditionalInstanceTypes)
}

func (s *DwClusterModelTestSuite) TestGetPollingTimeout() {
	timeout := s.rm.getPollingTimeout()
	s.Equal(90*time.Minute, timeout)
}

func (s *DwClusterModelTestSuite) TestGetCallFailureThreshold() {
	out := s.rm.getCallFailureThreshold()
	s.Equal(3, out)
}
