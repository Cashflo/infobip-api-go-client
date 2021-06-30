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

// TfaUpdateMessageRequest struct for TfaUpdateMessageRequest
type TfaUpdateMessageRequest struct {
	// Language code of language in which message text is written. It is used for reading the message when it is sent via voice. If no language is set, message will be read in `English`.
	Language *TfaLanguage `json:"language,omitempty"`
	// Text of a message that will be sent. Message text must contain `pinPlaceholder`.
	MessageText *string `json:"messageText,omitempty"`
	// PIN code length.
	PinLength *int32 `json:"pinLength,omitempty"`
	// Type of PIN code that will be generated and sent as part of 2FA message. You can set PIN type to numeric, alpha, alphanumeric or hex.
	PinType *TfaPinType `json:"pinType,omitempty"`
	// Region specific parameters, often specified by local laws. Use this if country or region that you are sending SMS to requires some extra parameters.
	Regional *TfaRegionalOptions `json:"regional,omitempty"`
	// In case PIN message is sent by Voice, DTMF code will enable replaying the message.
	RepeatDTMF *string `json:"repeatDTMF,omitempty"`
	// The name that will appear as the sender of the 2FA message (Example: CompanyName).
	SenderId *string `json:"senderId,omitempty"`
	// In case PIN message is sent by Voice, the speed of speech can be set for the message. Supported range is from `0.5` to `2`.
	SpeechRate *float64 `json:"speechRate,omitempty"`
}

// NewTfaUpdateMessageRequest instantiates a new TfaUpdateMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaUpdateMessageRequest() *TfaUpdateMessageRequest {
	this := TfaUpdateMessageRequest{}
	return &this
}

// NewTfaUpdateMessageRequestWithDefaults instantiates a new TfaUpdateMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaUpdateMessageRequestWithDefaults() *TfaUpdateMessageRequest {
	this := TfaUpdateMessageRequest{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetLanguage() TfaLanguage {
	if o == nil || o.Language == nil {
		var ret TfaLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetLanguageOk() (*TfaLanguage, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given TfaLanguage and assigns it to the Language field.
func (o *TfaUpdateMessageRequest) SetLanguage(v TfaLanguage) {
	o.Language = &v
}

// GetMessageText returns the MessageText field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetMessageText() string {
	if o == nil || o.MessageText == nil {
		var ret string
		return ret
	}
	return *o.MessageText
}

// GetMessageTextOk returns a tuple with the MessageText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetMessageTextOk() (*string, bool) {
	if o == nil || o.MessageText == nil {
		return nil, false
	}
	return o.MessageText, true
}

// HasMessageText returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasMessageText() bool {
	if o != nil && o.MessageText != nil {
		return true
	}

	return false
}

// SetMessageText gets a reference to the given string and assigns it to the MessageText field.
func (o *TfaUpdateMessageRequest) SetMessageText(v string) {
	o.MessageText = &v
}

// GetPinLength returns the PinLength field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetPinLength() int32 {
	if o == nil || o.PinLength == nil {
		var ret int32
		return ret
	}
	return *o.PinLength
}

// GetPinLengthOk returns a tuple with the PinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetPinLengthOk() (*int32, bool) {
	if o == nil || o.PinLength == nil {
		return nil, false
	}
	return o.PinLength, true
}

// HasPinLength returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasPinLength() bool {
	if o != nil && o.PinLength != nil {
		return true
	}

	return false
}

// SetPinLength gets a reference to the given int32 and assigns it to the PinLength field.
func (o *TfaUpdateMessageRequest) SetPinLength(v int32) {
	o.PinLength = &v
}

// GetPinType returns the PinType field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetPinType() TfaPinType {
	if o == nil || o.PinType == nil {
		var ret TfaPinType
		return ret
	}
	return *o.PinType
}

// GetPinTypeOk returns a tuple with the PinType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetPinTypeOk() (*TfaPinType, bool) {
	if o == nil || o.PinType == nil {
		return nil, false
	}
	return o.PinType, true
}

// HasPinType returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasPinType() bool {
	if o != nil && o.PinType != nil {
		return true
	}

	return false
}

// SetPinType gets a reference to the given TfaPinType and assigns it to the PinType field.
func (o *TfaUpdateMessageRequest) SetPinType(v TfaPinType) {
	o.PinType = &v
}

// GetRegional returns the Regional field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetRegional() TfaRegionalOptions {
	if o == nil || o.Regional == nil {
		var ret TfaRegionalOptions
		return ret
	}
	return *o.Regional
}

// GetRegionalOk returns a tuple with the Regional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetRegionalOk() (*TfaRegionalOptions, bool) {
	if o == nil || o.Regional == nil {
		return nil, false
	}
	return o.Regional, true
}

// HasRegional returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasRegional() bool {
	if o != nil && o.Regional != nil {
		return true
	}

	return false
}

// SetRegional gets a reference to the given TfaRegionalOptions and assigns it to the Regional field.
func (o *TfaUpdateMessageRequest) SetRegional(v TfaRegionalOptions) {
	o.Regional = &v
}

// GetRepeatDTMF returns the RepeatDTMF field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetRepeatDTMF() string {
	if o == nil || o.RepeatDTMF == nil {
		var ret string
		return ret
	}
	return *o.RepeatDTMF
}

// GetRepeatDTMFOk returns a tuple with the RepeatDTMF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetRepeatDTMFOk() (*string, bool) {
	if o == nil || o.RepeatDTMF == nil {
		return nil, false
	}
	return o.RepeatDTMF, true
}

// HasRepeatDTMF returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasRepeatDTMF() bool {
	if o != nil && o.RepeatDTMF != nil {
		return true
	}

	return false
}

// SetRepeatDTMF gets a reference to the given string and assigns it to the RepeatDTMF field.
func (o *TfaUpdateMessageRequest) SetRepeatDTMF(v string) {
	o.RepeatDTMF = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetSenderId() string {
	if o == nil || o.SenderId == nil {
		var ret string
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetSenderIdOk() (*string, bool) {
	if o == nil || o.SenderId == nil {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasSenderId() bool {
	if o != nil && o.SenderId != nil {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given string and assigns it to the SenderId field.
func (o *TfaUpdateMessageRequest) SetSenderId(v string) {
	o.SenderId = &v
}

// GetSpeechRate returns the SpeechRate field value if set, zero value otherwise.
func (o *TfaUpdateMessageRequest) GetSpeechRate() float64 {
	if o == nil || o.SpeechRate == nil {
		var ret float64
		return ret
	}
	return *o.SpeechRate
}

// GetSpeechRateOk returns a tuple with the SpeechRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaUpdateMessageRequest) GetSpeechRateOk() (*float64, bool) {
	if o == nil || o.SpeechRate == nil {
		return nil, false
	}
	return o.SpeechRate, true
}

// HasSpeechRate returns a boolean if a field has been set.
func (o *TfaUpdateMessageRequest) HasSpeechRate() bool {
	if o != nil && o.SpeechRate != nil {
		return true
	}

	return false
}

// SetSpeechRate gets a reference to the given float64 and assigns it to the SpeechRate field.
func (o *TfaUpdateMessageRequest) SetSpeechRate(v float64) {
	o.SpeechRate = &v
}

func (o TfaUpdateMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.MessageText != nil {
		toSerialize["messageText"] = o.MessageText
	}
	if o.PinLength != nil {
		toSerialize["pinLength"] = o.PinLength
	}
	if o.PinType != nil {
		toSerialize["pinType"] = o.PinType
	}
	if o.Regional != nil {
		toSerialize["regional"] = o.Regional
	}
	if o.RepeatDTMF != nil {
		toSerialize["repeatDTMF"] = o.RepeatDTMF
	}
	if o.SenderId != nil {
		toSerialize["senderId"] = o.SenderId
	}
	if o.SpeechRate != nil {
		toSerialize["speechRate"] = o.SpeechRate
	}
	return json.Marshal(toSerialize)
}

type NullableTfaUpdateMessageRequest struct {
	value *TfaUpdateMessageRequest
	isSet bool
}

func (v NullableTfaUpdateMessageRequest) Get() *TfaUpdateMessageRequest {
	return v.value
}

func (v *NullableTfaUpdateMessageRequest) Set(val *TfaUpdateMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaUpdateMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaUpdateMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaUpdateMessageRequest(val *TfaUpdateMessageRequest) *NullableTfaUpdateMessageRequest {
	return &NullableTfaUpdateMessageRequest{value: val, isSet: true}
}

func (v NullableTfaUpdateMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaUpdateMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
