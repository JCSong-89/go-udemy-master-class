package api

import "github.com/go-playground/validator/v10"

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	currency := fl.Field().String()
	switch currency {
	case "USD":
		return true
	}
	return false
}
