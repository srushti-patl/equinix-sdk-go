# ConnectionTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | transaction type | [optional] 
**State** | Pointer to **string** | transaction state | [optional] 
**InitiatedDateTime** | Pointer to **time.Time** | Created by Date and Time | [optional] 
**Duration** | Pointer to **string** | duration | [optional] 
**DurationSlo** | Pointer to **string** | for internal users | [optional] 
**Stage** | Pointer to [**TransactionStage**](TransactionStage.md) |  | [optional] 

## Methods

### NewConnectionTransaction

`func NewConnectionTransaction() *ConnectionTransaction`

NewConnectionTransaction instantiates a new ConnectionTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTransactionWithDefaults

`func NewConnectionTransactionWithDefaults() *ConnectionTransaction`

NewConnectionTransactionWithDefaults instantiates a new ConnectionTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *ConnectionTransaction) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectionTransaction) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectionTransaction) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ConnectionTransaction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetInitiatedDateTime

`func (o *ConnectionTransaction) GetInitiatedDateTime() time.Time`

GetInitiatedDateTime returns the InitiatedDateTime field if non-nil, zero value otherwise.

### GetInitiatedDateTimeOk

`func (o *ConnectionTransaction) GetInitiatedDateTimeOk() (*time.Time, bool)`

GetInitiatedDateTimeOk returns a tuple with the InitiatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedDateTime

`func (o *ConnectionTransaction) SetInitiatedDateTime(v time.Time)`

SetInitiatedDateTime sets InitiatedDateTime field to given value.

### HasInitiatedDateTime

`func (o *ConnectionTransaction) HasInitiatedDateTime() bool`

HasInitiatedDateTime returns a boolean if a field has been set.

### GetDuration

`func (o *ConnectionTransaction) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ConnectionTransaction) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ConnectionTransaction) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ConnectionTransaction) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationSlo

`func (o *ConnectionTransaction) GetDurationSlo() string`

GetDurationSlo returns the DurationSlo field if non-nil, zero value otherwise.

### GetDurationSloOk

`func (o *ConnectionTransaction) GetDurationSloOk() (*string, bool)`

GetDurationSloOk returns a tuple with the DurationSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSlo

`func (o *ConnectionTransaction) SetDurationSlo(v string)`

SetDurationSlo sets DurationSlo field to given value.

### HasDurationSlo

`func (o *ConnectionTransaction) HasDurationSlo() bool`

HasDurationSlo returns a boolean if a field has been set.

### GetStage

`func (o *ConnectionTransaction) GetStage() TransactionStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ConnectionTransaction) GetStageOk() (*TransactionStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ConnectionTransaction) SetStage(v TransactionStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ConnectionTransaction) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


