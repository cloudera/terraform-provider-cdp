// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environmentconfig

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	awsenv "github.com/cloudera/terraform-provider-cdp/resources/environments/environment/aws"
	azureenv "github.com/cloudera/terraform-provider-cdp/resources/environments/environment/azure"
	gcpenv "github.com/cloudera/terraform-provider-cdp/resources/environments/environment/gcp"
)

type EnvironmentConfigModel struct {
	Name  types.String                            `tfsdk:"name"`
	Crn   types.String                            `tfsdk:"crn"`
	Aws   *awsenv.ResourceModel                   `tfsdk:"aws"`
	Azure *azureenv.AzureEnvironmentResourceModel `tfsdk:"azure"`
	Gcp   *gcpenv.GcpEnvironmentResourceModel     `tfsdk:"gcp"`
}
