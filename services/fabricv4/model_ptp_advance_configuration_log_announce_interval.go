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

// PtpAdvanceConfigurationLogAnnounceInterval Logarithmic value that controls the rate of PTP Announce packets from the PTP time server. Default is 1 (1 packet every 2 seconds), Unit packets/second.
type PtpAdvanceConfigurationLogAnnounceInterval int32

// List of ptpAdvanceConfiguration_logAnnounceInterval
const (
	PTPADVANCECONFIGURATIONLOGANNOUNCEINTERVAL__MINUS_3 PtpAdvanceConfigurationLogAnnounceInterval = -3
	PTPADVANCECONFIGURATIONLOGANNOUNCEINTERVAL__MINUS_2 PtpAdvanceConfigurationLogAnnounceInterval = -2
	PTPADVANCECONFIGURATIONLOGANNOUNCEINTERVAL__MINUS_1 PtpAdvanceConfigurationLogAnnounceInterval = -1
	PTPADVANCECONFIGURATIONLOGANNOUNCEINTERVAL__0       PtpAdvanceConfigurationLogAnnounceInterval = 0
	PTPADVANCECONFIGURATIONLOGANNOUNCEINTERVAL__1       PtpAdvanceConfigurationLogAnnounceInterval = 1
)

// All allowed values of PtpAdvanceConfigurationLogAnnounceInterval enum
var AllowedPtpAdvanceConfigurationLogAnnounceIntervalEnumValues = []PtpAdvanceConfigurationLogAnnounceInterval{
	-3,
	-2,
	-1,
	0,
	1,
}

func (v *PtpAdvanceConfigurationLogAnnounceInterval) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PtpAdvanceConfigurationLogAnnounceInterval(value)
	for _, existing := range AllowedPtpAdvanceConfigurationLogAnnounceIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PtpAdvanceConfigurationLogAnnounceInterval", value)
}

// NewPtpAdvanceConfigurationLogAnnounceIntervalFromValue returns a pointer to a valid PtpAdvanceConfigurationLogAnnounceInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPtpAdvanceConfigurationLogAnnounceIntervalFromValue(v int32) (*PtpAdvanceConfigurationLogAnnounceInterval, error) {
	ev := PtpAdvanceConfigurationLogAnnounceInterval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PtpAdvanceConfigurationLogAnnounceInterval: valid values are %v", v, AllowedPtpAdvanceConfigurationLogAnnounceIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PtpAdvanceConfigurationLogAnnounceInterval) IsValid() bool {
	for _, existing := range AllowedPtpAdvanceConfigurationLogAnnounceIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ptpAdvanceConfiguration_logAnnounceInterval value
func (v PtpAdvanceConfigurationLogAnnounceInterval) Ptr() *PtpAdvanceConfigurationLogAnnounceInterval {
	return &v
}

type NullablePtpAdvanceConfigurationLogAnnounceInterval struct {
	value *PtpAdvanceConfigurationLogAnnounceInterval
	isSet bool
}

func (v NullablePtpAdvanceConfigurationLogAnnounceInterval) Get() *PtpAdvanceConfigurationLogAnnounceInterval {
	return v.value
}

func (v *NullablePtpAdvanceConfigurationLogAnnounceInterval) Set(val *PtpAdvanceConfigurationLogAnnounceInterval) {
	v.value = val
	v.isSet = true
}

func (v NullablePtpAdvanceConfigurationLogAnnounceInterval) IsSet() bool {
	return v.isSet
}

func (v *NullablePtpAdvanceConfigurationLogAnnounceInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtpAdvanceConfigurationLogAnnounceInterval(val *PtpAdvanceConfigurationLogAnnounceInterval) *NullablePtpAdvanceConfigurationLogAnnounceInterval {
	return &NullablePtpAdvanceConfigurationLogAnnounceInterval{value: val, isSet: true}
}

func (v NullablePtpAdvanceConfigurationLogAnnounceInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtpAdvanceConfigurationLogAnnounceInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
