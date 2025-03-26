# TransactionStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | Pointer to **string** | for internal users | [optional] 
**Type** | Pointer to **string** | stage type | [optional] 
**State** | Pointer to **string** | state | [optional] 
**InitiatedDateTime** | Pointer to **time.Time** | Created by Date and Time | [optional] 
**CompletedDateTime** | Pointer to **time.Time** | Created by Date and Time | [optional] 
**Duration** | Pointer to **string** | Connection name | [optional] 
**DurationSlo** | Pointer to **string** | for internal users | [optional] 

## Methods

### NewTransactionStage

`func NewTransactionStage() *TransactionStage`

NewTransactionStage instantiates a new TransactionStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStageWithDefaults

`func NewTransactionStageWithDefaults() *TransactionStage`

NewTransactionStageWithDefaults instantiates a new TransactionStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *TransactionStage) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *TransactionStage) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *TransactionStage) SetExternal(v string)`

SetExternal sets External field to given value.

### HasExternal

`func (o *TransactionStage) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetType

`func (o *TransactionStage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionStage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionStage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionStage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *TransactionStage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TransactionStage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TransactionStage) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TransactionStage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetInitiatedDateTime

`func (o *TransactionStage) GetInitiatedDateTime() time.Time`

GetInitiatedDateTime returns the InitiatedDateTime field if non-nil, zero value otherwise.

### GetInitiatedDateTimeOk

`func (o *TransactionStage) GetInitiatedDateTimeOk() (*time.Time, bool)`

GetInitiatedDateTimeOk returns a tuple with the InitiatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedDateTime

`func (o *TransactionStage) SetInitiatedDateTime(v time.Time)`

SetInitiatedDateTime sets InitiatedDateTime field to given value.

### HasInitiatedDateTime

`func (o *TransactionStage) HasInitiatedDateTime() bool`

HasInitiatedDateTime returns a boolean if a field has been set.

### GetCompletedDateTime

`func (o *TransactionStage) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *TransactionStage) GetCompletedDateTimeOk() (*time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *TransactionStage) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *TransactionStage) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### GetDuration

`func (o *TransactionStage) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TransactionStage) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TransactionStage) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TransactionStage) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationSlo

`func (o *TransactionStage) GetDurationSlo() string`

GetDurationSlo returns the DurationSlo field if non-nil, zero value otherwise.

### GetDurationSloOk

`func (o *TransactionStage) GetDurationSloOk() (*string, bool)`

GetDurationSloOk returns a tuple with the DurationSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSlo

`func (o *TransactionStage) SetDurationSlo(v string)`

SetDurationSlo sets DurationSlo field to given value.

### HasDurationSlo

`func (o *TransactionStage) HasDurationSlo() bool`

HasDurationSlo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


