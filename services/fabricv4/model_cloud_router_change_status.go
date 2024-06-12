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

// CloudRouterChangeStatus Current outcome of the change flow
type CloudRouterChangeStatus string

// List of CloudRouterChange_status
const (
	CLOUDROUTERCHANGESTATUS_COMPLETED CloudRouterChangeStatus = "COMPLETED"
	CLOUDROUTERCHANGESTATUS_FAILED    CloudRouterChangeStatus = "FAILED"
	CLOUDROUTERCHANGESTATUS_REQUESTED CloudRouterChangeStatus = "REQUESTED"
)

// All allowed values of CloudRouterChangeStatus enum
var AllowedCloudRouterChangeStatusEnumValues = []CloudRouterChangeStatus{
	"COMPLETED",
	"FAILED",
	"REQUESTED",
}

func (v *CloudRouterChangeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudRouterChangeStatus(value)
	for _, existing := range AllowedCloudRouterChangeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudRouterChangeStatus", value)
}

// NewCloudRouterChangeStatusFromValue returns a pointer to a valid CloudRouterChangeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudRouterChangeStatusFromValue(v string) (*CloudRouterChangeStatus, error) {
	ev := CloudRouterChangeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudRouterChangeStatus: valid values are %v", v, AllowedCloudRouterChangeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudRouterChangeStatus) IsValid() bool {
	for _, existing := range AllowedCloudRouterChangeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudRouterChange_status value
func (v CloudRouterChangeStatus) Ptr() *CloudRouterChangeStatus {
	return &v
}

type NullableCloudRouterChangeStatus struct {
	value *CloudRouterChangeStatus
	isSet bool
}

func (v NullableCloudRouterChangeStatus) Get() *CloudRouterChangeStatus {
	return v.value
}

func (v *NullableCloudRouterChangeStatus) Set(val *CloudRouterChangeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterChangeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterChangeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterChangeStatus(val *CloudRouterChangeStatus) *NullableCloudRouterChangeStatus {
	return &NullableCloudRouterChangeStatus{value: val, isSet: true}
}

func (v NullableCloudRouterChangeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterChangeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
