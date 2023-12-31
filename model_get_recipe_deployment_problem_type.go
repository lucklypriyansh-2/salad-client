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

// GetRecipeDeploymentProblemType the model 'GetRecipeDeploymentProblemType'
type GetRecipeDeploymentProblemType string


// All allowed values of GetRecipeDeploymentProblemType enum
var AllowedGetRecipeDeploymentProblemTypeEnumValues = []GetRecipeDeploymentProblemType{
	"null",
}

func (v *GetRecipeDeploymentProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetRecipeDeploymentProblemType(value)
	for _, existing := range AllowedGetRecipeDeploymentProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetRecipeDeploymentProblemType", value)
}

// NewGetRecipeDeploymentProblemTypeFromValue returns a pointer to a valid GetRecipeDeploymentProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetRecipeDeploymentProblemTypeFromValue(v string) (*GetRecipeDeploymentProblemType, error) {
	ev := GetRecipeDeploymentProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetRecipeDeploymentProblemType: valid values are %v", v, AllowedGetRecipeDeploymentProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetRecipeDeploymentProblemType) IsValid() bool {
	for _, existing := range AllowedGetRecipeDeploymentProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetRecipeDeploymentProblemType value
func (v GetRecipeDeploymentProblemType) Ptr() *GetRecipeDeploymentProblemType {
	return &v
}

type NullableGetRecipeDeploymentProblemType struct {
	value *GetRecipeDeploymentProblemType
	isSet bool
}

func (v NullableGetRecipeDeploymentProblemType) Get() *GetRecipeDeploymentProblemType {
	return v.value
}

func (v *NullableGetRecipeDeploymentProblemType) Set(val *GetRecipeDeploymentProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecipeDeploymentProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecipeDeploymentProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecipeDeploymentProblemType(val *GetRecipeDeploymentProblemType) *NullableGetRecipeDeploymentProblemType {
	return &NullableGetRecipeDeploymentProblemType{value: val, isSet: true}
}

func (v NullableGetRecipeDeploymentProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecipeDeploymentProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

