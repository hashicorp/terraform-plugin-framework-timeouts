package timeouts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
)

const (
	attributeNameCreate = "create"
	attributeNameRead   = "read"
	attributeNameUpdate = "update"
	attributeNameDelete = "delete"
)

// Opts is used as an argument to Block and Attributes to indicate which attributes
// should be created.
type Opts struct {
	Create bool
	Read   bool
	Update bool
	Delete bool
}

// Block returns a tfsdk.Block containing attributes for each of the fields
// in Opts which are set to true. Each attribute is defined as types.StringType
// and optional. A validator is used to verify that the value assigned to an
// attribute can be parsed as time.Duration.
func Block(ctx context.Context, opts Opts) tfsdk.Block {
	return tfsdk.Block{
		Typ:         TimeoutsType{},
		Attributes:  attributesMap(opts),
		NestingMode: tfsdk.BlockNestingModeSingle,
	}
}

// BlockAll returns a tfsdk.Block containing attributes for each of create, read,
// update and delete. Each attribute is defined as types.StringType and optional.
// A validator is used to verify that the value assigned to an attribute can be
// parsed as time.Duration.
func BlockAll(ctx context.Context) tfsdk.Block {
	return Block(ctx, Opts{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	})
}

// Attributes returns a tfsdk.Attribute containing a tfsdk.SingleNestedAttributes
// which contains attributes for each of the fields in Opts which are set to true.
// Each attribute is defined as types.StringType and optional. A validator is used
// to verify that the value assigned to an attribute can be parsed as time.Duration.
func Attributes(ctx context.Context, opts Opts) tfsdk.Attribute {
	return tfsdk.Attribute{
		Optional:   true,
		Attributes: tfsdk.SingleNestedAttributes(attributesMap(opts)),
	}
}

// AttributesAll returns a tfsdk.Attribute containing a tfsdk.SingleNestedAttributes
// which contains attributes for each of create, read, update and delete. Each
// attribute is defined as types.StringType and optional. A validator is used to
// verify that the value assigned to an attribute can be parsed as time.Duration.
func AttributesAll(ctx context.Context) tfsdk.Attribute {
	return Attributes(ctx, Opts{
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
	})
}

func attributesMap(opts Opts) map[string]tfsdk.Attribute {
	attributes := map[string]tfsdk.Attribute{}
	attribute := tfsdk.Attribute{
		Type:     types.StringType,
		Optional: true,
		Validators: []tfsdk.AttributeValidator{
			validators.TimeDuration(),
		},
	}

	if opts.Create {
		attributes[attributeNameCreate] = attribute
	}

	if opts.Read {
		attributes[attributeNameRead] = attribute
	}

	if opts.Update {
		attributes[attributeNameUpdate] = attribute
	}

	if opts.Delete {
		attributes[attributeNameDelete] = attribute
	}

	return attributes
}
