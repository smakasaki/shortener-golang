package validation

import (
	"fmt"
	"regexp"
)

type Rule interface {
	Validate() error
}

type RequiredRule struct {
	FieldName string
	Value     string
}

func ValidateRequired(fieldName, value string) Rule {
	return &RequiredRule{
		FieldName: fieldName,
		Value:     value,
	}
}

func (r *RequiredRule) Validate() error {
	if r.Value == "" {
		return fmt.Errorf("%s is required", r.FieldName)
	}
	return nil
}

type LengthRule struct {
	FieldName string
	Value     string
	Min       int
	Max       int
}

func ValidateLength(fieldName, value string, min, max int) Rule {
	return &LengthRule{
		FieldName: fieldName,
		Value:     value,
		Min:       min,
		Max:       max,
	}
}

func (r *LengthRule) Validate() error {
	length := len(r.Value)
	if length < r.Min || length > r.Max {
		return fmt.Errorf("%s must be between %d and %d characters", r.FieldName, r.Min, r.Max)
	}
	return nil
}

type EmailRule struct {
	FieldName string
	Value     string
}

func ValidateEmail(fieldName, value string) Rule {
	return &EmailRule{
		FieldName: fieldName,
		Value:     value,
	}
}

func (r *EmailRule) Validate() error {
	// Простой email-регулярное выражение для примера
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(r.Value) {
		return fmt.Errorf("%s is not a valid email", r.FieldName)
	}
	return nil
}
