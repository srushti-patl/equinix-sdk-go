/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the Filesystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filesystem{}

// Filesystem struct for Filesystem
type Filesystem struct {
	Mount                *Mount `json:"mount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Filesystem Filesystem

// NewFilesystem instantiates a new Filesystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystem() *Filesystem {
	this := Filesystem{}
	return &this
}

// NewFilesystemWithDefaults instantiates a new Filesystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemWithDefaults() *Filesystem {
	this := Filesystem{}
	return &this
}

// GetMount returns the Mount field value if set, zero value otherwise.
func (o *Filesystem) GetMount() Mount {
	if o == nil || IsNil(o.Mount) {
		var ret Mount
		return ret
	}
	return *o.Mount
}

// GetMountOk returns a tuple with the Mount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filesystem) GetMountOk() (*Mount, bool) {
	if o == nil || IsNil(o.Mount) {
		return nil, false
	}
	return o.Mount, true
}

// HasMount returns a boolean if a field has been set.
func (o *Filesystem) HasMount() bool {
	if o != nil && !IsNil(o.Mount) {
		return true
	}

	return false
}

// SetMount gets a reference to the given Mount and assigns it to the Mount field.
func (o *Filesystem) SetMount(v Mount) {
	o.Mount = &v
}

func (o Filesystem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filesystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mount) {
		toSerialize["mount"] = o.Mount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Filesystem) UnmarshalJSON(data []byte) (err error) {
	varFilesystem := _Filesystem{}

	err = json.Unmarshal(data, &varFilesystem)

	if err != nil {
		return err
	}

	*o = Filesystem(varFilesystem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesystem struct {
	value *Filesystem
	isSet bool
}

func (v NullableFilesystem) Get() *Filesystem {
	return v.value
}

func (v *NullableFilesystem) Set(val *Filesystem) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystem) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystem(val *Filesystem) *NullableFilesystem {
	return &NullableFilesystem{value: val, isSet: true}
}

func (v NullableFilesystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
