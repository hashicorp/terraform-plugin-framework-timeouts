package datasourcetimeouts_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts/datasource"
)

func TestBlock(t *testing.T) {
	t.Parallel()

	type testCase struct {
		opts     datasourcetimeouts.Opts
		expected schema.Block
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: datasourcetimeouts.Opts{},
			expected: schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{},
			},
		},
		"create-opts": {
			opts: datasourcetimeouts.Opts{
				Create: true,
			},
			expected: schema.SingleNestedBlock{
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
			opts: datasourcetimeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: schema.SingleNestedBlock{
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
			actual := datasourcetimeouts.Block(context.Background(), test.opts)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestBlockAll(t *testing.T) {
	t.Parallel()

	actual := datasourcetimeouts.BlockAll(context.Background())

	expected := schema.SingleNestedBlock{
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
		opts     datasourcetimeouts.Opts
		expected schema.Attribute
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: datasourcetimeouts.Opts{},
			expected: schema.SingleNestedAttribute{
				Optional:   true,
				Attributes: map[string]schema.Attribute{},
			},
		},
		"create-opts": {
			opts: datasourcetimeouts.Opts{
				Create: true,
			},
			expected: schema.SingleNestedAttribute{
				Optional: true,
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
			opts: datasourcetimeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: schema.SingleNestedAttribute{
				Optional: true,
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
			actual := datasourcetimeouts.Attributes(context.Background(), test.opts)

			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("unexpected block difference: %s", diff)
			}
		})
	}
}

func TestAttributesAll(t *testing.T) {
	t.Parallel()

	actual := datasourcetimeouts.AttributesAll(context.Background())

	expected := schema.SingleNestedAttribute{
		Optional: true,
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
