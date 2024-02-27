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

// ServiceProfileAccessPointCOLOType the model 'ServiceProfileAccessPointCOLOType'
type ServiceProfileAccessPointCOLOType string

// List of ServiceProfileAccessPointCOLO_type
const (
	SERVICEPROFILEACCESSPOINTCOLOTYPE_XF_PORT ServiceProfileAccessPointCOLOType = "XF_PORT"
)

// All allowed values of ServiceProfileAccessPointCOLOType enum
var AllowedServiceProfileAccessPointCOLOTypeEnumValues = []ServiceProfileAccessPointCOLOType{
	"XF_PORT",
}

func (v *ServiceProfileAccessPointCOLOType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceProfileAccessPointCOLOType(value)
	for _, existing := range AllowedServiceProfileAccessPointCOLOTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceProfileAccessPointCOLOType", value)
}

// NewServiceProfileAccessPointCOLOTypeFromValue returns a pointer to a valid ServiceProfileAccessPointCOLOType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceProfileAccessPointCOLOTypeFromValue(v string) (*ServiceProfileAccessPointCOLOType, error) {
	ev := ServiceProfileAccessPointCOLOType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceProfileAccessPointCOLOType: valid values are %v", v, AllowedServiceProfileAccessPointCOLOTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceProfileAccessPointCOLOType) IsValid() bool {
	for _, existing := range AllowedServiceProfileAccessPointCOLOTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceProfileAccessPointCOLO_type value
func (v ServiceProfileAccessPointCOLOType) Ptr() *ServiceProfileAccessPointCOLOType {
	return &v
}

type NullableServiceProfileAccessPointCOLOType struct {
	value *ServiceProfileAccessPointCOLOType
	isSet bool
}

func (v NullableServiceProfileAccessPointCOLOType) Get() *ServiceProfileAccessPointCOLOType {
	return v.value
}

func (v *NullableServiceProfileAccessPointCOLOType) Set(val *ServiceProfileAccessPointCOLOType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAccessPointCOLOType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAccessPointCOLOType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAccessPointCOLOType(val *ServiceProfileAccessPointCOLOType) *NullableServiceProfileAccessPointCOLOType {
	return &NullableServiceProfileAccessPointCOLOType{value: val, isSet: true}
}

func (v NullableServiceProfileAccessPointCOLOType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAccessPointCOLOType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
