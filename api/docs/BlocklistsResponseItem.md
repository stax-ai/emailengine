# BlocklistsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count of blocked addresses in this list | [optional] 
**ListId** | **string** | List ID | 

## Methods

### NewBlocklistsResponseItem

`func NewBlocklistsResponseItem(listId string, ) *BlocklistsResponseItem`

NewBlocklistsResponseItem instantiates a new BlocklistsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistsResponseItemWithDefaults

`func NewBlocklistsResponseItemWithDefaults() *BlocklistsResponseItem`

NewBlocklistsResponseItemWithDefaults instantiates a new BlocklistsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BlocklistsResponseItem) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BlocklistsResponseItem) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BlocklistsResponseItem) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BlocklistsResponseItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetListId

`func (o *BlocklistsResponseItem) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *BlocklistsResponseItem) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *BlocklistsResponseItem) SetListId(v string)`

SetListId sets ListId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


