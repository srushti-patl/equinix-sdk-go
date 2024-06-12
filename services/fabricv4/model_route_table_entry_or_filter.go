/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the RouteTableEntryOrFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteTableEntryOrFilter{}

// RouteTableEntryOrFilter struct for RouteTableEntryOrFilter
type RouteTableEntryOrFilter struct {
	Or                   []RouteTableEntrySimpleExpression `json:"or,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteTableEntryOrFilter RouteTableEntryOrFilter

// NewRouteTableEntryOrFilter instantiates a new RouteTableEntryOrFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteTableEntryOrFilter() *RouteTableEntryOrFilter {
	this := RouteTableEntryOrFilter{}
	return &this
}

// NewRouteTableEntryOrFilterWithDefaults instantiates a new RouteTableEntryOrFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteTableEntryOrFilterWithDefaults() *RouteTableEntryOrFilter {
	this := RouteTableEntryOrFilter{}
	return &this
}

// GetOr returns the Or field value if set, zero value otherwise.
func (o *RouteTableEntryOrFilter) GetOr() []RouteTableEntrySimpleExpression {
	if o == nil || IsNil(o.Or) {
		var ret []RouteTableEntrySimpleExpression
		return ret
	}
	return o.Or
}

// GetOrOk returns a tuple with the Or field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntryOrFilter) GetOrOk() ([]RouteTableEntrySimpleExpression, bool) {
	if o == nil || IsNil(o.Or) {
		return nil, false
	}
	return o.Or, true
}

// HasOr returns a boolean if a field has been set.
func (o *RouteTableEntryOrFilter) HasOr() bool {
	if o != nil && !IsNil(o.Or) {
		return true
	}

	return false
}

// SetOr gets a reference to the given []RouteTableEntrySimpleExpression and assigns it to the Or field.
func (o *RouteTableEntryOrFilter) SetOr(v []RouteTableEntrySimpleExpression) {
	o.Or = v
}

func (o RouteTableEntryOrFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteTableEntryOrFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Or) {
		toSerialize["or"] = o.Or
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteTableEntryOrFilter) UnmarshalJSON(data []byte) (err error) {
	varRouteTableEntryOrFilter := _RouteTableEntryOrFilter{}

	err = json.Unmarshal(data, &varRouteTableEntryOrFilter)

	if err != nil {
		return err
	}

	*o = RouteTableEntryOrFilter(varRouteTableEntryOrFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "or")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteTableEntryOrFilter struct {
	value *RouteTableEntryOrFilter
	isSet bool
}

func (v NullableRouteTableEntryOrFilter) Get() *RouteTableEntryOrFilter {
	return v.value
}

func (v *NullableRouteTableEntryOrFilter) Set(val *RouteTableEntryOrFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTableEntryOrFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTableEntryOrFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTableEntryOrFilter(val *RouteTableEntryOrFilter) *NullableRouteTableEntryOrFilter {
	return &NullableRouteTableEntryOrFilter{value: val, isSet: true}
}

func (v NullableRouteTableEntryOrFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTableEntryOrFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
