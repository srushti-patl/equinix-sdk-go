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

// RouteTableEntryProtocolType Route table entry protocol type
type RouteTableEntryProtocolType string

// List of RouteTableEntryProtocolType
const (
	ROUTETABLEENTRYPROTOCOLTYPE_BGP    RouteTableEntryProtocolType = "BGP"
	ROUTETABLEENTRYPROTOCOLTYPE_STATIC RouteTableEntryProtocolType = "STATIC"
	ROUTETABLEENTRYPROTOCOLTYPE_DIRECT RouteTableEntryProtocolType = "DIRECT"
)

// All allowed values of RouteTableEntryProtocolType enum
var AllowedRouteTableEntryProtocolTypeEnumValues = []RouteTableEntryProtocolType{
	"BGP",
	"STATIC",
	"DIRECT",
}

func (v *RouteTableEntryProtocolType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteTableEntryProtocolType(value)
	for _, existing := range AllowedRouteTableEntryProtocolTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteTableEntryProtocolType", value)
}

// NewRouteTableEntryProtocolTypeFromValue returns a pointer to a valid RouteTableEntryProtocolType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteTableEntryProtocolTypeFromValue(v string) (*RouteTableEntryProtocolType, error) {
	ev := RouteTableEntryProtocolType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteTableEntryProtocolType: valid values are %v", v, AllowedRouteTableEntryProtocolTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteTableEntryProtocolType) IsValid() bool {
	for _, existing := range AllowedRouteTableEntryProtocolTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteTableEntryProtocolType value
func (v RouteTableEntryProtocolType) Ptr() *RouteTableEntryProtocolType {
	return &v
}

type NullableRouteTableEntryProtocolType struct {
	value *RouteTableEntryProtocolType
	isSet bool
}

func (v NullableRouteTableEntryProtocolType) Get() *RouteTableEntryProtocolType {
	return v.value
}

func (v *NullableRouteTableEntryProtocolType) Set(val *RouteTableEntryProtocolType) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTableEntryProtocolType) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTableEntryProtocolType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTableEntryProtocolType(val *RouteTableEntryProtocolType) *NullableRouteTableEntryProtocolType {
	return &NullableRouteTableEntryProtocolType{value: val, isSet: true}
}

func (v NullableRouteTableEntryProtocolType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTableEntryProtocolType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
