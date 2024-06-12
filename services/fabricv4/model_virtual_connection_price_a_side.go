/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the VirtualConnectionPriceASide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualConnectionPriceASide{}

// VirtualConnectionPriceASide struct for VirtualConnectionPriceASide
type VirtualConnectionPriceASide struct {
	AccessPoint          *VirtualConnectionPriceASideAccessPoint `json:"accessPoint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualConnectionPriceASide VirtualConnectionPriceASide

// NewVirtualConnectionPriceASide instantiates a new VirtualConnectionPriceASide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualConnectionPriceASide() *VirtualConnectionPriceASide {
	this := VirtualConnectionPriceASide{}
	return &this
}

// NewVirtualConnectionPriceASideWithDefaults instantiates a new VirtualConnectionPriceASide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualConnectionPriceASideWithDefaults() *VirtualConnectionPriceASide {
	this := VirtualConnectionPriceASide{}
	return &this
}

// GetAccessPoint returns the AccessPoint field value if set, zero value otherwise.
func (o *VirtualConnectionPriceASide) GetAccessPoint() VirtualConnectionPriceASideAccessPoint {
	if o == nil || IsNil(o.AccessPoint) {
		var ret VirtualConnectionPriceASideAccessPoint
		return ret
	}
	return *o.AccessPoint
}

// GetAccessPointOk returns a tuple with the AccessPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualConnectionPriceASide) GetAccessPointOk() (*VirtualConnectionPriceASideAccessPoint, bool) {
	if o == nil || IsNil(o.AccessPoint) {
		return nil, false
	}
	return o.AccessPoint, true
}

// HasAccessPoint returns a boolean if a field has been set.
func (o *VirtualConnectionPriceASide) HasAccessPoint() bool {
	if o != nil && !IsNil(o.AccessPoint) {
		return true
	}

	return false
}

// SetAccessPoint gets a reference to the given VirtualConnectionPriceASideAccessPoint and assigns it to the AccessPoint field.
func (o *VirtualConnectionPriceASide) SetAccessPoint(v VirtualConnectionPriceASideAccessPoint) {
	o.AccessPoint = &v
}

func (o VirtualConnectionPriceASide) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualConnectionPriceASide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPoint) {
		toSerialize["accessPoint"] = o.AccessPoint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualConnectionPriceASide) UnmarshalJSON(data []byte) (err error) {
	varVirtualConnectionPriceASide := _VirtualConnectionPriceASide{}

	err = json.Unmarshal(data, &varVirtualConnectionPriceASide)

	if err != nil {
		return err
	}

	*o = VirtualConnectionPriceASide(varVirtualConnectionPriceASide)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessPoint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualConnectionPriceASide struct {
	value *VirtualConnectionPriceASide
	isSet bool
}

func (v NullableVirtualConnectionPriceASide) Get() *VirtualConnectionPriceASide {
	return v.value
}

func (v *NullableVirtualConnectionPriceASide) Set(val *VirtualConnectionPriceASide) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionPriceASide) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionPriceASide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionPriceASide(val *VirtualConnectionPriceASide) *NullableVirtualConnectionPriceASide {
	return &NullableVirtualConnectionPriceASide{value: val, isSet: true}
}

func (v NullableVirtualConnectionPriceASide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionPriceASide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
