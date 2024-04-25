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

// checks if the RouteFiltersDataProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFiltersDataProject{}

// RouteFiltersDataProject struct for RouteFiltersDataProject
type RouteFiltersDataProject struct {
	// Subscriber-assigned project ID
	ProjectId string `json:"projectId"`
	// Project URI
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFiltersDataProject RouteFiltersDataProject

// NewRouteFiltersDataProject instantiates a new RouteFiltersDataProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFiltersDataProject(projectId string) *RouteFiltersDataProject {
	this := RouteFiltersDataProject{}
	this.ProjectId = projectId
	return &this
}

// NewRouteFiltersDataProjectWithDefaults instantiates a new RouteFiltersDataProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFiltersDataProjectWithDefaults() *RouteFiltersDataProject {
	this := RouteFiltersDataProject{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *RouteFiltersDataProject) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RouteFiltersDataProject) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RouteFiltersDataProject) SetProjectId(v string) {
	o.ProjectId = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteFiltersDataProject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFiltersDataProject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteFiltersDataProject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteFiltersDataProject) SetHref(v string) {
	o.Href = &v
}

func (o RouteFiltersDataProject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFiltersDataProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFiltersDataProject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectId",
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

	varRouteFiltersDataProject := _RouteFiltersDataProject{}

	err = json.Unmarshal(data, &varRouteFiltersDataProject)

	if err != nil {
		return err
	}

	*o = RouteFiltersDataProject(varRouteFiltersDataProject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFiltersDataProject struct {
	value *RouteFiltersDataProject
	isSet bool
}

func (v NullableRouteFiltersDataProject) Get() *RouteFiltersDataProject {
	return v.value
}

func (v *NullableRouteFiltersDataProject) Set(val *RouteFiltersDataProject) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFiltersDataProject) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFiltersDataProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFiltersDataProject(val *RouteFiltersDataProject) *NullableRouteFiltersDataProject {
	return &NullableRouteFiltersDataProject{value: val, isSet: true}
}

func (v NullableRouteFiltersDataProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFiltersDataProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
