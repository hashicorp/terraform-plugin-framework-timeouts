package timeouts

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ attr.Value = Timeouts{}

type Timeouts struct {
	types.Object
}

// Type returns an ObjectType with the same attribute types as `o`.
func (t Timeouts) Type(_ context.Context) attr.Type {
	return TimeoutsType{
		types.ObjectType{AttrTypes: t.AttrTypes},
	}
}

// ToTerraformValue returns the data contained in the attr.Value as
// a tftypes.Value.
func (t Timeouts) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	if t.AttrTypes == nil {
		return tftypes.Value{}, fmt.Errorf("cannot convert Timeouts to tftypes.Value if AttrTypes field is not set")
	}
	attrTypes := map[string]tftypes.Type{}
	for attr, typ := range t.AttrTypes {
		attrTypes[attr] = typ.TerraformType(ctx)
	}
	objectType := tftypes.Object{AttributeTypes: attrTypes}
	if t.Unknown {
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	}
	if t.Null {
		return tftypes.NewValue(objectType, nil), nil
	}
	vals := map[string]tftypes.Value{}

	for k, v := range t.Attrs {
		val, err := v.ToTerraformValue(ctx)
		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}
		vals[k] = val
	}
	if err := tftypes.ValidateValue(objectType, vals); err != nil {
		return tftypes.NewValue(objectType, tftypes.UnknownValue), err
	}
	return tftypes.NewValue(objectType, vals), nil
}

// Equal returns true if the Object is considered semantically equal
// (same type and same value) to the attr.Value passed as an argument.
func (t Timeouts) Equal(c attr.Value) bool {
	other, ok := c.(Timeouts)
	if !ok {
		return false
	}
	if t.Unknown != other.Unknown {
		return false
	}
	if t.Null != other.Null {
		return false
	}
	if len(t.AttrTypes) != len(other.AttrTypes) {
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
	if len(t.Attrs) != len(other.Attrs) {
		return false
	}
	for k, v := range t.Attrs {
		attr, ok := other.Attrs[k]
		if !ok {
			return false
		}
		if !v.Equal(attr) {
			return false
		}
	}

	return true
}

func (t Timeouts) Create(ctx context.Context, def time.Duration) time.Duration {
	if _, ok := t.Attrs[attributeNameCreate]; !ok {
		return def
	}

	createTimeout := t.Attrs[attributeNameCreate]

	if createTimeout.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for create timeout
	// is a string, this function accepts any types.Object.
	if _, ok := createTimeout.(types.String); !ok {
		return def
	}

	// Although the schema validation guarantees that the type for create timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(createTimeout.(types.String).Value)
	if err != nil {
		return def
	}

	return duration
}
