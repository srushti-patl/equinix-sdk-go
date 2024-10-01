/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PtpAdvanceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PtpAdvanceConfiguration{}

// PtpAdvanceConfiguration PTP Advanced Configuration.
type PtpAdvanceConfiguration struct {
	TimeScale *PtpAdvanceConfigurationTimeScale `json:"timeScale,omitempty"`
	// The PTP domain value.
	Domain *int32 `json:"domain,omitempty"`
	// The priority1 value determines the best primary clock, Lower value indicates higher priority.
	Priority1 *int32 `json:"priority1,omitempty"`
	// The priority2 value differentiates and prioritizes the primary clock to avoid confusion when priority1-value is the same for different primary clocks in a network.
	Priority2           *int32                                      `json:"priority2,omitempty"`
	LogAnnounceInterval *PtpAdvanceConfigurationLogAnnounceInterval `json:"logAnnounceInterval,omitempty"`
	LogSyncInterval     *PtpAdvanceConfigurationLogSyncInterval     `json:"logSyncInterval,omitempty"`
	LogDelayReqInterval *PtpAdvanceConfigurationLogDelayReqInterval `json:"logDelayReqInterval,omitempty"`
	TransportMode       *PtpAdvanceConfigurationTransportMode       `json:"transportMode,omitempty"`
	// Unicast Grant Time in seconds. For Multicast and Hybrid transport modes, grant time defaults to 300 seconds. For Unicast mode, grant time can be between 30 to 7200.
	GrantTime            *int32 `json:"grantTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PtpAdvanceConfiguration PtpAdvanceConfiguration

// NewPtpAdvanceConfiguration instantiates a new PtpAdvanceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPtpAdvanceConfiguration() *PtpAdvanceConfiguration {
	this := PtpAdvanceConfiguration{}
	return &this
}

// NewPtpAdvanceConfigurationWithDefaults instantiates a new PtpAdvanceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPtpAdvanceConfigurationWithDefaults() *PtpAdvanceConfiguration {
	this := PtpAdvanceConfiguration{}
	return &this
}

// GetTimeScale returns the TimeScale field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetTimeScale() PtpAdvanceConfigurationTimeScale {
	if o == nil || IsNil(o.TimeScale) {
		var ret PtpAdvanceConfigurationTimeScale
		return ret
	}
	return *o.TimeScale
}

// GetTimeScaleOk returns a tuple with the TimeScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetTimeScaleOk() (*PtpAdvanceConfigurationTimeScale, bool) {
	if o == nil || IsNil(o.TimeScale) {
		return nil, false
	}
	return o.TimeScale, true
}

// HasTimeScale returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasTimeScale() bool {
	if o != nil && !IsNil(o.TimeScale) {
		return true
	}

	return false
}

// SetTimeScale gets a reference to the given PtpAdvanceConfigurationTimeScale and assigns it to the TimeScale field.
func (o *PtpAdvanceConfiguration) SetTimeScale(v PtpAdvanceConfigurationTimeScale) {
	o.TimeScale = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetDomain() int32 {
	if o == nil || IsNil(o.Domain) {
		var ret int32
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetDomainOk() (*int32, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given int32 and assigns it to the Domain field.
func (o *PtpAdvanceConfiguration) SetDomain(v int32) {
	o.Domain = &v
}

// GetPriority1 returns the Priority1 field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetPriority1() int32 {
	if o == nil || IsNil(o.Priority1) {
		var ret int32
		return ret
	}
	return *o.Priority1
}

// GetPriority1Ok returns a tuple with the Priority1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetPriority1Ok() (*int32, bool) {
	if o == nil || IsNil(o.Priority1) {
		return nil, false
	}
	return o.Priority1, true
}

// HasPriority1 returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasPriority1() bool {
	if o != nil && !IsNil(o.Priority1) {
		return true
	}

	return false
}

// SetPriority1 gets a reference to the given int32 and assigns it to the Priority1 field.
func (o *PtpAdvanceConfiguration) SetPriority1(v int32) {
	o.Priority1 = &v
}

// GetPriority2 returns the Priority2 field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetPriority2() int32 {
	if o == nil || IsNil(o.Priority2) {
		var ret int32
		return ret
	}
	return *o.Priority2
}

// GetPriority2Ok returns a tuple with the Priority2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetPriority2Ok() (*int32, bool) {
	if o == nil || IsNil(o.Priority2) {
		return nil, false
	}
	return o.Priority2, true
}

// HasPriority2 returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasPriority2() bool {
	if o != nil && !IsNil(o.Priority2) {
		return true
	}

	return false
}

// SetPriority2 gets a reference to the given int32 and assigns it to the Priority2 field.
func (o *PtpAdvanceConfiguration) SetPriority2(v int32) {
	o.Priority2 = &v
}

// GetLogAnnounceInterval returns the LogAnnounceInterval field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetLogAnnounceInterval() PtpAdvanceConfigurationLogAnnounceInterval {
	if o == nil || IsNil(o.LogAnnounceInterval) {
		var ret PtpAdvanceConfigurationLogAnnounceInterval
		return ret
	}
	return *o.LogAnnounceInterval
}

// GetLogAnnounceIntervalOk returns a tuple with the LogAnnounceInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetLogAnnounceIntervalOk() (*PtpAdvanceConfigurationLogAnnounceInterval, bool) {
	if o == nil || IsNil(o.LogAnnounceInterval) {
		return nil, false
	}
	return o.LogAnnounceInterval, true
}

// HasLogAnnounceInterval returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasLogAnnounceInterval() bool {
	if o != nil && !IsNil(o.LogAnnounceInterval) {
		return true
	}

	return false
}

// SetLogAnnounceInterval gets a reference to the given PtpAdvanceConfigurationLogAnnounceInterval and assigns it to the LogAnnounceInterval field.
func (o *PtpAdvanceConfiguration) SetLogAnnounceInterval(v PtpAdvanceConfigurationLogAnnounceInterval) {
	o.LogAnnounceInterval = &v
}

// GetLogSyncInterval returns the LogSyncInterval field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetLogSyncInterval() PtpAdvanceConfigurationLogSyncInterval {
	if o == nil || IsNil(o.LogSyncInterval) {
		var ret PtpAdvanceConfigurationLogSyncInterval
		return ret
	}
	return *o.LogSyncInterval
}

// GetLogSyncIntervalOk returns a tuple with the LogSyncInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetLogSyncIntervalOk() (*PtpAdvanceConfigurationLogSyncInterval, bool) {
	if o == nil || IsNil(o.LogSyncInterval) {
		return nil, false
	}
	return o.LogSyncInterval, true
}

// HasLogSyncInterval returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasLogSyncInterval() bool {
	if o != nil && !IsNil(o.LogSyncInterval) {
		return true
	}

	return false
}

// SetLogSyncInterval gets a reference to the given PtpAdvanceConfigurationLogSyncInterval and assigns it to the LogSyncInterval field.
func (o *PtpAdvanceConfiguration) SetLogSyncInterval(v PtpAdvanceConfigurationLogSyncInterval) {
	o.LogSyncInterval = &v
}

// GetLogDelayReqInterval returns the LogDelayReqInterval field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetLogDelayReqInterval() PtpAdvanceConfigurationLogDelayReqInterval {
	if o == nil || IsNil(o.LogDelayReqInterval) {
		var ret PtpAdvanceConfigurationLogDelayReqInterval
		return ret
	}
	return *o.LogDelayReqInterval
}

// GetLogDelayReqIntervalOk returns a tuple with the LogDelayReqInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetLogDelayReqIntervalOk() (*PtpAdvanceConfigurationLogDelayReqInterval, bool) {
	if o == nil || IsNil(o.LogDelayReqInterval) {
		return nil, false
	}
	return o.LogDelayReqInterval, true
}

// HasLogDelayReqInterval returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasLogDelayReqInterval() bool {
	if o != nil && !IsNil(o.LogDelayReqInterval) {
		return true
	}

	return false
}

// SetLogDelayReqInterval gets a reference to the given PtpAdvanceConfigurationLogDelayReqInterval and assigns it to the LogDelayReqInterval field.
func (o *PtpAdvanceConfiguration) SetLogDelayReqInterval(v PtpAdvanceConfigurationLogDelayReqInterval) {
	o.LogDelayReqInterval = &v
}

// GetTransportMode returns the TransportMode field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetTransportMode() PtpAdvanceConfigurationTransportMode {
	if o == nil || IsNil(o.TransportMode) {
		var ret PtpAdvanceConfigurationTransportMode
		return ret
	}
	return *o.TransportMode
}

// GetTransportModeOk returns a tuple with the TransportMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetTransportModeOk() (*PtpAdvanceConfigurationTransportMode, bool) {
	if o == nil || IsNil(o.TransportMode) {
		return nil, false
	}
	return o.TransportMode, true
}

// HasTransportMode returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasTransportMode() bool {
	if o != nil && !IsNil(o.TransportMode) {
		return true
	}

	return false
}

// SetTransportMode gets a reference to the given PtpAdvanceConfigurationTransportMode and assigns it to the TransportMode field.
func (o *PtpAdvanceConfiguration) SetTransportMode(v PtpAdvanceConfigurationTransportMode) {
	o.TransportMode = &v
}

// GetGrantTime returns the GrantTime field value if set, zero value otherwise.
func (o *PtpAdvanceConfiguration) GetGrantTime() int32 {
	if o == nil || IsNil(o.GrantTime) {
		var ret int32
		return ret
	}
	return *o.GrantTime
}

// GetGrantTimeOk returns a tuple with the GrantTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpAdvanceConfiguration) GetGrantTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.GrantTime) {
		return nil, false
	}
	return o.GrantTime, true
}

// HasGrantTime returns a boolean if a field has been set.
func (o *PtpAdvanceConfiguration) HasGrantTime() bool {
	if o != nil && !IsNil(o.GrantTime) {
		return true
	}

	return false
}

// SetGrantTime gets a reference to the given int32 and assigns it to the GrantTime field.
func (o *PtpAdvanceConfiguration) SetGrantTime(v int32) {
	o.GrantTime = &v
}

func (o PtpAdvanceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PtpAdvanceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeScale) {
		toSerialize["timeScale"] = o.TimeScale
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Priority1) {
		toSerialize["priority1"] = o.Priority1
	}
	if !IsNil(o.Priority2) {
		toSerialize["priority2"] = o.Priority2
	}
	if !IsNil(o.LogAnnounceInterval) {
		toSerialize["logAnnounceInterval"] = o.LogAnnounceInterval
	}
	if !IsNil(o.LogSyncInterval) {
		toSerialize["logSyncInterval"] = o.LogSyncInterval
	}
	if !IsNil(o.LogDelayReqInterval) {
		toSerialize["logDelayReqInterval"] = o.LogDelayReqInterval
	}
	if !IsNil(o.TransportMode) {
		toSerialize["transportMode"] = o.TransportMode
	}
	if !IsNil(o.GrantTime) {
		toSerialize["grantTime"] = o.GrantTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PtpAdvanceConfiguration) UnmarshalJSON(data []byte) (err error) {
	varPtpAdvanceConfiguration := _PtpAdvanceConfiguration{}

	err = json.Unmarshal(data, &varPtpAdvanceConfiguration)

	if err != nil {
		return err
	}

	*o = PtpAdvanceConfiguration(varPtpAdvanceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timeScale")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "priority1")
		delete(additionalProperties, "priority2")
		delete(additionalProperties, "logAnnounceInterval")
		delete(additionalProperties, "logSyncInterval")
		delete(additionalProperties, "logDelayReqInterval")
		delete(additionalProperties, "transportMode")
		delete(additionalProperties, "grantTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePtpAdvanceConfiguration struct {
	value *PtpAdvanceConfiguration
	isSet bool
}

func (v NullablePtpAdvanceConfiguration) Get() *PtpAdvanceConfiguration {
	return v.value
}

func (v *NullablePtpAdvanceConfiguration) Set(val *PtpAdvanceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePtpAdvanceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePtpAdvanceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtpAdvanceConfiguration(val *PtpAdvanceConfiguration) *NullablePtpAdvanceConfiguration {
	return &NullablePtpAdvanceConfiguration{value: val, isSet: true}
}

func (v NullablePtpAdvanceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtpAdvanceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
