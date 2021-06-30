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

// TfaApplicationRequest struct for TfaApplicationRequest
type TfaApplicationRequest struct {
	// Created 2FA application configuration.
	Configuration *TfaApplicationConfiguration `json:"configuration,omitempty"`
	// Indicates if the created application is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// 2FA application name.
	Name string `json:"name"`
}

// NewTfaApplicationRequest instantiates a new TfaApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaApplicationRequest(name string) *TfaApplicationRequest {
	this := TfaApplicationRequest{}
	this.Name = name
	return &this
}

// NewTfaApplicationRequestWithDefaults instantiates a new TfaApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaApplicationRequestWithDefaults() *TfaApplicationRequest {
	this := TfaApplicationRequest{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *TfaApplicationRequest) GetConfiguration() TfaApplicationConfiguration {
	if o == nil || o.Configuration == nil {
		var ret TfaApplicationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaApplicationRequest) GetConfigurationOk() (*TfaApplicationConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *TfaApplicationRequest) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given TfaApplicationConfiguration and assigns it to the Configuration field.
func (o *TfaApplicationRequest) SetConfiguration(v TfaApplicationConfiguration) {
	o.Configuration = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TfaApplicationRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaApplicationRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TfaApplicationRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TfaApplicationRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value
func (o *TfaApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TfaApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TfaApplicationRequest) SetName(v string) {
	o.Name = v
}

func (o TfaApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTfaApplicationRequest struct {
	value *TfaApplicationRequest
	isSet bool
}

func (v NullableTfaApplicationRequest) Get() *TfaApplicationRequest {
	return v.value
}

func (v *NullableTfaApplicationRequest) Set(val *TfaApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaApplicationRequest(val *TfaApplicationRequest) *NullableTfaApplicationRequest {
	return &NullableTfaApplicationRequest{value: val, isSet: true}
}

func (v NullableTfaApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
