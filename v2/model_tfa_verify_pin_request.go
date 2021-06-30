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

// TfaVerifyPinRequest struct for TfaVerifyPinRequest
type TfaVerifyPinRequest struct {
	// PIN code to verify
	Pin string `json:"pin"`
}

// NewTfaVerifyPinRequest instantiates a new TfaVerifyPinRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaVerifyPinRequest(pin string) *TfaVerifyPinRequest {
	this := TfaVerifyPinRequest{}
	this.Pin = pin
	return &this
}

// NewTfaVerifyPinRequestWithDefaults instantiates a new TfaVerifyPinRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaVerifyPinRequestWithDefaults() *TfaVerifyPinRequest {
	this := TfaVerifyPinRequest{}
	return &this
}

// GetPin returns the Pin field value
func (o *TfaVerifyPinRequest) GetPin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pin
}

// GetPinOk returns a tuple with the Pin field value
// and a boolean to check if the value has been set.
func (o *TfaVerifyPinRequest) GetPinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pin, true
}

// SetPin sets field value
func (o *TfaVerifyPinRequest) SetPin(v string) {
	o.Pin = v
}

func (o TfaVerifyPinRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableTfaVerifyPinRequest struct {
	value *TfaVerifyPinRequest
	isSet bool
}

func (v NullableTfaVerifyPinRequest) Get() *TfaVerifyPinRequest {
	return v.value
}

func (v *NullableTfaVerifyPinRequest) Set(val *TfaVerifyPinRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaVerifyPinRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaVerifyPinRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaVerifyPinRequest(val *TfaVerifyPinRequest) *NullableTfaVerifyPinRequest {
	return &NullableTfaVerifyPinRequest{value: val, isSet: true}
}

func (v NullableTfaVerifyPinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaVerifyPinRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
