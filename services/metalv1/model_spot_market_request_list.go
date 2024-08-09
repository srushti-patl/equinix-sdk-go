/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the SpotMarketRequestList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotMarketRequestList{}

// SpotMarketRequestList struct for SpotMarketRequestList
type SpotMarketRequestList struct {
	SpotMarketRequests   []SpotMarketRequest `json:"spot_market_requests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpotMarketRequestList SpotMarketRequestList

// NewSpotMarketRequestList instantiates a new SpotMarketRequestList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotMarketRequestList() *SpotMarketRequestList {
	this := SpotMarketRequestList{}
	return &this
}

// NewSpotMarketRequestListWithDefaults instantiates a new SpotMarketRequestList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotMarketRequestListWithDefaults() *SpotMarketRequestList {
	this := SpotMarketRequestList{}
	return &this
}

// GetSpotMarketRequests returns the SpotMarketRequests field value if set, zero value otherwise.
func (o *SpotMarketRequestList) GetSpotMarketRequests() []SpotMarketRequest {
	if o == nil || IsNil(o.SpotMarketRequests) {
		var ret []SpotMarketRequest
		return ret
	}
	return o.SpotMarketRequests
}

// GetSpotMarketRequestsOk returns a tuple with the SpotMarketRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestList) GetSpotMarketRequestsOk() ([]SpotMarketRequest, bool) {
	if o == nil || IsNil(o.SpotMarketRequests) {
		return nil, false
	}
	return o.SpotMarketRequests, true
}

// HasSpotMarketRequests returns a boolean if a field has been set.
func (o *SpotMarketRequestList) HasSpotMarketRequests() bool {
	if o != nil && !IsNil(o.SpotMarketRequests) {
		return true
	}

	return false
}

// SetSpotMarketRequests gets a reference to the given []SpotMarketRequest and assigns it to the SpotMarketRequests field.
func (o *SpotMarketRequestList) SetSpotMarketRequests(v []SpotMarketRequest) {
	o.SpotMarketRequests = v
}

func (o SpotMarketRequestList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotMarketRequestList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpotMarketRequests) {
		toSerialize["spot_market_requests"] = o.SpotMarketRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpotMarketRequestList) UnmarshalJSON(data []byte) (err error) {
	varSpotMarketRequestList := _SpotMarketRequestList{}

	err = json.Unmarshal(data, &varSpotMarketRequestList)

	if err != nil {
		return err
	}

	*o = SpotMarketRequestList(varSpotMarketRequestList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "spot_market_requests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpotMarketRequestList struct {
	value *SpotMarketRequestList
	isSet bool
}

func (v NullableSpotMarketRequestList) Get() *SpotMarketRequestList {
	return v.value
}

func (v *NullableSpotMarketRequestList) Set(val *SpotMarketRequestList) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotMarketRequestList) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotMarketRequestList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotMarketRequestList(val *SpotMarketRequestList) *NullableSpotMarketRequestList {
	return &NullableSpotMarketRequestList{value: val, isSet: true}
}

func (v NullableSpotMarketRequestList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotMarketRequestList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
