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

// checks if the ServiceProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileRequest{}

// ServiceProfileRequest Service Profile is a software definition for a named provider service and it's network connectivity requirements. This includes the basic marketing information and one or more sets of access points (a set per each access point type) fulfilling the provider service.
type ServiceProfileRequest struct {
	Project *Project `json:"project,omitempty"`
	// Service Profile URI response attribute
	Href *string                `json:"href,omitempty"`
	Type ServiceProfileTypeEnum `json:"type"`
	// Customer-assigned service profile name
	Name string `json:"name"`
	// Equinix-assigned service profile identifier
	Uuid *string `json:"uuid,omitempty"`
	// User-provided service description should be of maximum length 375
	Description string `json:"description"`
	// Recipients of notifications on service profile change
	Notifications          []SimplifiedNotification        `json:"notifications,omitempty"`
	Tags                   []string                        `json:"tags,omitempty"`
	Visibility             *ServiceProfileVisibilityEnum   `json:"visibility,omitempty"`
	AllowedEmails          []string                        `json:"allowedEmails,omitempty"`
	AccessPointTypeConfigs []ServiceProfileAccessPointType `json:"accessPointTypeConfigs,omitempty"`
	CustomFields           []CustomField                   `json:"customFields,omitempty"`
	MarketingInfo          *MarketingInfo                  `json:"marketingInfo,omitempty"`
	Ports                  []ServiceProfileAccessPointCOLO `json:"ports,omitempty"`
	VirtualDevices         []ServiceProfileAccessPointVD   `json:"virtualDevices,omitempty"`
	// Derived response attribute.
	Metros []ServiceMetro `json:"metros,omitempty"`
	// response attribute indicates whether the profile belongs to the same organization as the api-invoker.
	SelfProfile          *bool   `json:"selfProfile,omitempty"`
	ProjectId            *string `json:"projectId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProfileRequest ServiceProfileRequest

// NewServiceProfileRequest instantiates a new ServiceProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileRequest(type_ ServiceProfileTypeEnum, name string, description string) *ServiceProfileRequest {
	this := ServiceProfileRequest{}
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewServiceProfileRequestWithDefaults instantiates a new ServiceProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileRequestWithDefaults() *ServiceProfileRequest {
	this := ServiceProfileRequest{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ServiceProfileRequest) SetProject(v Project) {
	o.Project = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ServiceProfileRequest) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value
func (o *ServiceProfileRequest) GetType() ServiceProfileTypeEnum {
	if o == nil {
		var ret ServiceProfileTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetTypeOk() (*ServiceProfileTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceProfileRequest) SetType(v ServiceProfileTypeEnum) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ServiceProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceProfileRequest) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ServiceProfileRequest) SetUuid(v string) {
	o.Uuid = &v
}

// GetDescription returns the Description field value
func (o *ServiceProfileRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceProfileRequest) SetDescription(v string) {
	o.Description = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetNotifications() []SimplifiedNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []SimplifiedNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetNotificationsOk() ([]SimplifiedNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []SimplifiedNotification and assigns it to the Notifications field.
func (o *ServiceProfileRequest) SetNotifications(v []SimplifiedNotification) {
	o.Notifications = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceProfileRequest) SetTags(v []string) {
	o.Tags = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetVisibility() ServiceProfileVisibilityEnum {
	if o == nil || IsNil(o.Visibility) {
		var ret ServiceProfileVisibilityEnum
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetVisibilityOk() (*ServiceProfileVisibilityEnum, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given ServiceProfileVisibilityEnum and assigns it to the Visibility field.
func (o *ServiceProfileRequest) SetVisibility(v ServiceProfileVisibilityEnum) {
	o.Visibility = &v
}

// GetAllowedEmails returns the AllowedEmails field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetAllowedEmails() []string {
	if o == nil || IsNil(o.AllowedEmails) {
		var ret []string
		return ret
	}
	return o.AllowedEmails
}

// GetAllowedEmailsOk returns a tuple with the AllowedEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetAllowedEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedEmails) {
		return nil, false
	}
	return o.AllowedEmails, true
}

// HasAllowedEmails returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasAllowedEmails() bool {
	if o != nil && !IsNil(o.AllowedEmails) {
		return true
	}

	return false
}

// SetAllowedEmails gets a reference to the given []string and assigns it to the AllowedEmails field.
func (o *ServiceProfileRequest) SetAllowedEmails(v []string) {
	o.AllowedEmails = v
}

// GetAccessPointTypeConfigs returns the AccessPointTypeConfigs field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetAccessPointTypeConfigs() []ServiceProfileAccessPointType {
	if o == nil || IsNil(o.AccessPointTypeConfigs) {
		var ret []ServiceProfileAccessPointType
		return ret
	}
	return o.AccessPointTypeConfigs
}

// GetAccessPointTypeConfigsOk returns a tuple with the AccessPointTypeConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetAccessPointTypeConfigsOk() ([]ServiceProfileAccessPointType, bool) {
	if o == nil || IsNil(o.AccessPointTypeConfigs) {
		return nil, false
	}
	return o.AccessPointTypeConfigs, true
}

// HasAccessPointTypeConfigs returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasAccessPointTypeConfigs() bool {
	if o != nil && !IsNil(o.AccessPointTypeConfigs) {
		return true
	}

	return false
}

// SetAccessPointTypeConfigs gets a reference to the given []ServiceProfileAccessPointType and assigns it to the AccessPointTypeConfigs field.
func (o *ServiceProfileRequest) SetAccessPointTypeConfigs(v []ServiceProfileAccessPointType) {
	o.AccessPointTypeConfigs = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetCustomFields() []CustomField {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomField
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetCustomFieldsOk() ([]CustomField, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomField and assigns it to the CustomFields field.
func (o *ServiceProfileRequest) SetCustomFields(v []CustomField) {
	o.CustomFields = v
}

// GetMarketingInfo returns the MarketingInfo field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetMarketingInfo() MarketingInfo {
	if o == nil || IsNil(o.MarketingInfo) {
		var ret MarketingInfo
		return ret
	}
	return *o.MarketingInfo
}

// GetMarketingInfoOk returns a tuple with the MarketingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetMarketingInfoOk() (*MarketingInfo, bool) {
	if o == nil || IsNil(o.MarketingInfo) {
		return nil, false
	}
	return o.MarketingInfo, true
}

// HasMarketingInfo returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasMarketingInfo() bool {
	if o != nil && !IsNil(o.MarketingInfo) {
		return true
	}

	return false
}

// SetMarketingInfo gets a reference to the given MarketingInfo and assigns it to the MarketingInfo field.
func (o *ServiceProfileRequest) SetMarketingInfo(v MarketingInfo) {
	o.MarketingInfo = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetPorts() []ServiceProfileAccessPointCOLO {
	if o == nil || IsNil(o.Ports) {
		var ret []ServiceProfileAccessPointCOLO
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetPortsOk() ([]ServiceProfileAccessPointCOLO, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServiceProfileAccessPointCOLO and assigns it to the Ports field.
func (o *ServiceProfileRequest) SetPorts(v []ServiceProfileAccessPointCOLO) {
	o.Ports = v
}

// GetVirtualDevices returns the VirtualDevices field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetVirtualDevices() []ServiceProfileAccessPointVD {
	if o == nil || IsNil(o.VirtualDevices) {
		var ret []ServiceProfileAccessPointVD
		return ret
	}
	return o.VirtualDevices
}

// GetVirtualDevicesOk returns a tuple with the VirtualDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetVirtualDevicesOk() ([]ServiceProfileAccessPointVD, bool) {
	if o == nil || IsNil(o.VirtualDevices) {
		return nil, false
	}
	return o.VirtualDevices, true
}

// HasVirtualDevices returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasVirtualDevices() bool {
	if o != nil && !IsNil(o.VirtualDevices) {
		return true
	}

	return false
}

// SetVirtualDevices gets a reference to the given []ServiceProfileAccessPointVD and assigns it to the VirtualDevices field.
func (o *ServiceProfileRequest) SetVirtualDevices(v []ServiceProfileAccessPointVD) {
	o.VirtualDevices = v
}

// GetMetros returns the Metros field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetMetros() []ServiceMetro {
	if o == nil || IsNil(o.Metros) {
		var ret []ServiceMetro
		return ret
	}
	return o.Metros
}

// GetMetrosOk returns a tuple with the Metros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetMetrosOk() ([]ServiceMetro, bool) {
	if o == nil || IsNil(o.Metros) {
		return nil, false
	}
	return o.Metros, true
}

// HasMetros returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasMetros() bool {
	if o != nil && !IsNil(o.Metros) {
		return true
	}

	return false
}

// SetMetros gets a reference to the given []ServiceMetro and assigns it to the Metros field.
func (o *ServiceProfileRequest) SetMetros(v []ServiceMetro) {
	o.Metros = v
}

// GetSelfProfile returns the SelfProfile field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetSelfProfile() bool {
	if o == nil || IsNil(o.SelfProfile) {
		var ret bool
		return ret
	}
	return *o.SelfProfile
}

// GetSelfProfileOk returns a tuple with the SelfProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetSelfProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfProfile) {
		return nil, false
	}
	return o.SelfProfile, true
}

// HasSelfProfile returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasSelfProfile() bool {
	if o != nil && !IsNil(o.SelfProfile) {
		return true
	}

	return false
}

// SetSelfProfile gets a reference to the given bool and assigns it to the SelfProfile field.
func (o *ServiceProfileRequest) SetSelfProfile(v bool) {
	o.SelfProfile = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ServiceProfileRequest) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ServiceProfileRequest) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ServiceProfileRequest) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o ServiceProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.AllowedEmails) {
		toSerialize["allowedEmails"] = o.AllowedEmails
	}
	if !IsNil(o.AccessPointTypeConfigs) {
		toSerialize["accessPointTypeConfigs"] = o.AccessPointTypeConfigs
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.MarketingInfo) {
		toSerialize["marketingInfo"] = o.MarketingInfo
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.VirtualDevices) {
		toSerialize["virtualDevices"] = o.VirtualDevices
	}
	if !IsNil(o.Metros) {
		toSerialize["metros"] = o.Metros
	}
	if !IsNil(o.SelfProfile) {
		toSerialize["selfProfile"] = o.SelfProfile
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProfileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"description",
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

	varServiceProfileRequest := _ServiceProfileRequest{}

	err = json.Unmarshal(data, &varServiceProfileRequest)

	if err != nil {
		return err
	}

	*o = ServiceProfileRequest(varServiceProfileRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project")
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "visibility")
		delete(additionalProperties, "allowedEmails")
		delete(additionalProperties, "accessPointTypeConfigs")
		delete(additionalProperties, "customFields")
		delete(additionalProperties, "marketingInfo")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "virtualDevices")
		delete(additionalProperties, "metros")
		delete(additionalProperties, "selfProfile")
		delete(additionalProperties, "projectId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProfileRequest struct {
	value *ServiceProfileRequest
	isSet bool
}

func (v NullableServiceProfileRequest) Get() *ServiceProfileRequest {
	return v.value
}

func (v *NullableServiceProfileRequest) Set(val *ServiceProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileRequest(val *ServiceProfileRequest) *NullableServiceProfileRequest {
	return &NullableServiceProfileRequest{value: val, isSet: true}
}

func (v NullableServiceProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
