// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"github.com/cloudera/terraform-provider-cdp/skaff/resource"
	"github.com/spf13/cobra"
)

var (
	snakeName     string
	clearComments bool
	name          string
	force         bool
	includeTags   bool
)

var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Create scaffolding for a resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return resource.Create(name, snakeName, !clearComments, force, includeTags)
	},
}

func init() {
	rootCmd.AddCommand(resourceCmd)
	resourceCmd.Flags().StringVarP(&snakeName, "snakename", "s", "", "if skaff doesn't get it right, explicitly give name in snake case (e.g., db_vpc_instance)")
	resourceCmd.Flags().BoolVarP(&clearComments, "clear-comments", "c", false, "do not include instructional comments in source")
	resourceCmd.Flags().StringVarP(&name, "name", "n", "", "name of the entity")
	resourceCmd.Flags().BoolVarP(&force, "force", "f", false, "force creation, overwriting existing files")
	resourceCmd.Flags().BoolVarP(&includeTags, "include-tags", "t", false, "Indicate that this resource has tags and the code for tagging should be generated")
}
