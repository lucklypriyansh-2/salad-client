/*
SaladCloud Public API

The SaladCloud Public API.

API version: 1.0.0-alpha.56
Contact: support@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContainerGroupList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupList{}

// ContainerGroupList Represents a list of container groups
type ContainerGroupList struct {
	Items []ContainerGroup `json:"items"`
}

// NewContainerGroupList instantiates a new ContainerGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupList(items []ContainerGroup) *ContainerGroupList {
	this := ContainerGroupList{}
	this.Items = items
	return &this
}

// NewContainerGroupListWithDefaults instantiates a new ContainerGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupListWithDefaults() *ContainerGroupList {
	this := ContainerGroupList{}
	return &this
}

// GetItems returns the Items field value
func (o *ContainerGroupList) GetItems() []ContainerGroup {
	if o == nil {
		var ret []ContainerGroup
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupList) GetItemsOk() ([]ContainerGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ContainerGroupList) SetItems(v []ContainerGroup) {
	o.Items = v
}

func (o ContainerGroupList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableContainerGroupList struct {
	value *ContainerGroupList
	isSet bool
}

func (v NullableContainerGroupList) Get() *ContainerGroupList {
	return v.value
}

func (v *NullableContainerGroupList) Set(val *ContainerGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupList(val *ContainerGroupList) *NullableContainerGroupList {
	return &NullableContainerGroupList{value: val, isSet: true}
}

func (v NullableContainerGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

