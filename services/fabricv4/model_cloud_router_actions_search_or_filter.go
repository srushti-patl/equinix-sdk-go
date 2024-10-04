/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterActionsSearchOrFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterActionsSearchOrFilter{}

// CloudRouterActionsSearchOrFilter struct for CloudRouterActionsSearchOrFilter
type CloudRouterActionsSearchOrFilter struct {
	Or                   []CloudRouterActionsSearchExpression `json:"or,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterActionsSearchOrFilter CloudRouterActionsSearchOrFilter

// NewCloudRouterActionsSearchOrFilter instantiates a new CloudRouterActionsSearchOrFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterActionsSearchOrFilter() *CloudRouterActionsSearchOrFilter {
	this := CloudRouterActionsSearchOrFilter{}
	return &this
}

// NewCloudRouterActionsSearchOrFilterWithDefaults instantiates a new CloudRouterActionsSearchOrFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterActionsSearchOrFilterWithDefaults() *CloudRouterActionsSearchOrFilter {
	this := CloudRouterActionsSearchOrFilter{}
	return &this
}

// GetOr returns the Or field value if set, zero value otherwise.
func (o *CloudRouterActionsSearchOrFilter) GetOr() []CloudRouterActionsSearchExpression {
	if o == nil || IsNil(o.Or) {
		var ret []CloudRouterActionsSearchExpression
		return ret
	}
	return o.Or
}

// GetOrOk returns a tuple with the Or field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterActionsSearchOrFilter) GetOrOk() ([]CloudRouterActionsSearchExpression, bool) {
	if o == nil || IsNil(o.Or) {
		return nil, false
	}
	return o.Or, true
}

// HasOr returns a boolean if a field has been set.
func (o *CloudRouterActionsSearchOrFilter) HasOr() bool {
	if o != nil && !IsNil(o.Or) {
		return true
	}

	return false
}

// SetOr gets a reference to the given []CloudRouterActionsSearchExpression and assigns it to the Or field.
func (o *CloudRouterActionsSearchOrFilter) SetOr(v []CloudRouterActionsSearchExpression) {
	o.Or = v
}

func (o CloudRouterActionsSearchOrFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterActionsSearchOrFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Or) {
		toSerialize["or"] = o.Or
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterActionsSearchOrFilter) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterActionsSearchOrFilter := _CloudRouterActionsSearchOrFilter{}

	err = json.Unmarshal(data, &varCloudRouterActionsSearchOrFilter)

	if err != nil {
		return err
	}

	*o = CloudRouterActionsSearchOrFilter(varCloudRouterActionsSearchOrFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "or")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterActionsSearchOrFilter struct {
	value *CloudRouterActionsSearchOrFilter
	isSet bool
}

func (v NullableCloudRouterActionsSearchOrFilter) Get() *CloudRouterActionsSearchOrFilter {
	return v.value
}

func (v *NullableCloudRouterActionsSearchOrFilter) Set(val *CloudRouterActionsSearchOrFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterActionsSearchOrFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterActionsSearchOrFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterActionsSearchOrFilter(val *CloudRouterActionsSearchOrFilter) *NullableCloudRouterActionsSearchOrFilter {
	return &NullableCloudRouterActionsSearchOrFilter{value: val, isSet: true}
}

func (v NullableCloudRouterActionsSearchOrFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterActionsSearchOrFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
