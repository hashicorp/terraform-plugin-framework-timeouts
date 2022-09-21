package timeouts

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Create interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "create" that can be parsed then time.Duration is returned. If object.Attrs
// does not contain "create" the supplied default will be returned.
func Create(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	if _, ok := obj.Attrs[attributeNameCreate]; !ok {
		return def
	}

	createTimeout := obj.Attrs[attributeNameCreate]

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

// Read interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "read" that can be parsed then time.Duration is returned. If object.Attrs
// does not contain "read" the supplied default will be returned.
func Read(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	if _, ok := obj.Attrs[attributeNameRead]; !ok {
		return def
	}

	readTimeout := obj.Attrs[attributeNameRead]

	if readTimeout.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for read timeout
	// is a string, this function accepts any types.Object.
	if _, ok := readTimeout.(types.String); !ok {
		return def
	}

	// Although the schema validation guarantees that the type for read timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(readTimeout.(types.String).Value)
	if err != nil {
		return def
	}

	return duration
}

// Update interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "update" that can be parsed then time.Duration is returned. If object.Attrs
// does not contain "update" the supplied default will be returned.
func Update(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	if _, ok := obj.Attrs[attributeNameUpdate]; !ok {
		return def
	}

	updateTimeout := obj.Attrs[attributeNameUpdate]

	if updateTimeout.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for update timeout
	// is a string, this function accepts any types.Object.
	if _, ok := updateTimeout.(types.String); !ok {
		return def
	}

	// Although the schema validation guarantees that the type for update timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(updateTimeout.(types.String).Value)
	if err != nil {
		return def
	}

	return duration
}

// Delete interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "delete" that can be parsed then time.Duration is returned. If object.Attrs
// does not contain "delete" the supplied default will be returned.
func Delete(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	if _, ok := obj.Attrs[attributeNameDelete]; !ok {
		return def
	}

	deleteTimeout := obj.Attrs[attributeNameDelete]

	if deleteTimeout.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for delete timeout
	// is a string, this function accepts any types.Object.
	if _, ok := deleteTimeout.(types.String); !ok {
		return def
	}

	// Although the schema validation guarantees that the type for delete timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(deleteTimeout.(types.String).Value)
	if err != nil {
		return def
	}

	return duration
}
