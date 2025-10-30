package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type Exists struct {
}

// Signature The name of the rule.
func (receiver *Exists) Signature() string {
	return "exists"
}

// Passes Determine if the validation rule passes.
func (receiver *Exists) Passes(_ validation.Data, val any, options ...any) bool {

	// Get table name, field name, and request value
	tableName := options[0].(string)
	fieldName := options[1].(string)
	requestValue := val.(string)

	if len(requestValue) == 0 {
		return false
	}

	var count int64
	query := facades.Orm().Query().Table(tableName).Where(fieldName, requestValue)
	if len(options) > 2 {
		for i := 2; i < len(options); i++ {
			query = query.OrWhere(options[i].(string), requestValue)
		}
	}
	count, err := query.Count()
	if err != nil {
		return false
	}

	return count != 0
}

// Message Get the validation error message.
func (receiver *Exists) Message() string {
	return "The :attribute does not exist."
}