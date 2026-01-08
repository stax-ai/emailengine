# BlocklistListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]BlocklistListResponseItem**](BlocklistListResponseItem.md) |  | [optional] 
**ListId** | **string** | List ID | 
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 

## Methods

### NewBlocklistListResponse

`func NewBlocklistListResponse(listId string, ) *BlocklistListResponse`

NewBlocklistListResponse instantiates a new BlocklistListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistListResponseWithDefaults

`func NewBlocklistListResponseWithDefaults() *BlocklistListResponse`

NewBlocklistListResponseWithDefaults instantiates a new BlocklistListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *BlocklistListResponse) GetAddresses() []BlocklistListResponseItem`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BlocklistListResponse) GetAddressesOk() (*[]BlocklistListResponseItem, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BlocklistListResponse) SetAddresses(v []BlocklistListResponseItem)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BlocklistListResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetListId

`func (o *BlocklistListResponse) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *BlocklistListResponse) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *BlocklistListResponse) SetListId(v string)`

SetListId sets ListId field to given value.


### GetPage

`func (o *BlocklistListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BlocklistListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BlocklistListResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *BlocklistListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *BlocklistListResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *BlocklistListResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *BlocklistListResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *BlocklistListResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *BlocklistListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BlocklistListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BlocklistListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BlocklistListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


