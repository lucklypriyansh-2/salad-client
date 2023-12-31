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

// StartContainerGroupProblemType the model 'StartContainerGroupProblemType'
type StartContainerGroupProblemType string


// All allowed values of StartContainerGroupProblemType enum
var AllowedStartContainerGroupProblemTypeEnumValues = []StartContainerGroupProblemType{
	"null",
}

func (v *StartContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StartContainerGroupProblemType(value)
	for _, existing := range AllowedStartContainerGroupProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StartContainerGroupProblemType", value)
}

// NewStartContainerGroupProblemTypeFromValue returns a pointer to a valid StartContainerGroupProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStartContainerGroupProblemTypeFromValue(v string) (*StartContainerGroupProblemType, error) {
	ev := StartContainerGroupProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StartContainerGroupProblemType: valid values are %v", v, AllowedStartContainerGroupProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StartContainerGroupProblemType) IsValid() bool {
	for _, existing := range AllowedStartContainerGroupProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StartContainerGroupProblemType value
func (v StartContainerGroupProblemType) Ptr() *StartContainerGroupProblemType {
	return &v
}

type NullableStartContainerGroupProblemType struct {
	value *StartContainerGroupProblemType
	isSet bool
}

func (v NullableStartContainerGroupProblemType) Get() *StartContainerGroupProblemType {
	return v.value
}

func (v *NullableStartContainerGroupProblemType) Set(val *StartContainerGroupProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableStartContainerGroupProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableStartContainerGroupProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartContainerGroupProblemType(val *StartContainerGroupProblemType) *NullableStartContainerGroupProblemType {
	return &NullableStartContainerGroupProblemType{value: val, isSet: true}
}

func (v NullableStartContainerGroupProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

