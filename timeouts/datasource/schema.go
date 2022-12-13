package timeouts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/internal/validators"
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/timeouts/type"
)

const (
	attributeNameRead = "read"
)

// Opts is used as an argument to Block and Attributes to indicate which attributes
// should be created.
type Opts struct {
	Read bool
}

// Block returns a schema.Block containing attributes for `Read`, which is
// defined as types.StringType and optional. A validator is used to verify
// that the value assigned to `Read` can be parsed as time.Duration.
func Block(ctx context.Context) schema.Block {
	return schema.SingleNestedBlock{
		Attributes: attributesMap(),
		CustomType: timeouts.TimeoutsType{
			ObjectType: types.ObjectType{
				AttrTypes: attrTypesMap(),
			},
		},
	}
}

// Attributes returns a schema.SingleNestedAttribute which contains an
// attribute for `Read`, which is defined as types.StringType and optional.
// A validator is used to verify that the value assigned to an attribute
// can be parsed as time.Duration.
func Attributes(ctx context.Context) schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: attributesMap(),
		CustomType: timeouts.TimeoutsType{
			ObjectType: types.ObjectType{
				AttrTypes: attrTypesMap(),
			},
		},
		Optional: true,
	}
}

func attributesMap() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		attributeNameRead: schema.StringAttribute{
			Optional: true,
			Validators: []validator.String{
				validators.TimeDurationString(),
			},
		},
	}
}

func attrTypesMap() map[string]attr.Type {
	return map[string]attr.Type{
		attributeNameRead: types.StringType,
	}
}
