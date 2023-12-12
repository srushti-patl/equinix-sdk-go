/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// checks if the VrfIpReservationCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfIpReservationCreateInput{}

// VrfIpReservationCreateInput struct for VrfIpReservationCreateInput
type VrfIpReservationCreateInput struct {
	// The size of the VRF IP Reservation's subnet
	Cidr       int32                  `json:"cidr"`
	Customdata map[string]interface{} `json:"customdata,omitempty"`
	Details    *string                `json:"details,omitempty"`
	// The starting address for this VRF IP Reservation's subnet
	Network string   `json:"network"`
	Tags    []string `json:"tags,omitempty"`
	// Must be set to 'vrf'
	Type string `json:"type"`
	// The ID of the VRF in which this VRF IP Reservation is created. The VRF must have an existing IP Range that contains the requested subnet. This field may be aliased as just 'vrf'.
	VrfId                string `json:"vrf_id"`
	AdditionalProperties map[string]interface{}
}

type _VrfIpReservationCreateInput VrfIpReservationCreateInput

// NewVrfIpReservationCreateInput instantiates a new VrfIpReservationCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfIpReservationCreateInput(cidr int32, network string, type_ string, vrfId string) *VrfIpReservationCreateInput {
	this := VrfIpReservationCreateInput{}
	this.Cidr = cidr
	this.Network = network
	this.Type = type_
	this.VrfId = vrfId
	return &this
}

// NewVrfIpReservationCreateInputWithDefaults instantiates a new VrfIpReservationCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfIpReservationCreateInputWithDefaults() *VrfIpReservationCreateInput {
	this := VrfIpReservationCreateInput{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *VrfIpReservationCreateInput) GetCidr() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetCidrOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *VrfIpReservationCreateInput) SetCidr(v int32) {
	o.Cidr = v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *VrfIpReservationCreateInput) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *VrfIpReservationCreateInput) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *VrfIpReservationCreateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VrfIpReservationCreateInput) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VrfIpReservationCreateInput) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *VrfIpReservationCreateInput) SetDetails(v string) {
	o.Details = &v
}

// GetNetwork returns the Network field value
func (o *VrfIpReservationCreateInput) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *VrfIpReservationCreateInput) SetNetwork(v string) {
	o.Network = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfIpReservationCreateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfIpReservationCreateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfIpReservationCreateInput) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *VrfIpReservationCreateInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VrfIpReservationCreateInput) SetType(v string) {
	o.Type = v
}

// GetVrfId returns the VrfId field value
func (o *VrfIpReservationCreateInput) GetVrfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VrfId
}

// GetVrfIdOk returns a tuple with the VrfId field value
// and a boolean to check if the value has been set.
func (o *VrfIpReservationCreateInput) GetVrfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VrfId, true
}

// SetVrfId sets field value
func (o *VrfIpReservationCreateInput) SetVrfId(v string) {
	o.VrfId = v
}

func (o VrfIpReservationCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfIpReservationCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidr"] = o.Cidr
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	toSerialize["network"] = o.Network
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	toSerialize["vrf_id"] = o.VrfId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfIpReservationCreateInput) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cidr",
		"network",
		"type",
		"vrf_id",
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

	varVrfIpReservationCreateInput := _VrfIpReservationCreateInput{}

	err = json.Unmarshal(bytes, &varVrfIpReservationCreateInput)

	if err != nil {
		return err
	}

	*o = VrfIpReservationCreateInput(varVrfIpReservationCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "details")
		delete(additionalProperties, "network")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vrf_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfIpReservationCreateInput struct {
	value *VrfIpReservationCreateInput
	isSet bool
}

func (v NullableVrfIpReservationCreateInput) Get() *VrfIpReservationCreateInput {
	return v.value
}

func (v *NullableVrfIpReservationCreateInput) Set(val *VrfIpReservationCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfIpReservationCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfIpReservationCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfIpReservationCreateInput(val *VrfIpReservationCreateInput) *NullableVrfIpReservationCreateInput {
	return &NullableVrfIpReservationCreateInput{value: val, isSet: true}
}

func (v NullableVrfIpReservationCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfIpReservationCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
