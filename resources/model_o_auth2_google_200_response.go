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

// checks if the OAuth2Google200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Google200Response{}

// OAuth2Google200Response struct for OAuth2Google200Response
type OAuth2Google200Response struct {
	Data OAuth2 `json:"data"`
}

type _OAuth2Google200Response OAuth2Google200Response

// NewOAuth2Google200Response instantiates a new OAuth2Google200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Google200Response(data OAuth2) *OAuth2Google200Response {
	this := OAuth2Google200Response{}
	this.Data = data
	return &this
}

// NewOAuth2Google200ResponseWithDefaults instantiates a new OAuth2Google200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2Google200ResponseWithDefaults() *OAuth2Google200Response {
	this := OAuth2Google200Response{}
	return &this
}

// GetData returns the Data field value
func (o *OAuth2Google200Response) GetData() OAuth2 {
	if o == nil {
		var ret OAuth2
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OAuth2Google200Response) GetDataOk() (*OAuth2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OAuth2Google200Response) SetData(v OAuth2) {
	o.Data = v
}

func (o OAuth2Google200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Google200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *OAuth2Google200Response) UnmarshalJSON(data []byte) (err error) {
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

	varOAuth2Google200Response := _OAuth2Google200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2Google200Response)

	if err != nil {
		return err
	}

	*o = OAuth2Google200Response(varOAuth2Google200Response)

	return err
}

type NullableOAuth2Google200Response struct {
	value *OAuth2Google200Response
	isSet bool
}

func (v NullableOAuth2Google200Response) Get() *OAuth2Google200Response {
	return v.value
}

func (v *NullableOAuth2Google200Response) Set(val *OAuth2Google200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Google200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Google200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Google200Response(val *OAuth2Google200Response) *NullableOAuth2Google200Response {
	return &NullableOAuth2Google200Response{value: val, isSet: true}
}

func (v NullableOAuth2Google200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Google200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}