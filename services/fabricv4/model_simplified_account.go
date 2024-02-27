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
)

// checks if the SimplifiedAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedAccount{}

// SimplifiedAccount struct for SimplifiedAccount
type SimplifiedAccount struct {
	// Account number
	AccountNumber *int64 `json:"accountNumber,omitempty"`
	// Account name
	AccountName *string `json:"accountName,omitempty"`
	// Customer organization identifier
	OrgId *int64 `json:"orgId,omitempty"`
	// Customer organization name
	OrganizationName *string `json:"organizationName,omitempty"`
	// Global organization identifier
	GlobalOrgId *string `json:"globalOrgId,omitempty"`
	// Global organization name
	GlobalOrganizationName *string `json:"globalOrganizationName,omitempty"`
	// Account ucmId
	UcmId *string `json:"ucmId,omitempty"`
	// Account name
	GlobalCustId *string `json:"globalCustId,omitempty"`
	// Reseller account number
	ResellerAccountNumber *int64 `json:"resellerAccountNumber,omitempty"`
	// Reseller account name
	ResellerAccountName *string `json:"resellerAccountName,omitempty"`
	// Reseller account ucmId
	ResellerUcmId *string `json:"resellerUcmId,omitempty"`
	// Reseller customer organization identifier
	ResellerOrgId        *int64 `json:"resellerOrgId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedAccount SimplifiedAccount

// NewSimplifiedAccount instantiates a new SimplifiedAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedAccount() *SimplifiedAccount {
	this := SimplifiedAccount{}
	return &this
}

// NewSimplifiedAccountWithDefaults instantiates a new SimplifiedAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedAccountWithDefaults() *SimplifiedAccount {
	this := SimplifiedAccount{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetAccountNumber() int64 {
	if o == nil || IsNil(o.AccountNumber) {
		var ret int64
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetAccountNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given int64 and assigns it to the AccountNumber field.
func (o *SimplifiedAccount) SetAccountNumber(v int64) {
	o.AccountNumber = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *SimplifiedAccount) SetAccountName(v string) {
	o.AccountName = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetOrgId() int64 {
	if o == nil || IsNil(o.OrgId) {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetOrgIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *SimplifiedAccount) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *SimplifiedAccount) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetGlobalOrgId returns the GlobalOrgId field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetGlobalOrgId() string {
	if o == nil || IsNil(o.GlobalOrgId) {
		var ret string
		return ret
	}
	return *o.GlobalOrgId
}

// GetGlobalOrgIdOk returns a tuple with the GlobalOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetGlobalOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalOrgId) {
		return nil, false
	}
	return o.GlobalOrgId, true
}

// HasGlobalOrgId returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasGlobalOrgId() bool {
	if o != nil && !IsNil(o.GlobalOrgId) {
		return true
	}

	return false
}

// SetGlobalOrgId gets a reference to the given string and assigns it to the GlobalOrgId field.
func (o *SimplifiedAccount) SetGlobalOrgId(v string) {
	o.GlobalOrgId = &v
}

// GetGlobalOrganizationName returns the GlobalOrganizationName field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetGlobalOrganizationName() string {
	if o == nil || IsNil(o.GlobalOrganizationName) {
		var ret string
		return ret
	}
	return *o.GlobalOrganizationName
}

// GetGlobalOrganizationNameOk returns a tuple with the GlobalOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetGlobalOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalOrganizationName) {
		return nil, false
	}
	return o.GlobalOrganizationName, true
}

// HasGlobalOrganizationName returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasGlobalOrganizationName() bool {
	if o != nil && !IsNil(o.GlobalOrganizationName) {
		return true
	}

	return false
}

// SetGlobalOrganizationName gets a reference to the given string and assigns it to the GlobalOrganizationName field.
func (o *SimplifiedAccount) SetGlobalOrganizationName(v string) {
	o.GlobalOrganizationName = &v
}

// GetUcmId returns the UcmId field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetUcmId() string {
	if o == nil || IsNil(o.UcmId) {
		var ret string
		return ret
	}
	return *o.UcmId
}

// GetUcmIdOk returns a tuple with the UcmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetUcmIdOk() (*string, bool) {
	if o == nil || IsNil(o.UcmId) {
		return nil, false
	}
	return o.UcmId, true
}

// HasUcmId returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasUcmId() bool {
	if o != nil && !IsNil(o.UcmId) {
		return true
	}

	return false
}

// SetUcmId gets a reference to the given string and assigns it to the UcmId field.
func (o *SimplifiedAccount) SetUcmId(v string) {
	o.UcmId = &v
}

// GetGlobalCustId returns the GlobalCustId field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetGlobalCustId() string {
	if o == nil || IsNil(o.GlobalCustId) {
		var ret string
		return ret
	}
	return *o.GlobalCustId
}

// GetGlobalCustIdOk returns a tuple with the GlobalCustId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetGlobalCustIdOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalCustId) {
		return nil, false
	}
	return o.GlobalCustId, true
}

// HasGlobalCustId returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasGlobalCustId() bool {
	if o != nil && !IsNil(o.GlobalCustId) {
		return true
	}

	return false
}

// SetGlobalCustId gets a reference to the given string and assigns it to the GlobalCustId field.
func (o *SimplifiedAccount) SetGlobalCustId(v string) {
	o.GlobalCustId = &v
}

// GetResellerAccountNumber returns the ResellerAccountNumber field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetResellerAccountNumber() int64 {
	if o == nil || IsNil(o.ResellerAccountNumber) {
		var ret int64
		return ret
	}
	return *o.ResellerAccountNumber
}

// GetResellerAccountNumberOk returns a tuple with the ResellerAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetResellerAccountNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.ResellerAccountNumber) {
		return nil, false
	}
	return o.ResellerAccountNumber, true
}

// HasResellerAccountNumber returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasResellerAccountNumber() bool {
	if o != nil && !IsNil(o.ResellerAccountNumber) {
		return true
	}

	return false
}

// SetResellerAccountNumber gets a reference to the given int64 and assigns it to the ResellerAccountNumber field.
func (o *SimplifiedAccount) SetResellerAccountNumber(v int64) {
	o.ResellerAccountNumber = &v
}

// GetResellerAccountName returns the ResellerAccountName field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetResellerAccountName() string {
	if o == nil || IsNil(o.ResellerAccountName) {
		var ret string
		return ret
	}
	return *o.ResellerAccountName
}

// GetResellerAccountNameOk returns a tuple with the ResellerAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetResellerAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResellerAccountName) {
		return nil, false
	}
	return o.ResellerAccountName, true
}

// HasResellerAccountName returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasResellerAccountName() bool {
	if o != nil && !IsNil(o.ResellerAccountName) {
		return true
	}

	return false
}

// SetResellerAccountName gets a reference to the given string and assigns it to the ResellerAccountName field.
func (o *SimplifiedAccount) SetResellerAccountName(v string) {
	o.ResellerAccountName = &v
}

// GetResellerUcmId returns the ResellerUcmId field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetResellerUcmId() string {
	if o == nil || IsNil(o.ResellerUcmId) {
		var ret string
		return ret
	}
	return *o.ResellerUcmId
}

// GetResellerUcmIdOk returns a tuple with the ResellerUcmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetResellerUcmIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResellerUcmId) {
		return nil, false
	}
	return o.ResellerUcmId, true
}

// HasResellerUcmId returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasResellerUcmId() bool {
	if o != nil && !IsNil(o.ResellerUcmId) {
		return true
	}

	return false
}

// SetResellerUcmId gets a reference to the given string and assigns it to the ResellerUcmId field.
func (o *SimplifiedAccount) SetResellerUcmId(v string) {
	o.ResellerUcmId = &v
}

// GetResellerOrgId returns the ResellerOrgId field value if set, zero value otherwise.
func (o *SimplifiedAccount) GetResellerOrgId() int64 {
	if o == nil || IsNil(o.ResellerOrgId) {
		var ret int64
		return ret
	}
	return *o.ResellerOrgId
}

// GetResellerOrgIdOk returns a tuple with the ResellerOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAccount) GetResellerOrgIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ResellerOrgId) {
		return nil, false
	}
	return o.ResellerOrgId, true
}

// HasResellerOrgId returns a boolean if a field has been set.
func (o *SimplifiedAccount) HasResellerOrgId() bool {
	if o != nil && !IsNil(o.ResellerOrgId) {
		return true
	}

	return false
}

// SetResellerOrgId gets a reference to the given int64 and assigns it to the ResellerOrgId field.
func (o *SimplifiedAccount) SetResellerOrgId(v int64) {
	o.ResellerOrgId = &v
}

func (o SimplifiedAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.GlobalOrgId) {
		toSerialize["globalOrgId"] = o.GlobalOrgId
	}
	if !IsNil(o.GlobalOrganizationName) {
		toSerialize["globalOrganizationName"] = o.GlobalOrganizationName
	}
	if !IsNil(o.UcmId) {
		toSerialize["ucmId"] = o.UcmId
	}
	if !IsNil(o.GlobalCustId) {
		toSerialize["globalCustId"] = o.GlobalCustId
	}
	if !IsNil(o.ResellerAccountNumber) {
		toSerialize["resellerAccountNumber"] = o.ResellerAccountNumber
	}
	if !IsNil(o.ResellerAccountName) {
		toSerialize["resellerAccountName"] = o.ResellerAccountName
	}
	if !IsNil(o.ResellerUcmId) {
		toSerialize["resellerUcmId"] = o.ResellerUcmId
	}
	if !IsNil(o.ResellerOrgId) {
		toSerialize["resellerOrgId"] = o.ResellerOrgId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedAccount) UnmarshalJSON(data []byte) (err error) {
	varSimplifiedAccount := _SimplifiedAccount{}

	err = json.Unmarshal(data, &varSimplifiedAccount)

	if err != nil {
		return err
	}

	*o = SimplifiedAccount(varSimplifiedAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountNumber")
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "organizationName")
		delete(additionalProperties, "globalOrgId")
		delete(additionalProperties, "globalOrganizationName")
		delete(additionalProperties, "ucmId")
		delete(additionalProperties, "globalCustId")
		delete(additionalProperties, "resellerAccountNumber")
		delete(additionalProperties, "resellerAccountName")
		delete(additionalProperties, "resellerUcmId")
		delete(additionalProperties, "resellerOrgId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedAccount struct {
	value *SimplifiedAccount
	isSet bool
}

func (v NullableSimplifiedAccount) Get() *SimplifiedAccount {
	return v.value
}

func (v *NullableSimplifiedAccount) Set(val *SimplifiedAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedAccount(val *SimplifiedAccount) *NullableSimplifiedAccount {
	return &NullableSimplifiedAccount{value: val, isSet: true}
}

func (v NullableSimplifiedAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
