package timeouts

import (
	"context"
	"fmt"
	"sort"
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

// AttributeTypes returns the type's attribute types.
func (t TimeoutsType) AttributeTypes() map[string]attr.Type {
	return t.AttrTypes
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the AttributeType to something Terraform
// can understand.
func (t TimeoutsType) TerraformType(ctx context.Context) tftypes.Type {
	attributeTypes := map[string]tftypes.Type{}
	for k, v := range t.AttrTypes {
		attributeTypes[k] = v.TerraformType(ctx)
	}
	return tftypes.Object{
		AttributeTypes: attributeTypes,
	}
}

// ValueFromTerraform returns an attr.Value given a tftypes.Value.
// This is meant to convert the tftypes.Value into a more convenient Go
// type for the provider to consume the data with.
func (t TimeoutsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	object := Timeouts{
		types.Object{
			AttrTypes: t.AttrTypes,
		},
	}
	if in.Type() == nil {
		object.Null = true
		return object, nil
	}
	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}
	if !in.IsKnown() {
		object.Unknown = true
		return object, nil
	}
	if in.IsNull() {
		object.Null = true
		return object, nil
	}
	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}
	err := in.As(&val)
	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := object.AttrTypes[k].ValueFromTerraform(ctx, v)
		if err != nil {
			return nil, err
		}
		attributes[k] = a
	}
	object.Attrs = attributes
	return object, nil
}

// Equal returns true if `candidate` is also an ObjectType and has the same
// AttributeTypes.
func (t TimeoutsType) Equal(candidate attr.Type) bool {
	other, ok := candidate.(TimeoutsType)
	if !ok {
		return false
	}
	if len(other.AttrTypes) != len(t.AttrTypes) {
		return false
	}
	for k, v := range t.AttrTypes {
		attr, ok := other.AttrTypes[k]
		if !ok {
			return false
		}
		if !v.Equal(attr) {
			return false
		}
	}
	return true
}

// ApplyTerraform5AttributePathStep applies the given AttributePathStep to the
// object.
func (t TimeoutsType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	if _, ok := step.(tftypes.AttributeName); !ok {
		return nil, fmt.Errorf("cannot apply step %T to ObjectType", step)
	}

	return t.AttrTypes[string(step.(tftypes.AttributeName))], nil
}

// String returns a human-friendly description of the ObjectType.
func (t TimeoutsType) String() string {
	var res strings.Builder
	res.WriteString("types.TimeoutsType[")
	keys := make([]string, 0, len(t.AttrTypes))
	for k := range t.AttrTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for pos, key := range keys {
		if pos != 0 {
			res.WriteString(", ")
		}
		res.WriteString(`"` + key + `":`)
		res.WriteString(t.AttrTypes[key].String())
	}
	res.WriteString("]")
	return res.String()
}

// ValueType returns the Value type.
func (t TimeoutsType) ValueType(_ context.Context) attr.Value {
	return Timeouts{
		types.Object{
			AttrTypes: t.AttrTypes,
		},
	}
}
