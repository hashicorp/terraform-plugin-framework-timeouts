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

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
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
						"create": types.StringType,
						"read":   types.StringType,
						"update": types.StringType,
						"delete": types.StringType,
					},
				},
			},
			input: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"create": tftypes.String,
					"read":   tftypes.String,
					"update": tftypes.String,
					"delete": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"create": tftypes.NewValue(tftypes.String, "60m"),
				"read":   tftypes.NewValue(tftypes.String, "30m"),
				"update": tftypes.NewValue(tftypes.String, "10m"),
				"delete": tftypes.NewValue(tftypes.String, "25m"),
			}),
			expected: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"create": types.StringType,
						"read":   types.StringType,
						"update": types.StringType,
						"delete": types.StringType,
					},
					map[string]attr.Value{
						"create": types.StringValue("60m"),
						"read":   types.StringValue("30m"),
						"update": types.StringValue("10m"),
						"delete": types.StringValue("25m"),
					},
				),
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
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
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := test.receiver.Equal(test.input)
			if test.expected != got {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestTimeoutsValueCreate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		timeoutsValue   timeouts.Value
		expectedTimeout time.Duration
		expectedDiags   diag.Diagnostics
	}
	tests := map[string]testCase{
		"create": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"create": types.StringType,
					},
					map[string]attr.Value{
						"create": types.StringValue("10m"),
					},
				),
			},
			expectedTimeout: 10 * time.Minute,
			expectedDiags:   nil,
		},
		"create-not-set": {
			timeoutsValue: timeouts.Value{
				Object: types.Object{},
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Does Not Exist",
					`timeout for "create" does not exist`,
				),
			},
		},
		"create-not-parseable-as-time-duration": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"create": types.StringType,
					},
					map[string]attr.Value{
						"create": types.StringValue("10x"),
					},
				),
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Cannot Be Parsed",
					`timeout for "create" cannot be parsed, time: unknown unit "x" in duration "10x"`,
				),
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotTimeout, gotErr := test.timeoutsValue.Create(context.Background())

			if diff := cmp.Diff(gotTimeout, test.expectedTimeout); diff != "" {
				t.Errorf("unexpected timeout difference: %s", diff)
			}

			if diff := cmp.Diff(gotErr, test.expectedDiags, equateErrorMessage); diff != "" {
				t.Errorf("unexpected err difference: %s", diff)
			}
		})
	}
}

func TestTimeoutsValueRead(t *testing.T) {
	t.Parallel()

	type testCase struct {
		timeoutsValue   timeouts.Value
		expectedTimeout time.Duration
		expectedDiags   diag.Diagnostics
	}
	tests := map[string]testCase{
		"read": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"read": types.StringType,
					},
					map[string]attr.Value{
						"read": types.StringValue("10m"),
					},
				),
			},
			expectedTimeout: 10 * time.Minute,
			expectedDiags:   nil,
		},
		"read-not-set": {
			timeoutsValue: timeouts.Value{
				Object: types.Object{},
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Does Not Exist",
					`timeout for "read" does not exist`,
				),
			},
		},
		"read-not-parseable-as-time-duration": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"read": types.StringType,
					},
					map[string]attr.Value{
						"read": types.StringValue("10x"),
					},
				),
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Cannot Be Parsed",
					`timeout for "read" cannot be parsed, time: unknown unit "x" in duration "10x"`,
				),
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotTimeout, gotErr := test.timeoutsValue.Read(context.Background())

			if diff := cmp.Diff(gotTimeout, test.expectedTimeout); diff != "" {
				t.Errorf("unexpected timeout difference: %s", diff)
			}

			if diff := cmp.Diff(gotErr, test.expectedDiags, equateErrorMessage); diff != "" {
				t.Errorf("unexpected err difference: %s", diff)
			}
		})
	}
}

func TestTimeoutsValueUpdate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		timeoutsValue   timeouts.Value
		expectedTimeout time.Duration
		expectedDiags   diag.Diagnostics
	}
	tests := map[string]testCase{
		"update": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"update": types.StringType,
					},
					map[string]attr.Value{
						"update": types.StringValue("10m"),
					},
				),
			},
			expectedTimeout: 10 * time.Minute,
			expectedDiags:   nil,
		},
		"update-not-set": {
			timeoutsValue: timeouts.Value{
				Object: types.Object{},
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Does Not Exist",
					`timeout for "update" does not exist`,
				),
			},
		},
		"update-not-parseable-as-time-duration": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"update": types.StringType,
					},
					map[string]attr.Value{
						"update": types.StringValue("10x"),
					},
				),
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Cannot Be Parsed",
					`timeout for "update" cannot be parsed, time: unknown unit "x" in duration "10x"`,
				),
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotTimeout, gotErr := test.timeoutsValue.Update(context.Background())

			if diff := cmp.Diff(gotTimeout, test.expectedTimeout); diff != "" {
				t.Errorf("unexpected timeout difference: %s", diff)
			}

			if diff := cmp.Diff(gotErr, test.expectedDiags, equateErrorMessage); diff != "" {
				t.Errorf("unexpected err difference: %s", diff)
			}
		})
	}
}

func TestTimeoutsValueDelete(t *testing.T) {
	t.Parallel()

	type testCase struct {
		timeoutsValue   timeouts.Value
		expectedTimeout time.Duration
		expectedDiags   diag.Diagnostics
	}
	tests := map[string]testCase{
		"delete": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"delete": types.StringType,
					},
					map[string]attr.Value{
						"delete": types.StringValue("10m"),
					},
				),
			},
			expectedTimeout: 10 * time.Minute,
			expectedDiags:   nil,
		},
		"delete-not-set": {
			timeoutsValue: timeouts.Value{
				Object: types.Object{},
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Does Not Exist",
					`timeout for "delete" does not exist`,
				),
			},
		},
		"delete-not-parseable-as-time-duration": {
			timeoutsValue: timeouts.Value{
				Object: types.ObjectValueMust(
					map[string]attr.Type{
						"delete": types.StringType,
					},
					map[string]attr.Value{
						"delete": types.StringValue("10x"),
					},
				),
			},
			expectedTimeout: 20 * time.Minute,
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Timeout Cannot Be Parsed",
					`timeout for "delete" cannot be parsed, time: unknown unit "x" in duration "10x"`,
				),
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotTimeout, gotErr := test.timeoutsValue.Delete(context.Background())

			if diff := cmp.Diff(gotTimeout, test.expectedTimeout); diff != "" {
				t.Errorf("unexpected timeout difference: %s", diff)
			}

			if diff := cmp.Diff(gotErr, test.expectedDiags, equateErrorMessage); diff != "" {
				t.Errorf("unexpected err difference: %s", diff)
			}
		})
	}
}

// equateErrorMessage reports errors to be equal if both are nil
// or both have the same message.
var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	return x.Error() == y.Error()
})
