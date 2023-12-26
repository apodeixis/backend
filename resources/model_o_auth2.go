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

// checks if the OAuth2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2{}

// OAuth2 struct for OAuth2
type OAuth2 struct {
	Type       string                `json:"type"`
	Attributes OAuth2AllOfAttributes `json:"attributes"`
}

type _OAuth2 OAuth2

// NewOAuth2 instantiates a new OAuth2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2(type_ string, attributes OAuth2AllOfAttributes) *OAuth2 {
	this := OAuth2{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOAuth2WithDefaults instantiates a new OAuth2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2WithDefaults() *OAuth2 {
	this := OAuth2{}
	return &this
}

// GetType returns the Type field value
func (o *OAuth2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OAuth2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OAuth2) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OAuth2) GetAttributes() OAuth2AllOfAttributes {
	if o == nil {
		var ret OAuth2AllOfAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OAuth2) GetAttributesOk() (*OAuth2AllOfAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OAuth2) SetAttributes(v OAuth2AllOfAttributes) {
	o.Attributes = v
}

func (o OAuth2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *OAuth2) UnmarshalJSON(data []byte) (err error) {
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

	varOAuth2 := _OAuth2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2)

	if err != nil {
		return err
	}

	*o = OAuth2(varOAuth2)

	return err
}

type NullableOAuth2 struct {
	value *OAuth2
	isSet bool
}

func (v NullableOAuth2) Get() *OAuth2 {
	return v.value
}

func (v *NullableOAuth2) Set(val *OAuth2) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2(val *OAuth2) *NullableOAuth2 {
	return &NullableOAuth2{value: val, isSet: true}
}

func (v NullableOAuth2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}