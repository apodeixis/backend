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

// checks if the OAuth2SignInGoogle200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2SignInGoogle200Response{}

// OAuth2SignInGoogle200Response struct for OAuth2SignInGoogle200Response
type OAuth2SignInGoogle200Response struct {
	Data     User                            `json:"data"`
	Included []Login200ResponseIncludedInner `json:"included"`
}

type _OAuth2SignInGoogle200Response OAuth2SignInGoogle200Response

// NewOAuth2SignInGoogle200Response instantiates a new OAuth2SignInGoogle200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2SignInGoogle200Response(data User, included []Login200ResponseIncludedInner) *OAuth2SignInGoogle200Response {
	this := OAuth2SignInGoogle200Response{}
	this.Data = data
	this.Included = included
	return &this
}

// NewOAuth2SignInGoogle200ResponseWithDefaults instantiates a new OAuth2SignInGoogle200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2SignInGoogle200ResponseWithDefaults() *OAuth2SignInGoogle200Response {
	this := OAuth2SignInGoogle200Response{}
	return &this
}

// GetData returns the Data field value
func (o *OAuth2SignInGoogle200Response) GetData() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OAuth2SignInGoogle200Response) GetDataOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OAuth2SignInGoogle200Response) SetData(v User) {
	o.Data = v
}

// GetIncluded returns the Included field value
func (o *OAuth2SignInGoogle200Response) GetIncluded() []Login200ResponseIncludedInner {
	if o == nil {
		var ret []Login200ResponseIncludedInner
		return ret
	}

	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value
// and a boolean to check if the value has been set.
func (o *OAuth2SignInGoogle200Response) GetIncludedOk() ([]Login200ResponseIncludedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Included, true
}

// SetIncluded sets field value
func (o *OAuth2SignInGoogle200Response) SetIncluded(v []Login200ResponseIncludedInner) {
	o.Included = v
}

func (o OAuth2SignInGoogle200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2SignInGoogle200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["included"] = o.Included
	return toSerialize, nil
}

func (o *OAuth2SignInGoogle200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"included",
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

	varOAuth2SignInGoogle200Response := _OAuth2SignInGoogle200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2SignInGoogle200Response)

	if err != nil {
		return err
	}

	*o = OAuth2SignInGoogle200Response(varOAuth2SignInGoogle200Response)

	return err
}

type NullableOAuth2SignInGoogle200Response struct {
	value *OAuth2SignInGoogle200Response
	isSet bool
}

func (v NullableOAuth2SignInGoogle200Response) Get() *OAuth2SignInGoogle200Response {
	return v.value
}

func (v *NullableOAuth2SignInGoogle200Response) Set(val *OAuth2SignInGoogle200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2SignInGoogle200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2SignInGoogle200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2SignInGoogle200Response(val *OAuth2SignInGoogle200Response) *NullableOAuth2SignInGoogle200Response {
	return &NullableOAuth2SignInGoogle200Response{value: val, isSet: true}
}

func (v NullableOAuth2SignInGoogle200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2SignInGoogle200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
