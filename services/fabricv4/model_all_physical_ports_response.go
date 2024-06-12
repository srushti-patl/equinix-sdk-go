/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the AllPhysicalPortsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllPhysicalPortsResponse{}

// AllPhysicalPortsResponse GET All Physical Ports
type AllPhysicalPortsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// GET All Physical Ports
	Data                 []PhysicalPort `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllPhysicalPortsResponse AllPhysicalPortsResponse

// NewAllPhysicalPortsResponse instantiates a new AllPhysicalPortsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllPhysicalPortsResponse() *AllPhysicalPortsResponse {
	this := AllPhysicalPortsResponse{}
	return &this
}

// NewAllPhysicalPortsResponseWithDefaults instantiates a new AllPhysicalPortsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllPhysicalPortsResponseWithDefaults() *AllPhysicalPortsResponse {
	this := AllPhysicalPortsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AllPhysicalPortsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllPhysicalPortsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AllPhysicalPortsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *AllPhysicalPortsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AllPhysicalPortsResponse) GetData() []PhysicalPort {
	if o == nil || IsNil(o.Data) {
		var ret []PhysicalPort
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllPhysicalPortsResponse) GetDataOk() ([]PhysicalPort, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AllPhysicalPortsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PhysicalPort and assigns it to the Data field.
func (o *AllPhysicalPortsResponse) SetData(v []PhysicalPort) {
	o.Data = v
}

func (o AllPhysicalPortsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllPhysicalPortsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AllPhysicalPortsResponse) UnmarshalJSON(data []byte) (err error) {
	varAllPhysicalPortsResponse := _AllPhysicalPortsResponse{}

	err = json.Unmarshal(data, &varAllPhysicalPortsResponse)

	if err != nil {
		return err
	}

	*o = AllPhysicalPortsResponse(varAllPhysicalPortsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllPhysicalPortsResponse struct {
	value *AllPhysicalPortsResponse
	isSet bool
}

func (v NullableAllPhysicalPortsResponse) Get() *AllPhysicalPortsResponse {
	return v.value
}

func (v *NullableAllPhysicalPortsResponse) Set(val *AllPhysicalPortsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllPhysicalPortsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllPhysicalPortsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllPhysicalPortsResponse(val *AllPhysicalPortsResponse) *NullableAllPhysicalPortsResponse {
	return &NullableAllPhysicalPortsResponse{value: val, isSet: true}
}

func (v NullableAllPhysicalPortsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllPhysicalPortsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
