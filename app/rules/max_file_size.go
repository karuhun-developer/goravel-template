package rules

import (
	"mime/multipart"
	"strconv"

	"github.com/goravel/framework/contracts/validation"
)

type MaxFileSize struct {
}

// Signature The name of the rule.
func (receiver *MaxFileSize) Signature() string {
	return "max_file_size"
}

// Passes Determine if the validation rule passes.
func (receiver *MaxFileSize) Passes(data validation.Data, val any, options ...any) bool {
	// val is a file binary
	file := val.(multipart.FileHeader)

	// File size in KB
	fileSizeKb := file.Size / 1024

	// options[0] is max size in KB
	maxSizeKb := options[0].(string)
	maxSizeKbInt, err := strconv.ParseInt(maxSizeKb, 10, 64)

	// If error occurs during parsing, return false
	if err != nil {
		return false
	}

	return fileSizeKb <= maxSizeKbInt
}

// Message Get the validation error message.
func (receiver *MaxFileSize) Message() string {
	return "The :attribute must not be greater than :max kilobytes."
}