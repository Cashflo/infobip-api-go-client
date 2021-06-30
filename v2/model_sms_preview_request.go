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

// SmsPreviewRequest struct for SmsPreviewRequest
type SmsPreviewRequest struct {
	// Code for language character set of a message text. Possible values: `TR` for Turkish, `ES` for Spanish, `PT` for Portuguese and `AUTODETECT` to let platform pick character set automatically based on the message text.
	LanguageCode *string `json:"languageCode,omitempty"`
	// Message text to preview.
	Text string `json:"text"`
	// Conversion of a message text from one script to another. Possible values: `TURKISH`, `GREEK`, `CYRILLIC`, `SERBIAN_CYRILLIC`, `CENTRAL_EUROPEAN`, `BALTIC` and `NON_UNICODE`.
	Transliteration *string `json:"transliteration,omitempty"`
}

// NewSmsPreviewRequest instantiates a new SmsPreviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsPreviewRequest(text string) *SmsPreviewRequest {
	this := SmsPreviewRequest{}
	this.Text = text
	return &this
}

// NewSmsPreviewRequestWithDefaults instantiates a new SmsPreviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsPreviewRequestWithDefaults() *SmsPreviewRequest {
	this := SmsPreviewRequest{}
	return &this
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *SmsPreviewRequest) GetLanguageCode() string {
	if o == nil || o.LanguageCode == nil {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsPreviewRequest) GetLanguageCodeOk() (*string, bool) {
	if o == nil || o.LanguageCode == nil {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *SmsPreviewRequest) HasLanguageCode() bool {
	if o != nil && o.LanguageCode != nil {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *SmsPreviewRequest) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

// GetText returns the Text field value
func (o *SmsPreviewRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *SmsPreviewRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *SmsPreviewRequest) SetText(v string) {
	o.Text = v
}

// GetTransliteration returns the Transliteration field value if set, zero value otherwise.
func (o *SmsPreviewRequest) GetTransliteration() string {
	if o == nil || o.Transliteration == nil {
		var ret string
		return ret
	}
	return *o.Transliteration
}

// GetTransliterationOk returns a tuple with the Transliteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsPreviewRequest) GetTransliterationOk() (*string, bool) {
	if o == nil || o.Transliteration == nil {
		return nil, false
	}
	return o.Transliteration, true
}

// HasTransliteration returns a boolean if a field has been set.
func (o *SmsPreviewRequest) HasTransliteration() bool {
	if o != nil && o.Transliteration != nil {
		return true
	}

	return false
}

// SetTransliteration gets a reference to the given string and assigns it to the Transliteration field.
func (o *SmsPreviewRequest) SetTransliteration(v string) {
	o.Transliteration = &v
}

func (o SmsPreviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LanguageCode != nil {
		toSerialize["languageCode"] = o.LanguageCode
	}
	if true {
		toSerialize["text"] = o.Text
	}
	if o.Transliteration != nil {
		toSerialize["transliteration"] = o.Transliteration
	}
	return json.Marshal(toSerialize)
}

type NullableSmsPreviewRequest struct {
	value *SmsPreviewRequest
	isSet bool
}

func (v NullableSmsPreviewRequest) Get() *SmsPreviewRequest {
	return v.value
}

func (v *NullableSmsPreviewRequest) Set(val *SmsPreviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsPreviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsPreviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsPreviewRequest(val *SmsPreviewRequest) *NullableSmsPreviewRequest {
	return &NullableSmsPreviewRequest{value: val, isSet: true}
}

func (v NullableSmsPreviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsPreviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
