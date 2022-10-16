package timeouts

import (
	"context"
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
	return t.Object.ToTerraformValue(ctx)
}

// Equal returns true if the Object is considered semantically equal
// (same type and same value) to the attr.Value passed as an argument.
func (t Timeouts) Equal(c attr.Value) bool {
	_, ok := c.(Timeouts)
	if !ok {
		return false
	}

	return t.Object.Equal(c)
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
