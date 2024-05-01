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

// checks if the RoutingProtocolIpBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolIpBlock{}

// RoutingProtocolIpBlock struct for RoutingProtocolIpBlock
type RoutingProtocolIpBlock struct {
	CustomerRoute        CustomerRoute `json:"customerRoute"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolIpBlock RoutingProtocolIpBlock

// NewRoutingProtocolIpBlock instantiates a new RoutingProtocolIpBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolIpBlock(customerRoute CustomerRoute) *RoutingProtocolIpBlock {
	this := RoutingProtocolIpBlock{}
	this.CustomerRoute = customerRoute
	return &this
}

// NewRoutingProtocolIpBlockWithDefaults instantiates a new RoutingProtocolIpBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolIpBlockWithDefaults() *RoutingProtocolIpBlock {
	this := RoutingProtocolIpBlock{}
	return &this
}

// GetCustomerRoute returns the CustomerRoute field value
func (o *RoutingProtocolIpBlock) GetCustomerRoute() CustomerRoute {
	if o == nil {
		var ret CustomerRoute
		return ret
	}

	return o.CustomerRoute
}

// GetCustomerRouteOk returns a tuple with the CustomerRoute field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolIpBlock) GetCustomerRouteOk() (*CustomerRoute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerRoute, true
}

// SetCustomerRoute sets field value
func (o *RoutingProtocolIpBlock) SetCustomerRoute(v CustomerRoute) {
	o.CustomerRoute = v
}

func (o RoutingProtocolIpBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolIpBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerRoute"] = o.CustomerRoute

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolIpBlock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerRoute",
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

	varRoutingProtocolIpBlock := _RoutingProtocolIpBlock{}

	err = json.Unmarshal(data, &varRoutingProtocolIpBlock)

	if err != nil {
		return err
	}

	*o = RoutingProtocolIpBlock(varRoutingProtocolIpBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerRoute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolIpBlock struct {
	value *RoutingProtocolIpBlock
	isSet bool
}

func (v NullableRoutingProtocolIpBlock) Get() *RoutingProtocolIpBlock {
	return v.value
}

func (v *NullableRoutingProtocolIpBlock) Set(val *RoutingProtocolIpBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolIpBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolIpBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolIpBlock(val *RoutingProtocolIpBlock) *NullableRoutingProtocolIpBlock {
	return &NullableRoutingProtocolIpBlock{value: val, isSet: true}
}

func (v NullableRoutingProtocolIpBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolIpBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
