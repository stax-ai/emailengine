# ModifyMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPath** | Pointer to **[]string** | New mailbox path as an array or a string. If account is namespaced then namespace prefix is added by default. Optional. | [optional] 
**Path** | **string** | Mailbox folder path to modify | 
**Subscribed** | Pointer to **bool** | Change mailbox subscription status. Only applies to IMAP accounts, ignored for Gmail and Outlook. | [optional] 

## Methods

### NewModifyMailbox

`func NewModifyMailbox(path string, ) *ModifyMailbox`

NewModifyMailbox instantiates a new ModifyMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMailboxWithDefaults

`func NewModifyMailboxWithDefaults() *ModifyMailbox`

NewModifyMailboxWithDefaults instantiates a new ModifyMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPath

`func (o *ModifyMailbox) GetNewPath() []string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *ModifyMailbox) GetNewPathOk() (*[]string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *ModifyMailbox) SetNewPath(v []string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *ModifyMailbox) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### GetPath

`func (o *ModifyMailbox) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModifyMailbox) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModifyMailbox) SetPath(v string)`

SetPath sets Path field to given value.


### GetSubscribed

`func (o *ModifyMailbox) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *ModifyMailbox) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *ModifyMailbox) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *ModifyMailbox) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


