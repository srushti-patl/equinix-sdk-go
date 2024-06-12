/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PortSearchFieldName Possible field names to use on filters
type PortSearchFieldName string

// List of PortSearchFieldName
const (
	PORTSEARCHFIELDNAME_PROJECT_PROJECT_ID    PortSearchFieldName = "/project/projectId"
	PORTSEARCHFIELDNAME_SETTINGS_PRODUCT_CODE PortSearchFieldName = "/settings/productCode"
	PORTSEARCHFIELDNAME_STATE                 PortSearchFieldName = "/state"
)

// All allowed values of PortSearchFieldName enum
var AllowedPortSearchFieldNameEnumValues = []PortSearchFieldName{
	"/project/projectId",
	"/settings/productCode",
	"/state",
}

func (v *PortSearchFieldName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortSearchFieldName(value)
	for _, existing := range AllowedPortSearchFieldNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortSearchFieldName", value)
}

// NewPortSearchFieldNameFromValue returns a pointer to a valid PortSearchFieldName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortSearchFieldNameFromValue(v string) (*PortSearchFieldName, error) {
	ev := PortSearchFieldName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortSearchFieldName: valid values are %v", v, AllowedPortSearchFieldNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortSearchFieldName) IsValid() bool {
	for _, existing := range AllowedPortSearchFieldNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortSearchFieldName value
func (v PortSearchFieldName) Ptr() *PortSearchFieldName {
	return &v
}

type NullablePortSearchFieldName struct {
	value *PortSearchFieldName
	isSet bool
}

func (v NullablePortSearchFieldName) Get() *PortSearchFieldName {
	return v.value
}

func (v *NullablePortSearchFieldName) Set(val *PortSearchFieldName) {
	v.value = val
	v.isSet = true
}

func (v NullablePortSearchFieldName) IsSet() bool {
	return v.isSet
}

func (v *NullablePortSearchFieldName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortSearchFieldName(val *PortSearchFieldName) *NullablePortSearchFieldName {
	return &NullablePortSearchFieldName{value: val, isSet: true}
}

func (v NullablePortSearchFieldName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortSearchFieldName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
