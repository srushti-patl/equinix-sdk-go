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
	"fmt"
)

// checks if the RouteFiltersChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFiltersChange{}

// RouteFiltersChange Current state of latest Route Filter change
type RouteFiltersChange struct {
	// Uniquely identifies a change
	Uuid string                 `json:"uuid"`
	Type RouteFiltersChangeType `json:"type"`
	// Route Filter Change URI
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFiltersChange RouteFiltersChange

// NewRouteFiltersChange instantiates a new RouteFiltersChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFiltersChange(uuid string, type_ RouteFiltersChangeType) *RouteFiltersChange {
	this := RouteFiltersChange{}
	this.Uuid = uuid
	this.Type = type_
	return &this
}

// NewRouteFiltersChangeWithDefaults instantiates a new RouteFiltersChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFiltersChangeWithDefaults() *RouteFiltersChange {
	this := RouteFiltersChange{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *RouteFiltersChange) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RouteFiltersChange) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *RouteFiltersChange) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *RouteFiltersChange) GetType() RouteFiltersChangeType {
	if o == nil {
		var ret RouteFiltersChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteFiltersChange) GetTypeOk() (*RouteFiltersChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteFiltersChange) SetType(v RouteFiltersChangeType) {
	o.Type = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteFiltersChange) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFiltersChange) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteFiltersChange) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteFiltersChange) SetHref(v string) {
	o.Href = &v
}

func (o RouteFiltersChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFiltersChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["type"] = o.Type
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFiltersChange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"type",
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

	varRouteFiltersChange := _RouteFiltersChange{}

	err = json.Unmarshal(data, &varRouteFiltersChange)

	if err != nil {
		return err
	}

	*o = RouteFiltersChange(varRouteFiltersChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFiltersChange struct {
	value *RouteFiltersChange
	isSet bool
}

func (v NullableRouteFiltersChange) Get() *RouteFiltersChange {
	return v.value
}

func (v *NullableRouteFiltersChange) Set(val *RouteFiltersChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFiltersChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFiltersChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFiltersChange(val *RouteFiltersChange) *NullableRouteFiltersChange {
	return &NullableRouteFiltersChange{value: val, isSet: true}
}

func (v NullableRouteFiltersChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFiltersChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
