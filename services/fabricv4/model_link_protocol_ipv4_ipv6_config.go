/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the LinkProtocolIpv4Ipv6Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkProtocolIpv4Ipv6Config{}

// LinkProtocolIpv4Ipv6Config IPv4 or IPv6 specific configuration
type LinkProtocolIpv4Ipv6Config struct {
	// Link subnet prefix
	LinkPrefix *string `json:"linkPrefix,omitempty"`
	// Prefix datatype when linkPrefix not specified
	LocalIfaceIp *string `json:"localIfaceIp,omitempty"`
	// Equinix-side link interface address
	RemoteIfaceIp        *string `json:"remoteIfaceIp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkProtocolIpv4Ipv6Config LinkProtocolIpv4Ipv6Config

// NewLinkProtocolIpv4Ipv6Config instantiates a new LinkProtocolIpv4Ipv6Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkProtocolIpv4Ipv6Config() *LinkProtocolIpv4Ipv6Config {
	this := LinkProtocolIpv4Ipv6Config{}
	return &this
}

// NewLinkProtocolIpv4Ipv6ConfigWithDefaults instantiates a new LinkProtocolIpv4Ipv6Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkProtocolIpv4Ipv6ConfigWithDefaults() *LinkProtocolIpv4Ipv6Config {
	this := LinkProtocolIpv4Ipv6Config{}
	return &this
}

// GetLinkPrefix returns the LinkPrefix field value if set, zero value otherwise.
func (o *LinkProtocolIpv4Ipv6Config) GetLinkPrefix() string {
	if o == nil || IsNil(o.LinkPrefix) {
		var ret string
		return ret
	}
	return *o.LinkPrefix
}

// GetLinkPrefixOk returns a tuple with the LinkPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkProtocolIpv4Ipv6Config) GetLinkPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.LinkPrefix) {
		return nil, false
	}
	return o.LinkPrefix, true
}

// HasLinkPrefix returns a boolean if a field has been set.
func (o *LinkProtocolIpv4Ipv6Config) HasLinkPrefix() bool {
	if o != nil && !IsNil(o.LinkPrefix) {
		return true
	}

	return false
}

// SetLinkPrefix gets a reference to the given string and assigns it to the LinkPrefix field.
func (o *LinkProtocolIpv4Ipv6Config) SetLinkPrefix(v string) {
	o.LinkPrefix = &v
}

// GetLocalIfaceIp returns the LocalIfaceIp field value if set, zero value otherwise.
func (o *LinkProtocolIpv4Ipv6Config) GetLocalIfaceIp() string {
	if o == nil || IsNil(o.LocalIfaceIp) {
		var ret string
		return ret
	}
	return *o.LocalIfaceIp
}

// GetLocalIfaceIpOk returns a tuple with the LocalIfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkProtocolIpv4Ipv6Config) GetLocalIfaceIpOk() (*string, bool) {
	if o == nil || IsNil(o.LocalIfaceIp) {
		return nil, false
	}
	return o.LocalIfaceIp, true
}

// HasLocalIfaceIp returns a boolean if a field has been set.
func (o *LinkProtocolIpv4Ipv6Config) HasLocalIfaceIp() bool {
	if o != nil && !IsNil(o.LocalIfaceIp) {
		return true
	}

	return false
}

// SetLocalIfaceIp gets a reference to the given string and assigns it to the LocalIfaceIp field.
func (o *LinkProtocolIpv4Ipv6Config) SetLocalIfaceIp(v string) {
	o.LocalIfaceIp = &v
}

// GetRemoteIfaceIp returns the RemoteIfaceIp field value if set, zero value otherwise.
func (o *LinkProtocolIpv4Ipv6Config) GetRemoteIfaceIp() string {
	if o == nil || IsNil(o.RemoteIfaceIp) {
		var ret string
		return ret
	}
	return *o.RemoteIfaceIp
}

// GetRemoteIfaceIpOk returns a tuple with the RemoteIfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkProtocolIpv4Ipv6Config) GetRemoteIfaceIpOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteIfaceIp) {
		return nil, false
	}
	return o.RemoteIfaceIp, true
}

// HasRemoteIfaceIp returns a boolean if a field has been set.
func (o *LinkProtocolIpv4Ipv6Config) HasRemoteIfaceIp() bool {
	if o != nil && !IsNil(o.RemoteIfaceIp) {
		return true
	}

	return false
}

// SetRemoteIfaceIp gets a reference to the given string and assigns it to the RemoteIfaceIp field.
func (o *LinkProtocolIpv4Ipv6Config) SetRemoteIfaceIp(v string) {
	o.RemoteIfaceIp = &v
}

func (o LinkProtocolIpv4Ipv6Config) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkProtocolIpv4Ipv6Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkPrefix) {
		toSerialize["linkPrefix"] = o.LinkPrefix
	}
	if !IsNil(o.LocalIfaceIp) {
		toSerialize["localIfaceIp"] = o.LocalIfaceIp
	}
	if !IsNil(o.RemoteIfaceIp) {
		toSerialize["remoteIfaceIp"] = o.RemoteIfaceIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkProtocolIpv4Ipv6Config) UnmarshalJSON(data []byte) (err error) {
	varLinkProtocolIpv4Ipv6Config := _LinkProtocolIpv4Ipv6Config{}

	err = json.Unmarshal(data, &varLinkProtocolIpv4Ipv6Config)

	if err != nil {
		return err
	}

	*o = LinkProtocolIpv4Ipv6Config(varLinkProtocolIpv4Ipv6Config)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "linkPrefix")
		delete(additionalProperties, "localIfaceIp")
		delete(additionalProperties, "remoteIfaceIp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkProtocolIpv4Ipv6Config struct {
	value *LinkProtocolIpv4Ipv6Config
	isSet bool
}

func (v NullableLinkProtocolIpv4Ipv6Config) Get() *LinkProtocolIpv4Ipv6Config {
	return v.value
}

func (v *NullableLinkProtocolIpv4Ipv6Config) Set(val *LinkProtocolIpv4Ipv6Config) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkProtocolIpv4Ipv6Config) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkProtocolIpv4Ipv6Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkProtocolIpv4Ipv6Config(val *LinkProtocolIpv4Ipv6Config) *NullableLinkProtocolIpv4Ipv6Config {
	return &NullableLinkProtocolIpv4Ipv6Config{value: val, isSet: true}
}

func (v NullableLinkProtocolIpv4Ipv6Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkProtocolIpv4Ipv6Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
