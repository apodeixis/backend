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

// checks if the Post type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Post{}

// Post struct for Post
type Post struct {
	Id            string                  `json:"id"`
	Type          string                  `json:"type"`
	Attributes    PostAllOfAttributes     `json:"attributes"`
	Relationships *PostAllOfRelationships `json:"relationships,omitempty"`
}

type _Post Post

// NewPost instantiates a new Post object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPost(id string, type_ string, attributes PostAllOfAttributes) *Post {
	this := Post{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPostWithDefaults instantiates a new Post object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostWithDefaults() *Post {
	this := Post{}
	return &this
}

// GetId returns the Id field value
func (o *Post) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Post) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Post) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Post) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Post) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Post) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *Post) GetAttributes() PostAllOfAttributes {
	if o == nil {
		var ret PostAllOfAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *Post) GetAttributesOk() (*PostAllOfAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *Post) SetAttributes(v PostAllOfAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Post) GetRelationships() PostAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PostAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetRelationshipsOk() (*PostAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Post) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PostAllOfRelationships and assigns it to the Relationships field.
func (o *Post) SetRelationships(v PostAllOfRelationships) {
	o.Relationships = &v
}

func (o Post) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Post) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *Post) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
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

	varPost := _Post{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPost)

	if err != nil {
		return err
	}

	*o = Post(varPost)

	return err
}

type NullablePost struct {
	value *Post
	isSet bool
}

func (v NullablePost) Get() *Post {
	return v.value
}

func (v *NullablePost) Set(val *Post) {
	v.value = val
	v.isSet = true
}

func (v NullablePost) IsSet() bool {
	return v.isSet
}

func (v *NullablePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePost(val *Post) *NullablePost {
	return &NullablePost{value: val, isSet: true}
}

func (v NullablePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}