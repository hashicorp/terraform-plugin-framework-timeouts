// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timeouts_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/action/timeouts"
)

func TestTimeoutsTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver    timeouts.Type
		input       tftypes.Value
		expected    attr.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"basic-object": {
			receiver: timeouts.Type{
				ObjectType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"invoke": types.StringType,
					},
				},
			},
			input: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"invoke": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"invoke": tftypes.NewValue(tftypes.String, "30m"),
			}),
			expected: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"invoke": types.StringType,
					},
					map[string]attr.Value{
						"invoke": types.StringValue("30m"),
					},
				),
			},
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := test.receiver.ValueFromTerraform(context.Background(), test.input)
			if err != nil {
				if test.expectedErr == "" {
					t.Errorf("Unexpected error: %s", err.Error())
					return
				}
				if err.Error() != test.expectedErr {
					t.Errorf("Expected error to be %q, got %q", test.expectedErr, err.Error())
					return
				}
			}

			if diff := cmp.Diff(test.expected, got); diff != "" {
				t.Errorf("unexpected result (-expected, +got): %s", diff)
			}
		})
	}
}

func TestTimeoutsTypeEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver timeouts.Type
		input    attr.Type
		expected bool
	}
	tests := map[string]testCase{
		"equal": {
			receiver: timeouts.Type{ObjectType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"a": types.StringType,
				"b": types.NumberType,
				"c": types.BoolType,
				"d": types.ListType{
					ElemType: types.StringType,
				},
			}}},
			input: timeouts.Type{ObjectType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"a": types.StringType,
				"b": types.NumberType,
				"c": types.BoolType,
				"d": types.ListType{
					ElemType: types.StringType,
				},
			}}},
			expected: true,
		},
		"missing-attr": {
			receiver: timeouts.Type{ObjectType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"a": types.StringType,
				"b": types.NumberType,
				"c": types.BoolType,
				"d": types.ListType{
					ElemType: types.StringType,
				},
			}}},
			input: timeouts.Type{ObjectType: types.ObjectType{AttrTypes: map[string]attr.Type{
				"a": types.StringType,
				"b": types.NumberType,
				"d": types.ListType{
					ElemType: types.StringType,
				},
			}}},
			expected: false,
		},
	}
	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := test.receiver.Equal(test.input)
			if test.expected != got {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestTimeoutsValueInvoke(t *testing.T) {
	t.Parallel()

	type testCase struct {
		timeoutsValue   timeouts.Value
		expectedTimeout time.Duration
		expectedDiags   diag.Diagnostics
	}
	tests := map[string]testCase{
		"invoke": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"invoke": types.StringType,
					},
					map[string]attr.Value{
						"invoke": types.StringValue("10m"),
					},
				),
			},
			expectedTimeout: 10 * time.Minute,
			expectedDiags:   nil,
		},
		"invoke-not-set": {
			timeoutsValue: timeouts.Value{
				Object: types.Object{},
			},
			expectedTimeout: 20 * time.Minute,
		},
		"invoke-null": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"invoke": types.StringType,
					},
					map[string]attr.Value{
						"invoke": types.StringNull(),
					},
				),
			},
			expectedTimeout: 20 * time.Minute,
		},
		"invoke-unknown": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"invoke": types.StringType,
					},
					map[string]attr.Value{
						"invoke": types.StringUnknown(),
					},
				),
			},
			expectedTimeout: 20 * time.Minute,
		},
		"invoke-not-parseable-as-time-duration": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"invoke": types.StringType,
					},
					map[string]attr.Value{
						"invoke": types.StringValue("10x"),
					},
				),
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Cannot Be Parsed",
					`timeout for "invoke" cannot be parsed, time: unknown unit "x" in duration "10x"`,
				),
			},
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotTimeout, gotErr := test.timeoutsValue.Invoke(context.Background(), 20*time.Minute)

			if diff := cmp.Diff(gotTimeout, test.expectedTimeout); diff != "" {
				t.Errorf("unexpected timeout difference: %s", diff)
			}

			if diff := cmp.Diff(gotErr, test.expectedDiags); diff != "" {
				t.Errorf("unexpected err difference: %s", diff)
			}
		})
	}
}
