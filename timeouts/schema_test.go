package timeouts_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts"
)

func TestBlock(t *testing.T) {
	t.Parallel()

	type testCase struct {
		opts     timeouts.Opts
		expected tfsdk.Block
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: timeouts.Opts{},
			expected: tfsdk.Block{
				Attributes:  map[string]tfsdk.Attribute{},
				NestingMode: tfsdk.BlockNestingModeSingle,
			},
		},
		"create-opts": {
			opts: timeouts.Opts{
				Create: true,
			},
			expected: tfsdk.Block{
				Attributes: map[string]tfsdk.Attribute{
					"create": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validators.TimeDuration(),
						},
					},
				},
				NestingMode: tfsdk.BlockNestingModeSingle,
			},
		},
		"create-update-opts": {
			opts: timeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: tfsdk.Block{
				Attributes: map[string]tfsdk.Attribute{
					"create": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validators.TimeDuration(),
						},
					},
					"update": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validators.TimeDuration(),
						},
					},
				},
				NestingMode: tfsdk.BlockNestingModeSingle,
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

	expected := tfsdk.Block{
		Attributes: map[string]tfsdk.Attribute{
			"create": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
			"read": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
			"update": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
			"delete": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
		},
		NestingMode: tfsdk.BlockNestingModeSingle,
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected block difference: %s", diff)
	}
}

func TestAttributes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		opts     timeouts.Opts
		expected tfsdk.Attribute
	}
	tests := map[string]testCase{
		"empty-opts": {
			opts: timeouts.Opts{},
			expected: tfsdk.Attribute{
				Optional:   true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{}),
			},
		},
		"create-opts": {
			opts: timeouts.Opts{
				Create: true,
			},
			expected: tfsdk.Attribute{
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"create": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validators.TimeDuration(),
						},
					},
				}),
			},
		},
		"create-update-opts": {
			opts: timeouts.Opts{
				Create: true,
				Update: true,
			},
			expected: tfsdk.Attribute{
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"create": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validators.TimeDuration(),
						},
					},
					"update": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validators.TimeDuration(),
						},
					},
				}),
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

	expected := tfsdk.Attribute{
		Optional: true,
		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
			"create": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
			"read": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
			"update": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
			"delete": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					validators.TimeDuration(),
				},
			},
		}),
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("unexpected block difference: %s", diff)
	}
}
