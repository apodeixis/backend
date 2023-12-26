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

// checks if the EditUserAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditUserAllOfAttributes{}

// EditUserAllOfAttributes struct for EditUserAllOfAttributes
type EditUserAllOfAttributes struct {
	NewName *string `json:"new_name,omitempty"`
}

// NewEditUserAllOfAttributes instantiates a new EditUserAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditUserAllOfAttributes() *EditUserAllOfAttributes {
	this := EditUserAllOfAttributes{}
	return &this
}

// NewEditUserAllOfAttributesWithDefaults instantiates a new EditUserAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditUserAllOfAttributesWithDefaults() *EditUserAllOfAttributes {
	this := EditUserAllOfAttributes{}
	return &this
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *EditUserAllOfAttributes) GetNewName() string {
	if o == nil || IsNil(o.NewName) {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditUserAllOfAttributes) GetNewNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewName) {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *EditUserAllOfAttributes) HasNewName() bool {
	if o != nil && !IsNil(o.NewName) {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *EditUserAllOfAttributes) SetNewName(v string) {
	o.NewName = &v
}

func (o EditUserAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditUserAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewName) {
		toSerialize["new_name"] = o.NewName
	}
	return toSerialize, nil
}

type NullableEditUserAllOfAttributes struct {
	value *EditUserAllOfAttributes
	isSet bool
}

func (v NullableEditUserAllOfAttributes) Get() *EditUserAllOfAttributes {
	return v.value
}

func (v *NullableEditUserAllOfAttributes) Set(val *EditUserAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEditUserAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEditUserAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditUserAllOfAttributes(val *EditUserAllOfAttributes) *NullableEditUserAllOfAttributes {
	return &NullableEditUserAllOfAttributes{value: val, isSet: true}
}

func (v NullableEditUserAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditUserAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}