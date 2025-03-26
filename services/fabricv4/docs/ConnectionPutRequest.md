# ConnectionPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ConnectionType**](ConnectionType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided connection name | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 
**Bandwidth** | Pointer to **int32** | Connection bandwidth in Mbps | [optional] 
**Redundancy** | Pointer to [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | [optional] 
**ASide** | Pointer to [**ConnectionSide**](ConnectionSide.md) |  | [optional] 
**ZSide** | Pointer to [**ConnectionSide**](ConnectionSide.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**AdditionalInfo** | Pointer to [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Connection additional information | [optional] 

## Methods

### NewConnectionPutRequest

`func NewConnectionPutRequest() *ConnectionPutRequest`

NewConnectionPutRequest instantiates a new ConnectionPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPutRequestWithDefaults

`func NewConnectionPutRequestWithDefaults() *ConnectionPutRequest`

NewConnectionPutRequestWithDefaults instantiates a new ConnectionPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionPutRequest) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionPutRequest) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionPutRequest) SetType(v ConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionPutRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ConnectionPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *ConnectionPutRequest) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConnectionPutRequest) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConnectionPutRequest) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConnectionPutRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetNotifications

`func (o *ConnectionPutRequest) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ConnectionPutRequest) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ConnectionPutRequest) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ConnectionPutRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetBandwidth

`func (o *ConnectionPutRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ConnectionPutRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ConnectionPutRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ConnectionPutRequest) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetRedundancy

`func (o *ConnectionPutRequest) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *ConnectionPutRequest) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *ConnectionPutRequest) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *ConnectionPutRequest) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetASide

`func (o *ConnectionPutRequest) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *ConnectionPutRequest) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *ConnectionPutRequest) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.

### HasASide

`func (o *ConnectionPutRequest) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *ConnectionPutRequest) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *ConnectionPutRequest) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *ConnectionPutRequest) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *ConnectionPutRequest) HasZSide() bool`

HasZSide returns a boolean if a field has been set.

### GetProject

`func (o *ConnectionPutRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ConnectionPutRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ConnectionPutRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ConnectionPutRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *ConnectionPutRequest) GetAdditionalInfo() []ConnectionSideAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ConnectionPutRequest) GetAdditionalInfoOk() (*[]ConnectionSideAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ConnectionPutRequest) SetAdditionalInfo(v []ConnectionSideAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ConnectionPutRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


