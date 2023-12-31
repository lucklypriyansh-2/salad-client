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

// checks if the CreateRecipeDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecipeDeployment{}

// CreateRecipeDeployment Represents a request to create a new recipe deployment
type CreateRecipeDeployment struct {
	Name string `json:"name"`
	DisplayName NullableString `json:"display_name"`
	Replicas int32 `json:"replicas"`
	RecipeName string `json:"recipe_name"`
	Networking NullableCreateRecipeNetworking `json:"networking,omitempty"`
}

// NewCreateRecipeDeployment instantiates a new CreateRecipeDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecipeDeployment(name string, displayName NullableString, replicas int32, recipeName string) *CreateRecipeDeployment {
	this := CreateRecipeDeployment{}
	this.Name = name
	this.DisplayName = displayName
	this.Replicas = replicas
	this.RecipeName = recipeName
	return &this
}

// NewCreateRecipeDeploymentWithDefaults instantiates a new CreateRecipeDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecipeDeploymentWithDefaults() *CreateRecipeDeployment {
	this := CreateRecipeDeployment{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRecipeDeployment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRecipeDeployment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRecipeDeployment) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateRecipeDeployment) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRecipeDeployment) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// SetDisplayName sets field value
func (o *CreateRecipeDeployment) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// GetReplicas returns the Replicas field value
func (o *CreateRecipeDeployment) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *CreateRecipeDeployment) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value
func (o *CreateRecipeDeployment) SetReplicas(v int32) {
	o.Replicas = v
}

// GetRecipeName returns the RecipeName field value
func (o *CreateRecipeDeployment) GetRecipeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipeName
}

// GetRecipeNameOk returns a tuple with the RecipeName field value
// and a boolean to check if the value has been set.
func (o *CreateRecipeDeployment) GetRecipeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipeName, true
}

// SetRecipeName sets field value
func (o *CreateRecipeDeployment) SetRecipeName(v string) {
	o.RecipeName = v
}

// GetNetworking returns the Networking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRecipeDeployment) GetNetworking() CreateRecipeNetworking {
	if o == nil || IsNil(o.Networking.Get()) {
		var ret CreateRecipeNetworking
		return ret
	}
	return *o.Networking.Get()
}

// GetNetworkingOk returns a tuple with the Networking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRecipeDeployment) GetNetworkingOk() (*CreateRecipeNetworking, bool) {
	if o == nil {
		return nil, false
	}
	return o.Networking.Get(), o.Networking.IsSet()
}

// HasNetworking returns a boolean if a field has been set.
func (o *CreateRecipeDeployment) HasNetworking() bool {
	if o != nil && o.Networking.IsSet() {
		return true
	}

	return false
}

// SetNetworking gets a reference to the given NullableCreateRecipeNetworking and assigns it to the Networking field.
func (o *CreateRecipeDeployment) SetNetworking(v CreateRecipeNetworking) {
	o.Networking.Set(&v)
}
// SetNetworkingNil sets the value for Networking to be an explicit nil
func (o *CreateRecipeDeployment) SetNetworkingNil() {
	o.Networking.Set(nil)
}

// UnsetNetworking ensures that no value is present for Networking, not even an explicit nil
func (o *CreateRecipeDeployment) UnsetNetworking() {
	o.Networking.Unset()
}

func (o CreateRecipeDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecipeDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName.Get()
	toSerialize["replicas"] = o.Replicas
	toSerialize["recipe_name"] = o.RecipeName
	if o.Networking.IsSet() {
		toSerialize["networking"] = o.Networking.Get()
	}
	return toSerialize, nil
}

type NullableCreateRecipeDeployment struct {
	value *CreateRecipeDeployment
	isSet bool
}

func (v NullableCreateRecipeDeployment) Get() *CreateRecipeDeployment {
	return v.value
}

func (v *NullableCreateRecipeDeployment) Set(val *CreateRecipeDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecipeDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecipeDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecipeDeployment(val *CreateRecipeDeployment) *NullableCreateRecipeDeployment {
	return &NullableCreateRecipeDeployment{value: val, isSet: true}
}

func (v NullableCreateRecipeDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecipeDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


