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
	"fmt"
)

// CreateContainerGroupProblemType the model 'CreateContainerGroupProblemType'
type CreateContainerGroupProblemType string

// List of CreateContainerGroupProblemType
const (
	INVALID_LOGGING_CONFIGURATION CreateContainerGroupProblemType = "invalid_logging_configuration"
	INVALID_REGISTRY_AUTHENTICATION CreateContainerGroupProblemType = "invalid_registry_authentication"
)

// All allowed values of CreateContainerGroupProblemType enum
var AllowedCreateContainerGroupProblemTypeEnumValues = []CreateContainerGroupProblemType{
	"name_conflict",
	"created_quota_exceeded",
	"invalid_logging_configuration",
	"invalid_registry_authentication",
	"created_instance_quota_exceeded",
	"null",
}

func (v *CreateContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateContainerGroupProblemType(value)
	for _, existing := range AllowedCreateContainerGroupProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateContainerGroupProblemType", value)
}

// NewCreateContainerGroupProblemTypeFromValue returns a pointer to a valid CreateContainerGroupProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateContainerGroupProblemTypeFromValue(v string) (*CreateContainerGroupProblemType, error) {
	ev := CreateContainerGroupProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateContainerGroupProblemType: valid values are %v", v, AllowedCreateContainerGroupProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateContainerGroupProblemType) IsValid() bool {
	for _, existing := range AllowedCreateContainerGroupProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateContainerGroupProblemType value
func (v CreateContainerGroupProblemType) Ptr() *CreateContainerGroupProblemType {
	return &v
}

type NullableCreateContainerGroupProblemType struct {
	value *CreateContainerGroupProblemType
	isSet bool
}

func (v NullableCreateContainerGroupProblemType) Get() *CreateContainerGroupProblemType {
	return v.value
}

func (v *NullableCreateContainerGroupProblemType) Set(val *CreateContainerGroupProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerGroupProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerGroupProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerGroupProblemType(val *CreateContainerGroupProblemType) *NullableCreateContainerGroupProblemType {
	return &NullableCreateContainerGroupProblemType{value: val, isSet: true}
}

func (v NullableCreateContainerGroupProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

