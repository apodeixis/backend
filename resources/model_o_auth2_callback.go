/*
Apodeixis backend

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OAuth2Callback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Callback{}

// OAuth2Callback struct for OAuth2Callback
type OAuth2Callback struct {
	Type       string                        `json:"type"`
	Attributes OAuth2CallbackAllOfAttributes `json:"attributes"`
}

type _OAuth2Callback OAuth2Callback

// NewOAuth2Callback instantiates a new OAuth2Callback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Callback(type_ string, attributes OAuth2CallbackAllOfAttributes) *OAuth2Callback {
	this := OAuth2Callback{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOAuth2CallbackWithDefaults instantiates a new OAuth2Callback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2CallbackWithDefaults() *OAuth2Callback {
	this := OAuth2Callback{}
	return &this
}

// GetType returns the Type field value
func (o *OAuth2Callback) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OAuth2Callback) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OAuth2Callback) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OAuth2Callback) GetAttributes() OAuth2CallbackAllOfAttributes {
	if o == nil {
		var ret OAuth2CallbackAllOfAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OAuth2Callback) GetAttributesOk() (*OAuth2CallbackAllOfAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OAuth2Callback) SetAttributes(v OAuth2CallbackAllOfAttributes) {
	o.Attributes = v
}

func (o OAuth2Callback) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Callback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *OAuth2Callback) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOAuth2Callback := _OAuth2Callback{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2Callback)

	if err != nil {
		return err
	}

	*o = OAuth2Callback(varOAuth2Callback)

	return err
}

type NullableOAuth2Callback struct {
	value *OAuth2Callback
	isSet bool
}

func (v NullableOAuth2Callback) Get() *OAuth2Callback {
	return v.value
}

func (v *NullableOAuth2Callback) Set(val *OAuth2Callback) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Callback) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Callback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Callback(val *OAuth2Callback) *NullableOAuth2Callback {
	return &NullableOAuth2Callback{value: val, isSet: true}
}

func (v NullableOAuth2Callback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Callback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}