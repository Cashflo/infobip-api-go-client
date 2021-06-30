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

// SmsResponseDetails struct for SmsResponseDetails
type SmsResponseDetails struct {
	// The ID that uniquely identifies the message sent.
	MessageId *string `json:"messageId,omitempty"`
	// Indicates whether the message is successfully sent, not sent, delivered, not delivered, waiting for delivery or any other possible status.
	Status *SmsStatus `json:"status,omitempty"`
	// The message destination address.
	To *string `json:"to,omitempty"`
}

// NewSmsResponseDetails instantiates a new SmsResponseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsResponseDetails() *SmsResponseDetails {
	this := SmsResponseDetails{}
	return &this
}

// NewSmsResponseDetailsWithDefaults instantiates a new SmsResponseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsResponseDetailsWithDefaults() *SmsResponseDetails {
	this := SmsResponseDetails{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *SmsResponseDetails) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsResponseDetails) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *SmsResponseDetails) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *SmsResponseDetails) SetMessageId(v string) {
	o.MessageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SmsResponseDetails) GetStatus() SmsStatus {
	if o == nil || o.Status == nil {
		var ret SmsStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsResponseDetails) GetStatusOk() (*SmsStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SmsResponseDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SmsStatus and assigns it to the Status field.
func (o *SmsResponseDetails) SetStatus(v SmsStatus) {
	o.Status = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SmsResponseDetails) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsResponseDetails) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SmsResponseDetails) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *SmsResponseDetails) SetTo(v string) {
	o.To = &v
}

func (o SmsResponseDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableSmsResponseDetails struct {
	value *SmsResponseDetails
	isSet bool
}

func (v NullableSmsResponseDetails) Get() *SmsResponseDetails {
	return v.value
}

func (v *NullableSmsResponseDetails) Set(val *SmsResponseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsResponseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsResponseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsResponseDetails(val *SmsResponseDetails) *NullableSmsResponseDetails {
	return &NullableSmsResponseDetails{value: val, isSet: true}
}

func (v NullableSmsResponseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsResponseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}