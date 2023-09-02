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

// checks if the GpuClassesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GpuClassesList{}

// GpuClassesList Represents a list of GPU classes
type GpuClassesList struct {
	// The list of GPU classes
	Items []GpuClass `json:"items"`
}

// NewGpuClassesList instantiates a new GpuClassesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpuClassesList(items []GpuClass) *GpuClassesList {
	this := GpuClassesList{}
	this.Items = items
	return &this
}

// NewGpuClassesListWithDefaults instantiates a new GpuClassesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpuClassesListWithDefaults() *GpuClassesList {
	this := GpuClassesList{}
	return &this
}

// GetItems returns the Items field value
func (o *GpuClassesList) GetItems() []GpuClass {
	if o == nil {
		var ret []GpuClass
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GpuClassesList) GetItemsOk() ([]GpuClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GpuClassesList) SetItems(v []GpuClass) {
	o.Items = v
}

func (o GpuClassesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GpuClassesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableGpuClassesList struct {
	value *GpuClassesList
	isSet bool
}

func (v NullableGpuClassesList) Get() *GpuClassesList {
	return v.value
}

func (v *NullableGpuClassesList) Set(val *GpuClassesList) {
	v.value = val
	v.isSet = true
}

func (v NullableGpuClassesList) IsSet() bool {
	return v.isSet
}

func (v *NullableGpuClassesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpuClassesList(val *GpuClassesList) *NullableGpuClassesList {
	return &NullableGpuClassesList{value: val, isSet: true}
}

func (v NullableGpuClassesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpuClassesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

