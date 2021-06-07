// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: project/v1/project.proto

package projectv1

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

// Validate checks the field values on GetProjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProjectsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetNames()) < 1 {
		return GetProjectsRequestValidationError{
			field:  "Names",
			reason: "value must contain at least 1 item(s)",
		}
	}

	return nil
}

// GetProjectsRequestValidationError is the validation error returned by
// GetProjectsRequest.Validate if the designated constraints aren't met.
type GetProjectsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectsRequestValidationError) ErrorName() string {
	return "GetProjectsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectsRequestValidationError{}

// Validate checks the field values on GetProjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProjectsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetProjectsResponseValidationError{
					field:  fmt.Sprintf("Projects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetProjectsResponseValidationError is the validation error returned by
// GetProjectsResponse.Validate if the designated constraints aren't met.
type GetProjectsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectsResponseValidationError) ErrorName() string {
	return "GetProjectsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectsResponseValidationError{}
