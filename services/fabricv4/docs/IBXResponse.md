# IBXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]IBX**](IBX.md) | List of IBXs by Metro. | [optional] 

## Methods

### NewIBXResponse

`func NewIBXResponse() *IBXResponse`

NewIBXResponse instantiates a new IBXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBXResponseWithDefaults

`func NewIBXResponseWithDefaults() *IBXResponse`

NewIBXResponseWithDefaults instantiates a new IBXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *IBXResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IBXResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IBXResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *IBXResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *IBXResponse) GetData() []IBX`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IBXResponse) GetDataOk() (*[]IBX, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IBXResponse) SetData(v []IBX)`

SetData sets Data field to given value.

### HasData

`func (o *IBXResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


