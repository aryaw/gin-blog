package form

import (
    "unicode"
	"fmt"

    "github.com/go-playground/validator/v10"
)

var MustAlphaNum validator.Func = func(fl validator.FieldLevel) bool {
	var hasNumber, hasUpperCase, hasLowercase, hasSpecial bool

	inputValue := fl.Field().String()
	for _, c := range inputValue {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		case c == '#' || c == '|':
		    return false
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
		    hasSpecial = true
		}
	}
	fmt.Println(hasNumber && hasUpperCase && hasLowercase && hasSpecial)
    
	return hasNumber && hasUpperCase && hasLowercase && hasSpecial
}