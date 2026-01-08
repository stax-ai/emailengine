# MessageListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]AttachmentEntry**](AttachmentEntry.md) | List of attachments | [optional] 
**Bcc** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Cc** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Date** | Pointer to **time.Time** | Date when the message was received by the mail server | [optional] 
**Draft** | Pointer to **bool** | Whether this message is a draft | [optional] 
**EmailId** | Pointer to **string** | Globally unique email identifier (when supported by the email server) | [optional] 
**Flagged** | Pointer to **bool** | Whether this message is flagged/starred | [optional] 
**Flags** | Pointer to **[]string** | IMAP flags set on this message | [optional] 
**From** | Pointer to [**FromAddress**](FromAddress.md) |  | [optional] 
**Id** | Pointer to **string** | EmailEngine message identifier | [optional] 
**InReplyTo** | Pointer to **string** | Message-ID of the message this is replying to | [optional] 
**Labels** | Pointer to **[]string** | Gmail labels applied to this message | [optional] 
**MessageId** | Pointer to **string** | Message-ID header value | [optional] 
**Preview** | Pointer to **string** | Short preview of the message content | [optional] 
**ReplyTo** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Size** | Pointer to **int32** | Message size in bytes | [optional] 
**Subject** | Pointer to **string** | Message subject line | [optional] 
**Text** | Pointer to [**TextInfo**](TextInfo.md) |  | [optional] 
**ThreadId** | Pointer to **string** | Thread identifier for email conversations (when supported by the email server) | [optional] 
**To** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Uid** | Pointer to **int32** | IMAP UID (unique identifier within the mailbox) | [optional] 
**Unseen** | Pointer to **bool** | Whether this message is unread | [optional] 

## Methods

### NewMessageListEntry

`func NewMessageListEntry() *MessageListEntry`

NewMessageListEntry instantiates a new MessageListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageListEntryWithDefaults

`func NewMessageListEntryWithDefaults() *MessageListEntry`

NewMessageListEntryWithDefaults instantiates a new MessageListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *MessageListEntry) GetAttachments() []AttachmentEntry`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MessageListEntry) GetAttachmentsOk() (*[]AttachmentEntry, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MessageListEntry) SetAttachments(v []AttachmentEntry)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MessageListEntry) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBcc

`func (o *MessageListEntry) GetBcc() []RcptAddressEntry`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *MessageListEntry) GetBccOk() (*[]RcptAddressEntry, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *MessageListEntry) SetBcc(v []RcptAddressEntry)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *MessageListEntry) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCc

`func (o *MessageListEntry) GetCc() []RcptAddressEntry`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *MessageListEntry) GetCcOk() (*[]RcptAddressEntry, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *MessageListEntry) SetCc(v []RcptAddressEntry)`

SetCc sets Cc field to given value.

### HasCc

`func (o *MessageListEntry) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetDate

`func (o *MessageListEntry) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageListEntry) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageListEntry) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *MessageListEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDraft

`func (o *MessageListEntry) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *MessageListEntry) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *MessageListEntry) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *MessageListEntry) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetEmailId

`func (o *MessageListEntry) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *MessageListEntry) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *MessageListEntry) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.

### HasEmailId

`func (o *MessageListEntry) HasEmailId() bool`

HasEmailId returns a boolean if a field has been set.

### GetFlagged

`func (o *MessageListEntry) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *MessageListEntry) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *MessageListEntry) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *MessageListEntry) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetFlags

`func (o *MessageListEntry) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageListEntry) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageListEntry) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageListEntry) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetFrom

`func (o *MessageListEntry) GetFrom() FromAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageListEntry) GetFromOk() (*FromAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageListEntry) SetFrom(v FromAddress)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageListEntry) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *MessageListEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageListEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageListEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageListEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInReplyTo

`func (o *MessageListEntry) GetInReplyTo() string`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *MessageListEntry) GetInReplyToOk() (*string, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *MessageListEntry) SetInReplyTo(v string)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *MessageListEntry) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### GetLabels

`func (o *MessageListEntry) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MessageListEntry) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MessageListEntry) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MessageListEntry) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMessageId

`func (o *MessageListEntry) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageListEntry) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageListEntry) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageListEntry) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetPreview

`func (o *MessageListEntry) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *MessageListEntry) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *MessageListEntry) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *MessageListEntry) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetReplyTo

`func (o *MessageListEntry) GetReplyTo() []RcptAddressEntry`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MessageListEntry) GetReplyToOk() (*[]RcptAddressEntry, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MessageListEntry) SetReplyTo(v []RcptAddressEntry)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MessageListEntry) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSize

`func (o *MessageListEntry) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MessageListEntry) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MessageListEntry) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MessageListEntry) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSubject

`func (o *MessageListEntry) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MessageListEntry) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MessageListEntry) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MessageListEntry) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *MessageListEntry) GetText() TextInfo`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageListEntry) GetTextOk() (*TextInfo, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageListEntry) SetText(v TextInfo)`

SetText sets Text field to given value.

### HasText

`func (o *MessageListEntry) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThreadId

`func (o *MessageListEntry) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *MessageListEntry) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *MessageListEntry) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *MessageListEntry) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetTo

`func (o *MessageListEntry) GetTo() []RcptAddressEntry`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageListEntry) GetToOk() (*[]RcptAddressEntry, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageListEntry) SetTo(v []RcptAddressEntry)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageListEntry) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUid

`func (o *MessageListEntry) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MessageListEntry) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MessageListEntry) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MessageListEntry) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUnseen

`func (o *MessageListEntry) GetUnseen() bool`

GetUnseen returns the Unseen field if non-nil, zero value otherwise.

### GetUnseenOk

`func (o *MessageListEntry) GetUnseenOk() (*bool, bool)`

GetUnseenOk returns a tuple with the Unseen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseen

`func (o *MessageListEntry) SetUnseen(v bool)`

SetUnseen sets Unseen field to given value.

### HasUnseen

`func (o *MessageListEntry) HasUnseen() bool`

HasUnseen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


