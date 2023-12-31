/*
Apodeixis backend

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the UserAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAllOfRelationships{}

// UserAllOfRelationships struct for UserAllOfRelationships
type UserAllOfRelationships struct {
	Tokens *UserAllOfRelationshipsTokens `json:"tokens,omitempty"`
}

// NewUserAllOfRelationships instantiates a new UserAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAllOfRelationships() *UserAllOfRelationships {
	this := UserAllOfRelationships{}
	return &this
}

// NewUserAllOfRelationshipsWithDefaults instantiates a new UserAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAllOfRelationshipsWithDefaults() *UserAllOfRelationships {
	this := UserAllOfRelationships{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *UserAllOfRelationships) GetTokens() UserAllOfRelationshipsTokens {
	if o == nil || IsNil(o.Tokens) {
		var ret UserAllOfRelationshipsTokens
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAllOfRelationships) GetTokensOk() (*UserAllOfRelationshipsTokens, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *UserAllOfRelationships) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given UserAllOfRelationshipsTokens and assigns it to the Tokens field.
func (o *UserAllOfRelationships) SetTokens(v UserAllOfRelationshipsTokens) {
	o.Tokens = &v
}

func (o UserAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	return toSerialize, nil
}

type NullableUserAllOfRelationships struct {
	value *UserAllOfRelationships
	isSet bool
}

func (v NullableUserAllOfRelationships) Get() *UserAllOfRelationships {
	return v.value
}

func (v *NullableUserAllOfRelationships) Set(val *UserAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAllOfRelationships(val *UserAllOfRelationships) *NullableUserAllOfRelationships {
	return &NullableUserAllOfRelationships{value: val, isSet: true}
}

func (v NullableUserAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
