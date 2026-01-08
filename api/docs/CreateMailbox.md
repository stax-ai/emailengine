# CreateMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **[]string** | Mailbox path as an array or a string. If account is namespaced then namespace prefix is added by default. | [optional] 

## Methods

### NewCreateMailbox

`func NewCreateMailbox() *CreateMailbox`

NewCreateMailbox instantiates a new CreateMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMailboxWithDefaults

`func NewCreateMailboxWithDefaults() *CreateMailbox`

NewCreateMailboxWithDefaults instantiates a new CreateMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CreateMailbox) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateMailbox) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateMailbox) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateMailbox) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


