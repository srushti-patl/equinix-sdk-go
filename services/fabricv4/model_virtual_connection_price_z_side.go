/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the VirtualConnectionPriceZSide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualConnectionPriceZSide{}

// VirtualConnectionPriceZSide struct for VirtualConnectionPriceZSide
type VirtualConnectionPriceZSide struct {
	AccessPoint          *VirtualConnectionPriceZSideAccessPoint `json:"accessPoint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualConnectionPriceZSide VirtualConnectionPriceZSide

// NewVirtualConnectionPriceZSide instantiates a new VirtualConnectionPriceZSide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualConnectionPriceZSide() *VirtualConnectionPriceZSide {
	this := VirtualConnectionPriceZSide{}
	return &this
}

// NewVirtualConnectionPriceZSideWithDefaults instantiates a new VirtualConnectionPriceZSide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualConnectionPriceZSideWithDefaults() *VirtualConnectionPriceZSide {
	this := VirtualConnectionPriceZSide{}
	return &this
}

// GetAccessPoint returns the AccessPoint field value if set, zero value otherwise.
func (o *VirtualConnectionPriceZSide) GetAccessPoint() VirtualConnectionPriceZSideAccessPoint {
	if o == nil || IsNil(o.AccessPoint) {
		var ret VirtualConnectionPriceZSideAccessPoint
		return ret
	}
	return *o.AccessPoint
}

// GetAccessPointOk returns a tuple with the AccessPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualConnectionPriceZSide) GetAccessPointOk() (*VirtualConnectionPriceZSideAccessPoint, bool) {
	if o == nil || IsNil(o.AccessPoint) {
		return nil, false
	}
	return o.AccessPoint, true
}

// HasAccessPoint returns a boolean if a field has been set.
func (o *VirtualConnectionPriceZSide) HasAccessPoint() bool {
	if o != nil && !IsNil(o.AccessPoint) {
		return true
	}

	return false
}

// SetAccessPoint gets a reference to the given VirtualConnectionPriceZSideAccessPoint and assigns it to the AccessPoint field.
func (o *VirtualConnectionPriceZSide) SetAccessPoint(v VirtualConnectionPriceZSideAccessPoint) {
	o.AccessPoint = &v
}

func (o VirtualConnectionPriceZSide) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualConnectionPriceZSide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPoint) {
		toSerialize["accessPoint"] = o.AccessPoint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualConnectionPriceZSide) UnmarshalJSON(data []byte) (err error) {
	varVirtualConnectionPriceZSide := _VirtualConnectionPriceZSide{}

	err = json.Unmarshal(data, &varVirtualConnectionPriceZSide)

	if err != nil {
		return err
	}

	*o = VirtualConnectionPriceZSide(varVirtualConnectionPriceZSide)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessPoint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualConnectionPriceZSide struct {
	value *VirtualConnectionPriceZSide
	isSet bool
}

func (v NullableVirtualConnectionPriceZSide) Get() *VirtualConnectionPriceZSide {
	return v.value
}

func (v *NullableVirtualConnectionPriceZSide) Set(val *VirtualConnectionPriceZSide) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionPriceZSide) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionPriceZSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionPriceZSide(val *VirtualConnectionPriceZSide) *NullableVirtualConnectionPriceZSide {
	return &NullableVirtualConnectionPriceZSide{value: val, isSet: true}
}

func (v NullableVirtualConnectionPriceZSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionPriceZSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
