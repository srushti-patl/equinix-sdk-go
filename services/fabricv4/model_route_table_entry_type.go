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

// RouteTableEntryType Route table entry type
type RouteTableEntryType string

// List of RouteTableEntryType
const (
	ROUTETABLEENTRYTYPE_IPV4_BGP_ROUTE    RouteTableEntryType = "IPv4_BGP_ROUTE"
	ROUTETABLEENTRYTYPE_IPV4_STATIC_ROUTE RouteTableEntryType = "IPv4_STATIC_ROUTE"
	ROUTETABLEENTRYTYPE_IPV4_DIRECT_ROUTE RouteTableEntryType = "IPv4_DIRECT_ROUTE"
	ROUTETABLEENTRYTYPE_IPV6_BGP_ROUTE    RouteTableEntryType = "IPv6_BGP_ROUTE"
	ROUTETABLEENTRYTYPE_IPV6_STATIC_ROUTE RouteTableEntryType = "IPv6_STATIC_ROUTE"
	ROUTETABLEENTRYTYPE_IPV6_DIRECT_ROUTE RouteTableEntryType = "IPv6_DIRECT_ROUTE"
)

// All allowed values of RouteTableEntryType enum
var AllowedRouteTableEntryTypeEnumValues = []RouteTableEntryType{
	"IPv4_BGP_ROUTE",
	"IPv4_STATIC_ROUTE",
	"IPv4_DIRECT_ROUTE",
	"IPv6_BGP_ROUTE",
	"IPv6_STATIC_ROUTE",
	"IPv6_DIRECT_ROUTE",
}

func (v *RouteTableEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteTableEntryType(value)
	for _, existing := range AllowedRouteTableEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteTableEntryType", value)
}

// NewRouteTableEntryTypeFromValue returns a pointer to a valid RouteTableEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteTableEntryTypeFromValue(v string) (*RouteTableEntryType, error) {
	ev := RouteTableEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteTableEntryType: valid values are %v", v, AllowedRouteTableEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteTableEntryType) IsValid() bool {
	for _, existing := range AllowedRouteTableEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteTableEntryType value
func (v RouteTableEntryType) Ptr() *RouteTableEntryType {
	return &v
}

type NullableRouteTableEntryType struct {
	value *RouteTableEntryType
	isSet bool
}

func (v NullableRouteTableEntryType) Get() *RouteTableEntryType {
	return v.value
}

func (v *NullableRouteTableEntryType) Set(val *RouteTableEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTableEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTableEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTableEntryType(val *RouteTableEntryType) *NullableRouteTableEntryType {
	return &NullableRouteTableEntryType{value: val, isSet: true}
}

func (v NullableRouteTableEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTableEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
