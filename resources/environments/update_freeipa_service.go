package environments

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"reflect"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func updateFreeIpaRecipes(ctx context.Context, client *environmentsclient.Environments, planRecipes types.Set, stateRecipes types.Set, environment *string) error {
	recipesToBeAdded := make([]string, 0)
	recipesToBeRemoved := make([]string, 0)

	for _, planRecipe := range utils.FromSetValueToStringList(planRecipes) {
		shouldBeAdded := true
		for _, stateRecipe := range utils.FromSetValueToStringList(stateRecipes) {
			if planRecipe == stateRecipe {
				shouldBeAdded = false
				break
			}
		}
		if shouldBeAdded {
			recipesToBeAdded = append(recipesToBeAdded, planRecipe)
		}
	}

	for _, stateRecipe := range utils.FromSetValueToStringList(stateRecipes) {
		shouldBeRemoved := true
		for _, planRecipe := range utils.FromSetValueToStringList(planRecipes) {
			if stateRecipe == planRecipe {
				shouldBeRemoved = false
				break
			}
		}
		if shouldBeRemoved {
			recipesToBeRemoved = append(recipesToBeRemoved, stateRecipe)
		}
	}

	fmt.Println("recipesToBeAdded: ", recipesToBeAdded)
	fmt.Println("recipesToBeRemoved: ", recipesToBeRemoved)

	paramsToAttach := operations.NewAttachFreeIpaRecipesParamsWithContext(ctx)
	paramsToAttach.WithInput(&environmentsmodels.AttachFreeIpaRecipesRequest{
		Environment: environment,
		Recipes:     recipesToBeAdded,
	})
	_, err := client.Operations.AttachFreeIpaRecipes(paramsToAttach)

	paramsToDetach := operations.NewDetachFreeIpaRecipesParamsWithContext(ctx)
	paramsToDetach.WithInput(&environmentsmodels.DetachFreeIpaRecipesRequest{
		Environment: environment,
		Recipes:     recipesToBeRemoved,
	})
	_, err = client.Operations.DetachFreeIpaRecipes(paramsToDetach)
	return err
}

func updateFreeIpa(ctx context.Context, plan *awsEnvironmentResourceModel, state *awsEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	var freeIpaPlanDetails FreeIpaDetails
	plan.FreeIpa.As(ctx, &freeIpaPlanDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})

	var freeIpaStateDetails FreeIpaDetails
	state.FreeIpa.As(ctx, &freeIpaStateDetails, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})

	isStopped := false

	if !reflect.DeepEqual(freeIpaStateDetails.Recipes, freeIpaPlanDetails.Recipes) {
		if err := updateFreeIpaRecipes(ctx, client, freeIpaPlanDetails.Recipes, freeIpaStateDetails.Recipes, state.EnvironmentName.ValueStringPointer()); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "attaching FreeIPA recipes")
			return resp
		}
	}

	if !reflect.DeepEqual(freeIpaPlanDetails.InstanceType, freeIpaStateDetails.InstanceType) {
		err := stopAndWaitForEnvironment(ctx, plan.EnvironmentName.ValueString(), plan.PollingOptions, resp, client)
		if err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "stopping environment")
			return resp
		}
		isStopped = true
		params := operations.NewStartFreeIpaVerticalScalingParamsWithContext(ctx)
		params.WithInput(
			&environmentsmodels.StartFreeIpaVerticalScalingRequest{
				Environment: state.EnvironmentName.ValueStringPointer(),
				InstanceTemplate: &environmentsmodels.InstanceTemplate{
					InstanceType: freeIpaPlanDetails.InstanceType.ValueString(),
				},
			},
		)
		_, err = client.Operations.StartFreeIpaVerticalScaling(params)
		if err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "free IPA vertical scale")
			return resp
		}
		if err := waitForEnvironmentToBeStopped(state.EnvironmentName.ValueString(), timeoutOneHour, callFailureThreshold, client, ctx, state.PollingOptions); err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "create Environment failed")
			return resp
		}
	}

	if isStopped {
		err := startEnvironment(ctx, state, resp, client)
		if err != nil {
			utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "starting environment")
			return resp
		}
	}

	return resp
}
