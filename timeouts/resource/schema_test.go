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
	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts"
	resourcetimeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts/resource"
)

func TestBlock(t *testing.T) {
	t.Parallel()

	type testCase struct {
		opts     resourcetimeouts.Opts
		expected schema.Block
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: resourcetimeouts.Opts{},
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.TimeoutsType{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{},
					},
				},
				Attributes: map[string]schema.Attribute{},
			},
		},
		"create-opts": {
			opts: resourcetimeouts.Opts{
				Create: true,
			},
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.TimeoutsType{
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
							validators.TimeDurationString(),
						},
					},
				},
			},
		},
		"create-update-opts": {
			opts: resourcetimeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: schema.SingleNestedBlock{
				CustomType: timeouts.TimeoutsType{
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
							validators.TimeDurationString(),
						},
					},
					"update": schema.StringAttribute{
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
			actual := resourcetimeouts.Block(context.Background(), test.opts)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestBlockAll(t *testing.T) {
	t.Parallel()

	actual := resourcetimeouts.BlockAll(context.Background())

	expected := schema.SingleNestedBlock{
		CustomType: timeouts.TimeoutsType{
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
					validators.TimeDurationString(),
				},
			},
			"read": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDurationString(),
				},
			},
			"update": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDurationString(),
				},
			},
			"delete": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDurationString(),
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
		opts     resourcetimeouts.Opts
		expected schema.Attribute
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: resourcetimeouts.Opts{},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{},
				CustomType: timeouts.TimeoutsType{
					ObjectType: types.ObjectType{
						AttrTypes: map[string]attr.Type{},
					},
				},
				Optional: true,
			},
		},
		"create-opts": {
			opts: resourcetimeouts.Opts{
				Create: true,
			},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"create": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDurationString(),
						},
					},
				},
				CustomType: timeouts.TimeoutsType{
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
			opts: resourcetimeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"create": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDurationString(),
						},
					},
					"update": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							validators.TimeDurationString(),
						},
					},
				},
				CustomType: timeouts.TimeoutsType{
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
			actual := resourcetimeouts.Attributes(context.Background(), test.opts)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestAttributesAll(t *testing.T) {
	t.Parallel()

	actual := resourcetimeouts.AttributesAll(context.Background())

	expected := schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"create": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDurationString(),
				},
			},
			"read": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDurationString(),
				},
			},
			"update": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDurationString(),
				},
			},
			"delete": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					validators.TimeDurationString(),
				},
			},
		},
		CustomType: timeouts.TimeoutsType{
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
