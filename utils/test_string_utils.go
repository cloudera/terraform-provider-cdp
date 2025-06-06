// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package utils

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Concat concatenates input strings
func Concat(strs ...string) string {
	return strings.Join(strs, "")
}

// CheckStringEquals checks whether two strings are equal and returns nil if so, and an error otherwise.
func CheckStringEquals(name string, expected string, actual string) error {
	if expected != actual {
		return fmt.Errorf("%s name does not match, expected %q, actual: %q", name, expected, actual)
	}
	return nil
}

func ToSetValueFromStringList(sl []string) types.Set {
	if sl == nil {
		return types.SetNull(types.StringType)
	}
	elements := make([]attr.Value, len(sl))
	for i, s := range sl {
		elements[i] = types.StringValue(s)
	}

	setValue, _ := types.SetValue(types.StringType, elements)
	return setValue
}
