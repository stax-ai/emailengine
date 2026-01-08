# MessagesDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | [**SearchQuery**](SearchQuery.md) |  | 

## Methods

### NewMessagesDeleteRequest

`func NewMessagesDeleteRequest(search SearchQuery, ) *MessagesDeleteRequest`

NewMessagesDeleteRequest instantiates a new MessagesDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesDeleteRequestWithDefaults

`func NewMessagesDeleteRequestWithDefaults() *MessagesDeleteRequest`

NewMessagesDeleteRequestWithDefaults instantiates a new MessagesDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *MessagesDeleteRequest) GetSearch() SearchQuery`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *MessagesDeleteRequest) GetSearchOk() (*SearchQuery, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *MessagesDeleteRequest) SetSearch(v SearchQuery)`

SetSearch sets Search field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


