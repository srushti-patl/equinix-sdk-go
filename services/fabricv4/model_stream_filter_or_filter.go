/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the StreamFilterOrFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamFilterOrFilter{}

// StreamFilterOrFilter struct for StreamFilterOrFilter
type StreamFilterOrFilter struct {
	Or                   []StreamFilterSimpleExpression `json:"or,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StreamFilterOrFilter StreamFilterOrFilter

// NewStreamFilterOrFilter instantiates a new StreamFilterOrFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamFilterOrFilter() *StreamFilterOrFilter {
	this := StreamFilterOrFilter{}
	return &this
}

// NewStreamFilterOrFilterWithDefaults instantiates a new StreamFilterOrFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamFilterOrFilterWithDefaults() *StreamFilterOrFilter {
	this := StreamFilterOrFilter{}
	return &this
}

// GetOr returns the Or field value if set, zero value otherwise.
func (o *StreamFilterOrFilter) GetOr() []StreamFilterSimpleExpression {
	if o == nil || IsNil(o.Or) {
		var ret []StreamFilterSimpleExpression
		return ret
	}
	return o.Or
}

// GetOrOk returns a tuple with the Or field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamFilterOrFilter) GetOrOk() ([]StreamFilterSimpleExpression, bool) {
	if o == nil || IsNil(o.Or) {
		return nil, false
	}
	return o.Or, true
}

// HasOr returns a boolean if a field has been set.
func (o *StreamFilterOrFilter) HasOr() bool {
	if o != nil && !IsNil(o.Or) {
		return true
	}

	return false
}

// SetOr gets a reference to the given []StreamFilterSimpleExpression and assigns it to the Or field.
func (o *StreamFilterOrFilter) SetOr(v []StreamFilterSimpleExpression) {
	o.Or = v
}

func (o StreamFilterOrFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamFilterOrFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Or) {
		toSerialize["or"] = o.Or
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StreamFilterOrFilter) UnmarshalJSON(data []byte) (err error) {
	varStreamFilterOrFilter := _StreamFilterOrFilter{}

	err = json.Unmarshal(data, &varStreamFilterOrFilter)

	if err != nil {
		return err
	}

	*o = StreamFilterOrFilter(varStreamFilterOrFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "or")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStreamFilterOrFilter struct {
	value *StreamFilterOrFilter
	isSet bool
}

func (v NullableStreamFilterOrFilter) Get() *StreamFilterOrFilter {
	return v.value
}

func (v *NullableStreamFilterOrFilter) Set(val *StreamFilterOrFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamFilterOrFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamFilterOrFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamFilterOrFilter(val *StreamFilterOrFilter) *NullableStreamFilterOrFilter {
	return &NullableStreamFilterOrFilter{value: val, isSet: true}
}

func (v NullableStreamFilterOrFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamFilterOrFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
