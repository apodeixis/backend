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

// checks if the PostAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAllOfAttributes{}

// PostAllOfAttributes struct for PostAllOfAttributes
type PostAllOfAttributes struct {
	Title       string  `json:"title"`
	Body        string  `json:"body"`
	TxHash      *string `json:"tx_hash,omitempty"`
	TxTimestamp *int64  `json:"tx_timestamp,omitempty"`
	Starred     bool    `json:"starred"`
	Status      string  `json:"status"`
}

type _PostAllOfAttributes PostAllOfAttributes

// NewPostAllOfAttributes instantiates a new PostAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAllOfAttributes(title string, body string, starred bool, status string) *PostAllOfAttributes {
	this := PostAllOfAttributes{}
	this.Title = title
	this.Body = body
	this.Starred = starred
	this.Status = status
	return &this
}

// NewPostAllOfAttributesWithDefaults instantiates a new PostAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAllOfAttributesWithDefaults() *PostAllOfAttributes {
	this := PostAllOfAttributes{}
	return &this
}

// GetTitle returns the Title field value
func (o *PostAllOfAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PostAllOfAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PostAllOfAttributes) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value
func (o *PostAllOfAttributes) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *PostAllOfAttributes) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *PostAllOfAttributes) SetBody(v string) {
	o.Body = v
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *PostAllOfAttributes) GetTxHash() string {
	if o == nil || IsNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllOfAttributes) GetTxHashOk() (*string, bool) {
	if o == nil || IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *PostAllOfAttributes) HasTxHash() bool {
	if o != nil && !IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *PostAllOfAttributes) SetTxHash(v string) {
	o.TxHash = &v
}

// GetTxTimestamp returns the TxTimestamp field value if set, zero value otherwise.
func (o *PostAllOfAttributes) GetTxTimestamp() int64 {
	if o == nil || IsNil(o.TxTimestamp) {
		var ret int64
		return ret
	}
	return *o.TxTimestamp
}

// GetTxTimestampOk returns a tuple with the TxTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllOfAttributes) GetTxTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.TxTimestamp) {
		return nil, false
	}
	return o.TxTimestamp, true
}

// HasTxTimestamp returns a boolean if a field has been set.
func (o *PostAllOfAttributes) HasTxTimestamp() bool {
	if o != nil && !IsNil(o.TxTimestamp) {
		return true
	}

	return false
}

// SetTxTimestamp gets a reference to the given int64 and assigns it to the TxTimestamp field.
func (o *PostAllOfAttributes) SetTxTimestamp(v int64) {
	o.TxTimestamp = &v
}

// GetStarred returns the Starred field value
func (o *PostAllOfAttributes) GetStarred() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Starred
}

// GetStarredOk returns a tuple with the Starred field value
// and a boolean to check if the value has been set.
func (o *PostAllOfAttributes) GetStarredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Starred, true
}

// SetStarred sets field value
func (o *PostAllOfAttributes) SetStarred(v bool) {
	o.Starred = v
}

// GetStatus returns the Status field value
func (o *PostAllOfAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostAllOfAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PostAllOfAttributes) SetStatus(v string) {
	o.Status = v
}

func (o PostAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["body"] = o.Body
	if !IsNil(o.TxHash) {
		toSerialize["tx_hash"] = o.TxHash
	}
	if !IsNil(o.TxTimestamp) {
		toSerialize["tx_timestamp"] = o.TxTimestamp
	}
	toSerialize["starred"] = o.Starred
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *PostAllOfAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"body",
		"starred",
		"status",
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

	varPostAllOfAttributes := _PostAllOfAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostAllOfAttributes)

	if err != nil {
		return err
	}

	*o = PostAllOfAttributes(varPostAllOfAttributes)

	return err
}

type NullablePostAllOfAttributes struct {
	value *PostAllOfAttributes
	isSet bool
}

func (v NullablePostAllOfAttributes) Get() *PostAllOfAttributes {
	return v.value
}

func (v *NullablePostAllOfAttributes) Set(val *PostAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAllOfAttributes(val *PostAllOfAttributes) *NullablePostAllOfAttributes {
	return &NullablePostAllOfAttributes{value: val, isSet: true}
}

func (v NullablePostAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}