/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the FabricConnectionUuid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricConnectionUuid{}

// FabricConnectionUuid UUID of the Fabric Connection Instance
type FabricConnectionUuid struct {
	// uuid of the Fabric L2 connection
	Uuid string `json:"uuid"`
	// the href for the L2 connection
	Href                 *string `json:"href,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricConnectionUuid FabricConnectionUuid

// NewFabricConnectionUuid instantiates a new FabricConnectionUuid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricConnectionUuid(uuid string) *FabricConnectionUuid {
	this := FabricConnectionUuid{}
	this.Uuid = uuid
	return &this
}

// NewFabricConnectionUuidWithDefaults instantiates a new FabricConnectionUuid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricConnectionUuidWithDefaults() *FabricConnectionUuid {
	this := FabricConnectionUuid{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *FabricConnectionUuid) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *FabricConnectionUuid) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *FabricConnectionUuid) SetUuid(v string) {
	o.Uuid = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FabricConnectionUuid) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricConnectionUuid) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FabricConnectionUuid) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FabricConnectionUuid) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FabricConnectionUuid) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricConnectionUuid) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FabricConnectionUuid) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FabricConnectionUuid) SetType(v string) {
	o.Type = &v
}

func (o FabricConnectionUuid) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricConnectionUuid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricConnectionUuid) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
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

	varFabricConnectionUuid := _FabricConnectionUuid{}

	err = json.Unmarshal(data, &varFabricConnectionUuid)

	if err != nil {
		return err
	}

	*o = FabricConnectionUuid(varFabricConnectionUuid)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricConnectionUuid struct {
	value *FabricConnectionUuid
	isSet bool
}

func (v NullableFabricConnectionUuid) Get() *FabricConnectionUuid {
	return v.value
}

func (v *NullableFabricConnectionUuid) Set(val *FabricConnectionUuid) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricConnectionUuid) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricConnectionUuid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricConnectionUuid(val *FabricConnectionUuid) *NullableFabricConnectionUuid {
	return &NullableFabricConnectionUuid{value: val, isSet: true}
}

func (v NullableFabricConnectionUuid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricConnectionUuid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
