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
	"time"
)

// checks if the VrfVirtualCircuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfVirtualCircuit{}

// VrfVirtualCircuit struct for VrfVirtualCircuit
type VrfVirtualCircuit struct {
	// An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
	CustomerIp  *string `json:"customer_ip,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	// The MD5 password for the BGP peering in plaintext (not a checksum).
	Md5 *string `json:"md5,omitempty"`
	// An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
	MetalIp *string `json:"metal_ip,omitempty"`
	Name    *string `json:"name,omitempty"`
	Port    *Href   `json:"port,omitempty"`
	NniVlan *int32  `json:"nni_vlan,omitempty"`
	// The peer ASN that will be used with the VRF on the Virtual Circuit.
	PeerAsn *int32 `json:"peer_asn,omitempty"`
	Project *Href  `json:"project,omitempty"`
	// integer representing bps speed
	Speed *int32 `json:"speed,omitempty"`
	// The status changes of a VRF virtual circuit are generally the same as Virtual Circuits that aren't in a VRF. However, for VRF Virtual Circuits on Fabric VCs, the status will change to 'waiting_on_peering_details' once the Fabric service token associated with the virtual circuit has been redeemed on Fabric, and Metal has found the associated Fabric connection. At this point, users can update the subnet, MD5 password, customer IP and/or metal IP accordingly. For VRF Virtual Circuits on Dedicated Ports, we require all peering details to be set on creation of a VRF Virtual Circuit. The status will change to `changing_peering_details` whenever an active VRF Virtual Circuit has any of its peering details updated.
	Status *string `json:"status,omitempty"`
	// The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP.
	Subnet               *string    `json:"subnet,omitempty"`
	Tags                 []string   `json:"tags,omitempty"`
	Vrf                  *Vrf       `json:"vrf,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfVirtualCircuit VrfVirtualCircuit

// NewVrfVirtualCircuit instantiates a new VrfVirtualCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfVirtualCircuit() *VrfVirtualCircuit {
	this := VrfVirtualCircuit{}
	return &this
}

// NewVrfVirtualCircuitWithDefaults instantiates a new VrfVirtualCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfVirtualCircuitWithDefaults() *VrfVirtualCircuit {
	this := VrfVirtualCircuit{}
	return &this
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetCustomerIp() string {
	if o == nil || IsNil(o.CustomerIp) {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetCustomerIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIp) {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasCustomerIp() bool {
	if o != nil && !IsNil(o.CustomerIp) {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *VrfVirtualCircuit) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VrfVirtualCircuit) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VrfVirtualCircuit) SetId(v string) {
	o.Id = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *VrfVirtualCircuit) SetMd5(v string) {
	o.Md5 = &v
}

// GetMetalIp returns the MetalIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetMetalIp() string {
	if o == nil || IsNil(o.MetalIp) {
		var ret string
		return ret
	}
	return *o.MetalIp
}

// GetMetalIpOk returns a tuple with the MetalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetMetalIpOk() (*string, bool) {
	if o == nil || IsNil(o.MetalIp) {
		return nil, false
	}
	return o.MetalIp, true
}

// HasMetalIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasMetalIp() bool {
	if o != nil && !IsNil(o.MetalIp) {
		return true
	}

	return false
}

// SetMetalIp gets a reference to the given string and assigns it to the MetalIp field.
func (o *VrfVirtualCircuit) SetMetalIp(v string) {
	o.MetalIp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VrfVirtualCircuit) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetPort() Href {
	if o == nil || IsNil(o.Port) {
		var ret Href
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetPortOk() (*Href, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given Href and assigns it to the Port field.
func (o *VrfVirtualCircuit) SetPort(v Href) {
	o.Port = &v
}

// GetNniVlan returns the NniVlan field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetNniVlan() int32 {
	if o == nil || IsNil(o.NniVlan) {
		var ret int32
		return ret
	}
	return *o.NniVlan
}

// GetNniVlanOk returns a tuple with the NniVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetNniVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.NniVlan) {
		return nil, false
	}
	return o.NniVlan, true
}

// HasNniVlan returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasNniVlan() bool {
	if o != nil && !IsNil(o.NniVlan) {
		return true
	}

	return false
}

// SetNniVlan gets a reference to the given int32 and assigns it to the NniVlan field.
func (o *VrfVirtualCircuit) SetNniVlan(v int32) {
	o.NniVlan = &v
}

// GetPeerAsn returns the PeerAsn field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetPeerAsn() int32 {
	if o == nil || IsNil(o.PeerAsn) {
		var ret int32
		return ret
	}
	return *o.PeerAsn
}

// GetPeerAsnOk returns a tuple with the PeerAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetPeerAsnOk() (*int32, bool) {
	if o == nil || IsNil(o.PeerAsn) {
		return nil, false
	}
	return o.PeerAsn, true
}

// HasPeerAsn returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasPeerAsn() bool {
	if o != nil && !IsNil(o.PeerAsn) {
		return true
	}

	return false
}

// SetPeerAsn gets a reference to the given int32 and assigns it to the PeerAsn field.
func (o *VrfVirtualCircuit) SetPeerAsn(v int32) {
	o.PeerAsn = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetProject() Href {
	if o == nil || IsNil(o.Project) {
		var ret Href
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetProjectOk() (*Href, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Href and assigns it to the Project field.
func (o *VrfVirtualCircuit) SetProject(v Href) {
	o.Project = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *VrfVirtualCircuit) SetSpeed(v int32) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VrfVirtualCircuit) SetStatus(v string) {
	o.Status = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *VrfVirtualCircuit) SetSubnet(v string) {
	o.Subnet = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfVirtualCircuit) SetTags(v []string) {
	o.Tags = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetVrf() Vrf {
	if o == nil || IsNil(o.Vrf) {
		var ret Vrf
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetVrfOk() (*Vrf, bool) {
	if o == nil || IsNil(o.Vrf) {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasVrf() bool {
	if o != nil && !IsNil(o.Vrf) {
		return true
	}

	return false
}

// SetVrf gets a reference to the given Vrf and assigns it to the Vrf field.
func (o *VrfVirtualCircuit) SetVrf(v Vrf) {
	o.Vrf = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VrfVirtualCircuit) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VrfVirtualCircuit) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o VrfVirtualCircuit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfVirtualCircuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerIp) {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.MetalIp) {
		toSerialize["metal_ip"] = o.MetalIp
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.NniVlan) {
		toSerialize["nni_vlan"] = o.NniVlan
	}
	if !IsNil(o.PeerAsn) {
		toSerialize["peer_asn"] = o.PeerAsn
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Vrf) {
		toSerialize["vrf"] = o.Vrf
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfVirtualCircuit) UnmarshalJSON(bytes []byte) (err error) {
	varVrfVirtualCircuit := _VrfVirtualCircuit{}

	if err = json.Unmarshal(bytes, &varVrfVirtualCircuit); err == nil {
		*o = VrfVirtualCircuit(varVrfVirtualCircuit)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_ip")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "md5")
		delete(additionalProperties, "metal_ip")
		delete(additionalProperties, "name")
		delete(additionalProperties, "port")
		delete(additionalProperties, "nni_vlan")
		delete(additionalProperties, "peer_asn")
		delete(additionalProperties, "project")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subnet")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfVirtualCircuit struct {
	value *VrfVirtualCircuit
	isSet bool
}

func (v NullableVrfVirtualCircuit) Get() *VrfVirtualCircuit {
	return v.value
}

func (v *NullableVrfVirtualCircuit) Set(val *VrfVirtualCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfVirtualCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfVirtualCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfVirtualCircuit(val *VrfVirtualCircuit) *NullableVrfVirtualCircuit {
	return &NullableVrfVirtualCircuit{value: val, isSet: true}
}

func (v NullableVrfVirtualCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfVirtualCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
