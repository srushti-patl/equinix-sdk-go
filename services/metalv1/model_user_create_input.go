/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// checks if the UserCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCreateInput{}

// UserCreateInput struct for UserCreateInput
type UserCreateInput struct {
	Avatar               **os.File              `json:"avatar,omitempty"`
	CompanyName          *string                `json:"company_name,omitempty"`
	CompanyUrl           *string                `json:"company_url,omitempty"`
	Customdata           map[string]interface{} `json:"customdata,omitempty"`
	Emails               []EmailInput           `json:"emails"`
	FirstName            string                 `json:"first_name"`
	LastName             string                 `json:"last_name"`
	Level                *string                `json:"level,omitempty"`
	Password             *string                `json:"password,omitempty"`
	PhoneNumber          *string                `json:"phone_number,omitempty"`
	SocialAccounts       map[string]interface{} `json:"social_accounts,omitempty"`
	Timezone             *string                `json:"timezone,omitempty"`
	Title                *string                `json:"title,omitempty"`
	TwoFactorAuth        *string                `json:"two_factor_auth,omitempty"`
	VerifiedAt           *time.Time             `json:"verified_at,omitempty"`
	InvitationId         *string                `json:"invitation_id,omitempty"`
	Nonce                *string                `json:"nonce,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCreateInput UserCreateInput

// NewUserCreateInput instantiates a new UserCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateInput(emails []EmailInput, firstName string, lastName string) *UserCreateInput {
	this := UserCreateInput{}
	this.Emails = emails
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewUserCreateInputWithDefaults instantiates a new UserCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateInputWithDefaults() *UserCreateInput {
	this := UserCreateInput{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UserCreateInput) GetAvatar() *os.File {
	if o == nil || IsNil(o.Avatar) {
		var ret *os.File
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetAvatarOk() (**os.File, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UserCreateInput) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given *os.File and assigns it to the Avatar field.
func (o *UserCreateInput) SetAvatar(v *os.File) {
	o.Avatar = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UserCreateInput) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UserCreateInput) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UserCreateInput) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise.
func (o *UserCreateInput) GetCompanyUrl() string {
	if o == nil || IsNil(o.CompanyUrl) {
		var ret string
		return ret
	}
	return *o.CompanyUrl
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetCompanyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyUrl) {
		return nil, false
	}
	return o.CompanyUrl, true
}

// HasCompanyUrl returns a boolean if a field has been set.
func (o *UserCreateInput) HasCompanyUrl() bool {
	if o != nil && !IsNil(o.CompanyUrl) {
		return true
	}

	return false
}

// SetCompanyUrl gets a reference to the given string and assigns it to the CompanyUrl field.
func (o *UserCreateInput) SetCompanyUrl(v string) {
	o.CompanyUrl = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *UserCreateInput) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *UserCreateInput) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *UserCreateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetEmails returns the Emails field value
func (o *UserCreateInput) GetEmails() []EmailInput {
	if o == nil {
		var ret []EmailInput
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetEmailsOk() ([]EmailInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *UserCreateInput) SetEmails(v []EmailInput) {
	o.Emails = v
}

// GetFirstName returns the FirstName field value
func (o *UserCreateInput) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserCreateInput) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserCreateInput) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserCreateInput) SetLastName(v string) {
	o.LastName = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *UserCreateInput) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *UserCreateInput) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *UserCreateInput) SetLevel(v string) {
	o.Level = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserCreateInput) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserCreateInput) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserCreateInput) SetPassword(v string) {
	o.Password = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserCreateInput) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserCreateInput) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserCreateInput) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetSocialAccounts returns the SocialAccounts field value if set, zero value otherwise.
func (o *UserCreateInput) GetSocialAccounts() map[string]interface{} {
	if o == nil || IsNil(o.SocialAccounts) {
		var ret map[string]interface{}
		return ret
	}
	return o.SocialAccounts
}

// GetSocialAccountsOk returns a tuple with the SocialAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetSocialAccountsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SocialAccounts) {
		return map[string]interface{}{}, false
	}
	return o.SocialAccounts, true
}

// HasSocialAccounts returns a boolean if a field has been set.
func (o *UserCreateInput) HasSocialAccounts() bool {
	if o != nil && !IsNil(o.SocialAccounts) {
		return true
	}

	return false
}

// SetSocialAccounts gets a reference to the given map[string]interface{} and assigns it to the SocialAccounts field.
func (o *UserCreateInput) SetSocialAccounts(v map[string]interface{}) {
	o.SocialAccounts = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UserCreateInput) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UserCreateInput) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UserCreateInput) SetTimezone(v string) {
	o.Timezone = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserCreateInput) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserCreateInput) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserCreateInput) SetTitle(v string) {
	o.Title = &v
}

// GetTwoFactorAuth returns the TwoFactorAuth field value if set, zero value otherwise.
func (o *UserCreateInput) GetTwoFactorAuth() string {
	if o == nil || IsNil(o.TwoFactorAuth) {
		var ret string
		return ret
	}
	return *o.TwoFactorAuth
}

// GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetTwoFactorAuthOk() (*string, bool) {
	if o == nil || IsNil(o.TwoFactorAuth) {
		return nil, false
	}
	return o.TwoFactorAuth, true
}

// HasTwoFactorAuth returns a boolean if a field has been set.
func (o *UserCreateInput) HasTwoFactorAuth() bool {
	if o != nil && !IsNil(o.TwoFactorAuth) {
		return true
	}

	return false
}

// SetTwoFactorAuth gets a reference to the given string and assigns it to the TwoFactorAuth field.
func (o *UserCreateInput) SetTwoFactorAuth(v string) {
	o.TwoFactorAuth = &v
}

// GetVerifiedAt returns the VerifiedAt field value if set, zero value otherwise.
func (o *UserCreateInput) GetVerifiedAt() time.Time {
	if o == nil || IsNil(o.VerifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.VerifiedAt
}

// GetVerifiedAtOk returns a tuple with the VerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetVerifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VerifiedAt) {
		return nil, false
	}
	return o.VerifiedAt, true
}

// HasVerifiedAt returns a boolean if a field has been set.
func (o *UserCreateInput) HasVerifiedAt() bool {
	if o != nil && !IsNil(o.VerifiedAt) {
		return true
	}

	return false
}

// SetVerifiedAt gets a reference to the given time.Time and assigns it to the VerifiedAt field.
func (o *UserCreateInput) SetVerifiedAt(v time.Time) {
	o.VerifiedAt = &v
}

// GetInvitationId returns the InvitationId field value if set, zero value otherwise.
func (o *UserCreateInput) GetInvitationId() string {
	if o == nil || IsNil(o.InvitationId) {
		var ret string
		return ret
	}
	return *o.InvitationId
}

// GetInvitationIdOk returns a tuple with the InvitationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetInvitationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvitationId) {
		return nil, false
	}
	return o.InvitationId, true
}

// HasInvitationId returns a boolean if a field has been set.
func (o *UserCreateInput) HasInvitationId() bool {
	if o != nil && !IsNil(o.InvitationId) {
		return true
	}

	return false
}

// SetInvitationId gets a reference to the given string and assigns it to the InvitationId field.
func (o *UserCreateInput) SetInvitationId(v string) {
	o.InvitationId = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *UserCreateInput) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateInput) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *UserCreateInput) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *UserCreateInput) SetNonce(v string) {
	o.Nonce = &v
}

func (o UserCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.CompanyUrl) {
		toSerialize["company_url"] = o.CompanyUrl
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	toSerialize["emails"] = o.Emails
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.SocialAccounts) {
		toSerialize["social_accounts"] = o.SocialAccounts
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.TwoFactorAuth) {
		toSerialize["two_factor_auth"] = o.TwoFactorAuth
	}
	if !IsNil(o.VerifiedAt) {
		toSerialize["verified_at"] = o.VerifiedAt
	}
	if !IsNil(o.InvitationId) {
		toSerialize["invitation_id"] = o.InvitationId
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCreateInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emails",
		"first_name",
		"last_name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserCreateInput := _UserCreateInput{}

	err = json.Unmarshal(data, &varUserCreateInput)

	if err != nil {
		return err
	}

	*o = UserCreateInput(varUserCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avatar")
		delete(additionalProperties, "company_name")
		delete(additionalProperties, "company_url")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "emails")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "level")
		delete(additionalProperties, "password")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "social_accounts")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "title")
		delete(additionalProperties, "two_factor_auth")
		delete(additionalProperties, "verified_at")
		delete(additionalProperties, "invitation_id")
		delete(additionalProperties, "nonce")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCreateInput struct {
	value *UserCreateInput
	isSet bool
}

func (v NullableUserCreateInput) Get() *UserCreateInput {
	return v.value
}

func (v *NullableUserCreateInput) Set(val *UserCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateInput(val *UserCreateInput) *NullableUserCreateInput {
	return &NullableUserCreateInput{value: val, isSet: true}
}

func (v NullableUserCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
