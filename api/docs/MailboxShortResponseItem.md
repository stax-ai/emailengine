# MailboxShortResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delimiter** | Pointer to **string** | Hierarchy delimiter character used in paths | [optional] 
**Listed** | Pointer to **bool** | Whether this mailbox appears in LIST command results | [optional] 
**Name** | **string** | Display name of the mailbox | 
**ParentPath** | **string** | Path to the parent mailbox | 
**Path** | **string** | Full path to the mailbox | 
**SpecialUse** | Pointer to [**MailboxSpecialUse**](MailboxSpecialUse.md) |  | [optional] 
**Subscribed** | Pointer to **bool** | Whether the user is subscribed to this mailbox | [optional] 

## Methods

### NewMailboxShortResponseItem

`func NewMailboxShortResponseItem(name string, parentPath string, path string, ) *MailboxShortResponseItem`

NewMailboxShortResponseItem instantiates a new MailboxShortResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailboxShortResponseItemWithDefaults

`func NewMailboxShortResponseItemWithDefaults() *MailboxShortResponseItem`

NewMailboxShortResponseItemWithDefaults instantiates a new MailboxShortResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelimiter

`func (o *MailboxShortResponseItem) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *MailboxShortResponseItem) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *MailboxShortResponseItem) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *MailboxShortResponseItem) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetListed

`func (o *MailboxShortResponseItem) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *MailboxShortResponseItem) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *MailboxShortResponseItem) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *MailboxShortResponseItem) HasListed() bool`

HasListed returns a boolean if a field has been set.

### GetName

`func (o *MailboxShortResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MailboxShortResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MailboxShortResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetParentPath

`func (o *MailboxShortResponseItem) GetParentPath() string`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *MailboxShortResponseItem) GetParentPathOk() (*string, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *MailboxShortResponseItem) SetParentPath(v string)`

SetParentPath sets ParentPath field to given value.


### GetPath

`func (o *MailboxShortResponseItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MailboxShortResponseItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MailboxShortResponseItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetSpecialUse

`func (o *MailboxShortResponseItem) GetSpecialUse() MailboxSpecialUse`

GetSpecialUse returns the SpecialUse field if non-nil, zero value otherwise.

### GetSpecialUseOk

`func (o *MailboxShortResponseItem) GetSpecialUseOk() (*MailboxSpecialUse, bool)`

GetSpecialUseOk returns a tuple with the SpecialUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialUse

`func (o *MailboxShortResponseItem) SetSpecialUse(v MailboxSpecialUse)`

SetSpecialUse sets SpecialUse field to given value.

### HasSpecialUse

`func (o *MailboxShortResponseItem) HasSpecialUse() bool`

HasSpecialUse returns a boolean if a field has been set.

### GetSubscribed

`func (o *MailboxShortResponseItem) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *MailboxShortResponseItem) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *MailboxShortResponseItem) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *MailboxShortResponseItem) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


