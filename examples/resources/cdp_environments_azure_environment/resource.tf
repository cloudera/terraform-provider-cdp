## Copyright 2023 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

terraform {
  required_providers {
    cdp = {
      source = "cloudera/cdp"
    }
  }
}

provider "cdp" {
  cdp_profile                 = "mowdev"
  cdp_endpoint_url            = "https://cloudera.dps.mow-dev.cloudera.com"
  cdp_config_file             = "/Users/gmeszaros/.cdp/config"
  cdp_shared_credentials_file = "/Users/gmeszaros/.cdp/credentials"
  local_environment           = false
}

resource "cdp_environments_azure_environment" "my-azure-environment-flexible" {
  environment_name = "gmeszaros-tf-env-2"
  credential_name  = "dszabo-az-cred"
  region           = "westus2"
  security_access = {
    cidr = "0.0.0.0/0"
  }
  existing_network_params = {
    resource_group_name          = "cloud-daily"
    network_id                   = "cloud-daily"
    subnet_ids                   = ["cloud-daily.internal.0.westus2", "cloud-daily.internal.1.westus2"]
    flexible_server_subnet_ids   = ["subnet_10_124_116_0-23"]
    database_private_dns_zone_id = "/subscriptions/3ddda1c7-d1f5-4e7b-ac81-0523f483b3b3/resourceGroups/cloud-daily/providers/Microsoft.Network/privateDnsZones/flexible.cloud-daily.postgres.database.azure.com"
  }
  resource_group_name = "dszabo-cdp"
  public_key          = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCoACV6cXVqQskr4NPhlmR/rQ40p7xOGgKG3kS2GR0+iU3N5hK6NeK1AWecb9FxfiSSAS3Z9Nxpx/Ktq8PHWgIa5gsnHMHRvU5pPHHIQQbUc6/Ju7hxfX+t/rgX5y0q7o1UrT+m4HTP3bQESNg5ggRFnvgKeZtZf6xOpCh37ij4k2Whcc6K8z2nTPa+pRAjmYiYvxlGnOv3ga4yQyID35NuBnIOqZQPn3mVoIw60klUNbJ45l5gsr/hPHAMqO2+TnTNpajU0tecQED8iFlgR7293PSkwb94sPelQbey1Kr5s8WozqsAyA99kEpr91KHNYb9KhVC5E/Roxfx5R9+PxK1WCdx3cTiwc2NSrd+nQwc/DeE9DPqrZxmiHSnWxtRFnneh1H+I8pf/YDndTdHziHxoTz/IQ94pgKKX4yC+j2TRa4yvETtTVXfIilVjJtswXbpOwA2MSiU+a7y6tdhOyq3hhdLS49kiNNqH1ILnQ0PoEK9RchRPjtoFFUd10N5+20= cloudbreak"
  log_storage = {
    managed_identity      = "/subscriptions/3ddda1c7-d1f5-4e7b-ac81-0523f483b3b3/resourcegroups/dszabo-cdp/providers/Microsoft.ManagedIdentity/userAssignedIdentities/dszabo-logger"
    storage_location_base = "abfs://dszabocdpfs@dszabocdpstorage.dfs.core.windows.net"
  }
  use_public_ip                 = false
  enable_outbound_load_balancer = true
  encryption_at_host            = true
  freeipa = {
    instance_count_by_group = 2
    os = "centos7"
  }
  tags = {
    "made-with" : "CDP Terraform Provider"
  }
}

output "environment" {
  value = cdp_environments_azure_environment.my-azure-environment-flexible
}