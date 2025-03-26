# SimplifiedLinkProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LinkProtocolType**](LinkProtocolType.md) |  | [optional] 
**VlanTag** | Pointer to **int32** | vlanTag value specified for DOT1Q connections | [optional] 
**VlanTagMin** | Pointer to **int32** | vlanTag Min value specified for DOT1Q connections | [optional] 
**VlanTagMax** | Pointer to **int32** | vlanTag Max value specified for DOT1Q connections | [optional] 
**VlanSTag** | Pointer to **int32** | vlanSTag value specified for QINQ connections | [optional] 
**VlanCTag** | Pointer to **int32** | vlanCTag value specified for QINQ connections | [optional] 
**VlanCTagMin** | Pointer to **int32** | vlanCTag Minvalue specified for QINQ connections | [optional] 
**VlanCTagMax** | Pointer to **int32** | vlanCTag max value specified for QINQ connections | [optional] 
**Unit** | Pointer to **int32** |  | [optional] 
**Vni** | Pointer to **int32** |  | [optional] 
**IntUnit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSimplifiedLinkProtocol

`func NewSimplifiedLinkProtocol() *SimplifiedLinkProtocol`

NewSimplifiedLinkProtocol instantiates a new SimplifiedLinkProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedLinkProtocolWithDefaults

`func NewSimplifiedLinkProtocolWithDefaults() *SimplifiedLinkProtocol`

NewSimplifiedLinkProtocolWithDefaults instantiates a new SimplifiedLinkProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SimplifiedLinkProtocol) GetType() LinkProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedLinkProtocol) GetTypeOk() (*LinkProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedLinkProtocol) SetType(v LinkProtocolType)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedLinkProtocol) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlanTag

`func (o *SimplifiedLinkProtocol) GetVlanTag() int32`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *SimplifiedLinkProtocol) GetVlanTagOk() (*int32, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *SimplifiedLinkProtocol) SetVlanTag(v int32)`

SetVlanTag sets VlanTag field to given value.

### HasVlanTag

`func (o *SimplifiedLinkProtocol) HasVlanTag() bool`

HasVlanTag returns a boolean if a field has been set.

### GetVlanTagMin

`func (o *SimplifiedLinkProtocol) GetVlanTagMin() int32`

GetVlanTagMin returns the VlanTagMin field if non-nil, zero value otherwise.

### GetVlanTagMinOk

`func (o *SimplifiedLinkProtocol) GetVlanTagMinOk() (*int32, bool)`

GetVlanTagMinOk returns a tuple with the VlanTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMin

`func (o *SimplifiedLinkProtocol) SetVlanTagMin(v int32)`

SetVlanTagMin sets VlanTagMin field to given value.

### HasVlanTagMin

`func (o *SimplifiedLinkProtocol) HasVlanTagMin() bool`

HasVlanTagMin returns a boolean if a field has been set.

### GetVlanTagMax

`func (o *SimplifiedLinkProtocol) GetVlanTagMax() int32`

GetVlanTagMax returns the VlanTagMax field if non-nil, zero value otherwise.

### GetVlanTagMaxOk

`func (o *SimplifiedLinkProtocol) GetVlanTagMaxOk() (*int32, bool)`

GetVlanTagMaxOk returns a tuple with the VlanTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMax

`func (o *SimplifiedLinkProtocol) SetVlanTagMax(v int32)`

SetVlanTagMax sets VlanTagMax field to given value.

### HasVlanTagMax

`func (o *SimplifiedLinkProtocol) HasVlanTagMax() bool`

HasVlanTagMax returns a boolean if a field has been set.

### GetVlanSTag

`func (o *SimplifiedLinkProtocol) GetVlanSTag() int32`

GetVlanSTag returns the VlanSTag field if non-nil, zero value otherwise.

### GetVlanSTagOk

`func (o *SimplifiedLinkProtocol) GetVlanSTagOk() (*int32, bool)`

GetVlanSTagOk returns a tuple with the VlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSTag

`func (o *SimplifiedLinkProtocol) SetVlanSTag(v int32)`

SetVlanSTag sets VlanSTag field to given value.

### HasVlanSTag

`func (o *SimplifiedLinkProtocol) HasVlanSTag() bool`

HasVlanSTag returns a boolean if a field has been set.

### GetVlanCTag

`func (o *SimplifiedLinkProtocol) GetVlanCTag() int32`

GetVlanCTag returns the VlanCTag field if non-nil, zero value otherwise.

### GetVlanCTagOk

`func (o *SimplifiedLinkProtocol) GetVlanCTagOk() (*int32, bool)`

GetVlanCTagOk returns a tuple with the VlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTag

`func (o *SimplifiedLinkProtocol) SetVlanCTag(v int32)`

SetVlanCTag sets VlanCTag field to given value.

### HasVlanCTag

`func (o *SimplifiedLinkProtocol) HasVlanCTag() bool`

HasVlanCTag returns a boolean if a field has been set.

### GetVlanCTagMin

`func (o *SimplifiedLinkProtocol) GetVlanCTagMin() int32`

GetVlanCTagMin returns the VlanCTagMin field if non-nil, zero value otherwise.

### GetVlanCTagMinOk

`func (o *SimplifiedLinkProtocol) GetVlanCTagMinOk() (*int32, bool)`

GetVlanCTagMinOk returns a tuple with the VlanCTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMin

`func (o *SimplifiedLinkProtocol) SetVlanCTagMin(v int32)`

SetVlanCTagMin sets VlanCTagMin field to given value.

### HasVlanCTagMin

`func (o *SimplifiedLinkProtocol) HasVlanCTagMin() bool`

HasVlanCTagMin returns a boolean if a field has been set.

### GetVlanCTagMax

`func (o *SimplifiedLinkProtocol) GetVlanCTagMax() int32`

GetVlanCTagMax returns the VlanCTagMax field if non-nil, zero value otherwise.

### GetVlanCTagMaxOk

`func (o *SimplifiedLinkProtocol) GetVlanCTagMaxOk() (*int32, bool)`

GetVlanCTagMaxOk returns a tuple with the VlanCTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMax

`func (o *SimplifiedLinkProtocol) SetVlanCTagMax(v int32)`

SetVlanCTagMax sets VlanCTagMax field to given value.

### HasVlanCTagMax

`func (o *SimplifiedLinkProtocol) HasVlanCTagMax() bool`

HasVlanCTagMax returns a boolean if a field has been set.

### GetUnit

`func (o *SimplifiedLinkProtocol) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SimplifiedLinkProtocol) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SimplifiedLinkProtocol) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SimplifiedLinkProtocol) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVni

`func (o *SimplifiedLinkProtocol) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *SimplifiedLinkProtocol) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *SimplifiedLinkProtocol) SetVni(v int32)`

SetVni sets Vni field to given value.

### HasVni

`func (o *SimplifiedLinkProtocol) HasVni() bool`

HasVni returns a boolean if a field has been set.

### GetIntUnit

`func (o *SimplifiedLinkProtocol) GetIntUnit() int32`

GetIntUnit returns the IntUnit field if non-nil, zero value otherwise.

### GetIntUnitOk

`func (o *SimplifiedLinkProtocol) GetIntUnitOk() (*int32, bool)`

GetIntUnitOk returns a tuple with the IntUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntUnit

`func (o *SimplifiedLinkProtocol) SetIntUnit(v int32)`

SetIntUnit sets IntUnit field to given value.

### HasIntUnit

`func (o *SimplifiedLinkProtocol) HasIntUnit() bool`

HasIntUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


