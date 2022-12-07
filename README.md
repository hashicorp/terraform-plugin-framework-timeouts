[![PkgGoDev](https://pkg.go.dev/badge/github.com/hashicorp/terraform-plugin-framework-timeouts)](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework-timeouts)

# Terraform Plugin Framework Timeouts

terraform-plugin-framework-timeouts is a Go module containing convenience functions and types for timeouts for use with [terraform-plugin-framework](https://github.com/hashicorp/terraform-plugin-framework). It aims to provide simple access and usage of timeouts defined within configuration.

## Terraform Plugin Framework Compatibility

This Go module is typically kept up to date with the latest `terraform-plugin-framework` releases to ensure all timeouts functionality is available.

## Go Compatibility

This Go module follows `terraform-plugin-framework` Go compatibility.

Currently, that means Go **1.18** must be used when developing and testing code.

## Usage

Usage of this module requires the following changes in the provider code:

- [Schema Mutation](#schema-mutation)
- [Updating Models](#updating-models)
- [Accessing Timeouts in CRUD Functions](#accessing-timeouts-in-crud-functions)

### Schema Mutation

Timeouts can be defined using either nested blocks or nested attributes.

If you are writing a new provider using [terraform-plugin-framework](https://github.com/hashicorp/terraform-plugin-framework)
then we recommend using nested attributes.

If you are [migrating a provider from SDKv2 to the Framework](https://www.terraform.io/plugin/framework/migrating) and 
you are already using timeouts you can either continue to use block syntax, or switch to using nested attributes. 
However, switching to using nested attributes will require that practitioners that are using your provider update their
Terraform configuration.

#### Block

If your configuration is using a nested block to define timeouts, such as the following:

```terraform
resource "timeouts_example" "example" {
  /* ... */

  timeouts {
    create = "60m"
  }
}
```

You can use this module to mutate the `schema.Schema` as follows:

```go
func (t *exampleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    return tfsdk.Schema{
        /* ... */

        Blocks: map[string]tfsdk.Block{
            "timeouts": timeouts.Block(ctx, timeouts.Opts{
                Create: true,
            }),
        },
```

#### Attribute 

If your configuration is using nested attributes to define timeouts, such as the following:

```terraform
resource "timeouts_example" "example" {
  /* ... */

  timeouts = {
    create = "60m"
  }
}
```

You can use this module to mutate the `schema.Schema` as follows:

```go
func (t *exampleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    return tfsdk.Schema{
        Attributes: map[string]tfsdk.Attribute{
            /* ... */
            "timeouts": timeouts.Attributes(ctx, timeouts.Opts{
                Create: true,
            }),
        },
```

### Updating Models

In functions in which the config, state or plan is being unmarshalled, for instance, the `Create` function:

```go
func (r exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data exampleResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
```

The model that is being used, `exampleResourceData` in this example, will need to be modified to include a field for
timeouts which is of type `timeouts.TimeoutsValue`. For example:

```go
type exampleResourceData struct {
    /* ... */
    Timeouts    timeouts.TimeoutsValue `tfsdk:"timeouts"`
```

### Accessing Timeouts in CRUD Functions

Once the model has been populated with the config, state or plan the duration of the timeout can be accessed by calling
the appropriate helper function and then used to configure timeout behaviour, for instance:

```go
func (r exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data exampleResourceData

    diags := req.Plan.Get(ctx, &data)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    createTimeout, err := data.Timeouts.Create(ctx)
    if err != nil {
        // handle error
    }
	
    ctx, cancel := context.WithTimeout(ctx, createTimeout)
    defer cancel()
	
    /* ... */
}
```

## Contributing

See [`.github/CONTRIBUTING.md`](https://github.com/hashicorp/terraform-plugin-framework-timeouts/blob/main/.github/CONTRIBUTING.md)

## License

[Mozilla Public License v2.0](https://github.com/hashicorp/terraform-plugin-framework-timeouts/blob/main/LICENSE)
