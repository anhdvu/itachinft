package validator

import (
	"regexp"
	"slices"
)

var (
	EmailRegExp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]any
}

func New() *Validator {
	return &Validator{Errors: map[string]any{}}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, value string) {
	if _, found := v.Errors[key]; !found {
		v.Errors[key] = value
	}
}

func (v *Validator) Check(ok bool, key, value string) {
	if !ok {
		v.AddError(key, value)
	}
}

func PermittedValues[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

func Matches(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}

func Unique[T comparable](values []T) bool {
	m := map[T]bool{}
	for _, value := range values {
		m[value] = true
	}
	return len(m) == len(values)
}
