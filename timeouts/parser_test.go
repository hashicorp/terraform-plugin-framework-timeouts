package timeouts_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	def := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"create-not-present": {
			obj:      types.Object{},
			expected: def,
		},
		"create-is-null": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"create": types.String{Null: true},
				},
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
				},
			},
			expected: def,
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
			expected: def,
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
			expected: def,
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
			expected: def,
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
			actual := timeouts.Create(context.Background(), test.obj, def)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}

func TestRead(t *testing.T) {
	t.Parallel()

	def := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"read-not-present": {
			obj:      types.Object{},
			expected: def,
		},
		"read-is-null": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"read": types.String{Null: true},
				},
				AttrTypes: map[string]attr.Type{
					"read": types.StringType,
				},
			},
			expected: def,
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
			expected: def,
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
			expected: def,
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
			expected: def,
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
			actual := timeouts.Read(context.Background(), test.obj, def)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	def := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"update-not-present": {
			obj:      types.Object{},
			expected: def,
		},
		"update-is-null": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"update": types.String{Null: true},
				},
				AttrTypes: map[string]attr.Type{
					"update": types.StringType,
				},
			},
			expected: def,
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
			expected: def,
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
			expected: def,
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
			expected: def,
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
			actual := timeouts.Update(context.Background(), test.obj, def)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	def := 20 * time.Minute

	type testCase struct {
		obj      types.Object
		expected time.Duration
	}

	tests := map[string]testCase{
		"delete-not-present": {
			obj:      types.Object{},
			expected: def,
		},
		"delete-is-null": {
			obj: types.Object{
				Attrs: map[string]attr.Value{
					"delete": types.String{Null: true},
				},
				AttrTypes: map[string]attr.Type{
					"delete": types.StringType,
				},
			},
			expected: def,
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
			expected: def,
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
			expected: def,
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
			expected: def,
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
			actual := timeouts.Delete(context.Background(), test.obj, def)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected duration difference: %s", diff)
			}
		})
	}
}
