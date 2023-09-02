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

// checks if the ContainerGroupInstances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupInstances{}

// ContainerGroupInstances Represents a list of container group instances
type ContainerGroupInstances struct {
	Instances []ContainerGroupInstancesInstancesInner `json:"instances"`
}

// NewContainerGroupInstances instantiates a new ContainerGroupInstances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupInstances(instances []ContainerGroupInstancesInstancesInner) *ContainerGroupInstances {
	this := ContainerGroupInstances{}
	this.Instances = instances
	return &this
}

// NewContainerGroupInstancesWithDefaults instantiates a new ContainerGroupInstances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupInstancesWithDefaults() *ContainerGroupInstances {
	this := ContainerGroupInstances{}
	return &this
}

// GetInstances returns the Instances field value
func (o *ContainerGroupInstances) GetInstances() []ContainerGroupInstancesInstancesInner {
	if o == nil {
		var ret []ContainerGroupInstancesInstancesInner
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupInstances) GetInstancesOk() ([]ContainerGroupInstancesInstancesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *ContainerGroupInstances) SetInstances(v []ContainerGroupInstancesInstancesInner) {
	o.Instances = v
}

func (o ContainerGroupInstances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupInstances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instances"] = o.Instances
	return toSerialize, nil
}

type NullableContainerGroupInstances struct {
	value *ContainerGroupInstances
	isSet bool
}

func (v NullableContainerGroupInstances) Get() *ContainerGroupInstances {
	return v.value
}

func (v *NullableContainerGroupInstances) Set(val *ContainerGroupInstances) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupInstances(val *ContainerGroupInstances) *NullableContainerGroupInstances {
	return &NullableContainerGroupInstances{value: val, isSet: true}
}

func (v NullableContainerGroupInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

