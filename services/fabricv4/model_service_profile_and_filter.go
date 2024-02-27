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

// checks if the ServiceProfileAndFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileAndFilter{}

// ServiceProfileAndFilter struct for ServiceProfileAndFilter
type ServiceProfileAndFilter struct {
	And                  []ServiceProfileSimpleExpression `json:"and,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProfileAndFilter ServiceProfileAndFilter

// NewServiceProfileAndFilter instantiates a new ServiceProfileAndFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileAndFilter() *ServiceProfileAndFilter {
	this := ServiceProfileAndFilter{}
	return &this
}

// NewServiceProfileAndFilterWithDefaults instantiates a new ServiceProfileAndFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileAndFilterWithDefaults() *ServiceProfileAndFilter {
	this := ServiceProfileAndFilter{}
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *ServiceProfileAndFilter) GetAnd() []ServiceProfileSimpleExpression {
	if o == nil || IsNil(o.And) {
		var ret []ServiceProfileSimpleExpression
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAndFilter) GetAndOk() ([]ServiceProfileSimpleExpression, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *ServiceProfileAndFilter) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []ServiceProfileSimpleExpression and assigns it to the And field.
func (o *ServiceProfileAndFilter) SetAnd(v []ServiceProfileSimpleExpression) {
	o.And = v
}

func (o ServiceProfileAndFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileAndFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.And) {
		toSerialize["and"] = o.And
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProfileAndFilter) UnmarshalJSON(data []byte) (err error) {
	varServiceProfileAndFilter := _ServiceProfileAndFilter{}

	err = json.Unmarshal(data, &varServiceProfileAndFilter)

	if err != nil {
		return err
	}

	*o = ServiceProfileAndFilter(varServiceProfileAndFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "and")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProfileAndFilter struct {
	value *ServiceProfileAndFilter
	isSet bool
}

func (v NullableServiceProfileAndFilter) Get() *ServiceProfileAndFilter {
	return v.value
}

func (v *NullableServiceProfileAndFilter) Set(val *ServiceProfileAndFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAndFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAndFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAndFilter(val *ServiceProfileAndFilter) *NullableServiceProfileAndFilter {
	return &NullableServiceProfileAndFilter{value: val, isSet: true}
}

func (v NullableServiceProfileAndFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAndFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
