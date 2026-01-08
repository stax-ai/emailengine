# MessagesMoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Target mailbox folder path | 
**Search** | [**SearchQuery**](SearchQuery.md) |  | 

## Methods

### NewMessagesMoveRequest

`func NewMessagesMoveRequest(path string, search SearchQuery, ) *MessagesMoveRequest`

NewMessagesMoveRequest instantiates a new MessagesMoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesMoveRequestWithDefaults

`func NewMessagesMoveRequestWithDefaults() *MessagesMoveRequest`

NewMessagesMoveRequestWithDefaults instantiates a new MessagesMoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *MessagesMoveRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MessagesMoveRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MessagesMoveRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetSearch

`func (o *MessagesMoveRequest) GetSearch() SearchQuery`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *MessagesMoveRequest) GetSearchOk() (*SearchQuery, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *MessagesMoveRequest) SetSearch(v SearchQuery)`

SetSearch sets Search field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


