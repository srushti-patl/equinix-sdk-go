/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ServiceTokenSearchExpressionOperator the model 'ServiceTokenSearchExpressionOperator'
type ServiceTokenSearchExpressionOperator string

// List of ServiceTokenSearchExpression_operator
const (
	SERVICETOKENSEARCHEXPRESSIONOPERATOR_EQUAL ServiceTokenSearchExpressionOperator = "="
)

// All allowed values of ServiceTokenSearchExpressionOperator enum
var AllowedServiceTokenSearchExpressionOperatorEnumValues = []ServiceTokenSearchExpressionOperator{
	"=",
}

func (v *ServiceTokenSearchExpressionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceTokenSearchExpressionOperator(value)
	for _, existing := range AllowedServiceTokenSearchExpressionOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceTokenSearchExpressionOperator", value)
}

// NewServiceTokenSearchExpressionOperatorFromValue returns a pointer to a valid ServiceTokenSearchExpressionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTokenSearchExpressionOperatorFromValue(v string) (*ServiceTokenSearchExpressionOperator, error) {
	ev := ServiceTokenSearchExpressionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceTokenSearchExpressionOperator: valid values are %v", v, AllowedServiceTokenSearchExpressionOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceTokenSearchExpressionOperator) IsValid() bool {
	for _, existing := range AllowedServiceTokenSearchExpressionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceTokenSearchExpression_operator value
func (v ServiceTokenSearchExpressionOperator) Ptr() *ServiceTokenSearchExpressionOperator {
	return &v
}

type NullableServiceTokenSearchExpressionOperator struct {
	value *ServiceTokenSearchExpressionOperator
	isSet bool
}

func (v NullableServiceTokenSearchExpressionOperator) Get() *ServiceTokenSearchExpressionOperator {
	return v.value
}

func (v *NullableServiceTokenSearchExpressionOperator) Set(val *ServiceTokenSearchExpressionOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTokenSearchExpressionOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTokenSearchExpressionOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTokenSearchExpressionOperator(val *ServiceTokenSearchExpressionOperator) *NullableServiceTokenSearchExpressionOperator {
	return &NullableServiceTokenSearchExpressionOperator{value: val, isSet: true}
}

func (v NullableServiceTokenSearchExpressionOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTokenSearchExpressionOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
