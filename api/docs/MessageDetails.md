# MessageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]AttachmentEntry**](AttachmentEntry.md) | List of attachments | [optional] 
**Bcc** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Bounces** | Pointer to [**[]BounceEntry**](BounceEntry.md) |  | [optional] 
**Cc** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Date** | Pointer to **time.Time** | Date when the message was received by the mail server | [optional] 
**Draft** | Pointer to **bool** | Whether this message is a draft | [optional] 
**EmailId** | Pointer to **string** | Globally unique email identifier (when supported by the email server) | [optional] 
**Flagged** | Pointer to **bool** | Whether this message is flagged/starred | [optional] 
**Flags** | Pointer to **[]string** | IMAP flags set on this message | [optional] 
**From** | Pointer to [**FromAddress**](FromAddress.md) |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** | Raw email headers as key-value pairs (arrays contain multiple values for headers that appear multiple times). Not available for MS Graph API. | [optional] 
**Id** | Pointer to **string** | EmailEngine message identifier | [optional] 
**InReplyTo** | Pointer to **string** | Message-ID of the message this is replying to | [optional] 
**IsAutoReply** | Pointer to **bool** | Whether this message appears to be an automatic reply (out of office, vacation responder, etc.) | [optional] 
**Labels** | Pointer to **[]string** | Gmail labels applied to this message | [optional] 
**MessageId** | Pointer to **string** | Message-ID header value | [optional] 
**MessageSpecialUse** | Pointer to [**MessageSpecialUse**](MessageSpecialUse.md) |  | [optional] 
**ReplyTo** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Sender** | Pointer to [**FromAddress**](FromAddress.md) |  | [optional] 
**Size** | Pointer to **int32** | Message size in bytes | [optional] 
**SpecialUse** | Pointer to [**Model3**](Model3.md) |  | [optional] 
**Subject** | Pointer to **string** | Message subject line | [optional] 
**Text** | Pointer to [**TextInfoDetails**](TextInfoDetails.md) |  | [optional] 
**ThreadId** | Pointer to **string** | Thread identifier for email conversations (when supported by the email server) | [optional] 
**To** | Pointer to [**[]RcptAddressEntry**](RcptAddressEntry.md) | List of email addresses | [optional] 
**Uid** | Pointer to **int32** | IMAP UID (unique identifier within the mailbox) | [optional] 
**Unseen** | Pointer to **bool** | Whether this message is unread | [optional] 

## Methods

### NewMessageDetails

`func NewMessageDetails() *MessageDetails`

NewMessageDetails instantiates a new MessageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDetailsWithDefaults

`func NewMessageDetailsWithDefaults() *MessageDetails`

NewMessageDetailsWithDefaults instantiates a new MessageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *MessageDetails) GetAttachments() []AttachmentEntry`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MessageDetails) GetAttachmentsOk() (*[]AttachmentEntry, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MessageDetails) SetAttachments(v []AttachmentEntry)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MessageDetails) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBcc

`func (o *MessageDetails) GetBcc() []RcptAddressEntry`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *MessageDetails) GetBccOk() (*[]RcptAddressEntry, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *MessageDetails) SetBcc(v []RcptAddressEntry)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *MessageDetails) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetBounces

`func (o *MessageDetails) GetBounces() []BounceEntry`

GetBounces returns the Bounces field if non-nil, zero value otherwise.

### GetBouncesOk

`func (o *MessageDetails) GetBouncesOk() (*[]BounceEntry, bool)`

GetBouncesOk returns a tuple with the Bounces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounces

`func (o *MessageDetails) SetBounces(v []BounceEntry)`

SetBounces sets Bounces field to given value.

### HasBounces

`func (o *MessageDetails) HasBounces() bool`

HasBounces returns a boolean if a field has been set.

### GetCc

`func (o *MessageDetails) GetCc() []RcptAddressEntry`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *MessageDetails) GetCcOk() (*[]RcptAddressEntry, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *MessageDetails) SetCc(v []RcptAddressEntry)`

SetCc sets Cc field to given value.

### HasCc

`func (o *MessageDetails) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetDate

`func (o *MessageDetails) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageDetails) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageDetails) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *MessageDetails) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDraft

`func (o *MessageDetails) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *MessageDetails) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *MessageDetails) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *MessageDetails) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetEmailId

`func (o *MessageDetails) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *MessageDetails) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *MessageDetails) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.

### HasEmailId

`func (o *MessageDetails) HasEmailId() bool`

HasEmailId returns a boolean if a field has been set.

### GetFlagged

`func (o *MessageDetails) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *MessageDetails) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *MessageDetails) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *MessageDetails) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetFlags

`func (o *MessageDetails) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageDetails) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageDetails) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageDetails) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetFrom

`func (o *MessageDetails) GetFrom() FromAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MessageDetails) GetFromOk() (*FromAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MessageDetails) SetFrom(v FromAddress)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MessageDetails) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHeaders

`func (o *MessageDetails) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MessageDetails) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MessageDetails) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MessageDetails) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetId

`func (o *MessageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInReplyTo

`func (o *MessageDetails) GetInReplyTo() string`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *MessageDetails) GetInReplyToOk() (*string, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *MessageDetails) SetInReplyTo(v string)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *MessageDetails) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### GetIsAutoReply

`func (o *MessageDetails) GetIsAutoReply() bool`

GetIsAutoReply returns the IsAutoReply field if non-nil, zero value otherwise.

### GetIsAutoReplyOk

`func (o *MessageDetails) GetIsAutoReplyOk() (*bool, bool)`

GetIsAutoReplyOk returns a tuple with the IsAutoReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoReply

`func (o *MessageDetails) SetIsAutoReply(v bool)`

SetIsAutoReply sets IsAutoReply field to given value.

### HasIsAutoReply

`func (o *MessageDetails) HasIsAutoReply() bool`

HasIsAutoReply returns a boolean if a field has been set.

### GetLabels

`func (o *MessageDetails) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MessageDetails) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MessageDetails) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MessageDetails) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMessageId

`func (o *MessageDetails) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageDetails) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageDetails) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *MessageDetails) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageSpecialUse

`func (o *MessageDetails) GetMessageSpecialUse() MessageSpecialUse`

GetMessageSpecialUse returns the MessageSpecialUse field if non-nil, zero value otherwise.

### GetMessageSpecialUseOk

`func (o *MessageDetails) GetMessageSpecialUseOk() (*MessageSpecialUse, bool)`

GetMessageSpecialUseOk returns a tuple with the MessageSpecialUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSpecialUse

`func (o *MessageDetails) SetMessageSpecialUse(v MessageSpecialUse)`

SetMessageSpecialUse sets MessageSpecialUse field to given value.

### HasMessageSpecialUse

`func (o *MessageDetails) HasMessageSpecialUse() bool`

HasMessageSpecialUse returns a boolean if a field has been set.

### GetReplyTo

`func (o *MessageDetails) GetReplyTo() []RcptAddressEntry`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MessageDetails) GetReplyToOk() (*[]RcptAddressEntry, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MessageDetails) SetReplyTo(v []RcptAddressEntry)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MessageDetails) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSender

`func (o *MessageDetails) GetSender() FromAddress`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MessageDetails) GetSenderOk() (*FromAddress, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MessageDetails) SetSender(v FromAddress)`

SetSender sets Sender field to given value.

### HasSender

`func (o *MessageDetails) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSize

`func (o *MessageDetails) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MessageDetails) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MessageDetails) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MessageDetails) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSpecialUse

`func (o *MessageDetails) GetSpecialUse() Model3`

GetSpecialUse returns the SpecialUse field if non-nil, zero value otherwise.

### GetSpecialUseOk

`func (o *MessageDetails) GetSpecialUseOk() (*Model3, bool)`

GetSpecialUseOk returns a tuple with the SpecialUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialUse

`func (o *MessageDetails) SetSpecialUse(v Model3)`

SetSpecialUse sets SpecialUse field to given value.

### HasSpecialUse

`func (o *MessageDetails) HasSpecialUse() bool`

HasSpecialUse returns a boolean if a field has been set.

### GetSubject

`func (o *MessageDetails) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MessageDetails) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MessageDetails) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MessageDetails) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *MessageDetails) GetText() TextInfoDetails`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageDetails) GetTextOk() (*TextInfoDetails, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageDetails) SetText(v TextInfoDetails)`

SetText sets Text field to given value.

### HasText

`func (o *MessageDetails) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThreadId

`func (o *MessageDetails) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *MessageDetails) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *MessageDetails) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *MessageDetails) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetTo

`func (o *MessageDetails) GetTo() []RcptAddressEntry`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageDetails) GetToOk() (*[]RcptAddressEntry, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageDetails) SetTo(v []RcptAddressEntry)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageDetails) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUid

`func (o *MessageDetails) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MessageDetails) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MessageDetails) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MessageDetails) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUnseen

`func (o *MessageDetails) GetUnseen() bool`

GetUnseen returns the Unseen field if non-nil, zero value otherwise.

### GetUnseenOk

`func (o *MessageDetails) GetUnseenOk() (*bool, bool)`

GetUnseenOk returns a tuple with the Unseen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseen

`func (o *MessageDetails) SetUnseen(v bool)`

SetUnseen sets Unseen field to given value.

### HasUnseen

`func (o *MessageDetails) HasUnseen() bool`

HasUnseen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


