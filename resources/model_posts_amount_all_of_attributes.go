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

// checks if the PostsAmountAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostsAmountAllOfAttributes{}

// PostsAmountAllOfAttributes struct for PostsAmountAllOfAttributes
type PostsAmountAllOfAttributes struct {
	Amount int64 `json:"amount"`
}

type _PostsAmountAllOfAttributes PostsAmountAllOfAttributes

// NewPostsAmountAllOfAttributes instantiates a new PostsAmountAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostsAmountAllOfAttributes(amount int64) *PostsAmountAllOfAttributes {
	this := PostsAmountAllOfAttributes{}
	this.Amount = amount
	return &this
}

// NewPostsAmountAllOfAttributesWithDefaults instantiates a new PostsAmountAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostsAmountAllOfAttributesWithDefaults() *PostsAmountAllOfAttributes {
	this := PostsAmountAllOfAttributes{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PostsAmountAllOfAttributes) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PostsAmountAllOfAttributes) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PostsAmountAllOfAttributes) SetAmount(v int64) {
	o.Amount = v
}

func (o PostsAmountAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostsAmountAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *PostsAmountAllOfAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
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

	varPostsAmountAllOfAttributes := _PostsAmountAllOfAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostsAmountAllOfAttributes)

	if err != nil {
		return err
	}

	*o = PostsAmountAllOfAttributes(varPostsAmountAllOfAttributes)

	return err
}

type NullablePostsAmountAllOfAttributes struct {
	value *PostsAmountAllOfAttributes
	isSet bool
}

func (v NullablePostsAmountAllOfAttributes) Get() *PostsAmountAllOfAttributes {
	return v.value
}

func (v *NullablePostsAmountAllOfAttributes) Set(val *PostsAmountAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePostsAmountAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePostsAmountAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostsAmountAllOfAttributes(val *PostsAmountAllOfAttributes) *NullablePostsAmountAllOfAttributes {
	return &NullablePostsAmountAllOfAttributes{value: val, isSet: true}
}

func (v NullablePostsAmountAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostsAmountAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
