package rules

import (
	"fmt"

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

	// Option 3 is exceptions, but its optional
	var exceptions []any
	if len(options) > 2 {
		for i := 2; i < len(options); i++ {
			if s, ok := options[i].(string); ok {
				exceptions = append(exceptions, s)
			}
		}
	}

	var requestValue string
	switch v := val.(type) {
	case string:
		requestValue = v
	case float64:
		requestValue = fmt.Sprintf("%.0f", v)
	case int:
		requestValue = fmt.Sprintf("%d", v)
	default:
		return false
	}

	if len(requestValue) == 0 {
		return false
	}

	var count int64
	query := facades.Orm().Query().Table(tableName).Where(fieldName, requestValue)

	// Apply exceptions if any
	if len(exceptions) > 0 {
		query = query.WhereNotIn("id", exceptions)
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