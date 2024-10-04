/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the SimplifiedVirtualDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedVirtualDevice{}

// SimplifiedVirtualDevice struct for SimplifiedVirtualDevice
type SimplifiedVirtualDevice struct {
	// url to entity
	Href *string `json:"href,omitempty"`
	// Network Edge assigned Virtual Device Identifier
	Uuid *string                      `json:"uuid,omitempty"`
	Type *SimplifiedVirtualDeviceType `json:"type,omitempty"`
	// Customer-assigned Virtual Device name
	Name *string `json:"name,omitempty"`
	// Virtual Device Cluster Information
	Cluster              *string `json:"cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedVirtualDevice SimplifiedVirtualDevice

// NewSimplifiedVirtualDevice instantiates a new SimplifiedVirtualDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedVirtualDevice() *SimplifiedVirtualDevice {
	this := SimplifiedVirtualDevice{}
	return &this
}

// NewSimplifiedVirtualDeviceWithDefaults instantiates a new SimplifiedVirtualDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedVirtualDeviceWithDefaults() *SimplifiedVirtualDevice {
	this := SimplifiedVirtualDevice{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SimplifiedVirtualDevice) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedVirtualDevice) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SimplifiedVirtualDevice) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SimplifiedVirtualDevice) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SimplifiedVirtualDevice) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedVirtualDevice) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SimplifiedVirtualDevice) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SimplifiedVirtualDevice) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimplifiedVirtualDevice) GetType() SimplifiedVirtualDeviceType {
	if o == nil || IsNil(o.Type) {
		var ret SimplifiedVirtualDeviceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedVirtualDevice) GetTypeOk() (*SimplifiedVirtualDeviceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimplifiedVirtualDevice) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SimplifiedVirtualDeviceType and assigns it to the Type field.
func (o *SimplifiedVirtualDevice) SetType(v SimplifiedVirtualDeviceType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimplifiedVirtualDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedVirtualDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimplifiedVirtualDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimplifiedVirtualDevice) SetName(v string) {
	o.Name = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *SimplifiedVirtualDevice) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedVirtualDevice) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *SimplifiedVirtualDevice) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *SimplifiedVirtualDevice) SetCluster(v string) {
	o.Cluster = &v
}

func (o SimplifiedVirtualDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedVirtualDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedVirtualDevice) UnmarshalJSON(data []byte) (err error) {
	varSimplifiedVirtualDevice := _SimplifiedVirtualDevice{}

	err = json.Unmarshal(data, &varSimplifiedVirtualDevice)

	if err != nil {
		return err
	}

	*o = SimplifiedVirtualDevice(varSimplifiedVirtualDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "cluster")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedVirtualDevice struct {
	value *SimplifiedVirtualDevice
	isSet bool
}

func (v NullableSimplifiedVirtualDevice) Get() *SimplifiedVirtualDevice {
	return v.value
}

func (v *NullableSimplifiedVirtualDevice) Set(val *SimplifiedVirtualDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedVirtualDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedVirtualDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedVirtualDevice(val *SimplifiedVirtualDevice) *NullableSimplifiedVirtualDevice {
	return &NullableSimplifiedVirtualDevice{value: val, isSet: true}
}

func (v NullableSimplifiedVirtualDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedVirtualDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
