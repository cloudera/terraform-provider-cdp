## Copyright 2024 Cloudera. All Rights Reserved.
#
# This file is licensed under the Apache License Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. Refer to the License for the specific
# permissions and limitations governing your use of the file.

resource "cdp_recipe" "example" {
  name    = "example-tf-recipe"
  content = <<EOF
#!/bin/bash
echo 'some content'
EOF
  type    = "POST_SERVICE_DEPLOYMENT"
}

output "name" {
  value = cdp_recipe.example.name
}

output "content" {
  value = cdp_recipe.example.content
}

output "type" {
  value = cdp_recipe.example.type
}
