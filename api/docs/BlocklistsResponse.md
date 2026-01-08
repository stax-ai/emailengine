# BlocklistsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocklists** | Pointer to [**[]BlocklistsResponseItem**](BlocklistsResponseItem.md) |  | [optional] 
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 

## Methods

### NewBlocklistsResponse

`func NewBlocklistsResponse() *BlocklistsResponse`

NewBlocklistsResponse instantiates a new BlocklistsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistsResponseWithDefaults

`func NewBlocklistsResponseWithDefaults() *BlocklistsResponse`

NewBlocklistsResponseWithDefaults instantiates a new BlocklistsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocklists

`func (o *BlocklistsResponse) GetBlocklists() []BlocklistsResponseItem`

GetBlocklists returns the Blocklists field if non-nil, zero value otherwise.

### GetBlocklistsOk

`func (o *BlocklistsResponse) GetBlocklistsOk() (*[]BlocklistsResponseItem, bool)`

GetBlocklistsOk returns a tuple with the Blocklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocklists

`func (o *BlocklistsResponse) SetBlocklists(v []BlocklistsResponseItem)`

SetBlocklists sets Blocklists field to given value.

### HasBlocklists

`func (o *BlocklistsResponse) HasBlocklists() bool`

HasBlocklists returns a boolean if a field has been set.

### GetPage

`func (o *BlocklistsResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BlocklistsResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BlocklistsResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *BlocklistsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *BlocklistsResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *BlocklistsResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *BlocklistsResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *BlocklistsResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *BlocklistsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BlocklistsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BlocklistsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BlocklistsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


