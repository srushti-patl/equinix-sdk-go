/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the RoutingProtocolChangeDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolChangeDataResponse{}

// RoutingProtocolChangeDataResponse List of network changes
type RoutingProtocolChangeDataResponse struct {
	Pagination           *Pagination                 `json:"pagination,omitempty"`
	Data                 []RoutingProtocolChangeData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolChangeDataResponse RoutingProtocolChangeDataResponse

// NewRoutingProtocolChangeDataResponse instantiates a new RoutingProtocolChangeDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolChangeDataResponse() *RoutingProtocolChangeDataResponse {
	this := RoutingProtocolChangeDataResponse{}
	return &this
}

// NewRoutingProtocolChangeDataResponseWithDefaults instantiates a new RoutingProtocolChangeDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolChangeDataResponseWithDefaults() *RoutingProtocolChangeDataResponse {
	this := RoutingProtocolChangeDataResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *RoutingProtocolChangeDataResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolChangeDataResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *RoutingProtocolChangeDataResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *RoutingProtocolChangeDataResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RoutingProtocolChangeDataResponse) GetData() []RoutingProtocolChangeData {
	if o == nil || IsNil(o.Data) {
		var ret []RoutingProtocolChangeData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolChangeDataResponse) GetDataOk() ([]RoutingProtocolChangeData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RoutingProtocolChangeDataResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RoutingProtocolChangeData and assigns it to the Data field.
func (o *RoutingProtocolChangeDataResponse) SetData(v []RoutingProtocolChangeData) {
	o.Data = v
}

func (o RoutingProtocolChangeDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolChangeDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolChangeDataResponse) UnmarshalJSON(data []byte) (err error) {
	varRoutingProtocolChangeDataResponse := _RoutingProtocolChangeDataResponse{}

	err = json.Unmarshal(data, &varRoutingProtocolChangeDataResponse)

	if err != nil {
		return err
	}

	*o = RoutingProtocolChangeDataResponse(varRoutingProtocolChangeDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolChangeDataResponse struct {
	value *RoutingProtocolChangeDataResponse
	isSet bool
}

func (v NullableRoutingProtocolChangeDataResponse) Get() *RoutingProtocolChangeDataResponse {
	return v.value
}

func (v *NullableRoutingProtocolChangeDataResponse) Set(val *RoutingProtocolChangeDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolChangeDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolChangeDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolChangeDataResponse(val *RoutingProtocolChangeDataResponse) *NullableRoutingProtocolChangeDataResponse {
	return &NullableRoutingProtocolChangeDataResponse{value: val, isSet: true}
}

func (v NullableRoutingProtocolChangeDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolChangeDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
