package user

import "goyave.dev/goyave/v4/validation"

var (
	StoreRequest = validation.RuleSet{
		"email":    validation.List{"required", "email"},
		"password": validation.List{"required", "min:6"},
	}
)
