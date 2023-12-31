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

// checks if the CreateContainerRegistryAuthenticationBasic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerRegistryAuthenticationBasic{}

// CreateContainerRegistryAuthenticationBasic struct for CreateContainerRegistryAuthenticationBasic
type CreateContainerRegistryAuthenticationBasic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewCreateContainerRegistryAuthenticationBasic instantiates a new CreateContainerRegistryAuthenticationBasic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerRegistryAuthenticationBasic(username string, password string) *CreateContainerRegistryAuthenticationBasic {
	this := CreateContainerRegistryAuthenticationBasic{}
	this.Username = username
	this.Password = password
	return &this
}

// NewCreateContainerRegistryAuthenticationBasicWithDefaults instantiates a new CreateContainerRegistryAuthenticationBasic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerRegistryAuthenticationBasicWithDefaults() *CreateContainerRegistryAuthenticationBasic {
	this := CreateContainerRegistryAuthenticationBasic{}
	return &this
}

// GetUsername returns the Username field value
func (o *CreateContainerRegistryAuthenticationBasic) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateContainerRegistryAuthenticationBasic) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateContainerRegistryAuthenticationBasic) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CreateContainerRegistryAuthenticationBasic) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateContainerRegistryAuthenticationBasic) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateContainerRegistryAuthenticationBasic) SetPassword(v string) {
	o.Password = v
}

func (o CreateContainerRegistryAuthenticationBasic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainerRegistryAuthenticationBasic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableCreateContainerRegistryAuthenticationBasic struct {
	value *CreateContainerRegistryAuthenticationBasic
	isSet bool
}

func (v NullableCreateContainerRegistryAuthenticationBasic) Get() *CreateContainerRegistryAuthenticationBasic {
	return v.value
}

func (v *NullableCreateContainerRegistryAuthenticationBasic) Set(val *CreateContainerRegistryAuthenticationBasic) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerRegistryAuthenticationBasic) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerRegistryAuthenticationBasic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerRegistryAuthenticationBasic(val *CreateContainerRegistryAuthenticationBasic) *NullableCreateContainerRegistryAuthenticationBasic {
	return &NullableCreateContainerRegistryAuthenticationBasic{value: val, isSet: true}
}

func (v NullableCreateContainerRegistryAuthenticationBasic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerRegistryAuthenticationBasic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


