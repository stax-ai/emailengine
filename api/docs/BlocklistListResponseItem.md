# BlocklistListResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Created** | **time.Time** | The time this entry was added or updated | 
**MessageId** | Pointer to **string** | Message ID | [optional] 
**Reason** | Pointer to **string** | Why this entry was added | [optional] 
**Recipient** | **string** | Listed email address | 
**RemoteAddress** | Pointer to **string** | Which IP address triggered the entry | [optional] 
**Source** | Pointer to **string** | Which mechanism was used to add the entry | [optional] 
**UserAgent** | Pointer to **string** | Which user agent triggered the entry | [optional] 

## Methods

### NewBlocklistListResponseItem

`func NewBlocklistListResponseItem(account string, created time.Time, recipient string, ) *BlocklistListResponseItem`

NewBlocklistListResponseItem instantiates a new BlocklistListResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistListResponseItemWithDefaults

`func NewBlocklistListResponseItemWithDefaults() *BlocklistListResponseItem`

NewBlocklistListResponseItemWithDefaults instantiates a new BlocklistListResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BlocklistListResponseItem) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BlocklistListResponseItem) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BlocklistListResponseItem) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetCreated

`func (o *BlocklistListResponseItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BlocklistListResponseItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BlocklistListResponseItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetMessageId

`func (o *BlocklistListResponseItem) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *BlocklistListResponseItem) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *BlocklistListResponseItem) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *BlocklistListResponseItem) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetReason

`func (o *BlocklistListResponseItem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BlocklistListResponseItem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BlocklistListResponseItem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BlocklistListResponseItem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRecipient

`func (o *BlocklistListResponseItem) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *BlocklistListResponseItem) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *BlocklistListResponseItem) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetRemoteAddress

`func (o *BlocklistListResponseItem) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *BlocklistListResponseItem) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *BlocklistListResponseItem) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *BlocklistListResponseItem) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetSource

`func (o *BlocklistListResponseItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BlocklistListResponseItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BlocklistListResponseItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BlocklistListResponseItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUserAgent

`func (o *BlocklistListResponseItem) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *BlocklistListResponseItem) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *BlocklistListResponseItem) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *BlocklistListResponseItem) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


