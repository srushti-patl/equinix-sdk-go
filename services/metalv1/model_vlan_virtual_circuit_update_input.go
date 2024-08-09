/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the VlanVirtualCircuitUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VlanVirtualCircuitUpdateInput{}

// VlanVirtualCircuitUpdateInput struct for VlanVirtualCircuitUpdateInput
type VlanVirtualCircuitUpdateInput struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	// Speed can be changed only if it is an interconnection on a Dedicated Port
	Speed *string  `json:"speed,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	// A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project.
	Vnid                 *string `json:"vnid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VlanVirtualCircuitUpdateInput VlanVirtualCircuitUpdateInput

// NewVlanVirtualCircuitUpdateInput instantiates a new VlanVirtualCircuitUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlanVirtualCircuitUpdateInput() *VlanVirtualCircuitUpdateInput {
	this := VlanVirtualCircuitUpdateInput{}
	return &this
}

// NewVlanVirtualCircuitUpdateInputWithDefaults instantiates a new VlanVirtualCircuitUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanVirtualCircuitUpdateInputWithDefaults() *VlanVirtualCircuitUpdateInput {
	this := VlanVirtualCircuitUpdateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VlanVirtualCircuitUpdateInput) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VlanVirtualCircuitUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *VlanVirtualCircuitUpdateInput) SetSpeed(v string) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VlanVirtualCircuitUpdateInput) SetTags(v []string) {
	o.Tags = v
}

// GetVnid returns the Vnid field value if set, zero value otherwise.
func (o *VlanVirtualCircuitUpdateInput) GetVnid() string {
	if o == nil || IsNil(o.Vnid) {
		var ret string
		return ret
	}
	return *o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitUpdateInput) GetVnidOk() (*string, bool) {
	if o == nil || IsNil(o.Vnid) {
		return nil, false
	}
	return o.Vnid, true
}

// HasVnid returns a boolean if a field has been set.
func (o *VlanVirtualCircuitUpdateInput) HasVnid() bool {
	if o != nil && !IsNil(o.Vnid) {
		return true
	}

	return false
}

// SetVnid gets a reference to the given string and assigns it to the Vnid field.
func (o *VlanVirtualCircuitUpdateInput) SetVnid(v string) {
	o.Vnid = &v
}

func (o VlanVirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VlanVirtualCircuitUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Vnid) {
		toSerialize["vnid"] = o.Vnid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VlanVirtualCircuitUpdateInput) UnmarshalJSON(data []byte) (err error) {
	varVlanVirtualCircuitUpdateInput := _VlanVirtualCircuitUpdateInput{}

	err = json.Unmarshal(data, &varVlanVirtualCircuitUpdateInput)

	if err != nil {
		return err
	}

	*o = VlanVirtualCircuitUpdateInput(varVlanVirtualCircuitUpdateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "vnid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVlanVirtualCircuitUpdateInput struct {
	value *VlanVirtualCircuitUpdateInput
	isSet bool
}

func (v NullableVlanVirtualCircuitUpdateInput) Get() *VlanVirtualCircuitUpdateInput {
	return v.value
}

func (v *NullableVlanVirtualCircuitUpdateInput) Set(val *VlanVirtualCircuitUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanVirtualCircuitUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanVirtualCircuitUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanVirtualCircuitUpdateInput(val *VlanVirtualCircuitUpdateInput) *NullableVlanVirtualCircuitUpdateInput {
	return &NullableVlanVirtualCircuitUpdateInput{value: val, isSet: true}
}

func (v NullableVlanVirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanVirtualCircuitUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
