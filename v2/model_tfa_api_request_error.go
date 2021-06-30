/*
 * Infobip Client API Libraries OpenAPI Specification
 *
 * OpenAPI specification containing public endpoints supported in client API libraries.
 *
 * API version: 1.0.157
 * Contact: support@infobip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infobip

import (
	"encoding/json"
)

// TfaApiRequestError struct for TfaApiRequestError
type TfaApiRequestError struct {
	ServiceException *TfaApiRequestErrorDetails `json:"serviceException,omitempty"`
}

// NewTfaApiRequestError instantiates a new TfaApiRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaApiRequestError() *TfaApiRequestError {
	this := TfaApiRequestError{}
	return &this
}

// NewTfaApiRequestErrorWithDefaults instantiates a new TfaApiRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaApiRequestErrorWithDefaults() *TfaApiRequestError {
	this := TfaApiRequestError{}
	return &this
}

// GetServiceException returns the ServiceException field value if set, zero value otherwise.
func (o *TfaApiRequestError) GetServiceException() TfaApiRequestErrorDetails {
	if o == nil || o.ServiceException == nil {
		var ret TfaApiRequestErrorDetails
		return ret
	}
	return *o.ServiceException
}

// GetServiceExceptionOk returns a tuple with the ServiceException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaApiRequestError) GetServiceExceptionOk() (*TfaApiRequestErrorDetails, bool) {
	if o == nil || o.ServiceException == nil {
		return nil, false
	}
	return o.ServiceException, true
}

// HasServiceException returns a boolean if a field has been set.
func (o *TfaApiRequestError) HasServiceException() bool {
	if o != nil && o.ServiceException != nil {
		return true
	}

	return false
}

// SetServiceException gets a reference to the given TfaApiRequestErrorDetails and assigns it to the ServiceException field.
func (o *TfaApiRequestError) SetServiceException(v TfaApiRequestErrorDetails) {
	o.ServiceException = &v
}

func (o TfaApiRequestError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceException != nil {
		toSerialize["serviceException"] = o.ServiceException
	}
	return json.Marshal(toSerialize)
}

type NullableTfaApiRequestError struct {
	value *TfaApiRequestError
	isSet bool
}

func (v NullableTfaApiRequestError) Get() *TfaApiRequestError {
	return v.value
}

func (v *NullableTfaApiRequestError) Set(val *TfaApiRequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaApiRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaApiRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaApiRequestError(val *TfaApiRequestError) *NullableTfaApiRequestError {
	return &NullableTfaApiRequestError{value: val, isSet: true}
}

func (v NullableTfaApiRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaApiRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
