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

// PortLoaType Loa type
type PortLoaType string

// List of PortLoa_type
const (
	PORTLOATYPE_CTR_LOA PortLoaType = "CTR_LOA"
)

// All allowed values of PortLoaType enum
var AllowedPortLoaTypeEnumValues = []PortLoaType{
	"CTR_LOA",
}

func (v *PortLoaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortLoaType(value)
	for _, existing := range AllowedPortLoaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortLoaType", value)
}

// NewPortLoaTypeFromValue returns a pointer to a valid PortLoaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortLoaTypeFromValue(v string) (*PortLoaType, error) {
	ev := PortLoaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortLoaType: valid values are %v", v, AllowedPortLoaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortLoaType) IsValid() bool {
	for _, existing := range AllowedPortLoaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortLoa_type value
func (v PortLoaType) Ptr() *PortLoaType {
	return &v
}

type NullablePortLoaType struct {
	value *PortLoaType
	isSet bool
}

func (v NullablePortLoaType) Get() *PortLoaType {
	return v.value
}

func (v *NullablePortLoaType) Set(val *PortLoaType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortLoaType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortLoaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortLoaType(val *PortLoaType) *NullablePortLoaType {
	return &NullablePortLoaType{value: val, isSet: true}
}

func (v NullablePortLoaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortLoaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
