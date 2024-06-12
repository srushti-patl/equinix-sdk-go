/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the DirectConnectionIpv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectConnectionIpv4{}

// DirectConnectionIpv4 struct for DirectConnectionIpv4
type DirectConnectionIpv4 struct {
	// Equinix side Interface IP address
	EquinixIfaceIp       *string `json:"equinixIfaceIp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DirectConnectionIpv4 DirectConnectionIpv4

// NewDirectConnectionIpv4 instantiates a new DirectConnectionIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectConnectionIpv4() *DirectConnectionIpv4 {
	this := DirectConnectionIpv4{}
	return &this
}

// NewDirectConnectionIpv4WithDefaults instantiates a new DirectConnectionIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectConnectionIpv4WithDefaults() *DirectConnectionIpv4 {
	this := DirectConnectionIpv4{}
	return &this
}

// GetEquinixIfaceIp returns the EquinixIfaceIp field value if set, zero value otherwise.
func (o *DirectConnectionIpv4) GetEquinixIfaceIp() string {
	if o == nil || IsNil(o.EquinixIfaceIp) {
		var ret string
		return ret
	}
	return *o.EquinixIfaceIp
}

// GetEquinixIfaceIpOk returns a tuple with the EquinixIfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectConnectionIpv4) GetEquinixIfaceIpOk() (*string, bool) {
	if o == nil || IsNil(o.EquinixIfaceIp) {
		return nil, false
	}
	return o.EquinixIfaceIp, true
}

// HasEquinixIfaceIp returns a boolean if a field has been set.
func (o *DirectConnectionIpv4) HasEquinixIfaceIp() bool {
	if o != nil && !IsNil(o.EquinixIfaceIp) {
		return true
	}

	return false
}

// SetEquinixIfaceIp gets a reference to the given string and assigns it to the EquinixIfaceIp field.
func (o *DirectConnectionIpv4) SetEquinixIfaceIp(v string) {
	o.EquinixIfaceIp = &v
}

func (o DirectConnectionIpv4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectConnectionIpv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EquinixIfaceIp) {
		toSerialize["equinixIfaceIp"] = o.EquinixIfaceIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DirectConnectionIpv4) UnmarshalJSON(data []byte) (err error) {
	varDirectConnectionIpv4 := _DirectConnectionIpv4{}

	err = json.Unmarshal(data, &varDirectConnectionIpv4)

	if err != nil {
		return err
	}

	*o = DirectConnectionIpv4(varDirectConnectionIpv4)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "equinixIfaceIp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDirectConnectionIpv4 struct {
	value *DirectConnectionIpv4
	isSet bool
}

func (v NullableDirectConnectionIpv4) Get() *DirectConnectionIpv4 {
	return v.value
}

func (v *NullableDirectConnectionIpv4) Set(val *DirectConnectionIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectConnectionIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectConnectionIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectConnectionIpv4(val *DirectConnectionIpv4) *NullableDirectConnectionIpv4 {
	return &NullableDirectConnectionIpv4{value: val, isSet: true}
}

func (v NullableDirectConnectionIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectConnectionIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
