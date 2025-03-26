# GetRpOperationStatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubType** | Pointer to [**GetRpOperationStatusEventSubType**](GetRpOperationStatusEventSubType.md) |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** | event description | [optional] 
**RoutingProtocol** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 
**Connection** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 
**Router** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 

## Methods

### NewGetRpOperationStatusEvent

`func NewGetRpOperationStatusEvent() *GetRpOperationStatusEvent`

NewGetRpOperationStatusEvent instantiates a new GetRpOperationStatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRpOperationStatusEventWithDefaults

`func NewGetRpOperationStatusEventWithDefaults() *GetRpOperationStatusEvent`

NewGetRpOperationStatusEventWithDefaults instantiates a new GetRpOperationStatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubType

`func (o *GetRpOperationStatusEvent) GetSubType() GetRpOperationStatusEventSubType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetRpOperationStatusEvent) GetSubTypeOk() (*GetRpOperationStatusEventSubType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetRpOperationStatusEvent) SetSubType(v GetRpOperationStatusEventSubType)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetRpOperationStatusEvent) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSeverity

`func (o *GetRpOperationStatusEvent) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetRpOperationStatusEvent) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetRpOperationStatusEvent) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetRpOperationStatusEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetMessage

`func (o *GetRpOperationStatusEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetRpOperationStatusEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetRpOperationStatusEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetRpOperationStatusEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRoutingProtocol

`func (o *GetRpOperationStatusEvent) GetRoutingProtocol() EventMetadata`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *GetRpOperationStatusEvent) GetRoutingProtocolOk() (*EventMetadata, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *GetRpOperationStatusEvent) SetRoutingProtocol(v EventMetadata)`

SetRoutingProtocol sets RoutingProtocol field to given value.

### HasRoutingProtocol

`func (o *GetRpOperationStatusEvent) HasRoutingProtocol() bool`

HasRoutingProtocol returns a boolean if a field has been set.

### GetConnection

`func (o *GetRpOperationStatusEvent) GetConnection() EventMetadata`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *GetRpOperationStatusEvent) GetConnectionOk() (*EventMetadata, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *GetRpOperationStatusEvent) SetConnection(v EventMetadata)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *GetRpOperationStatusEvent) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetRouter

`func (o *GetRpOperationStatusEvent) GetRouter() EventMetadata`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *GetRpOperationStatusEvent) GetRouterOk() (*EventMetadata, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *GetRpOperationStatusEvent) SetRouter(v EventMetadata)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *GetRpOperationStatusEvent) HasRouter() bool`

HasRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


