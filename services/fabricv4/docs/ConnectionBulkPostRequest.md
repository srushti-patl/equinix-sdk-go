# ConnectionBulkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConnectionPostRequest**](ConnectionPostRequest.md) |  | [optional] 

## Methods

### NewConnectionBulkPostRequest

`func NewConnectionBulkPostRequest() *ConnectionBulkPostRequest`

NewConnectionBulkPostRequest instantiates a new ConnectionBulkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionBulkPostRequestWithDefaults

`func NewConnectionBulkPostRequestWithDefaults() *ConnectionBulkPostRequest`

NewConnectionBulkPostRequestWithDefaults instantiates a new ConnectionBulkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConnectionBulkPostRequest) GetData() []ConnectionPostRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectionBulkPostRequest) GetDataOk() (*[]ConnectionPostRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectionBulkPostRequest) SetData(v []ConnectionPostRequest)`

SetData sets Data field to given value.

### HasData

`func (o *ConnectionBulkPostRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


