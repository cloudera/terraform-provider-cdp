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

type proxyConfigurationResourceModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	Protocol     types.String `tfsdk:"protocol"`
	Host         types.String `tfsdk:"host"`
	Port         types.Int32  `tfsdk:"port"`
	NoProxyHosts types.Set    `tfsdk:"no_proxy_hosts"`
	User         types.String `tfsdk:"user"`
	Password     types.String `tfsdk:"password"`
}
