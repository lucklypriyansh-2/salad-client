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

// checks if the ContainerLoggingNewRelic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerLoggingNewRelic{}

// ContainerLoggingNewRelic struct for ContainerLoggingNewRelic
type ContainerLoggingNewRelic struct {
	Host string `json:"host"`
	IngestionKey string `json:"ingestion_key"`
}

// NewContainerLoggingNewRelic instantiates a new ContainerLoggingNewRelic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLoggingNewRelic(host string, ingestionKey string) *ContainerLoggingNewRelic {
	this := ContainerLoggingNewRelic{}
	this.Host = host
	this.IngestionKey = ingestionKey
	return &this
}

// NewContainerLoggingNewRelicWithDefaults instantiates a new ContainerLoggingNewRelic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLoggingNewRelicWithDefaults() *ContainerLoggingNewRelic {
	this := ContainerLoggingNewRelic{}
	return &this
}

// GetHost returns the Host field value
func (o *ContainerLoggingNewRelic) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingNewRelic) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ContainerLoggingNewRelic) SetHost(v string) {
	o.Host = v
}

// GetIngestionKey returns the IngestionKey field value
func (o *ContainerLoggingNewRelic) GetIngestionKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IngestionKey
}

// GetIngestionKeyOk returns a tuple with the IngestionKey field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingNewRelic) GetIngestionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IngestionKey, true
}

// SetIngestionKey sets field value
func (o *ContainerLoggingNewRelic) SetIngestionKey(v string) {
	o.IngestionKey = v
}

func (o ContainerLoggingNewRelic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerLoggingNewRelic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["ingestion_key"] = o.IngestionKey
	return toSerialize, nil
}

type NullableContainerLoggingNewRelic struct {
	value *ContainerLoggingNewRelic
	isSet bool
}

func (v NullableContainerLoggingNewRelic) Get() *ContainerLoggingNewRelic {
	return v.value
}

func (v *NullableContainerLoggingNewRelic) Set(val *ContainerLoggingNewRelic) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingNewRelic) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingNewRelic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingNewRelic(val *ContainerLoggingNewRelic) *NullableContainerLoggingNewRelic {
	return &NullableContainerLoggingNewRelic{value: val, isSet: true}
}

func (v NullableContainerLoggingNewRelic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingNewRelic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

