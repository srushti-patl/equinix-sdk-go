/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// BgpConfigRequestInputDeploymentType Wether the BGP deployment is local or global. Local deployments are configured immediately. Global deployments will need to be reviewed by Equinix Metal engineers.
type BgpConfigRequestInputDeploymentType string

// List of BgpConfigRequestInput_deployment_type
const (
	BGPCONFIGREQUESTINPUTDEPLOYMENTTYPE_LOCAL  BgpConfigRequestInputDeploymentType = "local"
	BGPCONFIGREQUESTINPUTDEPLOYMENTTYPE_GLOBAL BgpConfigRequestInputDeploymentType = "global"
)

// All allowed values of BgpConfigRequestInputDeploymentType enum
var AllowedBgpConfigRequestInputDeploymentTypeEnumValues = []BgpConfigRequestInputDeploymentType{
	"local",
	"global",
}

func (v *BgpConfigRequestInputDeploymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BgpConfigRequestInputDeploymentType(value)
	for _, existing := range AllowedBgpConfigRequestInputDeploymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BgpConfigRequestInputDeploymentType", value)
}

// NewBgpConfigRequestInputDeploymentTypeFromValue returns a pointer to a valid BgpConfigRequestInputDeploymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBgpConfigRequestInputDeploymentTypeFromValue(v string) (*BgpConfigRequestInputDeploymentType, error) {
	ev := BgpConfigRequestInputDeploymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BgpConfigRequestInputDeploymentType: valid values are %v", v, AllowedBgpConfigRequestInputDeploymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BgpConfigRequestInputDeploymentType) IsValid() bool {
	for _, existing := range AllowedBgpConfigRequestInputDeploymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BgpConfigRequestInput_deployment_type value
func (v BgpConfigRequestInputDeploymentType) Ptr() *BgpConfigRequestInputDeploymentType {
	return &v
}

type NullableBgpConfigRequestInputDeploymentType struct {
	value *BgpConfigRequestInputDeploymentType
	isSet bool
}

func (v NullableBgpConfigRequestInputDeploymentType) Get() *BgpConfigRequestInputDeploymentType {
	return v.value
}

func (v *NullableBgpConfigRequestInputDeploymentType) Set(val *BgpConfigRequestInputDeploymentType) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfigRequestInputDeploymentType) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfigRequestInputDeploymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfigRequestInputDeploymentType(val *BgpConfigRequestInputDeploymentType) *NullableBgpConfigRequestInputDeploymentType {
	return &NullableBgpConfigRequestInputDeploymentType{value: val, isSet: true}
}

func (v NullableBgpConfigRequestInputDeploymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfigRequestInputDeploymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
