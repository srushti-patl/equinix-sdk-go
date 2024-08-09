/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the BgpRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpRoute{}

// BgpRoute struct for BgpRoute
type BgpRoute struct {
	Exact                *bool   `json:"exact,omitempty"`
	Route                *string `json:"route,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpRoute BgpRoute

// NewBgpRoute instantiates a new BgpRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpRoute() *BgpRoute {
	this := BgpRoute{}
	return &this
}

// NewBgpRouteWithDefaults instantiates a new BgpRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpRouteWithDefaults() *BgpRoute {
	this := BgpRoute{}
	return &this
}

// GetExact returns the Exact field value if set, zero value otherwise.
func (o *BgpRoute) GetExact() bool {
	if o == nil || IsNil(o.Exact) {
		var ret bool
		return ret
	}
	return *o.Exact
}

// GetExactOk returns a tuple with the Exact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpRoute) GetExactOk() (*bool, bool) {
	if o == nil || IsNil(o.Exact) {
		return nil, false
	}
	return o.Exact, true
}

// HasExact returns a boolean if a field has been set.
func (o *BgpRoute) HasExact() bool {
	if o != nil && !IsNil(o.Exact) {
		return true
	}

	return false
}

// SetExact gets a reference to the given bool and assigns it to the Exact field.
func (o *BgpRoute) SetExact(v bool) {
	o.Exact = &v
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *BgpRoute) GetRoute() string {
	if o == nil || IsNil(o.Route) {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpRoute) GetRouteOk() (*string, bool) {
	if o == nil || IsNil(o.Route) {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *BgpRoute) HasRoute() bool {
	if o != nil && !IsNil(o.Route) {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *BgpRoute) SetRoute(v string) {
	o.Route = &v
}

func (o BgpRoute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exact) {
		toSerialize["exact"] = o.Exact
	}
	if !IsNil(o.Route) {
		toSerialize["route"] = o.Route
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpRoute) UnmarshalJSON(data []byte) (err error) {
	varBgpRoute := _BgpRoute{}

	err = json.Unmarshal(data, &varBgpRoute)

	if err != nil {
		return err
	}

	*o = BgpRoute(varBgpRoute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exact")
		delete(additionalProperties, "route")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpRoute struct {
	value *BgpRoute
	isSet bool
}

func (v NullableBgpRoute) Get() *BgpRoute {
	return v.value
}

func (v *NullableBgpRoute) Set(val *BgpRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpRoute(val *BgpRoute) *NullableBgpRoute {
	return &NullableBgpRoute{value: val, isSet: true}
}

func (v NullableBgpRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
