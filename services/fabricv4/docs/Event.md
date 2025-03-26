# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Event source reference | [optional] 
**Type** | Pointer to [**Type**](Type.md) |  | [optional] 
**Uuid** | Pointer to **string** | Event UUID | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | startDateTime | [optional] 
**AdditionalInfo** | Pointer to [**[]AdditionalInfoProperty**](AdditionalInfoProperty.md) | Connection additional information | [optional] 
**Connection** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 
**SubType** | Pointer to [**GetRpOperationStatusEventSubType**](GetRpOperationStatusEventSubType.md) |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** | event description | [optional] 
**RoutingProtocol** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 
**Router** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Event) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Event) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Event) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Event) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *Event) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Event) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Event) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Event) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Event) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Event) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Event) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Event) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *Event) GetAdditionalInfo() []AdditionalInfoProperty`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *Event) GetAdditionalInfoOk() (*[]AdditionalInfoProperty, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *Event) SetAdditionalInfo(v []AdditionalInfoProperty)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *Event) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetConnection

`func (o *Event) GetConnection() EventMetadata`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *Event) GetConnectionOk() (*EventMetadata, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *Event) SetConnection(v EventMetadata)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *Event) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetSubType

`func (o *Event) GetSubType() GetRpOperationStatusEventSubType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Event) GetSubTypeOk() (*GetRpOperationStatusEventSubType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Event) SetSubType(v GetRpOperationStatusEventSubType)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Event) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSeverity

`func (o *Event) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Event) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Event) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Event) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetMessage

`func (o *Event) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Event) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Event) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Event) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRoutingProtocol

`func (o *Event) GetRoutingProtocol() EventMetadata`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *Event) GetRoutingProtocolOk() (*EventMetadata, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *Event) SetRoutingProtocol(v EventMetadata)`

SetRoutingProtocol sets RoutingProtocol field to given value.

### HasRoutingProtocol

`func (o *Event) HasRoutingProtocol() bool`

HasRoutingProtocol returns a boolean if a field has been set.

### GetRouter

`func (o *Event) GetRouter() EventMetadata`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *Event) GetRouterOk() (*EventMetadata, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *Event) SetRouter(v EventMetadata)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *Event) HasRouter() bool`

HasRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


