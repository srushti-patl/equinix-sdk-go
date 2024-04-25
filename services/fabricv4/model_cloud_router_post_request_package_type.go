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

// CloudRouterPostRequestPackageType Cloud Router package type
type CloudRouterPostRequestPackageType string

// List of CloudRouterPostRequestPackage_type
const (
	CLOUDROUTERPOSTREQUESTPACKAGETYPE_ROUTER_PACKAGE CloudRouterPostRequestPackageType = "ROUTER_PACKAGE"
)

// All allowed values of CloudRouterPostRequestPackageType enum
var AllowedCloudRouterPostRequestPackageTypeEnumValues = []CloudRouterPostRequestPackageType{
	"ROUTER_PACKAGE",
}

func (v *CloudRouterPostRequestPackageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudRouterPostRequestPackageType(value)
	for _, existing := range AllowedCloudRouterPostRequestPackageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudRouterPostRequestPackageType", value)
}

// NewCloudRouterPostRequestPackageTypeFromValue returns a pointer to a valid CloudRouterPostRequestPackageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudRouterPostRequestPackageTypeFromValue(v string) (*CloudRouterPostRequestPackageType, error) {
	ev := CloudRouterPostRequestPackageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudRouterPostRequestPackageType: valid values are %v", v, AllowedCloudRouterPostRequestPackageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudRouterPostRequestPackageType) IsValid() bool {
	for _, existing := range AllowedCloudRouterPostRequestPackageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudRouterPostRequestPackage_type value
func (v CloudRouterPostRequestPackageType) Ptr() *CloudRouterPostRequestPackageType {
	return &v
}

type NullableCloudRouterPostRequestPackageType struct {
	value *CloudRouterPostRequestPackageType
	isSet bool
}

func (v NullableCloudRouterPostRequestPackageType) Get() *CloudRouterPostRequestPackageType {
	return v.value
}

func (v *NullableCloudRouterPostRequestPackageType) Set(val *CloudRouterPostRequestPackageType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterPostRequestPackageType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterPostRequestPackageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterPostRequestPackageType(val *CloudRouterPostRequestPackageType) *NullableCloudRouterPostRequestPackageType {
	return &NullableCloudRouterPostRequestPackageType{value: val, isSet: true}
}

func (v NullableCloudRouterPostRequestPackageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterPostRequestPackageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
