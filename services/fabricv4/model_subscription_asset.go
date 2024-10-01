/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the SubscriptionAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAsset{}

// SubscriptionAsset Asset information
type SubscriptionAsset struct {
	// Type of the subscription asset ( XF_ROUTER ,IP_VC, IPWAN_VC )
	Type    *string                        `json:"type,omitempty"`
	Package *SubscriptionRouterPackageType `json:"package,omitempty"`
	// Bandwidth of the asset in Mbps
	Bandwidth            *int32 `json:"bandwidth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionAsset SubscriptionAsset

// NewSubscriptionAsset instantiates a new SubscriptionAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAsset() *SubscriptionAsset {
	this := SubscriptionAsset{}
	return &this
}

// NewSubscriptionAssetWithDefaults instantiates a new SubscriptionAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAssetWithDefaults() *SubscriptionAsset {
	this := SubscriptionAsset{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubscriptionAsset) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAsset) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubscriptionAsset) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SubscriptionAsset) SetType(v string) {
	o.Type = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *SubscriptionAsset) GetPackage() SubscriptionRouterPackageType {
	if o == nil || IsNil(o.Package) {
		var ret SubscriptionRouterPackageType
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAsset) GetPackageOk() (*SubscriptionRouterPackageType, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *SubscriptionAsset) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given SubscriptionRouterPackageType and assigns it to the Package field.
func (o *SubscriptionAsset) SetPackage(v SubscriptionRouterPackageType) {
	o.Package = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *SubscriptionAsset) GetBandwidth() int32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAsset) GetBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *SubscriptionAsset) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *SubscriptionAsset) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

func (o SubscriptionAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionAsset) UnmarshalJSON(data []byte) (err error) {
	varSubscriptionAsset := _SubscriptionAsset{}

	err = json.Unmarshal(data, &varSubscriptionAsset)

	if err != nil {
		return err
	}

	*o = SubscriptionAsset(varSubscriptionAsset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "package")
		delete(additionalProperties, "bandwidth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionAsset struct {
	value *SubscriptionAsset
	isSet bool
}

func (v NullableSubscriptionAsset) Get() *SubscriptionAsset {
	return v.value
}

func (v *NullableSubscriptionAsset) Set(val *SubscriptionAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAsset(val *SubscriptionAsset) *NullableSubscriptionAsset {
	return &NullableSubscriptionAsset{value: val, isSet: true}
}

func (v NullableSubscriptionAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
