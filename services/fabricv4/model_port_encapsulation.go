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
)

// checks if the PortEncapsulation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortEncapsulation{}

// PortEncapsulation Port encapsulation configuration
type PortEncapsulation struct {
	Type *PortEncapsulationType `json:"type,omitempty"`
	// Port encapsulation tag protocol identifier
	TagProtocolId        *string `json:"tagProtocolId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortEncapsulation PortEncapsulation

// NewPortEncapsulation instantiates a new PortEncapsulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortEncapsulation() *PortEncapsulation {
	this := PortEncapsulation{}
	return &this
}

// NewPortEncapsulationWithDefaults instantiates a new PortEncapsulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortEncapsulationWithDefaults() *PortEncapsulation {
	this := PortEncapsulation{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PortEncapsulation) GetType() PortEncapsulationType {
	if o == nil || IsNil(o.Type) {
		var ret PortEncapsulationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEncapsulation) GetTypeOk() (*PortEncapsulationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PortEncapsulation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PortEncapsulationType and assigns it to the Type field.
func (o *PortEncapsulation) SetType(v PortEncapsulationType) {
	o.Type = &v
}

// GetTagProtocolId returns the TagProtocolId field value if set, zero value otherwise.
func (o *PortEncapsulation) GetTagProtocolId() string {
	if o == nil || IsNil(o.TagProtocolId) {
		var ret string
		return ret
	}
	return *o.TagProtocolId
}

// GetTagProtocolIdOk returns a tuple with the TagProtocolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEncapsulation) GetTagProtocolIdOk() (*string, bool) {
	if o == nil || IsNil(o.TagProtocolId) {
		return nil, false
	}
	return o.TagProtocolId, true
}

// HasTagProtocolId returns a boolean if a field has been set.
func (o *PortEncapsulation) HasTagProtocolId() bool {
	if o != nil && !IsNil(o.TagProtocolId) {
		return true
	}

	return false
}

// SetTagProtocolId gets a reference to the given string and assigns it to the TagProtocolId field.
func (o *PortEncapsulation) SetTagProtocolId(v string) {
	o.TagProtocolId = &v
}

func (o PortEncapsulation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortEncapsulation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TagProtocolId) {
		toSerialize["tagProtocolId"] = o.TagProtocolId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortEncapsulation) UnmarshalJSON(data []byte) (err error) {
	varPortEncapsulation := _PortEncapsulation{}

	err = json.Unmarshal(data, &varPortEncapsulation)

	if err != nil {
		return err
	}

	*o = PortEncapsulation(varPortEncapsulation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "tagProtocolId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortEncapsulation struct {
	value *PortEncapsulation
	isSet bool
}

func (v NullablePortEncapsulation) Get() *PortEncapsulation {
	return v.value
}

func (v *NullablePortEncapsulation) Set(val *PortEncapsulation) {
	v.value = val
	v.isSet = true
}

func (v NullablePortEncapsulation) IsSet() bool {
	return v.isSet
}

func (v *NullablePortEncapsulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortEncapsulation(val *PortEncapsulation) *NullablePortEncapsulation {
	return &NullablePortEncapsulation{value: val, isSet: true}
}

func (v NullablePortEncapsulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortEncapsulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
