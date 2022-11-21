# 0.2.0 (November 21, 2022)

NOTES:

* all: This Go module has been updated for deprecations in terraform-plugin-framework version 0.15.0 ([#11](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/11))
* all: This Go module has been updated to make it compatible with the breaking changes in terraform-plugin-framework version 0.16.0 ([#12](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/12))

# 0.1.0 (September 22, 2022)

FEATURES:
* Introduced `timeouts` package with `Block()`, `BlockAll()`, `Attributes()` and `AttributesAll()` schema mutation functions and `Create()`, `Read()`, `Update()` and `Delete()` object parsing functions ([#5](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/5))
* Introduced `validators` package with `TimeDuration()` function to obtain time duration validator ([#5](https://github.com/hashicorp/terraform-plugin-framework-timeouts/issues/5))

