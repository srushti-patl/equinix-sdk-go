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

// OpEnum the operation to be performed
type OpEnum string

// List of OpEnum
const (
	OPENUM_ADD     OpEnum = "add"
	OPENUM_REMOVE  OpEnum = "remove"
	OPENUM_REPLACE OpEnum = "replace"
)

// All allowed values of OpEnum enum
var AllowedOpEnumEnumValues = []OpEnum{
	"add",
	"remove",
	"replace",
}

func (v *OpEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OpEnum(value)
	for _, existing := range AllowedOpEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OpEnum", value)
}

// NewOpEnumFromValue returns a pointer to a valid OpEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOpEnumFromValue(v string) (*OpEnum, error) {
	ev := OpEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OpEnum: valid values are %v", v, AllowedOpEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OpEnum) IsValid() bool {
	for _, existing := range AllowedOpEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpEnum value
func (v OpEnum) Ptr() *OpEnum {
	return &v
}

type NullableOpEnum struct {
	value *OpEnum
	isSet bool
}

func (v NullableOpEnum) Get() *OpEnum {
	return v.value
}

func (v *NullableOpEnum) Set(val *OpEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOpEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOpEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpEnum(val *OpEnum) *NullableOpEnum {
	return &NullableOpEnum{value: val, isSet: true}
}

func (v NullableOpEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
