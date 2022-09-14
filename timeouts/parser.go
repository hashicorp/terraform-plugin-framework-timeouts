package timeouts

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Create interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "create" that can be parsed then *time.Duration is returned.
func Create(ctx context.Context, obj types.Object) (*time.Duration, diag.Diagnostics) {
	var diags diag.Diagnostics

	if _, ok := obj.Attrs[create]; !ok {
		diags.AddError(
			"Create Timeout Not Found",
			"Create timeout is not present within the timeouts")
		return nil, diags
	}

	createTimeout := obj.Attrs[create]

	// Although the schema mutation functions guarantee that the type for create timeout
	// is a string, this function accepts any types.Object.
	if _, ok := createTimeout.(types.String); !ok {
		diags.AddError(
			"Create Timeout Not String",
			"Create timeout must be a string")
		return nil, diags
	}

	// Although the schema validation guarantees that the type for create timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(createTimeout.(types.String).Value)
	if err != nil {
		diags.AddError(
			"Create Timeout Not Parseable",
			"Create timeout cannot be parsed as time.Duration")
		return nil, diags
	}

	return &duration, nil
}

// CreateDefault returns time.Duration generated from parsing the "create" value in obj.Attrs
// or the supplied default if "create" cannot be found or parsed.
func CreateDefault(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	duration, diags := Create(ctx, obj)

	if diags.HasError() {
		return def
	}

	return *duration
}

// Read interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "read" that can be parsed then *time.Duration is returned.
func Read(ctx context.Context, obj types.Object) (*time.Duration, diag.Diagnostics) {
	var diags diag.Diagnostics

	if _, ok := obj.Attrs[read]; !ok {
		diags.AddError(
			"Read Timeout Not Found",
			"Read timeout is not present within the timeouts")
		return nil, diags
	}

	readTimeout := obj.Attrs[read]

	// Although the schema mutation functions guarantee that the type for read timeout
	// is a string, this function accepts any types.Object.
	if _, ok := readTimeout.(types.String); !ok {
		diags.AddError(
			"Read Timeout Not String",
			"Read timeout must be a string")
		return nil, diags
	}

	// Although the schema validation guarantees that the type for read timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(readTimeout.(types.String).Value)
	if err != nil {
		diags.AddError(
			"Read Timeout Not Parseable",
			"Read timeout cannot be parsed as time.Duration")
		return nil, diags
	}

	return &duration, nil
}

// ReadDefault returns time.Duration generated from parsing the "read" value in obj.Attrs
// or the supplied default if "read" cannot be found or parsed.
func ReadDefault(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	duration, diags := Read(ctx, obj)

	if diags.HasError() {
		return def
	}

	return *duration
}

// Update interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "update" that can be parsed then *time.Duration is returned.
func Update(ctx context.Context, obj types.Object) (*time.Duration, diag.Diagnostics) {
	var diags diag.Diagnostics

	if _, ok := obj.Attrs[update]; !ok {
		diags.AddError(
			"Update Timeout Not Found",
			"Update timeout is not present within the timeouts")
		return nil, diags
	}

	updateTimeout := obj.Attrs[update]

	// Although the schema mutation functions guarantee that the type for update timeout
	// is a string, this function accepts any types.Object.
	if _, ok := updateTimeout.(types.String); !ok {
		diags.AddError(
			"Update Timeout Not String",
			"Update timeout must be a string")
		return nil, diags
	}

	// Although the schema validation guarantees that the type for update timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(updateTimeout.(types.String).Value)
	if err != nil {
		diags.AddError(
			"Update Timeout Not Parseable",
			"Update timeout cannot be parsed as time.Duration")
		return nil, diags
	}

	return &duration, nil
}

// UpdateDefault returns time.Duration generated from parsing the "update" value in obj.Attrs
// or the supplied default if "update" cannot be found or parsed.
func UpdateDefault(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	duration, diags := Update(ctx, obj)

	if diags.HasError() {
		return def
	}

	return *duration
}

// Delete interrogates the supplied types.Object and if the object.Attrs contains an
// entry for "delete" that can be parsed then *time.Duration is returned.
func Delete(ctx context.Context, obj types.Object) (*time.Duration, diag.Diagnostics) {
	var diags diag.Diagnostics

	if _, ok := obj.Attrs[del]; !ok {
		diags.AddError(
			"Delete Timeout Not Found",
			"Delete timeout is not present within the timeouts")
		return nil, diags
	}

	deleteTimeout := obj.Attrs[del]

	// Although the schema mutation functions guarantee that the type for delete timeout
	// is a string, this function accepts any types.Object.
	if _, ok := deleteTimeout.(types.String); !ok {
		diags.AddError(
			"Delete Timeout Not String",
			"Delete timeout must be a string")
		return nil, diags
	}

	// Although the schema validation guarantees that the type for delete timeout
	// is parseable as a time.Duration, this function accepts any types.Object.
	duration, err := time.ParseDuration(deleteTimeout.(types.String).Value)
	if err != nil {
		diags.AddError(
			"Delete Timeout Not Parseable",
			"Delete timeout cannot be parsed as time.Duration")
		return nil, diags
	}

	return &duration, nil
}

// DeleteDefault returns time.Duration generated from parsing the "delete" value in obj.Attrs
// or the supplied default if "delete" cannot be found or parsed.
func DeleteDefault(ctx context.Context, obj types.Object, def time.Duration) time.Duration {
	duration, diags := Delete(ctx, obj)

	if diags.HasError() {
		return def
	}

	return *duration
}
