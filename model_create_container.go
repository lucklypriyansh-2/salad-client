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

// checks if the CreateContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainer{}

// CreateContainer Represents a container
type CreateContainer struct {
	Image string `json:"image"`
	Resources ContainerResourceRequirements `json:"resources"`
	Command []string `json:"command,omitempty"`
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	Logging NullableContainerLogging `json:"logging,omitempty"`
	RegistryAuthentication NullableCreateContainerRegistryAuthentication `json:"registry_authentication,omitempty"`
}

// NewCreateContainer instantiates a new CreateContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainer(image string, resources ContainerResourceRequirements) *CreateContainer {
	this := CreateContainer{}
	this.Image = image
	this.Resources = resources
	return &this
}

// NewCreateContainerWithDefaults instantiates a new CreateContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerWithDefaults() *CreateContainer {
	this := CreateContainer{}
	return &this
}

// GetImage returns the Image field value
func (o *CreateContainer) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *CreateContainer) SetImage(v string) {
	o.Image = v
}

// GetResources returns the Resources field value
func (o *CreateContainer) GetResources() ContainerResourceRequirements {
	if o == nil {
		var ret ContainerResourceRequirements
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetResourcesOk() (*ContainerResourceRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *CreateContainer) SetResources(v ContainerResourceRequirements) {
	o.Resources = v
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainer) GetCommand() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainer) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *CreateContainer) HasCommand() bool {
	if o != nil && IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *CreateContainer) SetCommand(v []string) {
	o.Command = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *CreateContainer) GetEnvironmentVariables() map[string]string {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret map[string]string
		return ret
	}
	return *o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetEnvironmentVariablesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *CreateContainer) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given map[string]string and assigns it to the EnvironmentVariables field.
func (o *CreateContainer) SetEnvironmentVariables(v map[string]string) {
	o.EnvironmentVariables = &v
}

// GetLogging returns the Logging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainer) GetLogging() ContainerLogging {
	if o == nil || IsNil(o.Logging.Get()) {
		var ret ContainerLogging
		return ret
	}
	return *o.Logging.Get()
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainer) GetLoggingOk() (*ContainerLogging, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logging.Get(), o.Logging.IsSet()
}

// HasLogging returns a boolean if a field has been set.
func (o *CreateContainer) HasLogging() bool {
	if o != nil && o.Logging.IsSet() {
		return true
	}

	return false
}

// SetLogging gets a reference to the given NullableContainerLogging and assigns it to the Logging field.
func (o *CreateContainer) SetLogging(v ContainerLogging) {
	o.Logging.Set(&v)
}
// SetLoggingNil sets the value for Logging to be an explicit nil
func (o *CreateContainer) SetLoggingNil() {
	o.Logging.Set(nil)
}

// UnsetLogging ensures that no value is present for Logging, not even an explicit nil
func (o *CreateContainer) UnsetLogging() {
	o.Logging.Unset()
}

// GetRegistryAuthentication returns the RegistryAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainer) GetRegistryAuthentication() CreateContainerRegistryAuthentication {
	if o == nil || IsNil(o.RegistryAuthentication.Get()) {
		var ret CreateContainerRegistryAuthentication
		return ret
	}
	return *o.RegistryAuthentication.Get()
}

// GetRegistryAuthenticationOk returns a tuple with the RegistryAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainer) GetRegistryAuthenticationOk() (*CreateContainerRegistryAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistryAuthentication.Get(), o.RegistryAuthentication.IsSet()
}

// HasRegistryAuthentication returns a boolean if a field has been set.
func (o *CreateContainer) HasRegistryAuthentication() bool {
	if o != nil && o.RegistryAuthentication.IsSet() {
		return true
	}

	return false
}

// SetRegistryAuthentication gets a reference to the given NullableCreateContainerRegistryAuthentication and assigns it to the RegistryAuthentication field.
func (o *CreateContainer) SetRegistryAuthentication(v CreateContainerRegistryAuthentication) {
	o.RegistryAuthentication.Set(&v)
}
// SetRegistryAuthenticationNil sets the value for RegistryAuthentication to be an explicit nil
func (o *CreateContainer) SetRegistryAuthenticationNil() {
	o.RegistryAuthentication.Set(nil)
}

// UnsetRegistryAuthentication ensures that no value is present for RegistryAuthentication, not even an explicit nil
func (o *CreateContainer) UnsetRegistryAuthentication() {
	o.RegistryAuthentication.Unset()
}

func (o CreateContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["resources"] = o.Resources
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environment_variables"] = o.EnvironmentVariables
	}
	if o.Logging.IsSet() {
		toSerialize["logging"] = o.Logging.Get()
	}
	if o.RegistryAuthentication.IsSet() {
		toSerialize["registry_authentication"] = o.RegistryAuthentication.Get()
	}
	return toSerialize, nil
}

type NullableCreateContainer struct {
	value *CreateContainer
	isSet bool
}

func (v NullableCreateContainer) Get() *CreateContainer {
	return v.value
}

func (v *NullableCreateContainer) Set(val *CreateContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainer(val *CreateContainer) *NullableCreateContainer {
	return &NullableCreateContainer{value: val, isSet: true}
}

func (v NullableCreateContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

