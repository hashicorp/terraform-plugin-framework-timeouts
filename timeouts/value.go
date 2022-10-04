package timeouts

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ attr.Value = Timeouts{}

type Timeouts struct {
	// Unknown will be set to true if the entire object is an unknown value.
	// If only some of the elements in the object are unknown, their known or
	// unknown status will be represented however that attr.Value
	// surfaces that information. The Object's Unknown property only tracks
	// if the number of elements in a Object is known, not whether the
	// elements that are in the object are known.
	Unknown bool

	// Null will be set to true if the object is null, either because it was
	// omitted from the configuration, state, or plan, or because it was
	// explicitly set to null.
	Null bool

	Attrs map[string]attr.Value

	AttrTypes map[string]attr.Type
}

type TimeoutsAsOptions struct {
	// UnhandledNullAsEmpty controls what happens when As needs to put a
	// null value in a type that has no way to preserve that distinction.
	// When set to true, the type's empty value will be used.  When set to
	// false, an error will be returned.
	UnhandledNullAsEmpty bool

	// UnhandledUnknownAsEmpty controls what happens when As needs to put
	// an unknown value in a type that has no way to preserve that
	// distinction. When set to true, the type's empty value will be used.
	// When set to false, an error will be returned.
	UnhandledUnknownAsEmpty bool
}

// Type returns an ObjectType with the same attribute types as `o`.
func (t Timeouts) Type(_ context.Context) attr.Type {
	return TimeoutsType{AttrTypes: t.AttrTypes}
}

// ToTerraformValue returns the data contained in the attr.Value as
// a tftypes.Value.
func (t Timeouts) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	if t.AttrTypes == nil {
		return tftypes.Value{}, fmt.Errorf("cannot convert Object to tftypes.Value if AttrTypes field is not set")
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

// IsNull returns true if the Object represents a null value.
func (t Timeouts) IsNull() bool {
	return t.Null
}

// IsUnknown returns true if the Object represents a currently unknown value.
func (t Timeouts) IsUnknown() bool {
	return t.Unknown
}

// String returns a human-readable representation of the Object value.
// The string returned here is not protected by any compatibility guarantees,
// and is intended for logging and error reporting.
func (t Timeouts) String() string {
	if t.Unknown {
		return attr.UnknownValueString
	}

	if t.Null {
		return attr.NullValueString
	}

	// We want the output to be consistent, so we sort the output by key
	keys := make([]string, 0, len(t.Attrs))
	for k := range t.Attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var res strings.Builder

	res.WriteString("{")
	for i, k := range keys {
		if i != 0 {
			res.WriteString(",")
		}
		res.WriteString(fmt.Sprintf(`"%s":%s`, k, t.Attrs[k].String()))
	}
	res.WriteString("}")

	return res.String()
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
