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

// TfaMessage struct for TfaMessage
type TfaMessage struct {
	// 2FA application ID for which the requested message is created.
	ApplicationId *string `json:"applicationId,omitempty"`
	// Language code of language in which message text is written. It is used for reading the message when it is sent via voice. If no language is set, message will be read in `English`.
	Language *TfaLanguage `json:"language,omitempty"`
	// Message template ID.
	MessageId *string `json:"messageId,omitempty"`
	// Text of a message that will be sent. Message text must contain `pinPlaceholder`.
	MessageText *string `json:"messageText,omitempty"`
	// PIN code length.
	PinLength *int32 `json:"pinLength,omitempty"`
	// PIN code placeholder that will be replaced with generated PIN code.
	PinPlaceholder *string `json:"pinPlaceholder,omitempty"`
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

// NewTfaMessage instantiates a new TfaMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaMessage() *TfaMessage {
	this := TfaMessage{}
	return &this
}

// NewTfaMessageWithDefaults instantiates a new TfaMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaMessageWithDefaults() *TfaMessage {
	this := TfaMessage{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *TfaMessage) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *TfaMessage) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *TfaMessage) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TfaMessage) GetLanguage() TfaLanguage {
	if o == nil || o.Language == nil {
		var ret TfaLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetLanguageOk() (*TfaLanguage, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TfaMessage) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given TfaLanguage and assigns it to the Language field.
func (o *TfaMessage) SetLanguage(v TfaLanguage) {
	o.Language = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *TfaMessage) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *TfaMessage) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *TfaMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageText returns the MessageText field value if set, zero value otherwise.
func (o *TfaMessage) GetMessageText() string {
	if o == nil || o.MessageText == nil {
		var ret string
		return ret
	}
	return *o.MessageText
}

// GetMessageTextOk returns a tuple with the MessageText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetMessageTextOk() (*string, bool) {
	if o == nil || o.MessageText == nil {
		return nil, false
	}
	return o.MessageText, true
}

// HasMessageText returns a boolean if a field has been set.
func (o *TfaMessage) HasMessageText() bool {
	if o != nil && o.MessageText != nil {
		return true
	}

	return false
}

// SetMessageText gets a reference to the given string and assigns it to the MessageText field.
func (o *TfaMessage) SetMessageText(v string) {
	o.MessageText = &v
}

// GetPinLength returns the PinLength field value if set, zero value otherwise.
func (o *TfaMessage) GetPinLength() int32 {
	if o == nil || o.PinLength == nil {
		var ret int32
		return ret
	}
	return *o.PinLength
}

// GetPinLengthOk returns a tuple with the PinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetPinLengthOk() (*int32, bool) {
	if o == nil || o.PinLength == nil {
		return nil, false
	}
	return o.PinLength, true
}

// HasPinLength returns a boolean if a field has been set.
func (o *TfaMessage) HasPinLength() bool {
	if o != nil && o.PinLength != nil {
		return true
	}

	return false
}

// SetPinLength gets a reference to the given int32 and assigns it to the PinLength field.
func (o *TfaMessage) SetPinLength(v int32) {
	o.PinLength = &v
}

// GetPinPlaceholder returns the PinPlaceholder field value if set, zero value otherwise.
func (o *TfaMessage) GetPinPlaceholder() string {
	if o == nil || o.PinPlaceholder == nil {
		var ret string
		return ret
	}
	return *o.PinPlaceholder
}

// GetPinPlaceholderOk returns a tuple with the PinPlaceholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetPinPlaceholderOk() (*string, bool) {
	if o == nil || o.PinPlaceholder == nil {
		return nil, false
	}
	return o.PinPlaceholder, true
}

// HasPinPlaceholder returns a boolean if a field has been set.
func (o *TfaMessage) HasPinPlaceholder() bool {
	if o != nil && o.PinPlaceholder != nil {
		return true
	}

	return false
}

// SetPinPlaceholder gets a reference to the given string and assigns it to the PinPlaceholder field.
func (o *TfaMessage) SetPinPlaceholder(v string) {
	o.PinPlaceholder = &v
}

// GetPinType returns the PinType field value if set, zero value otherwise.
func (o *TfaMessage) GetPinType() TfaPinType {
	if o == nil || o.PinType == nil {
		var ret TfaPinType
		return ret
	}
	return *o.PinType
}

// GetPinTypeOk returns a tuple with the PinType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetPinTypeOk() (*TfaPinType, bool) {
	if o == nil || o.PinType == nil {
		return nil, false
	}
	return o.PinType, true
}

// HasPinType returns a boolean if a field has been set.
func (o *TfaMessage) HasPinType() bool {
	if o != nil && o.PinType != nil {
		return true
	}

	return false
}

// SetPinType gets a reference to the given TfaPinType and assigns it to the PinType field.
func (o *TfaMessage) SetPinType(v TfaPinType) {
	o.PinType = &v
}

// GetRegional returns the Regional field value if set, zero value otherwise.
func (o *TfaMessage) GetRegional() TfaRegionalOptions {
	if o == nil || o.Regional == nil {
		var ret TfaRegionalOptions
		return ret
	}
	return *o.Regional
}

// GetRegionalOk returns a tuple with the Regional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetRegionalOk() (*TfaRegionalOptions, bool) {
	if o == nil || o.Regional == nil {
		return nil, false
	}
	return o.Regional, true
}

// HasRegional returns a boolean if a field has been set.
func (o *TfaMessage) HasRegional() bool {
	if o != nil && o.Regional != nil {
		return true
	}

	return false
}

// SetRegional gets a reference to the given TfaRegionalOptions and assigns it to the Regional field.
func (o *TfaMessage) SetRegional(v TfaRegionalOptions) {
	o.Regional = &v
}

// GetRepeatDTMF returns the RepeatDTMF field value if set, zero value otherwise.
func (o *TfaMessage) GetRepeatDTMF() string {
	if o == nil || o.RepeatDTMF == nil {
		var ret string
		return ret
	}
	return *o.RepeatDTMF
}

// GetRepeatDTMFOk returns a tuple with the RepeatDTMF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetRepeatDTMFOk() (*string, bool) {
	if o == nil || o.RepeatDTMF == nil {
		return nil, false
	}
	return o.RepeatDTMF, true
}

// HasRepeatDTMF returns a boolean if a field has been set.
func (o *TfaMessage) HasRepeatDTMF() bool {
	if o != nil && o.RepeatDTMF != nil {
		return true
	}

	return false
}

// SetRepeatDTMF gets a reference to the given string and assigns it to the RepeatDTMF field.
func (o *TfaMessage) SetRepeatDTMF(v string) {
	o.RepeatDTMF = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *TfaMessage) GetSenderId() string {
	if o == nil || o.SenderId == nil {
		var ret string
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetSenderIdOk() (*string, bool) {
	if o == nil || o.SenderId == nil {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *TfaMessage) HasSenderId() bool {
	if o != nil && o.SenderId != nil {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given string and assigns it to the SenderId field.
func (o *TfaMessage) SetSenderId(v string) {
	o.SenderId = &v
}

// GetSpeechRate returns the SpeechRate field value if set, zero value otherwise.
func (o *TfaMessage) GetSpeechRate() float64 {
	if o == nil || o.SpeechRate == nil {
		var ret float64
		return ret
	}
	return *o.SpeechRate
}

// GetSpeechRateOk returns a tuple with the SpeechRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaMessage) GetSpeechRateOk() (*float64, bool) {
	if o == nil || o.SpeechRate == nil {
		return nil, false
	}
	return o.SpeechRate, true
}

// HasSpeechRate returns a boolean if a field has been set.
func (o *TfaMessage) HasSpeechRate() bool {
	if o != nil && o.SpeechRate != nil {
		return true
	}

	return false
}

// SetSpeechRate gets a reference to the given float64 and assigns it to the SpeechRate field.
func (o *TfaMessage) SetSpeechRate(v float64) {
	o.SpeechRate = &v
}

func (o TfaMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationId != nil {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.MessageText != nil {
		toSerialize["messageText"] = o.MessageText
	}
	if o.PinLength != nil {
		toSerialize["pinLength"] = o.PinLength
	}
	if o.PinPlaceholder != nil {
		toSerialize["pinPlaceholder"] = o.PinPlaceholder
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

type NullableTfaMessage struct {
	value *TfaMessage
	isSet bool
}

func (v NullableTfaMessage) Get() *TfaMessage {
	return v.value
}

func (v *NullableTfaMessage) Set(val *TfaMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaMessage(val *TfaMessage) *NullableTfaMessage {
	return &NullableTfaMessage{value: val, isSet: true}
}

func (v NullableTfaMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
