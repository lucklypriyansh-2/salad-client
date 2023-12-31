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

// checks if the ContainerGroupProbeGrpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupProbeGrpc{}

// ContainerGroupProbeGrpc struct for ContainerGroupProbeGrpc
type ContainerGroupProbeGrpc struct {
	Service string `json:"service"`
	Port int32 `json:"port"`
}

// NewContainerGroupProbeGrpc instantiates a new ContainerGroupProbeGrpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupProbeGrpc(service string, port int32) *ContainerGroupProbeGrpc {
	this := ContainerGroupProbeGrpc{}
	this.Service = service
	this.Port = port
	return &this
}

// NewContainerGroupProbeGrpcWithDefaults instantiates a new ContainerGroupProbeGrpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupProbeGrpcWithDefaults() *ContainerGroupProbeGrpc {
	this := ContainerGroupProbeGrpc{}
	return &this
}

// GetService returns the Service field value
func (o *ContainerGroupProbeGrpc) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeGrpc) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ContainerGroupProbeGrpc) SetService(v string) {
	o.Service = v
}

// GetPort returns the Port field value
func (o *ContainerGroupProbeGrpc) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeGrpc) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerGroupProbeGrpc) SetPort(v int32) {
	o.Port = v
}

func (o ContainerGroupProbeGrpc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupProbeGrpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service"] = o.Service
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableContainerGroupProbeGrpc struct {
	value *ContainerGroupProbeGrpc
	isSet bool
}

func (v NullableContainerGroupProbeGrpc) Get() *ContainerGroupProbeGrpc {
	return v.value
}

func (v *NullableContainerGroupProbeGrpc) Set(val *ContainerGroupProbeGrpc) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupProbeGrpc) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupProbeGrpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupProbeGrpc(val *ContainerGroupProbeGrpc) *NullableContainerGroupProbeGrpc {
	return &NullableContainerGroupProbeGrpc{value: val, isSet: true}
}

func (v NullableContainerGroupProbeGrpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupProbeGrpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


