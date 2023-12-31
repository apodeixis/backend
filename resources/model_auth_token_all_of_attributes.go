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

// checks if the AuthTokenAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenAllOfAttributes{}

// AuthTokenAllOfAttributes struct for AuthTokenAllOfAttributes
type AuthTokenAllOfAttributes struct {
	Access    string `json:"access"`
	ExpiresAt int64  `json:"expires_at"`
}

type _AuthTokenAllOfAttributes AuthTokenAllOfAttributes

// NewAuthTokenAllOfAttributes instantiates a new AuthTokenAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenAllOfAttributes(access string, expiresAt int64) *AuthTokenAllOfAttributes {
	this := AuthTokenAllOfAttributes{}
	this.Access = access
	this.ExpiresAt = expiresAt
	return &this
}

// NewAuthTokenAllOfAttributesWithDefaults instantiates a new AuthTokenAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenAllOfAttributesWithDefaults() *AuthTokenAllOfAttributes {
	this := AuthTokenAllOfAttributes{}
	return &this
}

// GetAccess returns the Access field value
func (o *AuthTokenAllOfAttributes) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *AuthTokenAllOfAttributes) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *AuthTokenAllOfAttributes) SetAccess(v string) {
	o.Access = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *AuthTokenAllOfAttributes) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *AuthTokenAllOfAttributes) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *AuthTokenAllOfAttributes) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

func (o AuthTokenAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access"] = o.Access
	toSerialize["expires_at"] = o.ExpiresAt
	return toSerialize, nil
}

func (o *AuthTokenAllOfAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access",
		"expires_at",
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

	varAuthTokenAllOfAttributes := _AuthTokenAllOfAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthTokenAllOfAttributes)

	if err != nil {
		return err
	}

	*o = AuthTokenAllOfAttributes(varAuthTokenAllOfAttributes)

	return err
}

type NullableAuthTokenAllOfAttributes struct {
	value *AuthTokenAllOfAttributes
	isSet bool
}

func (v NullableAuthTokenAllOfAttributes) Get() *AuthTokenAllOfAttributes {
	return v.value
}

func (v *NullableAuthTokenAllOfAttributes) Set(val *AuthTokenAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenAllOfAttributes(val *AuthTokenAllOfAttributes) *NullableAuthTokenAllOfAttributes {
	return &NullableAuthTokenAllOfAttributes{value: val, isSet: true}
}

func (v NullableAuthTokenAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
