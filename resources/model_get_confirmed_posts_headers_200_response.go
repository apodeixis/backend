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

// checks if the GetConfirmedPostsHeaders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConfirmedPostsHeaders200Response{}

// GetConfirmedPostsHeaders200Response struct for GetConfirmedPostsHeaders200Response
type GetConfirmedPostsHeaders200Response struct {
	Data     []PostHeader `json:"data"`
	Included []User       `json:"included,omitempty"`
}

type _GetConfirmedPostsHeaders200Response GetConfirmedPostsHeaders200Response

// NewGetConfirmedPostsHeaders200Response instantiates a new GetConfirmedPostsHeaders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConfirmedPostsHeaders200Response(data []PostHeader) *GetConfirmedPostsHeaders200Response {
	this := GetConfirmedPostsHeaders200Response{}
	this.Data = data
	return &this
}

// NewGetConfirmedPostsHeaders200ResponseWithDefaults instantiates a new GetConfirmedPostsHeaders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConfirmedPostsHeaders200ResponseWithDefaults() *GetConfirmedPostsHeaders200Response {
	this := GetConfirmedPostsHeaders200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetConfirmedPostsHeaders200Response) GetData() []PostHeader {
	if o == nil {
		var ret []PostHeader
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetConfirmedPostsHeaders200Response) GetDataOk() ([]PostHeader, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetConfirmedPostsHeaders200Response) SetData(v []PostHeader) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetConfirmedPostsHeaders200Response) GetIncluded() []User {
	if o == nil || IsNil(o.Included) {
		var ret []User
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConfirmedPostsHeaders200Response) GetIncludedOk() ([]User, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetConfirmedPostsHeaders200Response) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []User and assigns it to the Included field.
func (o *GetConfirmedPostsHeaders200Response) SetIncluded(v []User) {
	o.Included = v
}

func (o GetConfirmedPostsHeaders200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConfirmedPostsHeaders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetConfirmedPostsHeaders200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetConfirmedPostsHeaders200Response := _GetConfirmedPostsHeaders200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetConfirmedPostsHeaders200Response)

	if err != nil {
		return err
	}

	*o = GetConfirmedPostsHeaders200Response(varGetConfirmedPostsHeaders200Response)

	return err
}

type NullableGetConfirmedPostsHeaders200Response struct {
	value *GetConfirmedPostsHeaders200Response
	isSet bool
}

func (v NullableGetConfirmedPostsHeaders200Response) Get() *GetConfirmedPostsHeaders200Response {
	return v.value
}

func (v *NullableGetConfirmedPostsHeaders200Response) Set(val *GetConfirmedPostsHeaders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConfirmedPostsHeaders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConfirmedPostsHeaders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConfirmedPostsHeaders200Response(val *GetConfirmedPostsHeaders200Response) *NullableGetConfirmedPostsHeaders200Response {
	return &NullableGetConfirmedPostsHeaders200Response{value: val, isSet: true}
}

func (v NullableGetConfirmedPostsHeaders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConfirmedPostsHeaders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
