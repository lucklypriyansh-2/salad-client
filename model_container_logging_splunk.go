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

// checks if the ContainerLoggingSplunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerLoggingSplunk{}

// ContainerLoggingSplunk struct for ContainerLoggingSplunk
type ContainerLoggingSplunk struct {
	Host string `json:"host"`
	Token string `json:"token"`
}

// NewContainerLoggingSplunk instantiates a new ContainerLoggingSplunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLoggingSplunk(host string, token string) *ContainerLoggingSplunk {
	this := ContainerLoggingSplunk{}
	this.Host = host
	this.Token = token
	return &this
}

// NewContainerLoggingSplunkWithDefaults instantiates a new ContainerLoggingSplunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLoggingSplunkWithDefaults() *ContainerLoggingSplunk {
	this := ContainerLoggingSplunk{}
	return &this
}

// GetHost returns the Host field value
func (o *ContainerLoggingSplunk) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingSplunk) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ContainerLoggingSplunk) SetHost(v string) {
	o.Host = v
}

// GetToken returns the Token field value
func (o *ContainerLoggingSplunk) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingSplunk) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ContainerLoggingSplunk) SetToken(v string) {
	o.Token = v
}

func (o ContainerLoggingSplunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerLoggingSplunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableContainerLoggingSplunk struct {
	value *ContainerLoggingSplunk
	isSet bool
}

func (v NullableContainerLoggingSplunk) Get() *ContainerLoggingSplunk {
	return v.value
}

func (v *NullableContainerLoggingSplunk) Set(val *ContainerLoggingSplunk) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingSplunk) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingSplunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingSplunk(val *ContainerLoggingSplunk) *NullableContainerLoggingSplunk {
	return &NullableContainerLoggingSplunk{value: val, isSet: true}
}

func (v NullableContainerLoggingSplunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingSplunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


