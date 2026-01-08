# SubmitMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]UploadAttachment**](UploadAttachment.md) | List of attachments | [optional] 
**BaseUrl** | Pointer to **string** | Optional base URL for trackers. This URL must point to your EmailEngine instance. | [optional] 
**Bcc** | Pointer to [**[]BccAddress**](BccAddress.md) | List of BCC addresses | [optional] 
**Cc** | Pointer to [**[]CcAddress**](CcAddress.md) | List of CC addresses | [optional] 
**Copy** | Pointer to **NullableBool** | If set then either copies the message to the Sent Mail folder or not. If not set then uses the account&#39;s default setting. | [optional] 
**DeliveryAttempts** | Pointer to **int32** | How many delivery attempts to make until message is considered as failed | [optional] 
**DryRun** | Pointer to **bool** | If true, then EmailEngine does not send the email and returns an RFC822 formatted email file. Tracking information is not added to the email. | [optional] [default to false]
**Dsn** | Pointer to [**DSN**](DSN.md) |  | [optional] 
**Envelope** | Pointer to [**SMTPEnvelope**](SMTPEnvelope.md) |  | [optional] 
**From** | Pointer to [**FromAddress**](FromAddress.md) |  | [optional] 
**Gateway** | Pointer to **string** | Optional SMTP gateway ID for message routing | [optional] 
**Headers** | Pointer to **map[string]interface{}** | Custom Headers | [optional] 
**Html** | Pointer to **string** | HTML message content | [optional] 
**ListId** | Pointer to **string** | List ID for Mail Merge. Must use a subdomain name format. Lists are registered ad-hoc, so a new identifier defines a new list. | [optional] 
**LocalAddress** | Pointer to **string** | Optional local IP address to bind to when connecting to the SMTP server | [optional] 
**Locale** | Pointer to **string** | Optional locale | [optional] 
**MailMerge** | Pointer to [**[]MailMergeListEntry**](MailMergeListEntry.md) | Mail merge options. A separate email is generated for each recipient. Using mail merge disables &#x60;messageId&#x60;, &#x60;envelope&#x60;, &#x60;to&#x60;, &#x60;cc&#x60;, &#x60;bcc&#x60;, &#x60;render&#x60; keys for the message root. | [optional] 
**MessageId** | Pointer to **string** | Message ID | [optional] 
**PreviewText** | Pointer to **string** | Preview text shown in email clients after the subject line | [optional] 
**Proxy** | Pointer to **string** | Optional proxy URL to use when connecting to the SMTP server | [optional] 
**Raw** | Pointer to **string** | A Base64-encoded email message in RFC 822 format. If you provide other fields along with raw, those fields will override the corresponding values in the raw message. | [optional] 
**Reference** | Pointer to [**MessageReference**](MessageReference.md) |  | [optional] 
**Render** | Pointer to [**TemplateRender**](TemplateRender.md) |  | [optional] 
**ReplyTo** | Pointer to [**[]ReplyToAddress**](ReplyToAddress.md) | List of Reply-To addresses | [optional] 
**SendAt** | Pointer to **time.Time** | Send message at specified time | [optional] 
**SentMailPath** | Pointer to **string** | Upload sent message to this folder. By default the account&#39;s Sent Mail folder is used. | [optional] 
**Subject** | Pointer to **string** | Email subject line | [optional] 
**Template** | Pointer to **string** | Stored template ID to load the email content from | [optional] 
**Text** | Pointer to **string** | Plain text message content | [optional] 
**To** | Pointer to [**[]ToAddress**](ToAddress.md) | List of recipient addresses | [optional] 
**TrackClicks** | Pointer to **bool** | Should EmailEngine track clicks for this message | [optional] 
**TrackOpens** | Pointer to **bool** | Should EmailEngine track opens for this message | [optional] 
**Tz** | Pointer to **string** | Optional timezone | [optional] 

## Methods

### NewSubmitMessage

`func NewSubmitMessage() *SubmitMessage`

NewSubmitMessage instantiates a new SubmitMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitMessageWithDefaults

`func NewSubmitMessageWithDefaults() *SubmitMessage`

NewSubmitMessageWithDefaults instantiates a new SubmitMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *SubmitMessage) GetAttachments() []UploadAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SubmitMessage) GetAttachmentsOk() (*[]UploadAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SubmitMessage) SetAttachments(v []UploadAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SubmitMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBaseUrl

`func (o *SubmitMessage) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *SubmitMessage) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *SubmitMessage) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *SubmitMessage) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetBcc

`func (o *SubmitMessage) GetBcc() []BccAddress`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *SubmitMessage) GetBccOk() (*[]BccAddress, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *SubmitMessage) SetBcc(v []BccAddress)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *SubmitMessage) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCc

`func (o *SubmitMessage) GetCc() []CcAddress`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *SubmitMessage) GetCcOk() (*[]CcAddress, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *SubmitMessage) SetCc(v []CcAddress)`

SetCc sets Cc field to given value.

### HasCc

`func (o *SubmitMessage) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetCopy

`func (o *SubmitMessage) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *SubmitMessage) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *SubmitMessage) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *SubmitMessage) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### SetCopyNil

`func (o *SubmitMessage) SetCopyNil(b bool)`

 SetCopyNil sets the value for Copy to be an explicit nil

### UnsetCopy
`func (o *SubmitMessage) UnsetCopy()`

UnsetCopy ensures that no value is present for Copy, not even an explicit nil
### GetDeliveryAttempts

`func (o *SubmitMessage) GetDeliveryAttempts() int32`

GetDeliveryAttempts returns the DeliveryAttempts field if non-nil, zero value otherwise.

### GetDeliveryAttemptsOk

`func (o *SubmitMessage) GetDeliveryAttemptsOk() (*int32, bool)`

GetDeliveryAttemptsOk returns a tuple with the DeliveryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAttempts

`func (o *SubmitMessage) SetDeliveryAttempts(v int32)`

SetDeliveryAttempts sets DeliveryAttempts field to given value.

### HasDeliveryAttempts

`func (o *SubmitMessage) HasDeliveryAttempts() bool`

HasDeliveryAttempts returns a boolean if a field has been set.

### GetDryRun

`func (o *SubmitMessage) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SubmitMessage) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SubmitMessage) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *SubmitMessage) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetDsn

`func (o *SubmitMessage) GetDsn() DSN`

GetDsn returns the Dsn field if non-nil, zero value otherwise.

### GetDsnOk

`func (o *SubmitMessage) GetDsnOk() (*DSN, bool)`

GetDsnOk returns a tuple with the Dsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsn

`func (o *SubmitMessage) SetDsn(v DSN)`

SetDsn sets Dsn field to given value.

### HasDsn

`func (o *SubmitMessage) HasDsn() bool`

HasDsn returns a boolean if a field has been set.

### GetEnvelope

`func (o *SubmitMessage) GetEnvelope() SMTPEnvelope`

GetEnvelope returns the Envelope field if non-nil, zero value otherwise.

### GetEnvelopeOk

`func (o *SubmitMessage) GetEnvelopeOk() (*SMTPEnvelope, bool)`

GetEnvelopeOk returns a tuple with the Envelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvelope

`func (o *SubmitMessage) SetEnvelope(v SMTPEnvelope)`

SetEnvelope sets Envelope field to given value.

### HasEnvelope

`func (o *SubmitMessage) HasEnvelope() bool`

HasEnvelope returns a boolean if a field has been set.

### GetFrom

`func (o *SubmitMessage) GetFrom() FromAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SubmitMessage) GetFromOk() (*FromAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SubmitMessage) SetFrom(v FromAddress)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SubmitMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGateway

`func (o *SubmitMessage) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SubmitMessage) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SubmitMessage) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *SubmitMessage) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHeaders

`func (o *SubmitMessage) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SubmitMessage) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SubmitMessage) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SubmitMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHtml

`func (o *SubmitMessage) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *SubmitMessage) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *SubmitMessage) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *SubmitMessage) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetListId

`func (o *SubmitMessage) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *SubmitMessage) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *SubmitMessage) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *SubmitMessage) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetLocalAddress

`func (o *SubmitMessage) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *SubmitMessage) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *SubmitMessage) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *SubmitMessage) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocale

`func (o *SubmitMessage) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SubmitMessage) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SubmitMessage) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SubmitMessage) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMailMerge

`func (o *SubmitMessage) GetMailMerge() []MailMergeListEntry`

GetMailMerge returns the MailMerge field if non-nil, zero value otherwise.

### GetMailMergeOk

`func (o *SubmitMessage) GetMailMergeOk() (*[]MailMergeListEntry, bool)`

GetMailMergeOk returns a tuple with the MailMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailMerge

`func (o *SubmitMessage) SetMailMerge(v []MailMergeListEntry)`

SetMailMerge sets MailMerge field to given value.

### HasMailMerge

`func (o *SubmitMessage) HasMailMerge() bool`

HasMailMerge returns a boolean if a field has been set.

### GetMessageId

`func (o *SubmitMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SubmitMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SubmitMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SubmitMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetPreviewText

`func (o *SubmitMessage) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *SubmitMessage) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *SubmitMessage) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *SubmitMessage) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### GetProxy

`func (o *SubmitMessage) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SubmitMessage) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SubmitMessage) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SubmitMessage) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRaw

`func (o *SubmitMessage) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *SubmitMessage) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *SubmitMessage) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *SubmitMessage) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetReference

`func (o *SubmitMessage) GetReference() MessageReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SubmitMessage) GetReferenceOk() (*MessageReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SubmitMessage) SetReference(v MessageReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SubmitMessage) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetRender

`func (o *SubmitMessage) GetRender() TemplateRender`

GetRender returns the Render field if non-nil, zero value otherwise.

### GetRenderOk

`func (o *SubmitMessage) GetRenderOk() (*TemplateRender, bool)`

GetRenderOk returns a tuple with the Render field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRender

`func (o *SubmitMessage) SetRender(v TemplateRender)`

SetRender sets Render field to given value.

### HasRender

`func (o *SubmitMessage) HasRender() bool`

HasRender returns a boolean if a field has been set.

### GetReplyTo

`func (o *SubmitMessage) GetReplyTo() []ReplyToAddress`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *SubmitMessage) GetReplyToOk() (*[]ReplyToAddress, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *SubmitMessage) SetReplyTo(v []ReplyToAddress)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *SubmitMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSendAt

`func (o *SubmitMessage) GetSendAt() time.Time`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *SubmitMessage) GetSendAtOk() (*time.Time, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *SubmitMessage) SetSendAt(v time.Time)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *SubmitMessage) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### GetSentMailPath

`func (o *SubmitMessage) GetSentMailPath() string`

GetSentMailPath returns the SentMailPath field if non-nil, zero value otherwise.

### GetSentMailPathOk

`func (o *SubmitMessage) GetSentMailPathOk() (*string, bool)`

GetSentMailPathOk returns a tuple with the SentMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentMailPath

`func (o *SubmitMessage) SetSentMailPath(v string)`

SetSentMailPath sets SentMailPath field to given value.

### HasSentMailPath

`func (o *SubmitMessage) HasSentMailPath() bool`

HasSentMailPath returns a boolean if a field has been set.

### GetSubject

`func (o *SubmitMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SubmitMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SubmitMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SubmitMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplate

`func (o *SubmitMessage) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SubmitMessage) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SubmitMessage) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SubmitMessage) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetText

`func (o *SubmitMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SubmitMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SubmitMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SubmitMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *SubmitMessage) GetTo() []ToAddress`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SubmitMessage) GetToOk() (*[]ToAddress, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SubmitMessage) SetTo(v []ToAddress)`

SetTo sets To field to given value.

### HasTo

`func (o *SubmitMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTrackClicks

`func (o *SubmitMessage) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *SubmitMessage) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *SubmitMessage) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *SubmitMessage) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetTrackOpens

`func (o *SubmitMessage) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *SubmitMessage) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *SubmitMessage) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *SubmitMessage) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTz

`func (o *SubmitMessage) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *SubmitMessage) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *SubmitMessage) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *SubmitMessage) HasTz() bool`

HasTz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


