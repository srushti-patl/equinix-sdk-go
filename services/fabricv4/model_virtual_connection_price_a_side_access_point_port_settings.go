/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the VirtualConnectionPriceASideAccessPointPortSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualConnectionPriceASideAccessPointPortSettings{}

// VirtualConnectionPriceASideAccessPointPortSettings struct for VirtualConnectionPriceASideAccessPointPortSettings
type VirtualConnectionPriceASideAccessPointPortSettings struct {
	Buyout               *bool `json:"buyout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualConnectionPriceASideAccessPointPortSettings VirtualConnectionPriceASideAccessPointPortSettings

// NewVirtualConnectionPriceASideAccessPointPortSettings instantiates a new VirtualConnectionPriceASideAccessPointPortSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualConnectionPriceASideAccessPointPortSettings() *VirtualConnectionPriceASideAccessPointPortSettings {
	this := VirtualConnectionPriceASideAccessPointPortSettings{}
	var buyout bool = false
	this.Buyout = &buyout
	return &this
}

// NewVirtualConnectionPriceASideAccessPointPortSettingsWithDefaults instantiates a new VirtualConnectionPriceASideAccessPointPortSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualConnectionPriceASideAccessPointPortSettingsWithDefaults() *VirtualConnectionPriceASideAccessPointPortSettings {
	this := VirtualConnectionPriceASideAccessPointPortSettings{}
	var buyout bool = false
	this.Buyout = &buyout
	return &this
}

// GetBuyout returns the Buyout field value if set, zero value otherwise.
func (o *VirtualConnectionPriceASideAccessPointPortSettings) GetBuyout() bool {
	if o == nil || IsNil(o.Buyout) {
		var ret bool
		return ret
	}
	return *o.Buyout
}

// GetBuyoutOk returns a tuple with the Buyout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualConnectionPriceASideAccessPointPortSettings) GetBuyoutOk() (*bool, bool) {
	if o == nil || IsNil(o.Buyout) {
		return nil, false
	}
	return o.Buyout, true
}

// HasBuyout returns a boolean if a field has been set.
func (o *VirtualConnectionPriceASideAccessPointPortSettings) HasBuyout() bool {
	if o != nil && !IsNil(o.Buyout) {
		return true
	}

	return false
}

// SetBuyout gets a reference to the given bool and assigns it to the Buyout field.
func (o *VirtualConnectionPriceASideAccessPointPortSettings) SetBuyout(v bool) {
	o.Buyout = &v
}

func (o VirtualConnectionPriceASideAccessPointPortSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualConnectionPriceASideAccessPointPortSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Buyout) {
		toSerialize["buyout"] = o.Buyout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualConnectionPriceASideAccessPointPortSettings) UnmarshalJSON(data []byte) (err error) {
	varVirtualConnectionPriceASideAccessPointPortSettings := _VirtualConnectionPriceASideAccessPointPortSettings{}

	err = json.Unmarshal(data, &varVirtualConnectionPriceASideAccessPointPortSettings)

	if err != nil {
		return err
	}

	*o = VirtualConnectionPriceASideAccessPointPortSettings(varVirtualConnectionPriceASideAccessPointPortSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "buyout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualConnectionPriceASideAccessPointPortSettings struct {
	value *VirtualConnectionPriceASideAccessPointPortSettings
	isSet bool
}

func (v NullableVirtualConnectionPriceASideAccessPointPortSettings) Get() *VirtualConnectionPriceASideAccessPointPortSettings {
	return v.value
}

func (v *NullableVirtualConnectionPriceASideAccessPointPortSettings) Set(val *VirtualConnectionPriceASideAccessPointPortSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionPriceASideAccessPointPortSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionPriceASideAccessPointPortSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionPriceASideAccessPointPortSettings(val *VirtualConnectionPriceASideAccessPointPortSettings) *NullableVirtualConnectionPriceASideAccessPointPortSettings {
	return &NullableVirtualConnectionPriceASideAccessPointPortSettings{value: val, isSet: true}
}

func (v NullableVirtualConnectionPriceASideAccessPointPortSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionPriceASideAccessPointPortSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
