// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proxy/v1/proxy.proto

package proxyv1

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

// Validate checks the field values on RequestProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RequestProxyRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetService()) < 1 {
		return RequestProxyRequestValidationError{
			field:  "Service",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetHttpMethod()) < 1 {
		return RequestProxyRequestValidationError{
			field:  "HttpMethod",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		return RequestProxyRequestValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestProxyRequestValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RequestProxyRequestValidationError is the validation error returned by
// RequestProxyRequest.Validate if the designated constraints aren't met.
type RequestProxyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestProxyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestProxyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestProxyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestProxyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestProxyRequestValidationError) ErrorName() string {
	return "RequestProxyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RequestProxyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestProxyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestProxyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestProxyRequestValidationError{}

// Validate checks the field values on RequestProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RequestProxyResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for HttpStatus

	for key, val := range m.GetHeaders() {
		_ = val

		// no validation rules for Headers[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestProxyResponseValidationError{
					field:  fmt.Sprintf("Headers[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestProxyResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RequestProxyResponseValidationError is the validation error returned by
// RequestProxyResponse.Validate if the designated constraints aren't met.
type RequestProxyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestProxyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestProxyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestProxyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestProxyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestProxyResponseValidationError) ErrorName() string {
	return "RequestProxyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RequestProxyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestProxyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestProxyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestProxyResponseValidationError{}
