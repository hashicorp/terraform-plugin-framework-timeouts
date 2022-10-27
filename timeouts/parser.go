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
	createTimeoutValue, ok := obj.Attributes()[attributeNameCreate]

	if !ok {
		return def
	}

	if createTimeoutValue.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for create timeout
	// is a string, this function accepts any types.Object.
	createTimeoutString, ok := createTimeoutValue.(types.String)

	if !ok {
		return def
	}

	// Although the schema validation guarantees that the type for create timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(createTimeoutString.ValueString())
	if err != nil {
		return def
	}

	return duration
}

// Read interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "read" that can be parsed then time.Duration is returned. If object.Attrs
// does not contain "read" the supplied default will be returned.
func Read(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	readTimeoutValue, ok := obj.Attributes()[attributeNameRead]

	if !ok {
		return def
	}

	if readTimeoutValue.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for read timeout
	// is a string, this function accepts any types.Object.
	readTimeoutString, ok := readTimeoutValue.(types.String)

	if !ok {
		return def
	}

	// Although the schema validation guarantees that the type for read timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(readTimeoutString.ValueString())
	if err != nil {
		return def
	}

	return duration
}

// Update interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "update" that can be parsed then time.Duration is returned. If object.Attrs
// does not contain "update" the supplied default will be returned.
func Update(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	updateTimeoutValue, ok := obj.Attributes()[attributeNameUpdate]

	if !ok {
		return def
	}

	if updateTimeoutValue.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for update timeout
	// is a string, this function accepts any types.Object.
	updateTimeoutString, ok := updateTimeoutValue.(types.String)

	if !ok {
		return def
	}

	// Although the schema validation guarantees that the type for update timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(updateTimeoutString.ValueString())
	if err != nil {
		return def
	}

	return duration
}

// Delete interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "delete" that can be parsed then time.Duration is returned. If object.Attrs
// does not contain "delete" the supplied default will be returned.
func Delete(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	deleteTimeoutValue, ok := obj.Attributes()[attributeNameDelete]

	if !ok {
		return def
	}

	if deleteTimeoutValue.IsNull() {
		return def
	}

	// Although the schema mutation functions guarantee that the type for delete timeout
	// is a string, this function accepts any types.Object.
	deleteTimeoutString, ok := deleteTimeoutValue.(types.String)

	if !ok {
		return def
	}

	// Although the schema validation guarantees that the type for delete timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(deleteTimeoutString.ValueString())
	if err != nil {
		return def
	}

	return duration
}
