package timeouts

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	defaultTimeout      = time.Minute * 20
	attributeNameCreate = "create"
	attributeNameRead   = "read"
	attributeNameUpdate = "update"
	attributeNameDelete = "delete"
)

// TimeoutsType is an attribute type that represents timeouts.
type TimeoutsType struct {
	types.ObjectType
}

// ValueFromTerraform returns a TimeoutsValue given a tftypes.Value.
// TimeoutsValue embeds the types.Object value returned from calling ValueFromTerraform on the
// types.ObjectType embedded in TimeoutsType.
func (t TimeoutsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	val, err := t.ObjectType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	obj, ok := val.(types.Object)
	if !ok {
		return nil, fmt.Errorf("%T cannot be used as types.Object", val)
	}

	return TimeoutsValue{
		obj,
	}, err
}

// Equal returns true if `candidate` is also an TimeoutsType and has the same
// AttributeTypes.
func (t TimeoutsType) Equal(candidate attr.Type) bool {
	other, ok := candidate.(TimeoutsType)
	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

// TimeoutsValue represents an object containing values to be used as time.Duration for timeouts.
type TimeoutsValue struct {
	types.Object
}

// Equal returns true if the TimeoutsValue is considered semantically equal
// (same type and same value) to the attr.Value passed as an argument.
func (t TimeoutsValue) Equal(c attr.Value) bool {
	other, ok := c.(TimeoutsValue)

	if !ok {
		return false
	}

	return t.Object.Equal(other.Object)
}

// Create attempts to retrieve the "create" attribute and parse it as time.Duration.
// If any errors are generated they are returned along with the default timeout of 20 minutes.
func (t TimeoutsValue) Create(ctx context.Context) (time.Duration, error) {
	return t.getTimeout(ctx, attributeNameCreate)
}

// Read attempts to retrieve the "read" attribute and parse it as time.Duration.
// If any errors are generated they are returned along with the default timeout of 20 minutes.
func (t TimeoutsValue) Read(ctx context.Context) (time.Duration, error) {
	return t.getTimeout(ctx, attributeNameRead)
}

// Update attempts to retrieve the "update" attribute and parse it as time.Duration.
// If any errors are generated they are returned along with the default timeout of 20 minutes.
func (t TimeoutsValue) Update(ctx context.Context) (time.Duration, error) {
	return t.getTimeout(ctx, attributeNameUpdate)
}

// Delete attempts to retrieve the "delete" attribute and parse it as time.Duration.
// If any errors are generated they are returned along with the default timeout of 20 minutes.
func (t TimeoutsValue) Delete(ctx context.Context) (time.Duration, error) {
	return t.getTimeout(ctx, attributeNameDelete)
}

func (t TimeoutsValue) getTimeout(ctx context.Context, timeoutName string) (time.Duration, error) {
	value, ok := t.Object.Attributes()[timeoutName]
	if !ok {
		return defaultTimeout, fmt.Errorf("timeout for %q does not exist", timeoutName)
	}

	// No type assertion check is required as the schema guarantees that the object attributes
	// are types.String.
	createTimeout, err := time.ParseDuration(value.(types.String).ValueString())
	if err != nil {
		return defaultTimeout, err
	}

	return createTimeout, nil
}
