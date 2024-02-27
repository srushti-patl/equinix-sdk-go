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

// NetworkScope Network scope
type NetworkScope string

// List of NetworkScope
const (
	NETWORKSCOPE_REGIONAL NetworkScope = "REGIONAL"
	NETWORKSCOPE_GLOBAL   NetworkScope = "GLOBAL"
	NETWORKSCOPE_LOCAL    NetworkScope = "LOCAL"
)

// All allowed values of NetworkScope enum
var AllowedNetworkScopeEnumValues = []NetworkScope{
	"REGIONAL",
	"GLOBAL",
	"LOCAL",
}

func (v *NetworkScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkScope(value)
	for _, existing := range AllowedNetworkScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkScope", value)
}

// NewNetworkScopeFromValue returns a pointer to a valid NetworkScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkScopeFromValue(v string) (*NetworkScope, error) {
	ev := NetworkScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkScope: valid values are %v", v, AllowedNetworkScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkScope) IsValid() bool {
	for _, existing := range AllowedNetworkScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkScope value
func (v NetworkScope) Ptr() *NetworkScope {
	return &v
}

type NullableNetworkScope struct {
	value *NetworkScope
	isSet bool
}

func (v NullableNetworkScope) Get() *NetworkScope {
	return v.value
}

func (v *NullableNetworkScope) Set(val *NetworkScope) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkScope) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkScope(val *NetworkScope) *NullableNetworkScope {
	return &NullableNetworkScope{value: val, isSet: true}
}

func (v NullableNetworkScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
