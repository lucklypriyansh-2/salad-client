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

// checks if the ContainerGroupNetworking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupNetworking{}

// ContainerGroupNetworking Represents container group networking parameters
type ContainerGroupNetworking struct {
	Protocol ContainerNetworkingProtocol `json:"protocol"`
	Port int32 `json:"port"`
	Auth bool `json:"auth"`
	Dns string `json:"dns"`
}

// NewContainerGroupNetworking instantiates a new ContainerGroupNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupNetworking(protocol ContainerNetworkingProtocol, port int32, auth bool, dns string) *ContainerGroupNetworking {
	this := ContainerGroupNetworking{}
	this.Protocol = protocol
	this.Port = port
	this.Auth = auth
	this.Dns = dns
	return &this
}

// NewContainerGroupNetworkingWithDefaults instantiates a new ContainerGroupNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupNetworkingWithDefaults() *ContainerGroupNetworking {
	this := ContainerGroupNetworking{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *ContainerGroupNetworking) GetProtocol() ContainerNetworkingProtocol {
	if o == nil {
		var ret ContainerNetworkingProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupNetworking) GetProtocolOk() (*ContainerNetworkingProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *ContainerGroupNetworking) SetProtocol(v ContainerNetworkingProtocol) {
	o.Protocol = v
}

// GetPort returns the Port field value
func (o *ContainerGroupNetworking) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupNetworking) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerGroupNetworking) SetPort(v int32) {
	o.Port = v
}

// GetAuth returns the Auth field value
func (o *ContainerGroupNetworking) GetAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupNetworking) GetAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *ContainerGroupNetworking) SetAuth(v bool) {
	o.Auth = v
}

// GetDns returns the Dns field value
func (o *ContainerGroupNetworking) GetDns() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dns
}

// GetDnsOk returns a tuple with the Dns field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupNetworking) GetDnsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dns, true
}

// SetDns sets field value
func (o *ContainerGroupNetworking) SetDns(v string) {
	o.Dns = v
}

func (o ContainerGroupNetworking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupNetworking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["port"] = o.Port
	toSerialize["auth"] = o.Auth
	toSerialize["dns"] = o.Dns
	return toSerialize, nil
}

type NullableContainerGroupNetworking struct {
	value *ContainerGroupNetworking
	isSet bool
}

func (v NullableContainerGroupNetworking) Get() *ContainerGroupNetworking {
	return v.value
}

func (v *NullableContainerGroupNetworking) Set(val *ContainerGroupNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupNetworking(val *ContainerGroupNetworking) *NullableContainerGroupNetworking {
	return &NullableContainerGroupNetworking{value: val, isSet: true}
}

func (v NullableContainerGroupNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


