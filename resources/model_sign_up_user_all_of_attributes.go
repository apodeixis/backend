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

// checks if the SignUpUserAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignUpUserAllOfAttributes{}

// SignUpUserAllOfAttributes struct for SignUpUserAllOfAttributes
type SignUpUserAllOfAttributes struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type _SignUpUserAllOfAttributes SignUpUserAllOfAttributes

// NewSignUpUserAllOfAttributes instantiates a new SignUpUserAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignUpUserAllOfAttributes(name string, email string, password string) *SignUpUserAllOfAttributes {
	this := SignUpUserAllOfAttributes{}
	this.Name = name
	this.Email = email
	this.Password = password
	return &this
}

// NewSignUpUserAllOfAttributesWithDefaults instantiates a new SignUpUserAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignUpUserAllOfAttributesWithDefaults() *SignUpUserAllOfAttributes {
	this := SignUpUserAllOfAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *SignUpUserAllOfAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignUpUserAllOfAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignUpUserAllOfAttributes) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *SignUpUserAllOfAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SignUpUserAllOfAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SignUpUserAllOfAttributes) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *SignUpUserAllOfAttributes) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SignUpUserAllOfAttributes) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SignUpUserAllOfAttributes) SetPassword(v string) {
	o.Password = v
}

func (o SignUpUserAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignUpUserAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *SignUpUserAllOfAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"email",
		"password",
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

	varSignUpUserAllOfAttributes := _SignUpUserAllOfAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignUpUserAllOfAttributes)

	if err != nil {
		return err
	}

	*o = SignUpUserAllOfAttributes(varSignUpUserAllOfAttributes)

	return err
}

type NullableSignUpUserAllOfAttributes struct {
	value *SignUpUserAllOfAttributes
	isSet bool
}

func (v NullableSignUpUserAllOfAttributes) Get() *SignUpUserAllOfAttributes {
	return v.value
}

func (v *NullableSignUpUserAllOfAttributes) Set(val *SignUpUserAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSignUpUserAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSignUpUserAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignUpUserAllOfAttributes(val *SignUpUserAllOfAttributes) *NullableSignUpUserAllOfAttributes {
	return &NullableSignUpUserAllOfAttributes{value: val, isSet: true}
}

func (v NullableSignUpUserAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignUpUserAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}