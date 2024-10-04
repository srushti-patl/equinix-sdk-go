# Stream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Stream URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**State** | Pointer to **string** | state | [optional] 
**AssetCount** | Pointer to **int32** | stream asset count | [optional] 
**StreamSubscriptionCount** | Pointer to **int32** | stream subscription count | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Type** | Pointer to [**StreamPostRequestType**](StreamPostRequestType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided stream name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream description | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 

## Methods

### NewStream

`func NewStream() *Stream`

NewStream instantiates a new Stream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWithDefaults

`func NewStreamWithDefaults() *Stream`

NewStreamWithDefaults instantiates a new Stream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Stream) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Stream) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Stream) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Stream) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Stream) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Stream) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Stream) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Stream) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *Stream) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Stream) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Stream) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Stream) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAssetCount

`func (o *Stream) GetAssetCount() int32`

GetAssetCount returns the AssetCount field if non-nil, zero value otherwise.

### GetAssetCountOk

`func (o *Stream) GetAssetCountOk() (*int32, bool)`

GetAssetCountOk returns a tuple with the AssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCount

`func (o *Stream) SetAssetCount(v int32)`

SetAssetCount sets AssetCount field to given value.

### HasAssetCount

`func (o *Stream) HasAssetCount() bool`

HasAssetCount returns a boolean if a field has been set.

### GetStreamSubscriptionCount

`func (o *Stream) GetStreamSubscriptionCount() int32`

GetStreamSubscriptionCount returns the StreamSubscriptionCount field if non-nil, zero value otherwise.

### GetStreamSubscriptionCountOk

`func (o *Stream) GetStreamSubscriptionCountOk() (*int32, bool)`

GetStreamSubscriptionCountOk returns a tuple with the StreamSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSubscriptionCount

`func (o *Stream) SetStreamSubscriptionCount(v int32)`

SetStreamSubscriptionCount sets StreamSubscriptionCount field to given value.

### HasStreamSubscriptionCount

`func (o *Stream) HasStreamSubscriptionCount() bool`

HasStreamSubscriptionCount returns a boolean if a field has been set.

### GetChangelog

`func (o *Stream) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *Stream) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *Stream) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *Stream) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetType

`func (o *Stream) GetType() StreamPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Stream) GetTypeOk() (*StreamPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Stream) SetType(v StreamPostRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *Stream) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Stream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Stream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Stream) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Stream) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Stream) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Stream) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *Stream) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Stream) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Stream) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Stream) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


