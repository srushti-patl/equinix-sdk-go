/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the ConnectionRouteFiltersBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionRouteFiltersBase{}

// ConnectionRouteFiltersBase struct for ConnectionRouteFiltersBase
type ConnectionRouteFiltersBase struct {
	Direction            ConnectionRouteFiltersBaseDirection `json:"direction"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionRouteFiltersBase ConnectionRouteFiltersBase

// NewConnectionRouteFiltersBase instantiates a new ConnectionRouteFiltersBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionRouteFiltersBase(direction ConnectionRouteFiltersBaseDirection) *ConnectionRouteFiltersBase {
	this := ConnectionRouteFiltersBase{}
	this.Direction = direction
	return &this
}

// NewConnectionRouteFiltersBaseWithDefaults instantiates a new ConnectionRouteFiltersBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionRouteFiltersBaseWithDefaults() *ConnectionRouteFiltersBase {
	this := ConnectionRouteFiltersBase{}
	return &this
}

// GetDirection returns the Direction field value
func (o *ConnectionRouteFiltersBase) GetDirection() ConnectionRouteFiltersBaseDirection {
	if o == nil {
		var ret ConnectionRouteFiltersBaseDirection
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *ConnectionRouteFiltersBase) GetDirectionOk() (*ConnectionRouteFiltersBaseDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *ConnectionRouteFiltersBase) SetDirection(v ConnectionRouteFiltersBaseDirection) {
	o.Direction = v
}

func (o ConnectionRouteFiltersBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionRouteFiltersBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["direction"] = o.Direction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionRouteFiltersBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"direction",
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

	varConnectionRouteFiltersBase := _ConnectionRouteFiltersBase{}

	err = json.Unmarshal(data, &varConnectionRouteFiltersBase)

	if err != nil {
		return err
	}

	*o = ConnectionRouteFiltersBase(varConnectionRouteFiltersBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionRouteFiltersBase struct {
	value *ConnectionRouteFiltersBase
	isSet bool
}

func (v NullableConnectionRouteFiltersBase) Get() *ConnectionRouteFiltersBase {
	return v.value
}

func (v *NullableConnectionRouteFiltersBase) Set(val *ConnectionRouteFiltersBase) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteFiltersBase) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteFiltersBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteFiltersBase(val *ConnectionRouteFiltersBase) *NullableConnectionRouteFiltersBase {
	return &NullableConnectionRouteFiltersBase{value: val, isSet: true}
}

func (v NullableConnectionRouteFiltersBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteFiltersBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
