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
	"time"
)

// checks if the Quotas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Quotas{}

// Quotas Represents the organization quotas
type Quotas struct {
	ContainerGroupsQuotas ContainerGroupsQuotas `json:"container_groups_quotas"`
	// The time the resource was created
	CreateTime *time.Time `json:"create_time,omitempty"`
	RecipesQuotas RecipesQuotas `json:"recipes_quotas"`
	// The time the resource was last updated
	UpdateTime *time.Time `json:"update_time,omitempty"`
}

// NewQuotas instantiates a new Quotas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotas(containerGroupsQuotas ContainerGroupsQuotas, recipesQuotas RecipesQuotas) *Quotas {
	this := Quotas{}
	this.ContainerGroupsQuotas = containerGroupsQuotas
	this.RecipesQuotas = recipesQuotas
	return &this
}

// NewQuotasWithDefaults instantiates a new Quotas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotasWithDefaults() *Quotas {
	this := Quotas{}
	return &this
}

// GetContainerGroupsQuotas returns the ContainerGroupsQuotas field value
func (o *Quotas) GetContainerGroupsQuotas() ContainerGroupsQuotas {
	if o == nil {
		var ret ContainerGroupsQuotas
		return ret
	}

	return o.ContainerGroupsQuotas
}

// GetContainerGroupsQuotasOk returns a tuple with the ContainerGroupsQuotas field value
// and a boolean to check if the value has been set.
func (o *Quotas) GetContainerGroupsQuotasOk() (*ContainerGroupsQuotas, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerGroupsQuotas, true
}

// SetContainerGroupsQuotas sets field value
func (o *Quotas) SetContainerGroupsQuotas(v ContainerGroupsQuotas) {
	o.ContainerGroupsQuotas = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Quotas) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quotas) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Quotas) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Quotas) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetRecipesQuotas returns the RecipesQuotas field value
func (o *Quotas) GetRecipesQuotas() RecipesQuotas {
	if o == nil {
		var ret RecipesQuotas
		return ret
	}

	return o.RecipesQuotas
}

// GetRecipesQuotasOk returns a tuple with the RecipesQuotas field value
// and a boolean to check if the value has been set.
func (o *Quotas) GetRecipesQuotasOk() (*RecipesQuotas, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipesQuotas, true
}

// SetRecipesQuotas sets field value
func (o *Quotas) SetRecipesQuotas(v RecipesQuotas) {
	o.RecipesQuotas = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Quotas) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quotas) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Quotas) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *Quotas) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

func (o Quotas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Quotas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["container_groups_quotas"] = o.ContainerGroupsQuotas
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	toSerialize["recipes_quotas"] = o.RecipesQuotas
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableQuotas struct {
	value *Quotas
	isSet bool
}

func (v NullableQuotas) Get() *Quotas {
	return v.value
}

func (v *NullableQuotas) Set(val *Quotas) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotas) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotas(val *Quotas) *NullableQuotas {
	return &NullableQuotas{value: val, isSet: true}
}

func (v NullableQuotas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

