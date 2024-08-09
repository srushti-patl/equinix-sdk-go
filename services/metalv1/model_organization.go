/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"time"
)

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization struct for Organization
type Organization struct {
	Address        *Address               `json:"address,omitempty"`
	BillingAddress *Address               `json:"billing_address,omitempty"`
	CreatedAt      *time.Time             `json:"created_at,omitempty"`
	CreditAmount   *float32               `json:"credit_amount,omitempty"`
	Customdata     map[string]interface{} `json:"customdata,omitempty"`
	Description    *string                `json:"description,omitempty"`
	// Force to all members to have enabled the two factor authentication after that date, unless the value is null
	Enforce2faAt         *time.Time `json:"enforce_2fa_at,omitempty"`
	Id                   *string    `json:"id,omitempty"`
	Logo                 *string    `json:"logo,omitempty"`
	Members              []Href     `json:"members,omitempty"`
	Memberships          []Href     `json:"memberships,omitempty"`
	Name                 *string    `json:"name,omitempty"`
	Projects             []Href     `json:"projects,omitempty"`
	Terms                *int32     `json:"terms,omitempty"`
	Twitter              *string    `json:"twitter,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	Website              *string    `json:"website,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Organization Organization

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Organization) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Organization) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Organization) SetAddress(v Address) {
	o.Address = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *Organization) GetBillingAddress() Address {
	if o == nil || IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetBillingAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *Organization) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *Organization) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Organization) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Organization) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreditAmount returns the CreditAmount field value if set, zero value otherwise.
func (o *Organization) GetCreditAmount() float32 {
	if o == nil || IsNil(o.CreditAmount) {
		var ret float32
		return ret
	}
	return *o.CreditAmount
}

// GetCreditAmountOk returns a tuple with the CreditAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCreditAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditAmount) {
		return nil, false
	}
	return o.CreditAmount, true
}

// HasCreditAmount returns a boolean if a field has been set.
func (o *Organization) HasCreditAmount() bool {
	if o != nil && !IsNil(o.CreditAmount) {
		return true
	}

	return false
}

// SetCreditAmount gets a reference to the given float32 and assigns it to the CreditAmount field.
func (o *Organization) SetCreditAmount(v float32) {
	o.CreditAmount = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *Organization) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *Organization) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *Organization) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Organization) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Organization) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Organization) SetDescription(v string) {
	o.Description = &v
}

// GetEnforce2faAt returns the Enforce2faAt field value if set, zero value otherwise.
func (o *Organization) GetEnforce2faAt() time.Time {
	if o == nil || IsNil(o.Enforce2faAt) {
		var ret time.Time
		return ret
	}
	return *o.Enforce2faAt
}

// GetEnforce2faAtOk returns a tuple with the Enforce2faAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetEnforce2faAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Enforce2faAt) {
		return nil, false
	}
	return o.Enforce2faAt, true
}

// HasEnforce2faAt returns a boolean if a field has been set.
func (o *Organization) HasEnforce2faAt() bool {
	if o != nil && !IsNil(o.Enforce2faAt) {
		return true
	}

	return false
}

// SetEnforce2faAt gets a reference to the given time.Time and assigns it to the Enforce2faAt field.
func (o *Organization) SetEnforce2faAt(v time.Time) {
	o.Enforce2faAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Organization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Organization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Organization) SetId(v string) {
	o.Id = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *Organization) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *Organization) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *Organization) SetLogo(v string) {
	o.Logo = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Organization) GetMembers() []Href {
	if o == nil || IsNil(o.Members) {
		var ret []Href
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetMembersOk() ([]Href, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Organization) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Href and assigns it to the Members field.
func (o *Organization) SetMembers(v []Href) {
	o.Members = v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *Organization) GetMemberships() []Href {
	if o == nil || IsNil(o.Memberships) {
		var ret []Href
		return ret
	}
	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetMembershipsOk() ([]Href, bool) {
	if o == nil || IsNil(o.Memberships) {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *Organization) HasMemberships() bool {
	if o != nil && !IsNil(o.Memberships) {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []Href and assigns it to the Memberships field.
func (o *Organization) SetMemberships(v []Href) {
	o.Memberships = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organization) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organization) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organization) SetName(v string) {
	o.Name = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *Organization) GetProjects() []Href {
	if o == nil || IsNil(o.Projects) {
		var ret []Href
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetProjectsOk() ([]Href, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *Organization) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []Href and assigns it to the Projects field.
func (o *Organization) SetProjects(v []Href) {
	o.Projects = v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *Organization) GetTerms() int32 {
	if o == nil || IsNil(o.Terms) {
		var ret int32
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetTermsOk() (*int32, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *Organization) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given int32 and assigns it to the Terms field.
func (o *Organization) SetTerms(v int32) {
	o.Terms = &v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *Organization) GetTwitter() string {
	if o == nil || IsNil(o.Twitter) {
		var ret string
		return ret
	}
	return *o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetTwitterOk() (*string, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *Organization) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given string and assigns it to the Twitter field.
func (o *Organization) SetTwitter(v string) {
	o.Twitter = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Organization) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Organization) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Organization) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Organization) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Organization) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Organization) SetWebsite(v string) {
	o.Website = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreditAmount) {
		toSerialize["credit_amount"] = o.CreditAmount
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enforce2faAt) {
		toSerialize["enforce_2fa_at"] = o.Enforce2faAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Memberships) {
		toSerialize["memberships"] = o.Memberships
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !IsNil(o.Twitter) {
		toSerialize["twitter"] = o.Twitter
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Organization) UnmarshalJSON(data []byte) (err error) {
	varOrganization := _Organization{}

	err = json.Unmarshal(data, &varOrganization)

	if err != nil {
		return err
	}

	*o = Organization(varOrganization)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "billing_address")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "credit_amount")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enforce_2fa_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "members")
		delete(additionalProperties, "memberships")
		delete(additionalProperties, "name")
		delete(additionalProperties, "projects")
		delete(additionalProperties, "terms")
		delete(additionalProperties, "twitter")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "website")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
