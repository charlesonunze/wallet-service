// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pb/v1/wallet_service.proto

package walletpb

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

// Validate checks the field values on CreditUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreditUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreditUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreditUserRequestMultiError, or nil if none found.
func (m *CreditUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreditUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() != 0 {

	}

	if m.GetAmount() != 0 {

	}

	if len(errors) > 0 {
		return CreditUserRequestMultiError(errors)
	}
	return nil
}

// CreditUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreditUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreditUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreditUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreditUserRequestMultiError) AllErrors() []error { return m }

// CreditUserRequestValidationError is the validation error returned by
// CreditUserRequest.Validate if the designated constraints aren't met.
type CreditUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreditUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreditUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreditUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreditUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreditUserRequestValidationError) ErrorName() string {
	return "CreditUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreditUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreditUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreditUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreditUserRequestValidationError{}

// Validate checks the field values on DebitUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DebitUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DebitUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DebitUserRequestMultiError, or nil if none found.
func (m *DebitUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DebitUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() != 0 {

	}

	if m.GetAmount() != 0 {

	}

	if len(errors) > 0 {
		return DebitUserRequestMultiError(errors)
	}
	return nil
}

// DebitUserRequestMultiError is an error wrapping multiple validation errors
// returned by DebitUserRequest.ValidateAll() if the designated constraints
// aren't met.
type DebitUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DebitUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DebitUserRequestMultiError) AllErrors() []error { return m }

// DebitUserRequestValidationError is the validation error returned by
// DebitUserRequest.Validate if the designated constraints aren't met.
type DebitUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DebitUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DebitUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DebitUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DebitUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DebitUserRequestValidationError) ErrorName() string { return "DebitUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e DebitUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDebitUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DebitUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DebitUserRequestValidationError{}

// Validate checks the field values on GetUserBalanceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserBalanceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserBalanceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserBalanceRequestMultiError, or nil if none found.
func (m *GetUserBalanceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserBalanceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() != 0 {

	}

	if len(errors) > 0 {
		return GetUserBalanceRequestMultiError(errors)
	}
	return nil
}

// GetUserBalanceRequestMultiError is an error wrapping multiple validation
// errors returned by GetUserBalanceRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserBalanceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserBalanceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserBalanceRequestMultiError) AllErrors() []error { return m }

// GetUserBalanceRequestValidationError is the validation error returned by
// GetUserBalanceRequest.Validate if the designated constraints aren't met.
type GetUserBalanceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserBalanceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserBalanceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserBalanceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserBalanceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserBalanceRequestValidationError) ErrorName() string {
	return "GetUserBalanceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserBalanceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserBalanceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserBalanceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserBalanceRequestValidationError{}

// Validate checks the field values on GetUserBalanceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserBalanceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserBalanceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserBalanceResponseMultiError, or nil if none found.
func (m *GetUserBalanceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserBalanceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Balance

	if len(errors) > 0 {
		return GetUserBalanceResponseMultiError(errors)
	}
	return nil
}

// GetUserBalanceResponseMultiError is an error wrapping multiple validation
// errors returned by GetUserBalanceResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserBalanceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserBalanceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserBalanceResponseMultiError) AllErrors() []error { return m }

// GetUserBalanceResponseValidationError is the validation error returned by
// GetUserBalanceResponse.Validate if the designated constraints aren't met.
type GetUserBalanceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserBalanceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserBalanceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserBalanceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserBalanceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserBalanceResponseValidationError) ErrorName() string {
	return "GetUserBalanceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserBalanceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserBalanceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserBalanceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserBalanceResponseValidationError{}
