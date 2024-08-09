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

// VlanVirtualCircuitType the model 'VlanVirtualCircuitType'
type VlanVirtualCircuitType string

// List of VlanVirtualCircuit_type
const (
	VLANVIRTUALCIRCUITTYPE_VLAN VlanVirtualCircuitType = "vlan"
)

// All allowed values of VlanVirtualCircuitType enum
var AllowedVlanVirtualCircuitTypeEnumValues = []VlanVirtualCircuitType{
	"vlan",
}

func (v *VlanVirtualCircuitType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VlanVirtualCircuitType(value)
	for _, existing := range AllowedVlanVirtualCircuitTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VlanVirtualCircuitType", value)
}

// NewVlanVirtualCircuitTypeFromValue returns a pointer to a valid VlanVirtualCircuitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVlanVirtualCircuitTypeFromValue(v string) (*VlanVirtualCircuitType, error) {
	ev := VlanVirtualCircuitType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VlanVirtualCircuitType: valid values are %v", v, AllowedVlanVirtualCircuitTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VlanVirtualCircuitType) IsValid() bool {
	for _, existing := range AllowedVlanVirtualCircuitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VlanVirtualCircuit_type value
func (v VlanVirtualCircuitType) Ptr() *VlanVirtualCircuitType {
	return &v
}

type NullableVlanVirtualCircuitType struct {
	value *VlanVirtualCircuitType
	isSet bool
}

func (v NullableVlanVirtualCircuitType) Get() *VlanVirtualCircuitType {
	return v.value
}

func (v *NullableVlanVirtualCircuitType) Set(val *VlanVirtualCircuitType) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanVirtualCircuitType) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanVirtualCircuitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanVirtualCircuitType(val *VlanVirtualCircuitType) *NullableVlanVirtualCircuitType {
	return &NullableVlanVirtualCircuitType{value: val, isSet: true}
}

func (v NullableVlanVirtualCircuitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanVirtualCircuitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
