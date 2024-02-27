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

// CloudRouterPackageTypeType Cloud Router package type
type CloudRouterPackageTypeType string

// List of CloudRouterPackageType_type
const (
	CLOUDROUTERPACKAGETYPETYPE_ROUTER_PACKAGE CloudRouterPackageTypeType = "ROUTER_PACKAGE"
)

// All allowed values of CloudRouterPackageTypeType enum
var AllowedCloudRouterPackageTypeTypeEnumValues = []CloudRouterPackageTypeType{
	"ROUTER_PACKAGE",
}

func (v *CloudRouterPackageTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudRouterPackageTypeType(value)
	for _, existing := range AllowedCloudRouterPackageTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudRouterPackageTypeType", value)
}

// NewCloudRouterPackageTypeTypeFromValue returns a pointer to a valid CloudRouterPackageTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudRouterPackageTypeTypeFromValue(v string) (*CloudRouterPackageTypeType, error) {
	ev := CloudRouterPackageTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudRouterPackageTypeType: valid values are %v", v, AllowedCloudRouterPackageTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudRouterPackageTypeType) IsValid() bool {
	for _, existing := range AllowedCloudRouterPackageTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudRouterPackageType_type value
func (v CloudRouterPackageTypeType) Ptr() *CloudRouterPackageTypeType {
	return &v
}

type NullableCloudRouterPackageTypeType struct {
	value *CloudRouterPackageTypeType
	isSet bool
}

func (v NullableCloudRouterPackageTypeType) Get() *CloudRouterPackageTypeType {
	return v.value
}

func (v *NullableCloudRouterPackageTypeType) Set(val *CloudRouterPackageTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterPackageTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterPackageTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterPackageTypeType(val *CloudRouterPackageTypeType) *NullableCloudRouterPackageTypeType {
	return &NullableCloudRouterPackageTypeType{value: val, isSet: true}
}

func (v NullableCloudRouterPackageTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterPackageTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
