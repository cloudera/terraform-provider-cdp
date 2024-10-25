// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package databasecatalog

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type DwDatabaseCatalogModelTestSuite struct {
	suite.Suite
	rm *resourceModel
}

func TestDwModelDatabaseCatalogTestSuite(t *testing.T) {
	suite.Run(t, new(DwDatabaseCatalogModelTestSuite))
}

func (s *DwDatabaseCatalogModelTestSuite) SetupSuite() {
	req := resource.CreateRequest{
		Plan: tfsdk.Plan{
			Raw:    createRawCatalogResource(),
			Schema: testDatabaseCatalogSchema,
		},
	}
	rm := &resourceModel{}
	req.Plan.Get(context.Background(), &rm)
	s.rm = rm
}

func (s *DwDatabaseCatalogModelTestSuite) TestGetPollingTimeout() {
	timeout := s.rm.getPollingTimeout()
	s.Equal(90*time.Minute, timeout)
}

func (s *DwDatabaseCatalogModelTestSuite) TestGetCallFailureThreshold() {
	out := s.rm.getCallFailureThreshold()
	s.Equal(3, out)
}
