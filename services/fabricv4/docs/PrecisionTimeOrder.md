# PrecisionTimeOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number | [optional] 
**CustomerReferenceNumber** | Pointer to **string** | Customer reference number | [optional] 
**OrderNumber** | Pointer to **string** | Order Reference Number | [optional] 
**OrderStatus** | Pointer to **string** | Order status | [optional] 
**OrderType** | Pointer to **string** | Order channel type | [optional] 
**EffectiveDateTime** | Pointer to **time.Time** | Order Effective Date | [optional] 

## Methods

### NewPrecisionTimeOrder

`func NewPrecisionTimeOrder() *PrecisionTimeOrder`

NewPrecisionTimeOrder instantiates a new PrecisionTimeOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimeOrderWithDefaults

`func NewPrecisionTimeOrderWithDefaults() *PrecisionTimeOrder`

NewPrecisionTimeOrderWithDefaults instantiates a new PrecisionTimeOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchaseOrderNumber

`func (o *PrecisionTimeOrder) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *PrecisionTimeOrder) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *PrecisionTimeOrder) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *PrecisionTimeOrder) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetCustomerReferenceNumber

`func (o *PrecisionTimeOrder) GetCustomerReferenceNumber() string`

GetCustomerReferenceNumber returns the CustomerReferenceNumber field if non-nil, zero value otherwise.

### GetCustomerReferenceNumberOk

`func (o *PrecisionTimeOrder) GetCustomerReferenceNumberOk() (*string, bool)`

GetCustomerReferenceNumberOk returns a tuple with the CustomerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceNumber

`func (o *PrecisionTimeOrder) SetCustomerReferenceNumber(v string)`

SetCustomerReferenceNumber sets CustomerReferenceNumber field to given value.

### HasCustomerReferenceNumber

`func (o *PrecisionTimeOrder) HasCustomerReferenceNumber() bool`

HasCustomerReferenceNumber returns a boolean if a field has been set.

### GetOrderNumber

`func (o *PrecisionTimeOrder) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *PrecisionTimeOrder) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *PrecisionTimeOrder) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *PrecisionTimeOrder) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetOrderStatus

`func (o *PrecisionTimeOrder) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *PrecisionTimeOrder) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *PrecisionTimeOrder) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *PrecisionTimeOrder) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderType

`func (o *PrecisionTimeOrder) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *PrecisionTimeOrder) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *PrecisionTimeOrder) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *PrecisionTimeOrder) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetEffectiveDateTime

`func (o *PrecisionTimeOrder) GetEffectiveDateTime() time.Time`

GetEffectiveDateTime returns the EffectiveDateTime field if non-nil, zero value otherwise.

### GetEffectiveDateTimeOk

`func (o *PrecisionTimeOrder) GetEffectiveDateTimeOk() (*time.Time, bool)`

GetEffectiveDateTimeOk returns a tuple with the EffectiveDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDateTime

`func (o *PrecisionTimeOrder) SetEffectiveDateTime(v time.Time)`

SetEffectiveDateTime sets EffectiveDateTime field to given value.

### HasEffectiveDateTime

`func (o *PrecisionTimeOrder) HasEffectiveDateTime() bool`

HasEffectiveDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


