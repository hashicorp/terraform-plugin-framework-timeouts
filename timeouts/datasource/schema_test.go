package timeouts_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
	datasourcetimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts/datasource"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts/type"
)

func TestBlock(t *testing.T) {
	t.Parallel()

	type testCase struct {
		expected schema.Block
	}
	tests := map[string]testCase{
		"read": {
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.TimeoutsType{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"read": types.StringType,
						},
					},
				},
				Attributes: map[string]schema.Attribute{
					"read": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDurationString(),
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			actual := datasourcetimeouts.Block(context.Background())

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestAttributes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		expected schema.Attribute
	}
	tests := map[string]testCase{
		"read": {
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"read": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDurationString(),
						},
					},
				},
				CustomType: timeouts.TimeoutsType{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"read": types.StringType,
						},
					},
				},
				Optional: true,
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			actual := datasourcetimeouts.Attributes(context.Background())

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}
