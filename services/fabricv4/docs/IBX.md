# IBX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | The Canonical URL at which the resource resides. | [optional] 
**Type** | Pointer to **string** | Indicator of a IBX resource. | [optional] 
**Code** | Pointer to **string** | Code Assigned to an Equinix IBX data center in a specified metropolitan area. | [optional] 
**IsTimeServiceEnabled** | Pointer to **bool** | Indicates if Precision Time Service enabled in IBX or not. | [optional] 

## Methods

### NewIBX

`func NewIBX() *IBX`

NewIBX instantiates a new IBX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBXWithDefaults

`func NewIBXWithDefaults() *IBX`

NewIBXWithDefaults instantiates a new IBX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *IBX) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IBX) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IBX) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IBX) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *IBX) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IBX) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IBX) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IBX) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *IBX) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IBX) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IBX) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IBX) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIsTimeServiceEnabled

`func (o *IBX) GetIsTimeServiceEnabled() bool`

GetIsTimeServiceEnabled returns the IsTimeServiceEnabled field if non-nil, zero value otherwise.

### GetIsTimeServiceEnabledOk

`func (o *IBX) GetIsTimeServiceEnabledOk() (*bool, bool)`

GetIsTimeServiceEnabledOk returns a tuple with the IsTimeServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeServiceEnabled

`func (o *IBX) SetIsTimeServiceEnabled(v bool)`

SetIsTimeServiceEnabled sets IsTimeServiceEnabled field to given value.

### HasIsTimeServiceEnabled

`func (o *IBX) HasIsTimeServiceEnabled() bool`

HasIsTimeServiceEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


