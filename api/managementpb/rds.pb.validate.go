// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/rds.proto

package managementpb

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

// Validate checks the field values on DiscoverRDSInstance with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DiscoverRDSInstance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DiscoverRDSInstance with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DiscoverRDSInstanceMultiError, or nil if none found.
func (m *DiscoverRDSInstance) ValidateAll() error {
	return m.validate(true)
}

func (m *DiscoverRDSInstance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Region

	// no validation rules for Az

	// no validation rules for InstanceId

	// no validation rules for NodeModel

	// no validation rules for Address

	// no validation rules for Port

	// no validation rules for Engine

	// no validation rules for EngineVersion

	if len(errors) > 0 {
		return DiscoverRDSInstanceMultiError(errors)
	}

	return nil
}

// DiscoverRDSInstanceMultiError is an error wrapping multiple validation
// errors returned by DiscoverRDSInstance.ValidateAll() if the designated
// constraints aren't met.
type DiscoverRDSInstanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DiscoverRDSInstanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DiscoverRDSInstanceMultiError) AllErrors() []error { return m }

// DiscoverRDSInstanceValidationError is the validation error returned by
// DiscoverRDSInstance.Validate if the designated constraints aren't met.
type DiscoverRDSInstanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverRDSInstanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverRDSInstanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverRDSInstanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverRDSInstanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverRDSInstanceValidationError) ErrorName() string {
	return "DiscoverRDSInstanceValidationError"
}

// Error satisfies the builtin error interface
func (e DiscoverRDSInstanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverRDSInstance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverRDSInstanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverRDSInstanceValidationError{}

// Validate checks the field values on DiscoverRDSRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DiscoverRDSRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DiscoverRDSRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DiscoverRDSRequestMultiError, or nil if none found.
func (m *DiscoverRDSRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DiscoverRDSRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AwsAccessKey

	// no validation rules for AwsSecretKey

	if len(errors) > 0 {
		return DiscoverRDSRequestMultiError(errors)
	}

	return nil
}

// DiscoverRDSRequestMultiError is an error wrapping multiple validation errors
// returned by DiscoverRDSRequest.ValidateAll() if the designated constraints
// aren't met.
type DiscoverRDSRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DiscoverRDSRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DiscoverRDSRequestMultiError) AllErrors() []error { return m }

// DiscoverRDSRequestValidationError is the validation error returned by
// DiscoverRDSRequest.Validate if the designated constraints aren't met.
type DiscoverRDSRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverRDSRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverRDSRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverRDSRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverRDSRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverRDSRequestValidationError) ErrorName() string {
	return "DiscoverRDSRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DiscoverRDSRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverRDSRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverRDSRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverRDSRequestValidationError{}

// Validate checks the field values on DiscoverRDSResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DiscoverRDSResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DiscoverRDSResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DiscoverRDSResponseMultiError, or nil if none found.
func (m *DiscoverRDSResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DiscoverRDSResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRdsInstances() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DiscoverRDSResponseValidationError{
						field:  fmt.Sprintf("RdsInstances[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DiscoverRDSResponseValidationError{
						field:  fmt.Sprintf("RdsInstances[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DiscoverRDSResponseValidationError{
					field:  fmt.Sprintf("RdsInstances[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DiscoverRDSResponseMultiError(errors)
	}

	return nil
}

// DiscoverRDSResponseMultiError is an error wrapping multiple validation
// errors returned by DiscoverRDSResponse.ValidateAll() if the designated
// constraints aren't met.
type DiscoverRDSResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DiscoverRDSResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DiscoverRDSResponseMultiError) AllErrors() []error { return m }

// DiscoverRDSResponseValidationError is the validation error returned by
// DiscoverRDSResponse.Validate if the designated constraints aren't met.
type DiscoverRDSResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverRDSResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverRDSResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverRDSResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverRDSResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverRDSResponseValidationError) ErrorName() string {
	return "DiscoverRDSResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DiscoverRDSResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverRDSResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverRDSResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverRDSResponseValidationError{}

// Validate checks the field values on AddRDSRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddRDSRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddRDSRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddRDSRequestMultiError, or
// nil if none found.
func (m *AddRDSRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddRDSRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetRegion()) < 1 {
		err := AddRDSRequestValidationError{
			field:  "Region",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Az

	if utf8.RuneCountInString(m.GetInstanceId()) < 1 {
		err := AddRDSRequestValidationError{
			field:  "InstanceId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NodeModel

	if utf8.RuneCountInString(m.GetAddress()) < 1 {
		err := AddRDSRequestValidationError{
			field:  "Address",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPort() <= 0 {
		err := AddRDSRequestValidationError{
			field:  "Port",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Engine

	// no validation rules for NodeName

	// no validation rules for ServiceName

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := AddRDSRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Password

	// no validation rules for AwsAccessKey

	// no validation rules for AwsSecretKey

	// no validation rules for RdsExporter

	// no validation rules for QanMysqlPerfschema

	// no validation rules for CustomLabels

	// no validation rules for SkipConnectionCheck

	// no validation rules for Tls

	// no validation rules for TlsSkipVerify

	// no validation rules for DisableQueryExamples

	// no validation rules for TablestatsGroupTableLimit

	// no validation rules for DisableBasicMetrics

	// no validation rules for DisableEnhancedMetrics

	// no validation rules for MetricsMode

	// no validation rules for QanPostgresqlPgstatements

	// no validation rules for AgentPassword

	// no validation rules for Database

	// no validation rules for AutoDiscoveryLimit

	// no validation rules for DisableCommentsParsing

	// no validation rules for MaxPostgresqlExporterConnections

	if len(errors) > 0 {
		return AddRDSRequestMultiError(errors)
	}

	return nil
}

// AddRDSRequestMultiError is an error wrapping multiple validation errors
// returned by AddRDSRequest.ValidateAll() if the designated constraints
// aren't met.
type AddRDSRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddRDSRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddRDSRequestMultiError) AllErrors() []error { return m }

// AddRDSRequestValidationError is the validation error returned by
// AddRDSRequest.Validate if the designated constraints aren't met.
type AddRDSRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddRDSRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddRDSRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddRDSRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddRDSRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddRDSRequestValidationError) ErrorName() string { return "AddRDSRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddRDSRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddRDSRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddRDSRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddRDSRequestValidationError{}

// Validate checks the field values on AddRDSResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddRDSResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddRDSResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddRDSResponseMultiError,
// or nil if none found.
func (m *AddRDSResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddRDSResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRdsExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "RdsExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "RdsExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRdsExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "RdsExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMysql()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "Mysql",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "Mysql",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMysql()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "Mysql",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMysqldExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "MysqldExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "MysqldExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMysqldExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "MysqldExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetQanMysqlPerfschema()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "QanMysqlPerfschema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "QanMysqlPerfschema",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQanMysqlPerfschema()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "QanMysqlPerfschema",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TableCount

	if all {
		switch v := interface{}(m.GetPostgresql()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "Postgresql",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "Postgresql",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPostgresql()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "Postgresql",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPostgresqlExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "PostgresqlExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "PostgresqlExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPostgresqlExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "PostgresqlExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetQanPostgresqlPgstatements()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "QanPostgresqlPgstatements",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddRDSResponseValidationError{
					field:  "QanPostgresqlPgstatements",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQanPostgresqlPgstatements()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddRDSResponseValidationError{
				field:  "QanPostgresqlPgstatements",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddRDSResponseMultiError(errors)
	}

	return nil
}

// AddRDSResponseMultiError is an error wrapping multiple validation errors
// returned by AddRDSResponse.ValidateAll() if the designated constraints
// aren't met.
type AddRDSResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddRDSResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddRDSResponseMultiError) AllErrors() []error { return m }

// AddRDSResponseValidationError is the validation error returned by
// AddRDSResponse.Validate if the designated constraints aren't met.
type AddRDSResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddRDSResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddRDSResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddRDSResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddRDSResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddRDSResponseValidationError) ErrorName() string { return "AddRDSResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddRDSResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddRDSResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddRDSResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddRDSResponseValidationError{}
