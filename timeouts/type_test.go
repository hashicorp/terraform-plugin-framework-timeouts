package timeouts_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts"
)

func TestType_Create(t *testing.T) {
	t.Parallel()

	type testCase struct {
		obj          types.Object
		expected     *time.Duration
		expectedDiag diag.Diagnostic
	}

	tests := map[string]testCase{
		"create-not-present": {
			obj: types.Object{},
			expectedDiag: diag.NewErrorDiagnostic(
				"Create Timeout Not Found",
				"Create timeout is not present within the timeouts",
			),
		},
		"create-not-string": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.Bool{},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.BoolType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Create Timeout Not String",
				"Create timeout must be a string",
			),
		},
		"create-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Create Timeout Not Parseable",
				"Create timeout cannot be parsed as time.Duration",
			),
		},
		"create-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Create Timeout Not Parseable",
				"Create timeout cannot be parsed as time.Duration",
			),
		},
		"create-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
				},
			},
			expected: ptr(60 * time.Minute),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			expectedDiags := diag.Diagnostics{}
			expectedDiags.Append(test.expectedDiag)

			actual, diags := timeouts.Create(context.Background(), test.obj)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}

			if diff := cmp.Diff(diags, expectedDiags); diff != "" {
				t.Errorf("unexpected diags difference: %s", diff)
			}
		})
	}
}

func TestType_CreateDefault(t *testing.T) {
	t.Parallel()

	defaultTimeout := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"create-not-present": {
			expected: defaultTimeout,
		},
		"create-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"create-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"create-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
				},
			},
			expected: 60 * time.Minute,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			actual := timeouts.CreateDefault(context.Background(), test.obj, defaultTimeout)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}

func TestType_Read(t *testing.T) {
	t.Parallel()

	type testCase struct {
		obj          types.Object
		expected     *time.Duration
		expectedDiag diag.Diagnostic
	}

	tests := map[string]testCase{
		"read-not-present": {
			obj: types.Object{},
			expectedDiag: diag.NewErrorDiagnostic(
				"Read Timeout Not Found",
				"Read timeout is not present within the timeouts",
			),
		},
		"read-not-string": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.Bool{},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.BoolType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Read Timeout Not String",
				"Read timeout must be a string",
			),
		},
		"read-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Read Timeout Not Parseable",
				"Read timeout cannot be parsed as time.Duration",
			),
		},
		"read-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Read Timeout Not Parseable",
				"Read timeout cannot be parsed as time.Duration",
			),
		},
		"read-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.StringType,
				},
			},
			expected: ptr(60 * time.Minute),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			expectedDiags := diag.Diagnostics{}
			expectedDiags.Append(test.expectedDiag)

			actual, diags := timeouts.Read(context.Background(), test.obj)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}

			if diff := cmp.Diff(diags, expectedDiags); diff != "" {
				t.Errorf("unexpected diags difference: %s", diff)
			}
		})
	}
}

func TestType_ReadDefault(t *testing.T) {
	t.Parallel()

	defaultTimeout := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"read-not-present": {
			expected: defaultTimeout,
		},
		"read-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"read-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"read-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.StringType,
				},
			},
			expected: 60 * time.Minute,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			actual := timeouts.ReadDefault(context.Background(), test.obj, defaultTimeout)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}

func TestType_Update(t *testing.T) {
	t.Parallel()

	type testCase struct {
		obj          types.Object
		expected     *time.Duration
		expectedDiag diag.Diagnostic
	}

	tests := map[string]testCase{
		"update-not-present": {
			obj: types.Object{},
			expectedDiag: diag.NewErrorDiagnostic(
				"Update Timeout Not Found",
				"Update timeout is not present within the timeouts",
			),
		},
		"update-not-string": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.Bool{},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.BoolType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Update Timeout Not String",
				"Update timeout must be a string",
			),
		},
		"update-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Update Timeout Not Parseable",
				"Update timeout cannot be parsed as time.Duration",
			),
		},
		"update-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Update Timeout Not Parseable",
				"Update timeout cannot be parsed as time.Duration",
			),
		},
		"update-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.StringType,
				},
			},
			expected: ptr(60 * time.Minute),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			expectedDiags := diag.Diagnostics{}
			expectedDiags.Append(test.expectedDiag)

			actual, diags := timeouts.Update(context.Background(), test.obj)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}

			if diff := cmp.Diff(diags, expectedDiags); diff != "" {
				t.Errorf("unexpected diags difference: %s", diff)
			}
		})
	}
}

func TestType_UpdateDefault(t *testing.T) {
	t.Parallel()

	defaultTimeout := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"update-not-present": {
			expected: defaultTimeout,
		},
		"update-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"update-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"update-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.StringType,
				},
			},
			expected: 60 * time.Minute,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			actual := timeouts.UpdateDefault(context.Background(), test.obj, defaultTimeout)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}

func TestType_Delete(t *testing.T) {
	t.Parallel()

	type testCase struct {
		obj          types.Object
		expected     *time.Duration
		expectedDiag diag.Diagnostic
	}

	tests := map[string]testCase{
		"delete-not-present": {
			obj: types.Object{},
			expectedDiag: diag.NewErrorDiagnostic(
				"Delete Timeout Not Found",
				"Delete timeout is not present within the timeouts",
			),
		},
		"delete-not-string": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.Bool{},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.BoolType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Delete Timeout Not String",
				"Delete timeout must be a string",
			),
		},
		"delete-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Delete Timeout Not Parseable",
				"Delete timeout cannot be parsed as time.Duration",
			),
		},
		"delete-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.StringType,
				},
			},
			expectedDiag: diag.NewErrorDiagnostic(
				"Delete Timeout Not Parseable",
				"Delete timeout cannot be parsed as time.Duration",
			),
		},
		"delete-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.StringType,
				},
			},
			expected: ptr(60 * time.Minute),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			expectedDiags := diag.Diagnostics{}
			expectedDiags.Append(test.expectedDiag)

			actual, diags := timeouts.Delete(context.Background(), test.obj)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}

			if diff := cmp.Diff(diags, expectedDiags); diff != "" {
				t.Errorf("unexpected diags difference: %s", diff)
			}
		})
	}
}

func TestType_DeleteDefault(t *testing.T) {
	t.Parallel()

	defaultTimeout := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"delete-not-present": {
			expected: defaultTimeout,
		},
		"delete-not-parseable-empty": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.String{
						Value: "",
					},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"delete-not-parseable": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.String{
						Value: "60x",
					},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.StringType,
				},
			},
			expected: defaultTimeout,
		},
		"delete-valid": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.String{
						Value: "60m",
					},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.StringType,
				},
			},
			expected: 60 * time.Minute,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			actual := timeouts.DeleteDefault(context.Background(), test.obj, defaultTimeout)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
