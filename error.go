package com

// OleError stores COM errors.
type OleError struct {
	// HResult code or pointer.
	hr uintptr

	// Error description.
	description string

	// Parent COM error.
	subError error
}

// NewError creates new error with HResult.
func NewError(hr uintptr) *OleError {
	return &OleError{hr: hr}
}

// NewErrorWithDescription creates new COM error with HResult and description.
func NewErrorWithDescription(hr uintptr, description string) *OleError {
	return &OleError{hr: hr, description: description}
}

// NewErrorWithSubError creates new COM error with parent error.
func NewErrorWithSubError(hr uintptr, description string, err error) *OleError {
	return &OleError{hr: hr, description: description, subError: err}
}

// HResultToError may return an error from a HResult or nil, no error.
func HResultToError(hr uintptr, _ uintptr, _ error) (err error) {
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

// Code is the HResult.
func (v *OleError) Code() uintptr {
	return uintptr(v.hr)
}

// String description, either manually set or format message with error code.
func (v *OleError) String() string {
	if v.description != "" {
		return errstr(int(v.hr)) + " (" + v.description + ")"
	}
	return errstr(int(v.hr))
}

// Error implements error interface.
func (v *OleError) Error() string {
	return v.String()
}

// Description retrieves error summary, if there is one.
func (v *OleError) Description() string {
	return v.description
}

// SubError returns parent error, if there is one.
func (v *OleError) SubError() error {
	return v.subError
}
