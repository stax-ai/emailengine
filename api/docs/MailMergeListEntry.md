# MailMergeListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Message ID | [optional] 
**Params** | Pointer to **map[string]interface{}** | An object of variables for the template renderer | [optional] 
**SendAt** | Pointer to **time.Time** | Send message at specified time. Overrides message level &#x60;sendAt&#x60; value. | [optional] 
**To** | [**ToAddress**](ToAddress.md) |  | 

## Methods

### NewMailMergeListEntry

`func NewMailMergeListEntry(to ToAddress, ) *MailMergeListEntry`

NewMailMergeListEntry instantiates a new MailMergeListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailMergeListEntryWithDefaults

`func NewMailMergeListEntryWithDefaults() *MailMergeListEntry`

NewMailMergeListEntryWithDefaults instantiates a new MailMergeListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MailMergeListEntry) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MailMergeListEntry) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MailMergeListEntry) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MailMergeListEntry) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetParams

`func (o *MailMergeListEntry) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MailMergeListEntry) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MailMergeListEntry) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *MailMergeListEntry) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSendAt

`func (o *MailMergeListEntry) GetSendAt() time.Time`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *MailMergeListEntry) GetSendAtOk() (*time.Time, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *MailMergeListEntry) SetSendAt(v time.Time)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *MailMergeListEntry) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetTo

`func (o *MailMergeListEntry) GetTo() ToAddress`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MailMergeListEntry) GetToOk() (*ToAddress, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MailMergeListEntry) SetTo(v ToAddress)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


