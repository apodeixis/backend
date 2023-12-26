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

// checks if the AuthTokenKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenKey{}

// AuthTokenKey struct for AuthTokenKey
type AuthTokenKey struct {
	Type string `json:"type"`
}

type _AuthTokenKey AuthTokenKey

// NewAuthTokenKey instantiates a new AuthTokenKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenKey(type_ string) *AuthTokenKey {
	this := AuthTokenKey{}
	this.Type = type_
	return &this
}

// NewAuthTokenKeyWithDefaults instantiates a new AuthTokenKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenKeyWithDefaults() *AuthTokenKey {
	this := AuthTokenKey{}
	return &this
}

// GetType returns the Type field value
func (o *AuthTokenKey) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthTokenKey) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthTokenKey) SetType(v string) {
	o.Type = v
}

func (o AuthTokenKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AuthTokenKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varAuthTokenKey := _AuthTokenKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthTokenKey)

	if err != nil {
		return err
	}

	*o = AuthTokenKey(varAuthTokenKey)

	return err
}

type NullableAuthTokenKey struct {
	value *AuthTokenKey
	isSet bool
}

func (v NullableAuthTokenKey) Get() *AuthTokenKey {
	return v.value
}

func (v *NullableAuthTokenKey) Set(val *AuthTokenKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenKey(val *AuthTokenKey) *NullableAuthTokenKey {
	return &NullableAuthTokenKey{value: val, isSet: true}
}

func (v NullableAuthTokenKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
