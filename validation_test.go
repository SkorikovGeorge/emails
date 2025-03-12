package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// невалидные, которые проходят проверку
func TestSupposedToFail(t *testing.T) {
	tests := []struct {
		email string
	}{
		{"Joe Smith <email@example.com"},
		{".email@example.com"},
		{"email.@example.com"},
		{"email..email@example.com"},
		{"あいうえお@example.com"},
		{"email@example.com (Joe Smith)"},
		{"email@-example.com"},
		{"email@example.web"},
		{"email@111.222.333.44444"},
		{"email@example..com"},
		{"Abc..123@example.com"},
		{`”(),:;<>[\]@example.com`},
		{"just”not”right@example.com"},
		{`this\ is"really"not\allowed@example.com`},
		{"user@domain..aaaa;bbb,cc;d....com"},
	}
	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			err := IsValidEmail(test.email)
			assert.NoError(t, err)
		})
	}
}

// валидные, которые не проходят проверку
func TestSupposedToOK(t *testing.T) {
	tests := []struct {
		email string
	}{
		{"user@[IPv6:2001:db8::1]"},
		{`very.unusual."@".unusual.com@example.com`},
		{`very."(),:;<>[]".VERY."very@\\ "very".unusual@strange.example.com`},
	}
	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			err := IsValidEmail(test.email)
			assert.Error(t, err)
		})
	}
}
