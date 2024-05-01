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

// OrderSignatureSignatory the model 'OrderSignatureSignatory'
type OrderSignatureSignatory string

// List of OrderSignature_signatory
const (
	ORDERSIGNATURESIGNATORY_SELF     OrderSignatureSignatory = "SELF"
	ORDERSIGNATURESIGNATORY_DELEGATE OrderSignatureSignatory = "DELEGATE"
	ORDERSIGNATURESIGNATORY_SUPPORT  OrderSignatureSignatory = "SUPPORT"
)

// All allowed values of OrderSignatureSignatory enum
var AllowedOrderSignatureSignatoryEnumValues = []OrderSignatureSignatory{
	"SELF",
	"DELEGATE",
	"SUPPORT",
}

func (v *OrderSignatureSignatory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderSignatureSignatory(value)
	for _, existing := range AllowedOrderSignatureSignatoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderSignatureSignatory", value)
}

// NewOrderSignatureSignatoryFromValue returns a pointer to a valid OrderSignatureSignatory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderSignatureSignatoryFromValue(v string) (*OrderSignatureSignatory, error) {
	ev := OrderSignatureSignatory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderSignatureSignatory: valid values are %v", v, AllowedOrderSignatureSignatoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderSignatureSignatory) IsValid() bool {
	for _, existing := range AllowedOrderSignatureSignatoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderSignature_signatory value
func (v OrderSignatureSignatory) Ptr() *OrderSignatureSignatory {
	return &v
}

type NullableOrderSignatureSignatory struct {
	value *OrderSignatureSignatory
	isSet bool
}

func (v NullableOrderSignatureSignatory) Get() *OrderSignatureSignatory {
	return v.value
}

func (v *NullableOrderSignatureSignatory) Set(val *OrderSignatureSignatory) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSignatureSignatory) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSignatureSignatory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSignatureSignatory(val *OrderSignatureSignatory) *NullableOrderSignatureSignatory {
	return &NullableOrderSignatureSignatory{value: val, isSet: true}
}

func (v NullableOrderSignatureSignatory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSignatureSignatory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
