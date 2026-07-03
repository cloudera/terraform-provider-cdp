// Copyright 2026 Cloudera. All Rights Reserved.
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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var proxyConfigurationSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:          "'id' should exist",
		Field:         "id",
		Computed:      true,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'name' should exist",
		Field:         "name",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'description' should exist",
		Field:         "description",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'protocol' should exist",
		Field:         "protocol",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'host' should exist",
		Field:         "host",
		Computed:      false,
		Required:      true,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'port' should exist",
		Field:         "port",
		Computed:      false,
		Required:      true,
		AttributeType: schema.Int32Attribute{},
	},
	{
		Name:          "'no_proxy_hosts' should exist",
		Field:         "no_proxy_hosts",
		Computed:      false,
		Required:      false,
		AttributeType: schema.SetAttribute{},
	},
	{
		Name:          "'user' should exist",
		Field:         "user",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
	{
		Name:          "'password' should exist",
		Field:         "password",
		Computed:      false,
		Required:      false,
		AttributeType: schema.StringAttribute{},
	},
}

func TestProxyConfigurationSchemaAttributeNumber(t *testing.T) {
	expected := 9
	if len(ProxyConfigurationSchema.Attributes) != expected {
		t.Errorf("The number of fields in the ProxyConfigurationSchema should be: %d but it is: %d", expected, len(ProxyConfigurationSchema.Attributes))
		t.FailNow()
	}
}

func TestProxyConfigurationSchemaContainsElements(t *testing.T) {
	test.PerformResourceSchemaValidation(t, ProxyConfigurationSchema.Attributes, proxyConfigurationSchemaElements)
}
