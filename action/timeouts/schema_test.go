// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timeouts_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/action/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
)

func TestBlockWithOpts(t *testing.T) {
	t.Parallel()

	type testCase struct {
		opts     timeouts.Opts
		expected schema.Block
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: timeouts.Opts{},
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"invoke": types.StringType,
						},
					},
				},
				Attributes: map[string]schema.Attribute{
					"invoke": schema.StringAttribute{
						Optional: true,
						Description: `A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) ` +
							`consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are ` +
							`"s" (seconds), "m" (minutes), "h" (hours).`,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
			},
		},
		"invoke-opts-description": {
			opts: timeouts.Opts{
				InvokeDescription: "invoke description",
			},
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"invoke": types.StringType,
						},
					},
				},
				Attributes: map[string]schema.Attribute{
					"invoke": schema.StringAttribute{
						Optional:    true,
						Description: "invoke description",
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
			},
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := timeouts.BlockWithOpts(context.Background(), test.opts)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestBlock(t *testing.T) {
	t.Parallel()

	type testCase struct {
		expected schema.Block
	}
	tests := map[string]testCase{
		"invoke": {
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"invoke": types.StringType,
						},
					},
				},
				Attributes: map[string]schema.Attribute{
					"invoke": schema.StringAttribute{
						Optional: true,
						Description: `A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) ` +
							`consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are ` +
							`"s" (seconds), "m" (minutes), "h" (hours).`,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
			},
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := timeouts.Block(context.Background())

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestAttributesWithOpts(t *testing.T) {
	t.Parallel()

	type testCase struct {
		opts     timeouts.Opts
		expected schema.Attribute
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: timeouts.Opts{},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"invoke": schema.StringAttribute{
						Optional: true,
						Description: `A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) ` +
							`consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are ` +
							`"s" (seconds), "m" (minutes), "h" (hours).`,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"invoke": types.StringType,
						},
					},
				},
				Optional: true,
			},
		},
		"invoke-opts-description": {
			opts: timeouts.Opts{
				InvokeDescription: "invoke description",
			},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"invoke": schema.StringAttribute{
						Optional:    true,
						Description: "invoke description",
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"invoke": types.StringType,
						},
					},
				},
				Optional: true,
			},
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := timeouts.AttributesWithOpts(context.Background(), test.opts)

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
		"invoke": {
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"invoke": schema.StringAttribute{
						Optional: true,
						Description: `A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) ` +
							`consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are ` +
							`"s" (seconds), "m" (minutes), "h" (hours).`,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"invoke": types.StringType,
						},
					},
				},
				Optional: true,
			},
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := timeouts.Attributes(context.Background())

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}
