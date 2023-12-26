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

// checks if the OAuth2SignInGoogleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2SignInGoogleRequest{}

// OAuth2SignInGoogleRequest struct for OAuth2SignInGoogleRequest
type OAuth2SignInGoogleRequest struct {
	Data OAuth2Callback `json:"data"`
}

type _OAuth2SignInGoogleRequest OAuth2SignInGoogleRequest

// NewOAuth2SignInGoogleRequest instantiates a new OAuth2SignInGoogleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2SignInGoogleRequest(data OAuth2Callback) *OAuth2SignInGoogleRequest {
	this := OAuth2SignInGoogleRequest{}
	this.Data = data
	return &this
}

// NewOAuth2SignInGoogleRequestWithDefaults instantiates a new OAuth2SignInGoogleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2SignInGoogleRequestWithDefaults() *OAuth2SignInGoogleRequest {
	this := OAuth2SignInGoogleRequest{}
	return &this
}

// GetData returns the Data field value
func (o *OAuth2SignInGoogleRequest) GetData() OAuth2Callback {
	if o == nil {
		var ret OAuth2Callback
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OAuth2SignInGoogleRequest) GetDataOk() (*OAuth2Callback, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OAuth2SignInGoogleRequest) SetData(v OAuth2Callback) {
	o.Data = v
}

func (o OAuth2SignInGoogleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2SignInGoogleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *OAuth2SignInGoogleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varOAuth2SignInGoogleRequest := _OAuth2SignInGoogleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2SignInGoogleRequest)

	if err != nil {
		return err
	}

	*o = OAuth2SignInGoogleRequest(varOAuth2SignInGoogleRequest)

	return err
}

type NullableOAuth2SignInGoogleRequest struct {
	value *OAuth2SignInGoogleRequest
	isSet bool
}

func (v NullableOAuth2SignInGoogleRequest) Get() *OAuth2SignInGoogleRequest {
	return v.value
}

func (v *NullableOAuth2SignInGoogleRequest) Set(val *OAuth2SignInGoogleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2SignInGoogleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2SignInGoogleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2SignInGoogleRequest(val *OAuth2SignInGoogleRequest) *NullableOAuth2SignInGoogleRequest {
	return &NullableOAuth2SignInGoogleRequest{value: val, isSet: true}
}

func (v NullableOAuth2SignInGoogleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2SignInGoogleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
