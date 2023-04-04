// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: updatepb/update.proto

package updatepb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on StartUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StartUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StartUpdateRequestMultiError, or nil if none found.
func (m *StartUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StartUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetContainerId()) < 1 {
		err := StartUpdateRequestValidationError{
			field:  "ContainerId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StartUpdateRequestMultiError(errors)
	}

	return nil
}

// StartUpdateRequestMultiError is an error wrapping multiple validation errors
// returned by StartUpdateRequest.ValidateAll() if the designated constraints
// aren't met.
type StartUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartUpdateRequestMultiError) AllErrors() []error { return m }

// StartUpdateRequestValidationError is the validation error returned by
// StartUpdateRequest.Validate if the designated constraints aren't met.
type StartUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartUpdateRequestValidationError) ErrorName() string {
	return "StartUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StartUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartUpdateRequestValidationError{}

// Validate checks the field values on StartUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StartUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StartUpdateResponseMultiError, or nil if none found.
func (m *StartUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StartUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LogsToken

	if len(errors) > 0 {
		return StartUpdateResponseMultiError(errors)
	}

	return nil
}

// StartUpdateResponseMultiError is an error wrapping multiple validation
// errors returned by StartUpdateResponse.ValidateAll() if the designated
// constraints aren't met.
type StartUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartUpdateResponseMultiError) AllErrors() []error { return m }

// StartUpdateResponseValidationError is the validation error returned by
// StartUpdateResponse.Validate if the designated constraints aren't met.
type StartUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartUpdateResponseValidationError) ErrorName() string {
	return "StartUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StartUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartUpdateResponseValidationError{}

// Validate checks the field values on UpdateStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateStatusRequestMultiError, or nil if none found.
func (m *UpdateStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LogsToken

	// no validation rules for Offset

	if len(errors) > 0 {
		return UpdateStatusRequestMultiError(errors)
	}

	return nil
}

// UpdateStatusRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateStatusRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateStatusRequestMultiError) AllErrors() []error { return m }

// UpdateStatusRequestValidationError is the validation error returned by
// UpdateStatusRequest.Validate if the designated constraints aren't met.
type UpdateStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStatusRequestValidationError) ErrorName() string {
	return "UpdateStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStatusRequestValidationError{}

// Validate checks the field values on UpdateStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateStatusResponseMultiError, or nil if none found.
func (m *UpdateStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Offset

	// no validation rules for Done

	if len(errors) > 0 {
		return UpdateStatusResponseMultiError(errors)
	}

	return nil
}

// UpdateStatusResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateStatusResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateStatusResponseMultiError) AllErrors() []error { return m }

// UpdateStatusResponseValidationError is the validation error returned by
// UpdateStatusResponse.Validate if the designated constraints aren't met.
type UpdateStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStatusResponseValidationError) ErrorName() string {
	return "UpdateStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStatusResponseValidationError{}
