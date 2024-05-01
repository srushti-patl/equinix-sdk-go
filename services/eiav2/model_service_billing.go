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

// ServiceBilling Billing type of service
type ServiceBilling string

// List of ServiceBilling
const (
	SERVICEBILLING_FIXED       ServiceBilling = "FIXED"
	SERVICEBILLING_USAGE_BASED ServiceBilling = "USAGE_BASED"
	SERVICEBILLING_BURST_BASED ServiceBilling = "BURST_BASED"
)

// All allowed values of ServiceBilling enum
var AllowedServiceBillingEnumValues = []ServiceBilling{
	"FIXED",
	"USAGE_BASED",
	"BURST_BASED",
}

func (v *ServiceBilling) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceBilling(value)
	for _, existing := range AllowedServiceBillingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceBilling", value)
}

// NewServiceBillingFromValue returns a pointer to a valid ServiceBilling
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceBillingFromValue(v string) (*ServiceBilling, error) {
	ev := ServiceBilling(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceBilling: valid values are %v", v, AllowedServiceBillingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceBilling) IsValid() bool {
	for _, existing := range AllowedServiceBillingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceBilling value
func (v ServiceBilling) Ptr() *ServiceBilling {
	return &v
}

type NullableServiceBilling struct {
	value *ServiceBilling
	isSet bool
}

func (v NullableServiceBilling) Get() *ServiceBilling {
	return v.value
}

func (v *NullableServiceBilling) Set(val *ServiceBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBilling(val *ServiceBilling) *NullableServiceBilling {
	return &NullableServiceBilling{value: val, isSet: true}
}

func (v NullableServiceBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
