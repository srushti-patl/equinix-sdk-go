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

// RouteTableEntryFilter struct for RouteTableEntryFilter
type RouteTableEntryFilter struct {
	RouteTableEntryOrFilter         *RouteTableEntryOrFilter
	RouteTableEntrySimpleExpression *RouteTableEntrySimpleExpression
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RouteTableEntryFilter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RouteTableEntryOrFilter
	err = json.Unmarshal(data, &dst.RouteTableEntryOrFilter)
	if err == nil {
		jsonRouteTableEntryOrFilter, _ := json.Marshal(dst.RouteTableEntryOrFilter)
		if string(jsonRouteTableEntryOrFilter) == "{}" { // empty struct
			dst.RouteTableEntryOrFilter = nil
		} else {
			return nil // data stored in dst.RouteTableEntryOrFilter, return on the first match
		}
	} else {
		dst.RouteTableEntryOrFilter = nil
	}

	// try to unmarshal JSON data into RouteTableEntrySimpleExpression
	err = json.Unmarshal(data, &dst.RouteTableEntrySimpleExpression)
	if err == nil {
		jsonRouteTableEntrySimpleExpression, _ := json.Marshal(dst.RouteTableEntrySimpleExpression)
		if string(jsonRouteTableEntrySimpleExpression) == "{}" { // empty struct
			dst.RouteTableEntrySimpleExpression = nil
		} else {
			return nil // data stored in dst.RouteTableEntrySimpleExpression, return on the first match
		}
	} else {
		dst.RouteTableEntrySimpleExpression = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RouteTableEntryFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RouteTableEntryFilter) MarshalJSON() ([]byte, error) {
	if src.RouteTableEntryOrFilter != nil {
		return json.Marshal(&src.RouteTableEntryOrFilter)
	}

	if src.RouteTableEntrySimpleExpression != nil {
		return json.Marshal(&src.RouteTableEntrySimpleExpression)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRouteTableEntryFilter struct {
	value *RouteTableEntryFilter
	isSet bool
}

func (v NullableRouteTableEntryFilter) Get() *RouteTableEntryFilter {
	return v.value
}

func (v *NullableRouteTableEntryFilter) Set(val *RouteTableEntryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTableEntryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTableEntryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTableEntryFilter(val *RouteTableEntryFilter) *NullableRouteTableEntryFilter {
	return &NullableRouteTableEntryFilter{value: val, isSet: true}
}

func (v NullableRouteTableEntryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTableEntryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
