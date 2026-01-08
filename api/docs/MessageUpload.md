# MessageUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]UploadAttachment**](UploadAttachment.md) | List of attachments | [optional] 
**Bcc** | Pointer to [**[]Model5**](Model5.md) | List of addresses | [optional] 
**Cc** | Pointer to [**[]Model5**](Model5.md) | List of addresses | [optional] 
**Flags** | Pointer to **[]string** | Message flags | [optional] 
**From** | Pointer to [**FromAddress**](FromAddress.md) |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** | Custom Headers | [optional] 
**Html** | Pointer to **string** | Message HTML | [optional] 
**InternalDate** | Pointer to **time.Time** | Sets the internal date for this message | [optional] 
**Locale** | Pointer to **string** | Optional locale | [optional] 
**MessageId** | Pointer to **string** | Message ID | [optional] 
**Path** | **string** | Target mailbox folder path | 
**Raw** | Pointer to **string** | A Base64-encoded email message in RFC 822 format. If you provide other fields along with raw, those fields will override the corresponding values in the raw message. | [optional] 
**Reference** | Pointer to [**MessageReference**](MessageReference.md) |  | [optional] 
**Subject** | Pointer to **string** | Message subject | [optional] 
**Text** | Pointer to **string** | Message Text | [optional] 
**To** | Pointer to [**[]Model5**](Model5.md) | List of addresses | [optional] 
**Tz** | Pointer to **string** | Optional timezone | [optional] 

## Methods

### NewMessageUpload

`func NewMessageUpload(path string, ) *MessageUpload`

NewMessageUpload instantiates a new MessageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUploadWithDefaults

`func NewMessageUploadWithDefaults() *MessageUpload`

NewMessageUploadWithDefaults instantiates a new MessageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *MessageUpload) GetAttachments() []UploadAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MessageUpload) GetAttachmentsOk() (*[]UploadAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MessageUpload) SetAttachments(v []UploadAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MessageUpload) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBcc

`func (o *MessageUpload) GetBcc() []Model5`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *MessageUpload) GetBccOk() (*[]Model5, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *MessageUpload) SetBcc(v []Model5)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *MessageUpload) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCc

`func (o *MessageUpload) GetCc() []Model5`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *MessageUpload) GetCcOk() (*[]Model5, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *MessageUpload) SetCc(v []Model5)`

SetCc sets Cc field to given value.

### HasCc

`func (o *MessageUpload) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetFlags

`func (o *MessageUpload) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageUpload) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageUpload) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageUpload) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetFrom

`func (o *MessageUpload) GetFrom() FromAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageUpload) GetFromOk() (*FromAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageUpload) SetFrom(v FromAddress)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageUpload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHeaders

`func (o *MessageUpload) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MessageUpload) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MessageUpload) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MessageUpload) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHtml

`func (o *MessageUpload) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *MessageUpload) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *MessageUpload) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *MessageUpload) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetInternalDate

`func (o *MessageUpload) GetInternalDate() time.Time`

GetInternalDate returns the InternalDate field if non-nil, zero value otherwise.

### GetInternalDateOk

`func (o *MessageUpload) GetInternalDateOk() (*time.Time, bool)`

GetInternalDateOk returns a tuple with the InternalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDate

`func (o *MessageUpload) SetInternalDate(v time.Time)`

SetInternalDate sets InternalDate field to given value.

### HasInternalDate

`func (o *MessageUpload) HasInternalDate() bool`

HasInternalDate returns a boolean if a field has been set.

### GetLocale

`func (o *MessageUpload) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MessageUpload) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MessageUpload) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MessageUpload) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMessageId

`func (o *MessageUpload) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageUpload) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageUpload) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageUpload) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetPath

`func (o *MessageUpload) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MessageUpload) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MessageUpload) SetPath(v string)`

SetPath sets Path field to given value.


### GetRaw

`func (o *MessageUpload) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *MessageUpload) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *MessageUpload) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *MessageUpload) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetReference

`func (o *MessageUpload) GetReference() MessageReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MessageUpload) GetReferenceOk() (*MessageReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MessageUpload) SetReference(v MessageReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MessageUpload) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSubject

`func (o *MessageUpload) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MessageUpload) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MessageUpload) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MessageUpload) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *MessageUpload) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageUpload) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageUpload) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MessageUpload) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *MessageUpload) GetTo() []Model5`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageUpload) GetToOk() (*[]Model5, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageUpload) SetTo(v []Model5)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageUpload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTz

`func (o *MessageUpload) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *MessageUpload) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *MessageUpload) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *MessageUpload) HasTz() bool`

HasTz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


