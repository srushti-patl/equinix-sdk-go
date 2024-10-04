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

// ConnectionRouteTableEntryState the model 'ConnectionRouteTableEntryState'
type ConnectionRouteTableEntryState string

// List of ConnectionRouteTableEntry_state
const (
	CONNECTIONROUTETABLEENTRYSTATE_ACTIVE   ConnectionRouteTableEntryState = "ACTIVE"
	CONNECTIONROUTETABLEENTRYSTATE_INACTIVE ConnectionRouteTableEntryState = "INACTIVE"
)

// All allowed values of ConnectionRouteTableEntryState enum
var AllowedConnectionRouteTableEntryStateEnumValues = []ConnectionRouteTableEntryState{
	"ACTIVE",
	"INACTIVE",
}

func (v *ConnectionRouteTableEntryState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionRouteTableEntryState(value)
	for _, existing := range AllowedConnectionRouteTableEntryStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionRouteTableEntryState", value)
}

// NewConnectionRouteTableEntryStateFromValue returns a pointer to a valid ConnectionRouteTableEntryState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionRouteTableEntryStateFromValue(v string) (*ConnectionRouteTableEntryState, error) {
	ev := ConnectionRouteTableEntryState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionRouteTableEntryState: valid values are %v", v, AllowedConnectionRouteTableEntryStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionRouteTableEntryState) IsValid() bool {
	for _, existing := range AllowedConnectionRouteTableEntryStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionRouteTableEntry_state value
func (v ConnectionRouteTableEntryState) Ptr() *ConnectionRouteTableEntryState {
	return &v
}

type NullableConnectionRouteTableEntryState struct {
	value *ConnectionRouteTableEntryState
	isSet bool
}

func (v NullableConnectionRouteTableEntryState) Get() *ConnectionRouteTableEntryState {
	return v.value
}

func (v *NullableConnectionRouteTableEntryState) Set(val *ConnectionRouteTableEntryState) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteTableEntryState) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteTableEntryState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteTableEntryState(val *ConnectionRouteTableEntryState) *NullableConnectionRouteTableEntryState {
	return &NullableConnectionRouteTableEntryState{value: val, isSet: true}
}

func (v NullableConnectionRouteTableEntryState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteTableEntryState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
