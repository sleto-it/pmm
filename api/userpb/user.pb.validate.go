// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: userpb/user.proto

package userpb

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

// Validate checks the field values on UserDetailsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserDetailsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserDetailsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserDetailsRequestMultiError, or nil if none found.
func (m *UserDetailsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserDetailsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserDetailsRequestMultiError(errors)
	}

	return nil
}

// UserDetailsRequestMultiError is an error wrapping multiple validation errors
// returned by UserDetailsRequest.ValidateAll() if the designated constraints
// aren't met.
type UserDetailsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserDetailsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserDetailsRequestMultiError) AllErrors() []error { return m }

// UserDetailsRequestValidationError is the validation error returned by
// UserDetailsRequest.Validate if the designated constraints aren't met.
type UserDetailsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserDetailsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserDetailsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserDetailsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserDetailsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserDetailsRequestValidationError) ErrorName() string {
	return "UserDetailsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UserDetailsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserDetailsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserDetailsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserDetailsRequestValidationError{}

// Validate checks the field values on UserDetailsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserDetailsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserDetailsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserDetailsResponseMultiError, or nil if none found.
func (m *UserDetailsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserDetailsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ProductTourCompleted

	// no validation rules for AlertingTourCompleted

	// no validation rules for SnoozedPmmVersion

	if len(errors) > 0 {
		return UserDetailsResponseMultiError(errors)
	}

	return nil
}

// UserDetailsResponseMultiError is an error wrapping multiple validation
// errors returned by UserDetailsResponse.ValidateAll() if the designated
// constraints aren't met.
type UserDetailsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserDetailsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserDetailsResponseMultiError) AllErrors() []error { return m }

// UserDetailsResponseValidationError is the validation error returned by
// UserDetailsResponse.Validate if the designated constraints aren't met.
type UserDetailsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserDetailsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserDetailsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserDetailsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserDetailsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserDetailsResponseValidationError) ErrorName() string {
	return "UserDetailsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserDetailsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserDetailsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserDetailsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserDetailsResponseValidationError{}

// Validate checks the field values on UserUpdateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserUpdateRequestMultiError, or nil if none found.
func (m *UserUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProductTourCompleted

	// no validation rules for AlertingTourCompleted

	if all {
		switch v := interface{}(m.GetSnoozedPmmVersion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserUpdateRequestValidationError{
					field:  "SnoozedPmmVersion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserUpdateRequestValidationError{
					field:  "SnoozedPmmVersion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSnoozedPmmVersion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserUpdateRequestValidationError{
				field:  "SnoozedPmmVersion",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserUpdateRequestMultiError(errors)
	}

	return nil
}

// UserUpdateRequestMultiError is an error wrapping multiple validation errors
// returned by UserUpdateRequest.ValidateAll() if the designated constraints
// aren't met.
type UserUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserUpdateRequestMultiError) AllErrors() []error { return m }

// UserUpdateRequestValidationError is the validation error returned by
// UserUpdateRequest.Validate if the designated constraints aren't met.
type UserUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserUpdateRequestValidationError) ErrorName() string {
	return "UserUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UserUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserUpdateRequestValidationError{}

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersRequestMultiError, or nil if none found.
func (m *ListUsersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUsersRequestMultiError(errors)
	}

	return nil
}

// ListUsersRequestMultiError is an error wrapping multiple validation errors
// returned by ListUsersRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUsersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersRequestMultiError) AllErrors() []error { return m }

// ListUsersRequestValidationError is the validation error returned by
// ListUsersRequest.Validate if the designated constraints aren't met.
type ListUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersRequestValidationError) ErrorName() string { return "ListUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersRequestValidationError{}

// Validate checks the field values on ListUsersResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUsersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersResponseMultiError, or nil if none found.
func (m *ListUsersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUsersResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUsersResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListUsersResponseMultiError(errors)
	}

	return nil
}

// ListUsersResponseMultiError is an error wrapping multiple validation errors
// returned by ListUsersResponse.ValidateAll() if the designated constraints
// aren't met.
type ListUsersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersResponseMultiError) AllErrors() []error { return m }

// ListUsersResponseValidationError is the validation error returned by
// ListUsersResponse.Validate if the designated constraints aren't met.
type ListUsersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersResponseValidationError) ErrorName() string {
	return "ListUsersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersResponseValidationError{}

// Validate checks the field values on ListUsersResponse_UserDetail with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUsersResponse_UserDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUsersResponse_UserDetail with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUsersResponse_UserDetailMultiError, or nil if none found.
func (m *ListUsersResponse_UserDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUsersResponse_UserDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return ListUsersResponse_UserDetailMultiError(errors)
	}

	return nil
}

// ListUsersResponse_UserDetailMultiError is an error wrapping multiple
// validation errors returned by ListUsersResponse_UserDetail.ValidateAll() if
// the designated constraints aren't met.
type ListUsersResponse_UserDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUsersResponse_UserDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUsersResponse_UserDetailMultiError) AllErrors() []error { return m }

// ListUsersResponse_UserDetailValidationError is the validation error returned
// by ListUsersResponse_UserDetail.Validate if the designated constraints
// aren't met.
type ListUsersResponse_UserDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersResponse_UserDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersResponse_UserDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersResponse_UserDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersResponse_UserDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersResponse_UserDetailValidationError) ErrorName() string {
	return "ListUsersResponse_UserDetailValidationError"
}

// Error satisfies the builtin error interface
func (e ListUsersResponse_UserDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersResponse_UserDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersResponse_UserDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersResponse_UserDetailValidationError{}
