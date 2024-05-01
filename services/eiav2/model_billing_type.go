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
)

// BillingType Type of billing
type BillingType string

// List of BillingType
const (
	BILLINGTYPE_FIXED       BillingType = "FIXED"
	BILLINGTYPE_USAGE_BASED BillingType = "USAGE_BASED"
	BILLINGTYPE_BURST_BASED BillingType = "BURST_BASED"
)

// All allowed values of BillingType enum
var AllowedBillingTypeEnumValues = []BillingType{
	"FIXED",
	"USAGE_BASED",
	"BURST_BASED",
}

func (v *BillingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingType(value)
	for _, existing := range AllowedBillingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingType", value)
}

// NewBillingTypeFromValue returns a pointer to a valid BillingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingTypeFromValue(v string) (*BillingType, error) {
	ev := BillingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingType: valid values are %v", v, AllowedBillingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingType) IsValid() bool {
	for _, existing := range AllowedBillingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingType value
func (v BillingType) Ptr() *BillingType {
	return &v
}

type NullableBillingType struct {
	value *BillingType
	isSet bool
}

func (v NullableBillingType) Get() *BillingType {
	return v.value
}

func (v *NullableBillingType) Set(val *BillingType) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingType) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingType(val *BillingType) *NullableBillingType {
	return &NullableBillingType{value: val, isSet: true}
}

func (v NullableBillingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
