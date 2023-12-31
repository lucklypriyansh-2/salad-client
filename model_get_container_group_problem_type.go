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

// GetContainerGroupProblemType the model 'GetContainerGroupProblemType'
type GetContainerGroupProblemType string

// All allowed values of GetContainerGroupProblemType enum
var AllowedGetContainerGroupProblemTypeEnumValues = []GetContainerGroupProblemType{
	"null",
}

func (v *GetContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetContainerGroupProblemType(value)
	for _, existing := range AllowedGetContainerGroupProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetContainerGroupProblemType", value)
}

// NewGetContainerGroupProblemTypeFromValue returns a pointer to a valid GetContainerGroupProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetContainerGroupProblemTypeFromValue(v string) (*GetContainerGroupProblemType, error) {
	ev := GetContainerGroupProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetContainerGroupProblemType: valid values are %v", v, AllowedGetContainerGroupProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetContainerGroupProblemType) IsValid() bool {
	for _, existing := range AllowedGetContainerGroupProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetContainerGroupProblemType value
func (v GetContainerGroupProblemType) Ptr() *GetContainerGroupProblemType {
	return &v
}

type NullableGetContainerGroupProblemType struct {
	value *GetContainerGroupProblemType
	isSet bool
}

func (v NullableGetContainerGroupProblemType) Get() *GetContainerGroupProblemType {
	return v.value
}

func (v *NullableGetContainerGroupProblemType) Set(val *GetContainerGroupProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContainerGroupProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContainerGroupProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContainerGroupProblemType(val *GetContainerGroupProblemType) *NullableGetContainerGroupProblemType {
	return &NullableGetContainerGroupProblemType{value: val, isSet: true}
}

func (v NullableGetContainerGroupProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

