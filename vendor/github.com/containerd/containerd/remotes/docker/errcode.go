/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package docker

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ErrorCoder is the base interface for ErrorCode and Error allowing
// users of each to just call ErrorCode to get the real ID of each
type ErrorCoder interface {
	ErrorCode() ErrorCode
}

// ErrorCode represents the error type. The errors are serialized via strings
// and the integer format may change and should *never* be exported.
type ErrorCode int

var _ error = ErrorCode(0)

// ErrorCode just returns itself
func (ec ErrorCode) ErrorCode() ErrorCode {
	return ec
}

// Error returns the ID/Value
func (ec ErrorCode) Error() string {
	// NOTE(stevvooe): Cannot use message here since it may have unpopulated args.
	return strings.ToLower(strings.Replace(ec.String(), "_", " ", -1))
}

// Descriptor returns the descriptor for the error code.
func (ec ErrorCode) Descriptor() ErrorDescriptor {
	d, ok := errorCodeToDescriptors[ec]

	if !ok {
		return ErrorCodeUnknown.Descriptor()
	}

	return d
}

// String returns the canonical identifier for this error code.
func (ec ErrorCode) String() string {
	return ec.Descriptor().Value
}

// Message returned the human-readable error message for this error code.
func (ec ErrorCode) Message() string {
	return ec.Descriptor().Message
}

// MarshalText encodes the receiver into UTF-8-encoded text and returns the
// result.
func (ec ErrorCode) MarshalText() (text []byte, err error) {
	return []byte(ec.String()), nil
}

// UnmarshalText decodes the form generated by MarshalText.
func (ec *ErrorCode) UnmarshalText(text []byte) error {
	desc, ok := idToDescriptors[string(text)]

	if !ok {
		desc = ErrorCodeUnknown.Descriptor()
	}

	*ec = desc.Code

	return nil
}

// WithMessage creates a new Error struct based on the passed-in info and
// overrides the Message property.
func (ec ErrorCode) WithMessage(message string) Error {
	return Error{
		Code:    ec,
		Message: message,
	}
}

// WithDetail creates a new Error struct based on the passed-in info and
// set the Detail property appropriately
func (ec ErrorCode) WithDetail(detail interface{}) Error {
	return Error{
		Code:    ec,
		Message: ec.Message(),
	}.WithDetail(detail)
}

// WithArgs creates a new Error struct and sets the Args slice
func (ec ErrorCode) WithArgs(args ...interface{}) Error {
	return Error{
		Code:    ec,
		Message: ec.Message(),
	}.WithArgs(args...)
}

// Error provides a wrapper around ErrorCode with extra Details provided.
type Error struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`

	// TODO(duglin): See if we need an "args" property so we can do the
	// variable substitution right before showing the message to the user
}

var _ error = Error{}

// ErrorCode returns the ID/Value of this Error
func (e Error) ErrorCode() ErrorCode {
	return e.Code
}

// Error returns a human readable representation of the error.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code.Error(), e.Message)
}

// WithDetail will return a new Error, based on the current one, but with
// some Detail info added
func (e Error) WithDetail(detail interface{}) Error {
	return Error{
		Code:    e.Code,
		Message: e.Message,
		Detail:  detail,
	}
}

// WithArgs uses the passed-in list of interface{} as the substitution
// variables in the Error's Message string, but returns a new Error
func (e Error) WithArgs(args ...interface{}) Error {
	return Error{
		Code:    e.Code,
		Message: fmt.Sprintf(e.Code.Message(), args...),
		Detail:  e.Detail,
	}
}

// ErrorDescriptor provides relevant information about a given error code.
type ErrorDescriptor struct {
	// Code is the error code that this descriptor describes.
	Code ErrorCode

	// Value provides a unique, string key, often captilized with
	// underscores, to identify the error code. This value is used as the
	// keyed value when serializing api errors.
	Value string

	// Message is a short, human readable description of the error condition
	// included in API responses.
	Message string

	// Description provides a complete account of the errors purpose, suitable
	// for use in documentation.
	Description string

	// HTTPStatusCode provides the http status code that is associated with
	// this error condition.
	HTTPStatusCode int
}

// ParseErrorCode returns the value by the string error code.
// `ErrorCodeUnknown` will be returned if the error is not known.
func ParseErrorCode(value string) ErrorCode {
	ed, ok := idToDescriptors[value]
	if ok {
		return ed.Code
	}

	return ErrorCodeUnknown
}

// Errors provides the envelope for multiple errors and a few sugar methods
// for use within the application.
type Errors []error

var _ error = Errors{}

func (errs Errors) Error() string {
	switch len(errs) {
	case 0:
		return "<nil>"
	case 1:
		return errs[0].Error()
	default:
		msg := "errors:\n"
		for _, err := range errs {
			msg += err.Error() + "\n"
		}
		return msg
	}
}

// Len returns the current number of errors.
func (errs Errors) Len() int {
	return len(errs)
}

// MarshalJSON converts slice of error, ErrorCode or Error into a
// slice of Error - then serializes
func (errs Errors) MarshalJSON() ([]byte, error) {
	var tmpErrs struct {
		Errors []Error `json:"errors,omitempty"`
	}

	for _, daErr := range errs {
		var err Error

		switch daErr := daErr.(type) {
		case ErrorCode:
			err = daErr.WithDetail(nil)
		case Error:
			err = daErr
		default:
			err = ErrorCodeUnknown.WithDetail(daErr)

		}

		// If the Error struct was setup and they forgot to set the
		// Message field (meaning its "") then grab it from the ErrCode
		msg := err.Message
		if msg == "" {
			msg = err.Code.Message()
		}

		tmpErrs.Errors = append(tmpErrs.Errors, Error{
			Code:    err.Code,
			Message: msg,
			Detail:  err.Detail,
		})
	}

	return json.Marshal(tmpErrs)
}

// UnmarshalJSON deserializes []Error and then converts it into slice of
// Error or ErrorCode
func (errs *Errors) UnmarshalJSON(data []byte) error {
	var tmpErrs struct {
		Errors []Error
	}

	if err := json.Unmarshal(data, &tmpErrs); err != nil {
		return err
	}

	var newErrs Errors
	for _, daErr := range tmpErrs.Errors {
		// If Message is empty or exactly matches the Code's message string
		// then just use the Code, no need for a full Error struct
		if daErr.Detail == nil && (daErr.Message == "" || daErr.Message == daErr.Code.Message()) {
			// Error's w/o details get converted to ErrorCode
			newErrs = append(newErrs, daErr.Code)
		} else {
			// Error's w/ details are untouched
			newErrs = append(newErrs, Error{
				Code:    daErr.Code,
				Message: daErr.Message,
				Detail:  daErr.Detail,
			})
		}
	}

	*errs = newErrs
	return nil
}
