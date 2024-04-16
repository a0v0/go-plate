package utils

import (
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

// IsValidUsername validate username against regex
//
//	-> The length must be between 1 and 50 characters.
//	-> The accepted characters are : a-z 0-9 dot(.) underscore(_).
//	-> It's not allowed to have two or more consecutive dots in a row.
//	-> It's not allowed to start or end the username with a dot.
//
// source: https://regexr.com/3cg7r
func IsValidUsername(username string) bool {
	v := validator.New()
	err := v.Var(username, "lowercase,gte=1,lte=50,startsnotwith=.,endsnotwith=.")
	if err != nil {
		return false
	}

	// check for two dots in a row and whitespace
	re := regexp2.MustCompile(`^(?!.*\.\.)(?!.*\.$)(?!.*\s)`, 0)
	isValid, _ := re.MatchString(username)
	return isValid
}
