package timeouts

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ attr.Type = TimeoutsType{}
)

type TimeoutsType struct {
	types.ObjectType
}

// WithAttributeTypes returns a new copy of the type with its attribute types
// set.
func (t TimeoutsType) WithAttributeTypes(typs map[string]attr.Type) attr.TypeWithAttributeTypes {
	return TimeoutsType{
		types.ObjectType{
			AttrTypes: typs,
		},
	}
}

// ValueFromTerraform returns an attr.Value given a tftypes.Value.
// This is meant to convert the tftypes.Value into a more convenient Go
// type for the provider to consume the data with.
func (t TimeoutsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	obj, err := t.ObjectType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	return Timeouts{
		obj.(types.Object),
	}, nil
}

// Equal returns true if `candidate` is also an ObjectType and has the same
// AttributeTypes.
func (t TimeoutsType) Equal(candidate attr.Type) bool {
	_, ok := candidate.(TimeoutsType)
	if !ok {
		return false
	}

	return t.ObjectType.Equal(candidate.(TimeoutsType).ObjectType)
}

// String returns a human-friendly description of the ObjectType.
func (t TimeoutsType) String() string {
	return strings.Replace(t.ObjectType.String(), "types.ObjectType[", "timeouts.TimeoutsType[", 1)
}

// ValueType returns the Value type.
func (t TimeoutsType) ValueType(_ context.Context) attr.Value {
	return Timeouts{
		types.Object{
			AttrTypes: t.AttrTypes,
		},
	}
}
