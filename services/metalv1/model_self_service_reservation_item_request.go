/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the SelfServiceReservationItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceReservationItemRequest{}

// SelfServiceReservationItemRequest struct for SelfServiceReservationItemRequest
type SelfServiceReservationItemRequest struct {
	Amount               *float32 `json:"amount,omitempty"`
	MetroId              *string  `json:"metro_id,omitempty"`
	PlanId               *string  `json:"plan_id,omitempty"`
	Quantity             *int32   `json:"quantity,omitempty"`
	Term                 *string  `json:"term,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceReservationItemRequest SelfServiceReservationItemRequest

// NewSelfServiceReservationItemRequest instantiates a new SelfServiceReservationItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceReservationItemRequest() *SelfServiceReservationItemRequest {
	this := SelfServiceReservationItemRequest{}
	return &this
}

// NewSelfServiceReservationItemRequestWithDefaults instantiates a new SelfServiceReservationItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceReservationItemRequestWithDefaults() *SelfServiceReservationItemRequest {
	this := SelfServiceReservationItemRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SelfServiceReservationItemRequest) SetAmount(v float32) {
	o.Amount = &v
}

// GetMetroId returns the MetroId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetMetroId() string {
	if o == nil || IsNil(o.MetroId) {
		var ret string
		return ret
	}
	return *o.MetroId
}

// GetMetroIdOk returns a tuple with the MetroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetMetroIdOk() (*string, bool) {
	if o == nil || IsNil(o.MetroId) {
		return nil, false
	}
	return o.MetroId, true
}

// HasMetroId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasMetroId() bool {
	if o != nil && !IsNil(o.MetroId) {
		return true
	}

	return false
}

// SetMetroId gets a reference to the given string and assigns it to the MetroId field.
func (o *SelfServiceReservationItemRequest) SetMetroId(v string) {
	o.MetroId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SelfServiceReservationItemRequest) SetPlanId(v string) {
	o.PlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SelfServiceReservationItemRequest) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SelfServiceReservationItemRequest) SetTerm(v string) {
	o.Term = &v
}

func (o SelfServiceReservationItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceReservationItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.MetroId) {
		toSerialize["metro_id"] = o.MetroId
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceReservationItemRequest) UnmarshalJSON(data []byte) (err error) {
	varSelfServiceReservationItemRequest := _SelfServiceReservationItemRequest{}

	err = json.Unmarshal(data, &varSelfServiceReservationItemRequest)

	if err != nil {
		return err
	}

	*o = SelfServiceReservationItemRequest(varSelfServiceReservationItemRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "metro_id")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "term")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceReservationItemRequest struct {
	value *SelfServiceReservationItemRequest
	isSet bool
}

func (v NullableSelfServiceReservationItemRequest) Get() *SelfServiceReservationItemRequest {
	return v.value
}

func (v *NullableSelfServiceReservationItemRequest) Set(val *SelfServiceReservationItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceReservationItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceReservationItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceReservationItemRequest(val *SelfServiceReservationItemRequest) *NullableSelfServiceReservationItemRequest {
	return &NullableSelfServiceReservationItemRequest{value: val, isSet: true}
}

func (v NullableSelfServiceReservationItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceReservationItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
