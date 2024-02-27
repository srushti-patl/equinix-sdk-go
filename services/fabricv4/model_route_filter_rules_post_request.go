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

// checks if the RouteFilterRulesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFilterRulesPostRequest{}

// RouteFilterRulesPostRequest Create Route Filter Rule POST request
type RouteFilterRulesPostRequest struct {
	// Route Filter Rule configuration
	Data                 []RouteFilterRulesBase `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFilterRulesPostRequest RouteFilterRulesPostRequest

// NewRouteFilterRulesPostRequest instantiates a new RouteFilterRulesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFilterRulesPostRequest() *RouteFilterRulesPostRequest {
	this := RouteFilterRulesPostRequest{}
	return &this
}

// NewRouteFilterRulesPostRequestWithDefaults instantiates a new RouteFilterRulesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFilterRulesPostRequestWithDefaults() *RouteFilterRulesPostRequest {
	this := RouteFilterRulesPostRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RouteFilterRulesPostRequest) GetData() []RouteFilterRulesBase {
	if o == nil || IsNil(o.Data) {
		var ret []RouteFilterRulesBase
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesPostRequest) GetDataOk() ([]RouteFilterRulesBase, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RouteFilterRulesPostRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RouteFilterRulesBase and assigns it to the Data field.
func (o *RouteFilterRulesPostRequest) SetData(v []RouteFilterRulesBase) {
	o.Data = v
}

func (o RouteFilterRulesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFilterRulesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFilterRulesPostRequest) UnmarshalJSON(data []byte) (err error) {
	varRouteFilterRulesPostRequest := _RouteFilterRulesPostRequest{}

	err = json.Unmarshal(data, &varRouteFilterRulesPostRequest)

	if err != nil {
		return err
	}

	*o = RouteFilterRulesPostRequest(varRouteFilterRulesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFilterRulesPostRequest struct {
	value *RouteFilterRulesPostRequest
	isSet bool
}

func (v NullableRouteFilterRulesPostRequest) Get() *RouteFilterRulesPostRequest {
	return v.value
}

func (v *NullableRouteFilterRulesPostRequest) Set(val *RouteFilterRulesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterRulesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterRulesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterRulesPostRequest(val *RouteFilterRulesPostRequest) *NullableRouteFilterRulesPostRequest {
	return &NullableRouteFilterRulesPostRequest{value: val, isSet: true}
}

func (v NullableRouteFilterRulesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterRulesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
