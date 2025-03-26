# EventConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] 
**Type** | Pointer to **string** | Connection type | [optional] 
**Uuid** | Pointer to **string** | Connection uuid | [optional] 
**Name** | Pointer to **string** | Connection name | [optional] 
**Transaction** | Pointer to [**ConnectionTransaction**](ConnectionTransaction.md) |  | [optional] 

## Methods

### NewEventConnection

`func NewEventConnection() *EventConnection`

NewEventConnection instantiates a new EventConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventConnectionWithDefaults

`func NewEventConnectionWithDefaults() *EventConnection`

NewEventConnectionWithDefaults instantiates a new EventConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *EventConnection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *EventConnection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *EventConnection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *EventConnection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *EventConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *EventConnection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EventConnection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EventConnection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EventConnection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *EventConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTransaction

`func (o *EventConnection) GetTransaction() ConnectionTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *EventConnection) GetTransactionOk() (*ConnectionTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *EventConnection) SetTransaction(v ConnectionTransaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *EventConnection) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


