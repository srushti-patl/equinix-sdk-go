/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// NetworkEquinixStatus Network status
type NetworkEquinixStatus string

// List of NetworkEquinixStatus
const (
	NETWORKEQUINIXSTATUS_PROVISIONING      NetworkEquinixStatus = "PROVISIONING"
	NETWORKEQUINIXSTATUS_PROVISIONED       NetworkEquinixStatus = "PROVISIONED"
	NETWORKEQUINIXSTATUS_NOT_PROVISIONED   NetworkEquinixStatus = "NOT_PROVISIONED"
	NETWORKEQUINIXSTATUS_DEPROVISIONING    NetworkEquinixStatus = "DEPROVISIONING"
	NETWORKEQUINIXSTATUS_DEPROVISIONED     NetworkEquinixStatus = "DEPROVISIONED"
	NETWORKEQUINIXSTATUS_NOT_DEPROVISIONED NetworkEquinixStatus = "NOT_DEPROVISIONED"
)

// All allowed values of NetworkEquinixStatus enum
var AllowedNetworkEquinixStatusEnumValues = []NetworkEquinixStatus{
	"PROVISIONING",
	"PROVISIONED",
	"NOT_PROVISIONED",
	"DEPROVISIONING",
	"DEPROVISIONED",
	"NOT_DEPROVISIONED",
}

func (v *NetworkEquinixStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkEquinixStatus(value)
	for _, existing := range AllowedNetworkEquinixStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkEquinixStatus", value)
}

// NewNetworkEquinixStatusFromValue returns a pointer to a valid NetworkEquinixStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkEquinixStatusFromValue(v string) (*NetworkEquinixStatus, error) {
	ev := NetworkEquinixStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkEquinixStatus: valid values are %v", v, AllowedNetworkEquinixStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkEquinixStatus) IsValid() bool {
	for _, existing := range AllowedNetworkEquinixStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkEquinixStatus value
func (v NetworkEquinixStatus) Ptr() *NetworkEquinixStatus {
	return &v
}

type NullableNetworkEquinixStatus struct {
	value *NetworkEquinixStatus
	isSet bool
}

func (v NullableNetworkEquinixStatus) Get() *NetworkEquinixStatus {
	return v.value
}

func (v *NullableNetworkEquinixStatus) Set(val *NetworkEquinixStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkEquinixStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkEquinixStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkEquinixStatus(val *NetworkEquinixStatus) *NullableNetworkEquinixStatus {
	return &NullableNetworkEquinixStatus{value: val, isSet: true}
}

func (v NullableNetworkEquinixStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkEquinixStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
