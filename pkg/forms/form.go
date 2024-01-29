package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](" +
	"?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Form struct {
	url.Values
	FormErrors formErrors
}

func NewForm(data url.Values) *Form {
	return &Form{
		data,
		formErrors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field) // Using embedded url.Values.Get method
		if strings.TrimSpace(value) == "" {
			f.FormErrors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field) // Using embedded url.Values.Get method
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.FormErrors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters", d))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field) // Using embedded url.Values.Get method
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.FormErrors.Add(field, "This field is invalid")
}

func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.FormErrors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)",
			d))
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.FormErrors.Add(field, "This field is invalid")
	}
}

func (f *Form) Valid() bool {
	return len(f.FormErrors) == 0
}
