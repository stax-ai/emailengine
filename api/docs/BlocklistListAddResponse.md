# BlocklistListAddResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to **bool** | Was the address added to the list | [optional] 
**Success** | Pointer to **bool** | Was the request successful | [optional] 

## Methods

### NewBlocklistListAddResponse

`func NewBlocklistListAddResponse() *BlocklistListAddResponse`

NewBlocklistListAddResponse instantiates a new BlocklistListAddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistListAddResponseWithDefaults

`func NewBlocklistListAddResponseWithDefaults() *BlocklistListAddResponse`

NewBlocklistListAddResponseWithDefaults instantiates a new BlocklistListAddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *BlocklistListAddResponse) GetAdded() bool`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *BlocklistListAddResponse) GetAddedOk() (*bool, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *BlocklistListAddResponse) SetAdded(v bool)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *BlocklistListAddResponse) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetSuccess

`func (o *BlocklistListAddResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BlocklistListAddResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BlocklistListAddResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BlocklistListAddResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


