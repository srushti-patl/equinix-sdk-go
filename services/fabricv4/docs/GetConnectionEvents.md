# GetConnectionEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to [**[]AdditionalInfoProperty**](AdditionalInfoProperty.md) | Connection additional information | [optional] 
**Connection** | Pointer to [**EventConnection**](EventConnection.md) |  | [optional] 

## Methods

### NewGetConnectionEvents

`func NewGetConnectionEvents() *GetConnectionEvents`

NewGetConnectionEvents instantiates a new GetConnectionEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionEventsWithDefaults

`func NewGetConnectionEventsWithDefaults() *GetConnectionEvents`

NewGetConnectionEventsWithDefaults instantiates a new GetConnectionEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *GetConnectionEvents) GetAdditionalInfo() []AdditionalInfoProperty`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *GetConnectionEvents) GetAdditionalInfoOk() (*[]AdditionalInfoProperty, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *GetConnectionEvents) SetAdditionalInfo(v []AdditionalInfoProperty)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *GetConnectionEvents) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetConnection

`func (o *GetConnectionEvents) GetConnection() EventConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *GetConnectionEvents) GetConnectionOk() (*EventConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *GetConnectionEvents) SetConnection(v EventConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *GetConnectionEvents) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


