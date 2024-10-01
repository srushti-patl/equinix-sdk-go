/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PhysicalPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhysicalPort{}

// PhysicalPort Physical Port specification
type PhysicalPort struct {
	Type *PhysicalPortType `json:"type,omitempty"`
	// Equinix assigned response attribute for Physical Port Id
	// Deprecated
	Id *int32 `json:"id,omitempty"`
	// Equinix assigned response attribute for an absolute URL that is the subject of the link's context.
	Href    *string            `json:"href,omitempty"`
	State   *PortState         `json:"state,omitempty"`
	Account *SimplifiedAccount `json:"account,omitempty"`
	// Physical Port Speed in Mbps
	InterfaceSpeed *int32 `json:"interfaceSpeed,omitempty"`
	// Physical Port Interface Type
	InterfaceType *string `json:"interfaceType,omitempty"`
	// Equinix assigned response attribute for physical port identifier
	Uuid             *string               `json:"uuid,omitempty"`
	Tether           *PortTether           `json:"tether,omitempty"`
	DemarcationPoint *PortDemarcationPoint `json:"demarcationPoint,omitempty"`
	Settings         *PhysicalPortSettings `json:"settings,omitempty"`
	Interface        *PortInterface        `json:"interface,omitempty"`
	// Notification preferences
	Notifications []PortNotification `json:"notifications,omitempty"`
	// Physical Port additional information
	AdditionalInfo []PortAdditionalInfo `json:"additionalInfo,omitempty"`
	Order          *PortOrder           `json:"order,omitempty"`
	Operation      *PortOperation       `json:"operation,omitempty"`
	// Port Loas
	Loas                 []PortLoa `json:"loas,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PhysicalPort PhysicalPort

// NewPhysicalPort instantiates a new PhysicalPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalPort() *PhysicalPort {
	this := PhysicalPort{}
	return &this
}

// NewPhysicalPortWithDefaults instantiates a new PhysicalPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalPortWithDefaults() *PhysicalPort {
	this := PhysicalPort{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PhysicalPort) GetType() PhysicalPortType {
	if o == nil || IsNil(o.Type) {
		var ret PhysicalPortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetTypeOk() (*PhysicalPortType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PhysicalPort) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PhysicalPortType and assigns it to the Type field.
func (o *PhysicalPort) SetType(v PhysicalPortType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
// Deprecated
func (o *PhysicalPort) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhysicalPort) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PhysicalPort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
// Deprecated
func (o *PhysicalPort) SetId(v int32) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PhysicalPort) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PhysicalPort) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PhysicalPort) SetHref(v string) {
	o.Href = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PhysicalPort) GetState() PortState {
	if o == nil || IsNil(o.State) {
		var ret PortState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetStateOk() (*PortState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PhysicalPort) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given PortState and assigns it to the State field.
func (o *PhysicalPort) SetState(v PortState) {
	o.State = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *PhysicalPort) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *PhysicalPort) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *PhysicalPort) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetInterfaceSpeed returns the InterfaceSpeed field value if set, zero value otherwise.
func (o *PhysicalPort) GetInterfaceSpeed() int32 {
	if o == nil || IsNil(o.InterfaceSpeed) {
		var ret int32
		return ret
	}
	return *o.InterfaceSpeed
}

// GetInterfaceSpeedOk returns a tuple with the InterfaceSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetInterfaceSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.InterfaceSpeed) {
		return nil, false
	}
	return o.InterfaceSpeed, true
}

// HasInterfaceSpeed returns a boolean if a field has been set.
func (o *PhysicalPort) HasInterfaceSpeed() bool {
	if o != nil && !IsNil(o.InterfaceSpeed) {
		return true
	}

	return false
}

// SetInterfaceSpeed gets a reference to the given int32 and assigns it to the InterfaceSpeed field.
func (o *PhysicalPort) SetInterfaceSpeed(v int32) {
	o.InterfaceSpeed = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *PhysicalPort) GetInterfaceType() string {
	if o == nil || IsNil(o.InterfaceType) {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceType) {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *PhysicalPort) HasInterfaceType() bool {
	if o != nil && !IsNil(o.InterfaceType) {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *PhysicalPort) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *PhysicalPort) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *PhysicalPort) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *PhysicalPort) SetUuid(v string) {
	o.Uuid = &v
}

// GetTether returns the Tether field value if set, zero value otherwise.
func (o *PhysicalPort) GetTether() PortTether {
	if o == nil || IsNil(o.Tether) {
		var ret PortTether
		return ret
	}
	return *o.Tether
}

// GetTetherOk returns a tuple with the Tether field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetTetherOk() (*PortTether, bool) {
	if o == nil || IsNil(o.Tether) {
		return nil, false
	}
	return o.Tether, true
}

// HasTether returns a boolean if a field has been set.
func (o *PhysicalPort) HasTether() bool {
	if o != nil && !IsNil(o.Tether) {
		return true
	}

	return false
}

// SetTether gets a reference to the given PortTether and assigns it to the Tether field.
func (o *PhysicalPort) SetTether(v PortTether) {
	o.Tether = &v
}

// GetDemarcationPoint returns the DemarcationPoint field value if set, zero value otherwise.
func (o *PhysicalPort) GetDemarcationPoint() PortDemarcationPoint {
	if o == nil || IsNil(o.DemarcationPoint) {
		var ret PortDemarcationPoint
		return ret
	}
	return *o.DemarcationPoint
}

// GetDemarcationPointOk returns a tuple with the DemarcationPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetDemarcationPointOk() (*PortDemarcationPoint, bool) {
	if o == nil || IsNil(o.DemarcationPoint) {
		return nil, false
	}
	return o.DemarcationPoint, true
}

// HasDemarcationPoint returns a boolean if a field has been set.
func (o *PhysicalPort) HasDemarcationPoint() bool {
	if o != nil && !IsNil(o.DemarcationPoint) {
		return true
	}

	return false
}

// SetDemarcationPoint gets a reference to the given PortDemarcationPoint and assigns it to the DemarcationPoint field.
func (o *PhysicalPort) SetDemarcationPoint(v PortDemarcationPoint) {
	o.DemarcationPoint = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *PhysicalPort) GetSettings() PhysicalPortSettings {
	if o == nil || IsNil(o.Settings) {
		var ret PhysicalPortSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetSettingsOk() (*PhysicalPortSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PhysicalPort) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given PhysicalPortSettings and assigns it to the Settings field.
func (o *PhysicalPort) SetSettings(v PhysicalPortSettings) {
	o.Settings = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *PhysicalPort) GetInterface() PortInterface {
	if o == nil || IsNil(o.Interface) {
		var ret PortInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetInterfaceOk() (*PortInterface, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *PhysicalPort) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given PortInterface and assigns it to the Interface field.
func (o *PhysicalPort) SetInterface(v PortInterface) {
	o.Interface = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *PhysicalPort) GetNotifications() []PortNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []PortNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetNotificationsOk() ([]PortNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *PhysicalPort) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []PortNotification and assigns it to the Notifications field.
func (o *PhysicalPort) SetNotifications(v []PortNotification) {
	o.Notifications = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *PhysicalPort) GetAdditionalInfo() []PortAdditionalInfo {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []PortAdditionalInfo
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetAdditionalInfoOk() ([]PortAdditionalInfo, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *PhysicalPort) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []PortAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *PhysicalPort) SetAdditionalInfo(v []PortAdditionalInfo) {
	o.AdditionalInfo = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PhysicalPort) GetOrder() PortOrder {
	if o == nil || IsNil(o.Order) {
		var ret PortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetOrderOk() (*PortOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PhysicalPort) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given PortOrder and assigns it to the Order field.
func (o *PhysicalPort) SetOrder(v PortOrder) {
	o.Order = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *PhysicalPort) GetOperation() PortOperation {
	if o == nil || IsNil(o.Operation) {
		var ret PortOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetOperationOk() (*PortOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *PhysicalPort) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given PortOperation and assigns it to the Operation field.
func (o *PhysicalPort) SetOperation(v PortOperation) {
	o.Operation = &v
}

// GetLoas returns the Loas field value if set, zero value otherwise.
func (o *PhysicalPort) GetLoas() []PortLoa {
	if o == nil || IsNil(o.Loas) {
		var ret []PortLoa
		return ret
	}
	return o.Loas
}

// GetLoasOk returns a tuple with the Loas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPort) GetLoasOk() ([]PortLoa, bool) {
	if o == nil || IsNil(o.Loas) {
		return nil, false
	}
	return o.Loas, true
}

// HasLoas returns a boolean if a field has been set.
func (o *PhysicalPort) HasLoas() bool {
	if o != nil && !IsNil(o.Loas) {
		return true
	}

	return false
}

// SetLoas gets a reference to the given []PortLoa and assigns it to the Loas field.
func (o *PhysicalPort) SetLoas(v []PortLoa) {
	o.Loas = v
}

func (o PhysicalPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhysicalPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.InterfaceSpeed) {
		toSerialize["interfaceSpeed"] = o.InterfaceSpeed
	}
	if !IsNil(o.InterfaceType) {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Tether) {
		toSerialize["tether"] = o.Tether
	}
	if !IsNil(o.DemarcationPoint) {
		toSerialize["demarcationPoint"] = o.DemarcationPoint
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Loas) {
		toSerialize["loas"] = o.Loas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PhysicalPort) UnmarshalJSON(data []byte) (err error) {
	varPhysicalPort := _PhysicalPort{}

	err = json.Unmarshal(data, &varPhysicalPort)

	if err != nil {
		return err
	}

	*o = PhysicalPort(varPhysicalPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "href")
		delete(additionalProperties, "state")
		delete(additionalProperties, "account")
		delete(additionalProperties, "interfaceSpeed")
		delete(additionalProperties, "interfaceType")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "tether")
		delete(additionalProperties, "demarcationPoint")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "interface")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "order")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "loas")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePhysicalPort struct {
	value *PhysicalPort
	isSet bool
}

func (v NullablePhysicalPort) Get() *PhysicalPort {
	return v.value
}

func (v *NullablePhysicalPort) Set(val *PhysicalPort) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalPort) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalPort(val *PhysicalPort) *NullablePhysicalPort {
	return &NullablePhysicalPort{value: val, isSet: true}
}

func (v NullablePhysicalPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
