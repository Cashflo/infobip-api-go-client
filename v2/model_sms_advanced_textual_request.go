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

// SmsAdvancedTextualRequest struct for SmsAdvancedTextualRequest
type SmsAdvancedTextualRequest struct {
	// The ID which uniquely identifies the request. Bulk ID will be received only when you send a message to more than one destination address.
	BulkId   *string              `json:"bulkId,omitempty"`
	Messages *[]SmsTextualMessage `json:"messages,omitempty"`
	// Limit the sending speed for message bulks. In some use cases, you might want to reduce message sending speed if your message call to action involves visiting a website, calling your contact center or similar recipient activity, in which you can handle a limited amount of load. This setting helps you to spread the delivery of the messages over a longer period, allowing your systems or agents to handle incoming traffic in real-time, resulting in better customer satisfaction.
	SendingSpeedLimit *SmsSendingSpeedLimit `json:"sendingSpeedLimit,omitempty"`
	Tracking          *SmsTracking          `json:"tracking,omitempty"`
}

// NewSmsAdvancedTextualRequest instantiates a new SmsAdvancedTextualRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsAdvancedTextualRequest() *SmsAdvancedTextualRequest {
	this := SmsAdvancedTextualRequest{}
	return &this
}

// NewSmsAdvancedTextualRequestWithDefaults instantiates a new SmsAdvancedTextualRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsAdvancedTextualRequestWithDefaults() *SmsAdvancedTextualRequest {
	this := SmsAdvancedTextualRequest{}
	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *SmsAdvancedTextualRequest) GetBulkId() string {
	if o == nil || o.BulkId == nil {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsAdvancedTextualRequest) GetBulkIdOk() (*string, bool) {
	if o == nil || o.BulkId == nil {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *SmsAdvancedTextualRequest) HasBulkId() bool {
	if o != nil && o.BulkId != nil {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *SmsAdvancedTextualRequest) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *SmsAdvancedTextualRequest) GetMessages() []SmsTextualMessage {
	if o == nil || o.Messages == nil {
		var ret []SmsTextualMessage
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsAdvancedTextualRequest) GetMessagesOk() (*[]SmsTextualMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *SmsAdvancedTextualRequest) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []SmsTextualMessage and assigns it to the Messages field.
func (o *SmsAdvancedTextualRequest) SetMessages(v []SmsTextualMessage) {
	o.Messages = &v
}

// GetSendingSpeedLimit returns the SendingSpeedLimit field value if set, zero value otherwise.
func (o *SmsAdvancedTextualRequest) GetSendingSpeedLimit() SmsSendingSpeedLimit {
	if o == nil || o.SendingSpeedLimit == nil {
		var ret SmsSendingSpeedLimit
		return ret
	}
	return *o.SendingSpeedLimit
}

// GetSendingSpeedLimitOk returns a tuple with the SendingSpeedLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsAdvancedTextualRequest) GetSendingSpeedLimitOk() (*SmsSendingSpeedLimit, bool) {
	if o == nil || o.SendingSpeedLimit == nil {
		return nil, false
	}
	return o.SendingSpeedLimit, true
}

// HasSendingSpeedLimit returns a boolean if a field has been set.
func (o *SmsAdvancedTextualRequest) HasSendingSpeedLimit() bool {
	if o != nil && o.SendingSpeedLimit != nil {
		return true
	}

	return false
}

// SetSendingSpeedLimit gets a reference to the given SmsSendingSpeedLimit and assigns it to the SendingSpeedLimit field.
func (o *SmsAdvancedTextualRequest) SetSendingSpeedLimit(v SmsSendingSpeedLimit) {
	o.SendingSpeedLimit = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *SmsAdvancedTextualRequest) GetTracking() SmsTracking {
	if o == nil || o.Tracking == nil {
		var ret SmsTracking
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsAdvancedTextualRequest) GetTrackingOk() (*SmsTracking, bool) {
	if o == nil || o.Tracking == nil {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *SmsAdvancedTextualRequest) HasTracking() bool {
	if o != nil && o.Tracking != nil {
		return true
	}

	return false
}

// SetTracking gets a reference to the given SmsTracking and assigns it to the Tracking field.
func (o *SmsAdvancedTextualRequest) SetTracking(v SmsTracking) {
	o.Tracking = &v
}

func (o SmsAdvancedTextualRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BulkId != nil {
		toSerialize["bulkId"] = o.BulkId
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.SendingSpeedLimit != nil {
		toSerialize["sendingSpeedLimit"] = o.SendingSpeedLimit
	}
	if o.Tracking != nil {
		toSerialize["tracking"] = o.Tracking
	}
	return json.Marshal(toSerialize)
}

type NullableSmsAdvancedTextualRequest struct {
	value *SmsAdvancedTextualRequest
	isSet bool
}

func (v NullableSmsAdvancedTextualRequest) Get() *SmsAdvancedTextualRequest {
	return v.value
}

func (v *NullableSmsAdvancedTextualRequest) Set(val *SmsAdvancedTextualRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsAdvancedTextualRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsAdvancedTextualRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsAdvancedTextualRequest(val *SmsAdvancedTextualRequest) *NullableSmsAdvancedTextualRequest {
	return &NullableSmsAdvancedTextualRequest{value: val, isSet: true}
}

func (v NullableSmsAdvancedTextualRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsAdvancedTextualRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
