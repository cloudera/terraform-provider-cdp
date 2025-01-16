// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package recipe

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/client/operations"
	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.ResourceWithConfigure   = &recipeResource{}
	_ resource.ResourceWithImportState = &recipeResource{}
)

type recipeResource struct {
	client *cdp.Client
}

func NewRecipeResource() resource.Resource {
	return &recipeResource{}
}

func (r *recipeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_recipe"
}

func (r *recipeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *recipeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = recipeSchema
}

func (r *recipeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan recipeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Datahub

	params := operations.NewCreateRecipeParamsWithContext(ctx)
	content, processErr := processInput(plan.Content.ValueString())
	if processErr != nil {
		utils.AddRecipeDiagnosticsError(processErr, &resp.Diagnostics, "create recipe")
		return
	}
	params.WithInput(&datahubmodels.CreateRecipeRequest{
		RecipeContent: &content,
		RecipeName:    plan.Name.ValueStringPointer(),
		Description:   plan.Description.ValueString(),
		Type:          plan.Type.ValueStringPointer(),
	})

	respOk, err := client.Operations.CreateRecipe(params)
	if err != nil {
		utils.AddRecipeDiagnosticsError(err, &resp.Diagnostics, "create recipe")
		return
	}

	plan.Crn = types.StringPointerValue(respOk.GetPayload().Recipe.Crn)
	plan.ID = types.StringPointerValue(respOk.GetPayload().Recipe.RecipeName)

	// Save plan into Terraform state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *recipeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state recipeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recipe, err := FindRecipeByName(ctx, r.client, state.Name.ValueString())
	if err != nil {
		utils.AddRecipeDiagnosticsError(err, &resp.Diagnostics, "read AWS Credential")
		return
	}
	if recipe == nil {
		resp.State.RemoveResource(ctx) // deleted
		return
	}

	// Overwrite items with refreshed state
	state.ID = types.StringPointerValue(recipe.RecipeName)
	state.Name = types.StringPointerValue(recipe.RecipeName)
	state.Crn = types.StringPointerValue(recipe.Crn)
	if recipe.Description != "" {
		state.Description = types.StringValue(recipe.Description)
	} else {
		state.Description = types.StringNull()
	}
	if !isPath(state.Content.ValueString()) {
		state.Content = types.StringValue(recipe.RecipeContent)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *recipeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Error(ctx, "Update not supported for recipe resource")
}

func (r *recipeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state recipeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	credentialName := state.ID.ValueString()
	params := operations.NewDeleteRecipesParamsWithContext(ctx)
	params.WithInput(&datahubmodels.DeleteRecipesRequest{RecipeNames: []string{credentialName}})
	_, err := r.client.Datahub.Operations.DeleteRecipes(params)
	if err != nil {
		utils.AddRecipeDiagnosticsError(err, &resp.Diagnostics, "delete recipe")
		return
	}
}

func (r *recipeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func FindRecipeByName(ctx context.Context, cdpClient *cdp.Client, recipeName string) (*datahubmodels.Recipe, error) {
	params := operations.NewDescribeRecipeParamsWithContext(ctx)
	params.WithInput(&datahubmodels.DescribeRecipeRequest{RecipeName: &recipeName})
	resp, err := cdpClient.Datahub.Operations.DescribeRecipe(params)
	if err != nil {
		return nil, err
	}
	if resp.GetPayload() == nil {
		return nil, fmt.Errorf("recipe not found")
	} else {
		return resp.GetPayload().Recipe, nil
	}
}
