/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// checks if the DedicatedPortCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DedicatedPortCreateInput{}

// DedicatedPortCreateInput struct for DedicatedPortCreateInput
type DedicatedPortCreateInput struct {
	// The billing account name of the Equinix Fabric account.
	BillingAccountName *string `json:"billing_account_name,omitempty"`
	// The preferred email used for communication and notifications about the Equinix Fabric interconnection. Required when using a Project API key. Optional and defaults to the primary user email address when using a User API key.
	ContactEmail *string `json:"contact_email,omitempty"`
	Description  *string `json:"description,omitempty"`
	// A Metro ID or code. For interconnections with Dedicated Ports, this will be the location of the issued Dedicated Ports.
	Metro   string                        `json:"metro"`
	Mode    *DedicatedPortCreateInputMode `json:"mode,omitempty"`
	Name    string                        `json:"name"`
	Project *string                       `json:"project,omitempty"`
	// Either 'primary' or 'redundant'.
	Redundancy string `json:"redundancy"`
	// A interconnection speed, in bps, mbps, or gbps. For Dedicated Ports, this can be 10Gbps or 100Gbps.
	Speed *int32                       `json:"speed,omitempty"`
	Tags  []string                     `json:"tags,omitempty"`
	Type  DedicatedPortCreateInputType `json:"type"`
	// The intended use case of the dedicated port.
	UseCase              *string `json:"use_case,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DedicatedPortCreateInput DedicatedPortCreateInput

// NewDedicatedPortCreateInput instantiates a new DedicatedPortCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedPortCreateInput(metro string, name string, redundancy string, type_ DedicatedPortCreateInputType) *DedicatedPortCreateInput {
	this := DedicatedPortCreateInput{}
	this.Metro = metro
	this.Name = name
	this.Redundancy = redundancy
	this.Type = type_
	return &this
}

// NewDedicatedPortCreateInputWithDefaults instantiates a new DedicatedPortCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDedicatedPortCreateInputWithDefaults() *DedicatedPortCreateInput {
	this := DedicatedPortCreateInput{}
	return &this
}

// GetBillingAccountName returns the BillingAccountName field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetBillingAccountName() string {
	if o == nil || IsNil(o.BillingAccountName) {
		var ret string
		return ret
	}
	return *o.BillingAccountName
}

// GetBillingAccountNameOk returns a tuple with the BillingAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetBillingAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.BillingAccountName) {
		return nil, false
	}
	return o.BillingAccountName, true
}

// HasBillingAccountName returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasBillingAccountName() bool {
	if o != nil && !IsNil(o.BillingAccountName) {
		return true
	}

	return false
}

// SetBillingAccountName gets a reference to the given string and assigns it to the BillingAccountName field.
func (o *DedicatedPortCreateInput) SetBillingAccountName(v string) {
	o.BillingAccountName = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *DedicatedPortCreateInput) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DedicatedPortCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetMetro returns the Metro field value
func (o *DedicatedPortCreateInput) GetMetro() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metro
}

// GetMetroOk returns a tuple with the Metro field value
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetMetroOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metro, true
}

// SetMetro sets field value
func (o *DedicatedPortCreateInput) SetMetro(v string) {
	o.Metro = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetMode() DedicatedPortCreateInputMode {
	if o == nil || IsNil(o.Mode) {
		var ret DedicatedPortCreateInputMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetModeOk() (*DedicatedPortCreateInputMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given DedicatedPortCreateInputMode and assigns it to the Mode field.
func (o *DedicatedPortCreateInput) SetMode(v DedicatedPortCreateInputMode) {
	o.Mode = &v
}

// GetName returns the Name field value
func (o *DedicatedPortCreateInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DedicatedPortCreateInput) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *DedicatedPortCreateInput) SetProject(v string) {
	o.Project = &v
}

// GetRedundancy returns the Redundancy field value
func (o *DedicatedPortCreateInput) GetRedundancy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetRedundancyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redundancy, true
}

// SetRedundancy sets field value
func (o *DedicatedPortCreateInput) SetRedundancy(v string) {
	o.Redundancy = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *DedicatedPortCreateInput) SetSpeed(v int32) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DedicatedPortCreateInput) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *DedicatedPortCreateInput) GetType() DedicatedPortCreateInputType {
	if o == nil {
		var ret DedicatedPortCreateInputType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetTypeOk() (*DedicatedPortCreateInputType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DedicatedPortCreateInput) SetType(v DedicatedPortCreateInputType) {
	o.Type = v
}

// GetUseCase returns the UseCase field value if set, zero value otherwise.
func (o *DedicatedPortCreateInput) GetUseCase() string {
	if o == nil || IsNil(o.UseCase) {
		var ret string
		return ret
	}
	return *o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedPortCreateInput) GetUseCaseOk() (*string, bool) {
	if o == nil || IsNil(o.UseCase) {
		return nil, false
	}
	return o.UseCase, true
}

// HasUseCase returns a boolean if a field has been set.
func (o *DedicatedPortCreateInput) HasUseCase() bool {
	if o != nil && !IsNil(o.UseCase) {
		return true
	}

	return false
}

// SetUseCase gets a reference to the given string and assigns it to the UseCase field.
func (o *DedicatedPortCreateInput) SetUseCase(v string) {
	o.UseCase = &v
}

func (o DedicatedPortCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DedicatedPortCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAccountName) {
		toSerialize["billing_account_name"] = o.BillingAccountName
	}
	if !IsNil(o.ContactEmail) {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["metro"] = o.Metro
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	toSerialize["redundancy"] = o.Redundancy
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.UseCase) {
		toSerialize["use_case"] = o.UseCase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DedicatedPortCreateInput) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metro",
		"name",
		"redundancy",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDedicatedPortCreateInput := _DedicatedPortCreateInput{}

	err = json.Unmarshal(bytes, &varDedicatedPortCreateInput)

	if err != nil {
		return err
	}

	*o = DedicatedPortCreateInput(varDedicatedPortCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "billing_account_name")
		delete(additionalProperties, "contact_email")
		delete(additionalProperties, "description")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "name")
		delete(additionalProperties, "project")
		delete(additionalProperties, "redundancy")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "type")
		delete(additionalProperties, "use_case")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDedicatedPortCreateInput struct {
	value *DedicatedPortCreateInput
	isSet bool
}

func (v NullableDedicatedPortCreateInput) Get() *DedicatedPortCreateInput {
	return v.value
}

func (v *NullableDedicatedPortCreateInput) Set(val *DedicatedPortCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDedicatedPortCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDedicatedPortCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDedicatedPortCreateInput(val *DedicatedPortCreateInput) *NullableDedicatedPortCreateInput {
	return &NullableDedicatedPortCreateInput{value: val, isSet: true}
}

func (v NullableDedicatedPortCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDedicatedPortCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
