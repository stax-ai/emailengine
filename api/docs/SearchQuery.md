# SearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answered** | Pointer to **bool** | Filter for messages that have been replied to. Only supported for IMAP. | [optional] 
**Bcc** | Pointer to **string** | Search in Bcc addresses. Not supported for MS Graph API. | [optional] 
**Before** | Pointer to **string** | Find messages received before this date | [optional] 
**Body** | Pointer to **string** | Search in message body content | [optional] 
**Cc** | Pointer to **string** | Search in Cc addresses. Not supported for MS Graph API. | [optional] 
**Deleted** | Pointer to **bool** | Filter for messages marked for deletion. Only supported for IMAP. | [optional] 
**Draft** | Pointer to **bool** | Filter for draft messages | [optional] 
**EmailId** | Pointer to **string** | Globally unique email identifier (when supported by the email server) | [optional] 
**EmailIds** | Pointer to **[]string** | List of specific email IDs to fetch. When provided, other search criteria are ignored. Useful for bulk operations on known messages. | [optional] 
**Flagged** | Pointer to **bool** | Filter for flagged/starred messages | [optional] 
**From** | Pointer to **string** | Search in From addresses | [optional] 
**GmailRaw** | Pointer to **string** | Gmail search syntax (only works with Gmail accounts) | [optional] 
**Header** | Pointer to **map[string]interface{}** | Search specific email headers | [optional] 
**Larger** | Pointer to **int32** | Find messages larger than specified size in bytes. Not supported for MS Graph API. | [optional] 
**Modseq** | Pointer to **int32** | Find messages with modification sequence higher than specified value. Only supported for IMAP with CONDSTORE. | [optional] 
**Seen** | Pointer to **bool** | Filter for read messages | [optional] 
**SentBefore** | Pointer to **string** | Find messages sent before this date | [optional] 
**SentSince** | Pointer to **string** | Find messages sent after this date | [optional] 
**Seq** | Pointer to **string** | Sequence number range (e.g., \&quot;1:10\&quot; or \&quot;1,3,5\&quot;). Only supported for IMAP. | [optional] 
**Since** | Pointer to **string** | Find messages received after this date | [optional] 
**Smaller** | Pointer to **int32** | Find messages smaller than specified size in bytes. Not supported for MS Graph API. | [optional] 
**Subject** | Pointer to **string** | Search in message subject | [optional] 
**ThreadId** | Pointer to **string** | Thread identifier for email conversations (when supported by the email server) | [optional] 
**To** | Pointer to **string** | Search in To addresses. Not supported for MS Graph API. | [optional] 
**Uid** | Pointer to **string** | UID range (e.g., \&quot;100:200\&quot; or \&quot;150,200,250\&quot;). Only supported for IMAP. | [optional] 
**Unseen** | Pointer to **bool** | Filter for unread messages | [optional] 

## Methods

### NewSearchQuery

`func NewSearchQuery() *SearchQuery`

NewSearchQuery instantiates a new SearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryWithDefaults

`func NewSearchQueryWithDefaults() *SearchQuery`

NewSearchQueryWithDefaults instantiates a new SearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswered

`func (o *SearchQuery) GetAnswered() bool`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *SearchQuery) GetAnsweredOk() (*bool, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *SearchQuery) SetAnswered(v bool)`

SetAnswered sets Answered field to given value.

### HasAnswered

`func (o *SearchQuery) HasAnswered() bool`

HasAnswered returns a boolean if a field has been set.

### GetBcc

`func (o *SearchQuery) GetBcc() string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *SearchQuery) GetBccOk() (*string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *SearchQuery) SetBcc(v string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *SearchQuery) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetBefore

`func (o *SearchQuery) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *SearchQuery) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *SearchQuery) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *SearchQuery) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetBody

`func (o *SearchQuery) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SearchQuery) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SearchQuery) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SearchQuery) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCc

`func (o *SearchQuery) GetCc() string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *SearchQuery) GetCcOk() (*string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *SearchQuery) SetCc(v string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *SearchQuery) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetDeleted

`func (o *SearchQuery) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SearchQuery) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SearchQuery) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *SearchQuery) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDraft

`func (o *SearchQuery) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *SearchQuery) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *SearchQuery) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *SearchQuery) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetEmailId

`func (o *SearchQuery) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *SearchQuery) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *SearchQuery) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.

### HasEmailId

`func (o *SearchQuery) HasEmailId() bool`

HasEmailId returns a boolean if a field has been set.

### GetEmailIds

`func (o *SearchQuery) GetEmailIds() []string`

GetEmailIds returns the EmailIds field if non-nil, zero value otherwise.

### GetEmailIdsOk

`func (o *SearchQuery) GetEmailIdsOk() (*[]string, bool)`

GetEmailIdsOk returns a tuple with the EmailIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIds

`func (o *SearchQuery) SetEmailIds(v []string)`

SetEmailIds sets EmailIds field to given value.

### HasEmailIds

`func (o *SearchQuery) HasEmailIds() bool`

HasEmailIds returns a boolean if a field has been set.

### GetFlagged

`func (o *SearchQuery) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *SearchQuery) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *SearchQuery) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *SearchQuery) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetFrom

`func (o *SearchQuery) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchQuery) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchQuery) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchQuery) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGmailRaw

`func (o *SearchQuery) GetGmailRaw() string`

GetGmailRaw returns the GmailRaw field if non-nil, zero value otherwise.

### GetGmailRawOk

`func (o *SearchQuery) GetGmailRawOk() (*string, bool)`

GetGmailRawOk returns a tuple with the GmailRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmailRaw

`func (o *SearchQuery) SetGmailRaw(v string)`

SetGmailRaw sets GmailRaw field to given value.

### HasGmailRaw

`func (o *SearchQuery) HasGmailRaw() bool`

HasGmailRaw returns a boolean if a field has been set.

### GetHeader

`func (o *SearchQuery) GetHeader() map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SearchQuery) GetHeaderOk() (*map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SearchQuery) SetHeader(v map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SearchQuery) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetLarger

`func (o *SearchQuery) GetLarger() int32`

GetLarger returns the Larger field if non-nil, zero value otherwise.

### GetLargerOk

`func (o *SearchQuery) GetLargerOk() (*int32, bool)`

GetLargerOk returns a tuple with the Larger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLarger

`func (o *SearchQuery) SetLarger(v int32)`

SetLarger sets Larger field to given value.

### HasLarger

`func (o *SearchQuery) HasLarger() bool`

HasLarger returns a boolean if a field has been set.

### GetModseq

`func (o *SearchQuery) GetModseq() int32`

GetModseq returns the Modseq field if non-nil, zero value otherwise.

### GetModseqOk

`func (o *SearchQuery) GetModseqOk() (*int32, bool)`

GetModseqOk returns a tuple with the Modseq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModseq

`func (o *SearchQuery) SetModseq(v int32)`

SetModseq sets Modseq field to given value.

### HasModseq

`func (o *SearchQuery) HasModseq() bool`

HasModseq returns a boolean if a field has been set.

### GetSeen

`func (o *SearchQuery) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *SearchQuery) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *SearchQuery) SetSeen(v bool)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *SearchQuery) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetSentBefore

`func (o *SearchQuery) GetSentBefore() string`

GetSentBefore returns the SentBefore field if non-nil, zero value otherwise.

### GetSentBeforeOk

`func (o *SearchQuery) GetSentBeforeOk() (*string, bool)`

GetSentBeforeOk returns a tuple with the SentBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentBefore

`func (o *SearchQuery) SetSentBefore(v string)`

SetSentBefore sets SentBefore field to given value.

### HasSentBefore

`func (o *SearchQuery) HasSentBefore() bool`

HasSentBefore returns a boolean if a field has been set.

### GetSentSince

`func (o *SearchQuery) GetSentSince() string`

GetSentSince returns the SentSince field if non-nil, zero value otherwise.

### GetSentSinceOk

`func (o *SearchQuery) GetSentSinceOk() (*string, bool)`

GetSentSinceOk returns a tuple with the SentSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentSince

`func (o *SearchQuery) SetSentSince(v string)`

SetSentSince sets SentSince field to given value.

### HasSentSince

`func (o *SearchQuery) HasSentSince() bool`

HasSentSince returns a boolean if a field has been set.

### GetSeq

`func (o *SearchQuery) GetSeq() string`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *SearchQuery) GetSeqOk() (*string, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *SearchQuery) SetSeq(v string)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *SearchQuery) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSince

`func (o *SearchQuery) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *SearchQuery) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *SearchQuery) SetSince(v string)`

SetSince sets Since field to given value.

### HasSince

`func (o *SearchQuery) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetSmaller

`func (o *SearchQuery) GetSmaller() int32`

GetSmaller returns the Smaller field if non-nil, zero value otherwise.

### GetSmallerOk

`func (o *SearchQuery) GetSmallerOk() (*int32, bool)`

GetSmallerOk returns a tuple with the Smaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmaller

`func (o *SearchQuery) SetSmaller(v int32)`

SetSmaller sets Smaller field to given value.

### HasSmaller

`func (o *SearchQuery) HasSmaller() bool`

HasSmaller returns a boolean if a field has been set.

### GetSubject

`func (o *SearchQuery) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SearchQuery) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SearchQuery) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SearchQuery) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetThreadId

`func (o *SearchQuery) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *SearchQuery) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *SearchQuery) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *SearchQuery) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetTo

`func (o *SearchQuery) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SearchQuery) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SearchQuery) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SearchQuery) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUid

`func (o *SearchQuery) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SearchQuery) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SearchQuery) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SearchQuery) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUnseen

`func (o *SearchQuery) GetUnseen() bool`

GetUnseen returns the Unseen field if non-nil, zero value otherwise.

### GetUnseenOk

`func (o *SearchQuery) GetUnseenOk() (*bool, bool)`

GetUnseenOk returns a tuple with the Unseen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseen

`func (o *SearchQuery) SetUnseen(v bool)`

SetUnseen sets Unseen field to given value.

### HasUnseen

`func (o *SearchQuery) HasUnseen() bool`

HasUnseen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


