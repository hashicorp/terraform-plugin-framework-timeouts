## 0.4.0 (June 21, 2023)

NOTES:

* This Go module has been updated to Go 1.19 per the [Go support policy](https://golang.org/doc/devel/release.html#policy). Any consumers building on earlier Go versions may experience errors. ([#40](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/40))

ENHANCEMENTS:

* datasource/timeouts: Add default description for read ([#51](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/51))
* resource/timeouts: Add default description for create, delete, read and update ([#51](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/51))
* resource/timeouts: Add opts for `CreateDescription`, `ReadDescription`, `UpdateDescription` and `DeleteDescription` to allow overriding of default description ([#51](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/51))
* datasource/timeouts: Add `BlockWithOpts()` and `AttributesWithOpts()` functions to allow overriding of default description ([#51](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/51))

## 0.3.1 (February 13, 2023)

BUG FIXES:

* datasource/timeouts: Use default for null and unknown ([#35](https://github.com/hashicorp/terraform-plugin-framework-timeouts/pull/35)). ([#35](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/35))
* resource/timeouts: Use default for null and unknown ([#35](https://github.com/hashicorp/terraform-plugin-framework-timeouts/pull/35)). ([#35](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/35))

# 0.3.0 (December 21, 2022)

BREAKING CHANGES:
* all: The `Attributes() tfsdk.Attribute` method has been removed. Use the resource `Attributes() schema.Attribute` or data source `Attributes() schema.Attribute` function instead. ([#18](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/18))
* all: The `AttributesAll() tfsdk.Attribute` method has been removed. Use the resource `AttributesAll() schema.Attribute` or data source `Attributes() schema.Attribute` function instead. ([#18](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/18))
* all: The `Block() tfsdk.Block` method has been removed. Use the resource `Block() schema.Block` or data source `Block() schema.Block` function instead. ([#18](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/18))
* all: The `BlockAll() tfsdk.Block` method has been removed. Use the resource `BlockAll() schema.Block` or data source `Block() schema.Block` function instead. ([#18](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/18))

FEATURES:
* Introduced `datasource/timeouts` package for use with datasource schema ([#18](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/18))
* Introduced `resource/timeouts` package for use with resource schema ([#18](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/18))


# 0.2.0 (November 21, 2022)

NOTES:

* all: This Go module has been updated for deprecations in terraform-plugin-framework version 0.15.0 ([#11](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/11))
* all: This Go module has been updated to make it compatible with the breaking changes in terraform-plugin-framework version 0.16.0 ([#12](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/12))

# 0.1.0 (September 22, 2022)

FEATURES:
* Introduced `timeouts` package with `Block()`, `BlockAll()`, `Attributes()` and `AttributesAll()` schema mutation functions and `Create()`, `Read()`, `Update()` and `Delete()` object parsing functions ([#5](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/5))
* Introduced `validators` package with `TimeDuration()` function to obtain time duration validator ([#5](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/5))

