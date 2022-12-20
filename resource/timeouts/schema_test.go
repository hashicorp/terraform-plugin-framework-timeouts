package timeouts_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
)

func TestBlock(t *testing.T) {
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
						AttrTypes: map[string]attr.Type{},
					},
				},
				Attributes: map[string]schema.Attribute{},
			},
		},
		"create-opts": {
			opts: timeouts.Opts{
				Create: true,
			},
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"create": types.StringType,
						},
					},
				},
				Attributes: map[string]schema.Attribute{
					"create": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
			},
		},
		"create-update-opts": {
			opts: timeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"create": types.StringType,
							"update": types.StringType,
						},
					},
				},
				Attributes: map[string]schema.Attribute{
					"create": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
					"update": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			actual := timeouts.Block(context.Background(), test.opts)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestBlockAll(t *testing.T) {
	t.Parallel()

	actual := timeouts.BlockAll(context.Background())

	expected := schema.SingleNestedBlock{
		CustomType: timeouts.Type{
			ObjectType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
					"read":   types.StringType,
					"update": types.StringType,
					"delete": types.StringType,
				},
			},
		},
		Attributes: map[string]schema.Attribute{
			"create": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
			"read": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
			"update": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
			"delete": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
		},
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected block difference: %s", diff)
	}
}

func TestAttributes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		opts     timeouts.Opts
		expected schema.Attribute
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: timeouts.Opts{},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{},
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{},
					},
				},
				Optional: true,
			},
		},
		"create-opts": {
			opts: timeouts.Opts{
				Create: true,
			},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"create": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"create": types.StringType,
						},
					},
				},
				Optional: true,
			},
		},
		"create-update-opts": {
			opts: timeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"create": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
					"update": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDuration(),
						},
					},
				},
				CustomType: timeouts.Type{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"create": types.StringType,
							"update": types.StringType,
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
			actual := timeouts.Attributes(context.Background(), test.opts)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestAttributesAll(t *testing.T) {
	t.Parallel()

	actual := timeouts.AttributesAll(context.Background())

	expected := schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"create": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
			"read": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
			"update": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
			"delete": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDuration(),
				},
			},
		},
		CustomType: timeouts.Type{
			ObjectType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"create": types.StringType,
					"read":   types.StringType,
					"update": types.StringType,
					"delete": types.StringType,
				},
			},
		},
		Optional: true,
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected block difference: %s", diff)
	}
}
