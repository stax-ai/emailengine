# MessagesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | [**SearchQuery**](SearchQuery.md) |  | 
**Update** | Pointer to [**MessageUpdate**](MessageUpdate.md) |  | [optional] 

## Methods

### NewMessagesUpdateRequest

`func NewMessagesUpdateRequest(search SearchQuery, ) *MessagesUpdateRequest`

NewMessagesUpdateRequest instantiates a new MessagesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesUpdateRequestWithDefaults

`func NewMessagesUpdateRequestWithDefaults() *MessagesUpdateRequest`

NewMessagesUpdateRequestWithDefaults instantiates a new MessagesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *MessagesUpdateRequest) GetSearch() SearchQuery`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *MessagesUpdateRequest) GetSearchOk() (*SearchQuery, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *MessagesUpdateRequest) SetSearch(v SearchQuery)`

SetSearch sets Search field to given value.


### GetUpdate

`func (o *MessagesUpdateRequest) GetUpdate() MessageUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *MessagesUpdateRequest) GetUpdateOk() (*MessageUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *MessagesUpdateRequest) SetUpdate(v MessageUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *MessagesUpdateRequest) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


