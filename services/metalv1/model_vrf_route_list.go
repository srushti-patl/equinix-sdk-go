/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the VrfRouteList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfRouteList{}

// VrfRouteList struct for VrfRouteList
type VrfRouteList struct {
	Routes               []VrfRoute `json:"routes,omitempty"`
	Meta                 *Meta      `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfRouteList VrfRouteList

// NewVrfRouteList instantiates a new VrfRouteList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfRouteList() *VrfRouteList {
	this := VrfRouteList{}
	return &this
}

// NewVrfRouteListWithDefaults instantiates a new VrfRouteList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfRouteListWithDefaults() *VrfRouteList {
	this := VrfRouteList{}
	return &this
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *VrfRouteList) GetRoutes() []VrfRoute {
	if o == nil || IsNil(o.Routes) {
		var ret []VrfRoute
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRouteList) GetRoutesOk() ([]VrfRoute, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *VrfRouteList) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []VrfRoute and assigns it to the Routes field.
func (o *VrfRouteList) SetRoutes(v []VrfRoute) {
	o.Routes = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *VrfRouteList) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfRouteList) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *VrfRouteList) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *VrfRouteList) SetMeta(v Meta) {
	o.Meta = &v
}

func (o VrfRouteList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfRouteList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfRouteList) UnmarshalJSON(data []byte) (err error) {
	varVrfRouteList := _VrfRouteList{}

	err = json.Unmarshal(data, &varVrfRouteList)

	if err != nil {
		return err
	}

	*o = VrfRouteList(varVrfRouteList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "routes")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfRouteList struct {
	value *VrfRouteList
	isSet bool
}

func (v NullableVrfRouteList) Get() *VrfRouteList {
	return v.value
}

func (v *NullableVrfRouteList) Set(val *VrfRouteList) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfRouteList) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfRouteList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfRouteList(val *VrfRouteList) *NullableVrfRouteList {
	return &NullableVrfRouteList{value: val, isSet: true}
}

func (v NullableVrfRouteList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfRouteList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
