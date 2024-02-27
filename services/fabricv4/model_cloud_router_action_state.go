/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// CloudRouterActionState Cloud Router action state
type CloudRouterActionState string

// List of CloudRouterActionState
const (
	CLOUDROUTERACTIONSTATE_DONE    CloudRouterActionState = "DONE"
	CLOUDROUTERACTIONSTATE_FAILED  CloudRouterActionState = "FAILED"
	CLOUDROUTERACTIONSTATE_PENDING CloudRouterActionState = "PENDING"
)

// All allowed values of CloudRouterActionState enum
var AllowedCloudRouterActionStateEnumValues = []CloudRouterActionState{
	"DONE",
	"FAILED",
	"PENDING",
}

func (v *CloudRouterActionState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudRouterActionState(value)
	for _, existing := range AllowedCloudRouterActionStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudRouterActionState", value)
}

// NewCloudRouterActionStateFromValue returns a pointer to a valid CloudRouterActionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudRouterActionStateFromValue(v string) (*CloudRouterActionState, error) {
	ev := CloudRouterActionState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudRouterActionState: valid values are %v", v, AllowedCloudRouterActionStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudRouterActionState) IsValid() bool {
	for _, existing := range AllowedCloudRouterActionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudRouterActionState value
func (v CloudRouterActionState) Ptr() *CloudRouterActionState {
	return &v
}

type NullableCloudRouterActionState struct {
	value *CloudRouterActionState
	isSet bool
}

func (v NullableCloudRouterActionState) Get() *CloudRouterActionState {
	return v.value
}

func (v *NullableCloudRouterActionState) Set(val *CloudRouterActionState) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterActionState) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterActionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterActionState(val *CloudRouterActionState) *NullableCloudRouterActionState {
	return &NullableCloudRouterActionState{value: val, isSet: true}
}

func (v NullableCloudRouterActionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterActionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
