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

// checks if the Container type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Container{}

// Container Represents a container
type Container struct {
	Image string `json:"image"`
	Resources ContainerResourceRequirements `json:"resources"`
	Command []string `json:"command"`
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	Logging NullableContainerLogging `json:"logging,omitempty"`
}

// NewContainer instantiates a new Container object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainer(image string, resources ContainerResourceRequirements, command []string) *Container {
	this := Container{}
	this.Image = image
	this.Resources = resources
	this.Command = command
	return &this
}

// NewContainerWithDefaults instantiates a new Container object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerWithDefaults() *Container {
	this := Container{}
	return &this
}

// GetImage returns the Image field value
func (o *Container) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *Container) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *Container) SetImage(v string) {
	o.Image = v
}

// GetResources returns the Resources field value
func (o *Container) GetResources() ContainerResourceRequirements {
	if o == nil {
		var ret ContainerResourceRequirements
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *Container) GetResourcesOk() (*ContainerResourceRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *Container) SetResources(v ContainerResourceRequirements) {
	o.Resources = v
}

// GetCommand returns the Command field value
func (o *Container) GetCommand() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *Container) GetCommandOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Command, true
}

// SetCommand sets field value
func (o *Container) SetCommand(v []string) {
	o.Command = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *Container) GetEnvironmentVariables() map[string]string {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret map[string]string
		return ret
	}
	return *o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetEnvironmentVariablesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *Container) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given map[string]string and assigns it to the EnvironmentVariables field.
func (o *Container) SetEnvironmentVariables(v map[string]string) {
	o.EnvironmentVariables = &v
}

// GetLogging returns the Logging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetLogging() ContainerLogging {
	if o == nil || IsNil(o.Logging.Get()) {
		var ret ContainerLogging
		return ret
	}
	return *o.Logging.Get()
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetLoggingOk() (*ContainerLogging, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logging.Get(), o.Logging.IsSet()
}

// HasLogging returns a boolean if a field has been set.
func (o *Container) HasLogging() bool {
	if o != nil && o.Logging.IsSet() {
		return true
	}

	return false
}

// SetLogging gets a reference to the given NullableContainerLogging and assigns it to the Logging field.
func (o *Container) SetLogging(v ContainerLogging) {
	o.Logging.Set(&v)
}
// SetLoggingNil sets the value for Logging to be an explicit nil
func (o *Container) SetLoggingNil() {
	o.Logging.Set(nil)
}

// UnsetLogging ensures that no value is present for Logging, not even an explicit nil
func (o *Container) UnsetLogging() {
	o.Logging.Unset()
}

func (o Container) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Container) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["resources"] = o.Resources
	toSerialize["command"] = o.Command
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environment_variables"] = o.EnvironmentVariables
	}
	if o.Logging.IsSet() {
		toSerialize["logging"] = o.Logging.Get()
	}
	return toSerialize, nil
}

type NullableContainer struct {
	value *Container
	isSet bool
}

func (v NullableContainer) Get() *Container {
	return v.value
}

func (v *NullableContainer) Set(val *Container) {
	v.value = val
	v.isSet = true
}

func (v NullableContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainer(val *Container) *NullableContainer {
	return &NullableContainer{value: val, isSet: true}
}

func (v NullableContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

