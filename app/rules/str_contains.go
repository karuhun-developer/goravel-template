package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/support/str"
)

type StrContains struct {
}

// Signature The name of the rule.
func (receiver *StrContains) Signature() string {
	return "str_contains"
}

// Passes Determine if the validation rule passes.
func (receiver *StrContains) Passes(data validation.Data, val any, options ...any) bool {
	requestValue := val.(string)

	// Str contains
	strContains := options[0].(string)

	// Valid
	return str.Of(requestValue).Contains(strContains)
}

// Message Get the validation error message.
func (receiver *StrContains) Message() string {
	return "The :attribute must contain one of the following substrings."
}
