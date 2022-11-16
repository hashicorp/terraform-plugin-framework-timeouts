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
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"create": types.StringType,
				},
				map[string]attr.Value{
					"create": types.StringNull(),
				},
			),
			expected: def,
		},
		"create-not-string": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"create": types.BoolType,
				},
				map[string]attr.Value{
					"create": types.Bool{},
				},
			),
			expected: def,
		},
		"create-not-parseable-empty": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"create": types.StringType,
				},
				map[string]attr.Value{
					"create": types.StringValue(""),
				},
			),
			expected: def,
		},
		"create-not-parseable": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"create": types.StringType,
				},
				map[string]attr.Value{
					"create": types.StringValue("60x"),
				},
			),
			expected: def,
		},
		"create-valid": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"create": types.StringType,
				},
				map[string]attr.Value{
					"create": types.StringValue("60m"),
				},
			),
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
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"read": types.StringType,
				},
				map[string]attr.Value{
					"read": types.StringNull(),
				},
			),
			expected: def,
		},
		"read-not-string": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"read": types.BoolType,
				},
				map[string]attr.Value{
					"read": types.BoolValue(true),
				},
			),
			expected: def,
		},
		"read-not-parseable-empty": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"read": types.StringType,
				},
				map[string]attr.Value{
					"read": types.StringValue(""),
				},
			),
			expected: def,
		},
		"read-not-parseable": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"read": types.StringType,
				},
				map[string]attr.Value{
					"read": types.StringValue("60x"),
				},
			),
			expected: def,
		},
		"read-valid": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"read": types.StringType,
				},
				map[string]attr.Value{
					"read": types.StringValue("60m"),
				},
			),
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
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"update": types.StringType,
				},
				map[string]attr.Value{
					"update": types.StringNull(),
				},
			),
			expected: def,
		},
		"update-not-string": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"update": types.BoolType,
				},
				map[string]attr.Value{
					"update": types.BoolValue(true),
				},
			),
			expected: def,
		},
		"update-not-parseable-empty": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"update": types.StringType,
				},
				map[string]attr.Value{
					"update": types.StringValue(""),
				},
			),
			expected: def,
		},
		"update-not-parseable": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"update": types.StringType,
				},
				map[string]attr.Value{
					"update": types.StringValue("60x"),
				},
			),
			expected: def,
		},
		"update-valid": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"update": types.StringType,
				},
				map[string]attr.Value{
					"update": types.StringValue("60m"),
				},
			),
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
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"delete": types.StringType,
				},
				map[string]attr.Value{
					"delete": types.StringNull(),
				},
			),
			expected: def,
		},
		"delete-not-string": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"delete": types.BoolType,
				},
				map[string]attr.Value{
					"delete": types.BoolValue(true),
				},
			),
			expected: def,
		},
		"delete-not-parseable-empty": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"delete": types.StringType,
				},
				map[string]attr.Value{
					"delete": types.StringValue(""),
				},
			),
			expected: def,
		},
		"delete-not-parseable": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"delete": types.StringType,
				},
				map[string]attr.Value{
					"delete": types.StringValue("60x"),
				},
			),
			expected: def,
		},
		"delete-valid": {
			obj: types.ObjectValueMust(
				map[string]attr.Type{
					"delete": types.StringType,
				},
				map[string]attr.Value{
					"delete": types.StringValue("60m"),
				},
			),
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
