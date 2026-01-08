# CreateMailboxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Was the mailbox created | [optional] 
**MailboxId** | Pointer to **string** | Mailbox ID (if server has support) | [optional] 
**Path** | **string** | Full path to mailbox | 

## Methods

### NewCreateMailboxResponse

`func NewCreateMailboxResponse(path string, ) *CreateMailboxResponse`

NewCreateMailboxResponse instantiates a new CreateMailboxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMailboxResponseWithDefaults

`func NewCreateMailboxResponseWithDefaults() *CreateMailboxResponse`

NewCreateMailboxResponseWithDefaults instantiates a new CreateMailboxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CreateMailboxResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateMailboxResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateMailboxResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateMailboxResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetMailboxId

`func (o *CreateMailboxResponse) GetMailboxId() string`

GetMailboxId returns the MailboxId field if non-nil, zero value otherwise.

### GetMailboxIdOk

`func (o *CreateMailboxResponse) GetMailboxIdOk() (*string, bool)`

GetMailboxIdOk returns a tuple with the MailboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxId

`func (o *CreateMailboxResponse) SetMailboxId(v string)`

SetMailboxId sets MailboxId field to given value.

### HasMailboxId

`func (o *CreateMailboxResponse) HasMailboxId() bool`

HasMailboxId returns a boolean if a field has been set.

### GetPath

`func (o *CreateMailboxResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateMailboxResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateMailboxResponse) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


