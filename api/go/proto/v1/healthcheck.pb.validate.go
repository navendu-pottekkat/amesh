// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/v1/healthcheck.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}
	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckResponseMultiError, or nil if none found.
func (m *HealthCheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Text

	if len(errors) > 0 {
		return HealthCheckResponseMultiError(errors)
	}
	return nil
}

// HealthCheckResponseMultiError is an error wrapping multiple validation
// errors returned by HealthCheckResponse.ValidateAll() if the designated
// constraints aren't met.
type HealthCheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckResponseMultiError) AllErrors() []error { return m }

// HealthCheckResponseValidationError is the validation error returned by
// HealthCheckResponse.Validate if the designated constraints aren't met.
type HealthCheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckResponseValidationError) ErrorName() string {
	return "HealthCheckResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckResponseValidationError{}
