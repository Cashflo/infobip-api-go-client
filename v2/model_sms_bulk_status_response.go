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

// SmsBulkStatusResponse struct for SmsBulkStatusResponse
type SmsBulkStatusResponse struct {
	BulkId *string        `json:"bulkId,omitempty"`
	Status *SmsBulkStatus `json:"status,omitempty"`
}

// NewSmsBulkStatusResponse instantiates a new SmsBulkStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsBulkStatusResponse() *SmsBulkStatusResponse {
	this := SmsBulkStatusResponse{}
	return &this
}

// NewSmsBulkStatusResponseWithDefaults instantiates a new SmsBulkStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsBulkStatusResponseWithDefaults() *SmsBulkStatusResponse {
	this := SmsBulkStatusResponse{}
	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *SmsBulkStatusResponse) GetBulkId() string {
	if o == nil || o.BulkId == nil {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsBulkStatusResponse) GetBulkIdOk() (*string, bool) {
	if o == nil || o.BulkId == nil {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *SmsBulkStatusResponse) HasBulkId() bool {
	if o != nil && o.BulkId != nil {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *SmsBulkStatusResponse) SetBulkId(v string) {
	o.BulkId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SmsBulkStatusResponse) GetStatus() SmsBulkStatus {
	if o == nil || o.Status == nil {
		var ret SmsBulkStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsBulkStatusResponse) GetStatusOk() (*SmsBulkStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SmsBulkStatusResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SmsBulkStatus and assigns it to the Status field.
func (o *SmsBulkStatusResponse) SetStatus(v SmsBulkStatus) {
	o.Status = &v
}

func (o SmsBulkStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BulkId != nil {
		toSerialize["bulkId"] = o.BulkId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSmsBulkStatusResponse struct {
	value *SmsBulkStatusResponse
	isSet bool
}

func (v NullableSmsBulkStatusResponse) Get() *SmsBulkStatusResponse {
	return v.value
}

func (v *NullableSmsBulkStatusResponse) Set(val *SmsBulkStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsBulkStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsBulkStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsBulkStatusResponse(val *SmsBulkStatusResponse) *NullableSmsBulkStatusResponse {
	return &NullableSmsBulkStatusResponse{value: val, isSet: true}
}

func (v NullableSmsBulkStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsBulkStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
