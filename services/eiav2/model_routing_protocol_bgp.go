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
	"reflect"
	"strings"
)

// checks if the RoutingProtocolBgp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolBgp{}

// RoutingProtocolBgp struct for RoutingProtocolBgp
type RoutingProtocolBgp struct {
	RoutingProtocolReadModel
	// Customer Autonomous System Number
	CustomerAsn      int64                                           `json:"customerAsn"`
	CustomerAsnRange *BgpRoutingProtocolRequestAllOfCustomerAsnRange `json:"customerAsnRange,omitempty"`
	// Equinix Autonomous System Number
	EquinixAsn int64 `json:"equinixAsn"`
	// BGP authentication key
	BgpAuthKey           *string      `json:"bgpAuthKey,omitempty"`
	ExportPolicy         ExportPolicy `json:"exportPolicy"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolBgp RoutingProtocolBgp

// NewRoutingProtocolBgp instantiates a new RoutingProtocolBgp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolBgp(customerAsn int64, equinixAsn int64, exportPolicy ExportPolicy, type_ RoutingProtocolType, name string, changeLog ChangeLog, links []Link) *RoutingProtocolBgp {
	this := RoutingProtocolBgp{}
	this.Type = type_
	this.Name = name
	this.ChangeLog = changeLog
	this.Links = links
	this.CustomerAsn = customerAsn
	this.EquinixAsn = equinixAsn
	this.ExportPolicy = exportPolicy
	return &this
}

// NewRoutingProtocolBgpWithDefaults instantiates a new RoutingProtocolBgp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolBgpWithDefaults() *RoutingProtocolBgp {
	this := RoutingProtocolBgp{}
	return &this
}

// GetCustomerAsn returns the CustomerAsn field value
func (o *RoutingProtocolBgp) GetCustomerAsn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CustomerAsn
}

// GetCustomerAsnOk returns a tuple with the CustomerAsn field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBgp) GetCustomerAsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAsn, true
}

// SetCustomerAsn sets field value
func (o *RoutingProtocolBgp) SetCustomerAsn(v int64) {
	o.CustomerAsn = v
}

// GetCustomerAsnRange returns the CustomerAsnRange field value if set, zero value otherwise.
func (o *RoutingProtocolBgp) GetCustomerAsnRange() BgpRoutingProtocolRequestAllOfCustomerAsnRange {
	if o == nil || IsNil(o.CustomerAsnRange) {
		var ret BgpRoutingProtocolRequestAllOfCustomerAsnRange
		return ret
	}
	return *o.CustomerAsnRange
}

// GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBgp) GetCustomerAsnRangeOk() (*BgpRoutingProtocolRequestAllOfCustomerAsnRange, bool) {
	if o == nil || IsNil(o.CustomerAsnRange) {
		return nil, false
	}
	return o.CustomerAsnRange, true
}

// HasCustomerAsnRange returns a boolean if a field has been set.
func (o *RoutingProtocolBgp) HasCustomerAsnRange() bool {
	if o != nil && !IsNil(o.CustomerAsnRange) {
		return true
	}

	return false
}

// SetCustomerAsnRange gets a reference to the given BgpRoutingProtocolRequestAllOfCustomerAsnRange and assigns it to the CustomerAsnRange field.
func (o *RoutingProtocolBgp) SetCustomerAsnRange(v BgpRoutingProtocolRequestAllOfCustomerAsnRange) {
	o.CustomerAsnRange = &v
}

// GetEquinixAsn returns the EquinixAsn field value
func (o *RoutingProtocolBgp) GetEquinixAsn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EquinixAsn
}

// GetEquinixAsnOk returns a tuple with the EquinixAsn field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBgp) GetEquinixAsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EquinixAsn, true
}

// SetEquinixAsn sets field value
func (o *RoutingProtocolBgp) SetEquinixAsn(v int64) {
	o.EquinixAsn = v
}

// GetBgpAuthKey returns the BgpAuthKey field value if set, zero value otherwise.
func (o *RoutingProtocolBgp) GetBgpAuthKey() string {
	if o == nil || IsNil(o.BgpAuthKey) {
		var ret string
		return ret
	}
	return *o.BgpAuthKey
}

// GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBgp) GetBgpAuthKeyOk() (*string, bool) {
	if o == nil || IsNil(o.BgpAuthKey) {
		return nil, false
	}
	return o.BgpAuthKey, true
}

// HasBgpAuthKey returns a boolean if a field has been set.
func (o *RoutingProtocolBgp) HasBgpAuthKey() bool {
	if o != nil && !IsNil(o.BgpAuthKey) {
		return true
	}

	return false
}

// SetBgpAuthKey gets a reference to the given string and assigns it to the BgpAuthKey field.
func (o *RoutingProtocolBgp) SetBgpAuthKey(v string) {
	o.BgpAuthKey = &v
}

// GetExportPolicy returns the ExportPolicy field value
func (o *RoutingProtocolBgp) GetExportPolicy() ExportPolicy {
	if o == nil {
		var ret ExportPolicy
		return ret
	}

	return o.ExportPolicy
}

// GetExportPolicyOk returns a tuple with the ExportPolicy field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBgp) GetExportPolicyOk() (*ExportPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportPolicy, true
}

// SetExportPolicy sets field value
func (o *RoutingProtocolBgp) SetExportPolicy(v ExportPolicy) {
	o.ExportPolicy = v
}

func (o RoutingProtocolBgp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolBgp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedRoutingProtocolReadModel, errRoutingProtocolReadModel := json.Marshal(o.RoutingProtocolReadModel)
	if errRoutingProtocolReadModel != nil {
		return map[string]interface{}{}, errRoutingProtocolReadModel
	}
	errRoutingProtocolReadModel = json.Unmarshal([]byte(serializedRoutingProtocolReadModel), &toSerialize)
	if errRoutingProtocolReadModel != nil {
		return map[string]interface{}{}, errRoutingProtocolReadModel
	}
	toSerialize["customerAsn"] = o.CustomerAsn
	if !IsNil(o.CustomerAsnRange) {
		toSerialize["customerAsnRange"] = o.CustomerAsnRange
	}
	toSerialize["equinixAsn"] = o.EquinixAsn
	if !IsNil(o.BgpAuthKey) {
		toSerialize["bgpAuthKey"] = o.BgpAuthKey
	}
	toSerialize["exportPolicy"] = o.ExportPolicy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolBgp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerAsn",
		"equinixAsn",
		"exportPolicy",
		"type",
		"name",
		"changeLog",
		"links",
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

	type RoutingProtocolBgpWithoutEmbeddedStruct struct {
		// Customer Autonomous System Number
		CustomerAsn      int64                                           `json:"customerAsn"`
		CustomerAsnRange *BgpRoutingProtocolRequestAllOfCustomerAsnRange `json:"customerAsnRange,omitempty"`
		// Equinix Autonomous System Number
		EquinixAsn int64 `json:"equinixAsn"`
		// BGP authentication key
		BgpAuthKey   *string      `json:"bgpAuthKey,omitempty"`
		ExportPolicy ExportPolicy `json:"exportPolicy"`
	}

	varRoutingProtocolBgpWithoutEmbeddedStruct := RoutingProtocolBgpWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRoutingProtocolBgpWithoutEmbeddedStruct)
	if err == nil {
		varRoutingProtocolBgp := _RoutingProtocolBgp{}
		varRoutingProtocolBgp.CustomerAsn = varRoutingProtocolBgpWithoutEmbeddedStruct.CustomerAsn
		varRoutingProtocolBgp.CustomerAsnRange = varRoutingProtocolBgpWithoutEmbeddedStruct.CustomerAsnRange
		varRoutingProtocolBgp.EquinixAsn = varRoutingProtocolBgpWithoutEmbeddedStruct.EquinixAsn
		varRoutingProtocolBgp.BgpAuthKey = varRoutingProtocolBgpWithoutEmbeddedStruct.BgpAuthKey
		varRoutingProtocolBgp.ExportPolicy = varRoutingProtocolBgpWithoutEmbeddedStruct.ExportPolicy
		*o = RoutingProtocolBgp(varRoutingProtocolBgp)
	} else {
		return err
	}

	varRoutingProtocolBgp := _RoutingProtocolBgp{}

	err = json.Unmarshal(data, &varRoutingProtocolBgp)
	if err == nil {
		o.RoutingProtocolReadModel = varRoutingProtocolBgp.RoutingProtocolReadModel
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerAsn")
		delete(additionalProperties, "customerAsnRange")
		delete(additionalProperties, "equinixAsn")
		delete(additionalProperties, "bgpAuthKey")
		delete(additionalProperties, "exportPolicy")

		// remove fields from embedded structs
		reflectRoutingProtocolReadModel := reflect.ValueOf(o.RoutingProtocolReadModel)
		for i := 0; i < reflectRoutingProtocolReadModel.Type().NumField(); i++ {
			t := reflectRoutingProtocolReadModel.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolBgp struct {
	value *RoutingProtocolBgp
	isSet bool
}

func (v NullableRoutingProtocolBgp) Get() *RoutingProtocolBgp {
	return v.value
}

func (v *NullableRoutingProtocolBgp) Set(val *RoutingProtocolBgp) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolBgp) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolBgp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolBgp(val *RoutingProtocolBgp) *NullableRoutingProtocolBgp {
	return &NullableRoutingProtocolBgp{value: val, isSet: true}
}

func (v NullableRoutingProtocolBgp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolBgp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
