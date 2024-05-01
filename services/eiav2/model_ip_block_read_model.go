/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the IpBlockReadModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpBlockReadModel{}

// IpBlockReadModel struct for IpBlockReadModel
type IpBlockReadModel struct {
	// Ip block URI
	Href                 *string              `json:"href,omitempty"`
	Uuid                 *string              `json:"uuid,omitempty"`
	Type                 IpBlockReadModelType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _IpBlockReadModel IpBlockReadModel

// NewIpBlockReadModel instantiates a new IpBlockReadModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpBlockReadModel(type_ IpBlockReadModelType) *IpBlockReadModel {
	this := IpBlockReadModel{}
	this.Type = type_
	return &this
}

// NewIpBlockReadModelWithDefaults instantiates a new IpBlockReadModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpBlockReadModelWithDefaults() *IpBlockReadModel {
	this := IpBlockReadModel{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *IpBlockReadModel) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlockReadModel) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *IpBlockReadModel) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *IpBlockReadModel) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *IpBlockReadModel) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlockReadModel) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *IpBlockReadModel) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *IpBlockReadModel) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value
func (o *IpBlockReadModel) GetType() IpBlockReadModelType {
	if o == nil {
		var ret IpBlockReadModelType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IpBlockReadModel) GetTypeOk() (*IpBlockReadModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IpBlockReadModel) SetType(v IpBlockReadModelType) {
	o.Type = v
}

func (o IpBlockReadModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpBlockReadModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IpBlockReadModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIpBlockReadModel := _IpBlockReadModel{}

	err = json.Unmarshal(data, &varIpBlockReadModel)

	if err != nil {
		return err
	}

	*o = IpBlockReadModel(varIpBlockReadModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpBlockReadModel struct {
	value *IpBlockReadModel
	isSet bool
}

func (v NullableIpBlockReadModel) Get() *IpBlockReadModel {
	return v.value
}

func (v *NullableIpBlockReadModel) Set(val *IpBlockReadModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlockReadModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlockReadModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlockReadModel(val *IpBlockReadModel) *NullableIpBlockReadModel {
	return &NullableIpBlockReadModel{value: val, isSet: true}
}

func (v NullableIpBlockReadModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlockReadModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
