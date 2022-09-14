package validators_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
)

func TestTimeDuration(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		expectError bool
	}

	tests := map[string]testCase{
		"unknown": {
			val: types.String{Unknown: true},
		},
		"null": {
			val: types.String{Null: true},
		},
		"valid": {
			val: types.String{Value: "20m"},
		},
		"invalid": {
			val:         types.String{Value: "20x"},
			expectError: true,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			request := tfsdk.ValidateAttributeRequest{
				AttributePath:           path.Root("test"),
				AttributePathExpression: path.MatchRoot("test"),
				AttributeConfig:         test.val,
			}

			response := tfsdk.ValidateAttributeResponse{}

			validators.TimeDuration().Validate(context.Background(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
