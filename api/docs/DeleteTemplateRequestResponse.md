# DeleteTemplateRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Deleted** | Pointer to **bool** | Were the templates flushed | [optional] [default to true]

## Methods

### NewDeleteTemplateRequestResponse

`func NewDeleteTemplateRequestResponse(account string, ) *DeleteTemplateRequestResponse`

NewDeleteTemplateRequestResponse instantiates a new DeleteTemplateRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTemplateRequestResponseWithDefaults

`func NewDeleteTemplateRequestResponseWithDefaults() *DeleteTemplateRequestResponse`

NewDeleteTemplateRequestResponseWithDefaults instantiates a new DeleteTemplateRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DeleteTemplateRequestResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DeleteTemplateRequestResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DeleteTemplateRequestResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetDeleted

`func (o *DeleteTemplateRequestResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteTemplateRequestResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteTemplateRequestResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteTemplateRequestResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


