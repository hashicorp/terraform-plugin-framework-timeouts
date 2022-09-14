package validators

import (
	"context"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ tfsdk.AttributeValidator = timeDurationValidator{}

// timeDurationValidator validates that a string Attribute's value is parseable as time.Duration.
type timeDurationValidator struct {
}

// Description describes the validation in plain text formatting.
func (validator timeDurationValidator) Description(_ context.Context) string {
	return "string must be parseable as time.Duration"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator timeDurationValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator timeDurationValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s := request.AttributeConfig.(types.String)

	if s.Unknown || s.Null {
		return
	}

	if _, err := time.ParseDuration(s.Value); err != nil {
		response.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			request.AttributePath,
			"Invalid Attribute Value Time Duration",
			capitalize(validator.Description(ctx))+", got: "+s.Value))
		return
	}
}

// TimeDuration returns an AttributeValidator which ensures that any configured
// attribute value:
//
//   - Is parseable as time duration.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func TimeDuration() tfsdk.AttributeValidator {
	return timeDurationValidator{}
}

// capitalize will uppercase the first letter in a UTF-8 string.
func capitalize(str string) string {
	if str == "" {
		return ""
	}

	firstRune, size := utf8.DecodeRuneInString(str)

	return string(unicode.ToUpper(firstRune)) + str[size:]
}
