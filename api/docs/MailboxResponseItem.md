# MailboxResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | Pointer to **string** | Hierarchy delimiter character used in paths | [optional] 
**Listed** | Pointer to **bool** | Whether this mailbox appears in LIST command results | [optional] 
**Messages** | Pointer to **int32** | Total number of messages in the mailbox | [optional] 
**Name** | **string** | Display name of the mailbox | 
**NoInferiors** | Pointer to **bool** | Whether this mailbox can contain child mailboxes | [optional] 
**ParentPath** | **string** | Path to the parent mailbox | 
**Path** | **string** | Full path to the mailbox | 
**SpecialUse** | Pointer to [**MailboxSpecialUse**](MailboxSpecialUse.md) |  | [optional] 
**SpecialUseSource** | Pointer to [**MailboxSpecialUseSource**](MailboxSpecialUseSource.md) |  | [optional] 
**Status** | Pointer to [**MailboxResponseStatus**](MailboxResponseStatus.md) |  | [optional] 
**Subscribed** | Pointer to **bool** | Whether the user is subscribed to this mailbox | [optional] 
**UidNext** | Pointer to **int32** | Next UID value that will be assigned | [optional] 

## Methods

### NewMailboxResponseItem

`func NewMailboxResponseItem(name string, parentPath string, path string, ) *MailboxResponseItem`

NewMailboxResponseItem instantiates a new MailboxResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxResponseItemWithDefaults

`func NewMailboxResponseItemWithDefaults() *MailboxResponseItem`

NewMailboxResponseItemWithDefaults instantiates a new MailboxResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *MailboxResponseItem) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *MailboxResponseItem) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *MailboxResponseItem) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *MailboxResponseItem) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetListed

`func (o *MailboxResponseItem) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *MailboxResponseItem) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *MailboxResponseItem) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *MailboxResponseItem) HasListed() bool`

HasListed returns a boolean if a field has been set.

### GetMessages

`func (o *MailboxResponseItem) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MailboxResponseItem) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MailboxResponseItem) SetMessages(v int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MailboxResponseItem) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetName

`func (o *MailboxResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MailboxResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MailboxResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetNoInferiors

`func (o *MailboxResponseItem) GetNoInferiors() bool`

GetNoInferiors returns the NoInferiors field if non-nil, zero value otherwise.

### GetNoInferiorsOk

`func (o *MailboxResponseItem) GetNoInferiorsOk() (*bool, bool)`

GetNoInferiorsOk returns a tuple with the NoInferiors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoInferiors

`func (o *MailboxResponseItem) SetNoInferiors(v bool)`

SetNoInferiors sets NoInferiors field to given value.

### HasNoInferiors

`func (o *MailboxResponseItem) HasNoInferiors() bool`

HasNoInferiors returns a boolean if a field has been set.

### GetParentPath

`func (o *MailboxResponseItem) GetParentPath() string`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *MailboxResponseItem) GetParentPathOk() (*string, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *MailboxResponseItem) SetParentPath(v string)`

SetParentPath sets ParentPath field to given value.


### GetPath

`func (o *MailboxResponseItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MailboxResponseItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MailboxResponseItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetSpecialUse

`func (o *MailboxResponseItem) GetSpecialUse() MailboxSpecialUse`

GetSpecialUse returns the SpecialUse field if non-nil, zero value otherwise.

### GetSpecialUseOk

`func (o *MailboxResponseItem) GetSpecialUseOk() (*MailboxSpecialUse, bool)`

GetSpecialUseOk returns a tuple with the SpecialUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialUse

`func (o *MailboxResponseItem) SetSpecialUse(v MailboxSpecialUse)`

SetSpecialUse sets SpecialUse field to given value.

### HasSpecialUse

`func (o *MailboxResponseItem) HasSpecialUse() bool`

HasSpecialUse returns a boolean if a field has been set.

### GetSpecialUseSource

`func (o *MailboxResponseItem) GetSpecialUseSource() MailboxSpecialUseSource`

GetSpecialUseSource returns the SpecialUseSource field if non-nil, zero value otherwise.

### GetSpecialUseSourceOk

`func (o *MailboxResponseItem) GetSpecialUseSourceOk() (*MailboxSpecialUseSource, bool)`

GetSpecialUseSourceOk returns a tuple with the SpecialUseSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialUseSource

`func (o *MailboxResponseItem) SetSpecialUseSource(v MailboxSpecialUseSource)`

SetSpecialUseSource sets SpecialUseSource field to given value.

### HasSpecialUseSource

`func (o *MailboxResponseItem) HasSpecialUseSource() bool`

HasSpecialUseSource returns a boolean if a field has been set.

### GetStatus

`func (o *MailboxResponseItem) GetStatus() MailboxResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MailboxResponseItem) GetStatusOk() (*MailboxResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MailboxResponseItem) SetStatus(v MailboxResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MailboxResponseItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscribed

`func (o *MailboxResponseItem) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *MailboxResponseItem) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *MailboxResponseItem) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *MailboxResponseItem) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.

### GetUidNext

`func (o *MailboxResponseItem) GetUidNext() int32`

GetUidNext returns the UidNext field if non-nil, zero value otherwise.

### GetUidNextOk

`func (o *MailboxResponseItem) GetUidNextOk() (*int32, bool)`

GetUidNextOk returns a tuple with the UidNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidNext

`func (o *MailboxResponseItem) SetUidNext(v int32)`

SetUidNext sets UidNext field to given value.

### HasUidNext

`func (o *MailboxResponseItem) HasUidNext() bool`

HasUidNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


