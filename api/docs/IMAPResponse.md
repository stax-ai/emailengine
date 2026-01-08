# IMAPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveMailPath** | Pointer to **NullableString** | Custom folder path for archived messages. Defaults to auto-detected \&quot;Archive\&quot; folder. Set to null to use default. | [optional] 
**Auth** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**Disabled** | Pointer to **bool** | Temporarily disable IMAP operations for this account | [optional] 
**DraftsMailPath** | Pointer to **NullableString** | Custom folder path for draft messages. Defaults to auto-detected \&quot;Drafts\&quot; folder. Set to null to use default. | [optional] 
**Host** | Pointer to **string** | IMAP server hostname | [optional] 
**JunkMailPath** | Pointer to **NullableString** | Custom folder path for spam/junk messages. Defaults to auto-detected \&quot;Junk\&quot; folder. Set to null to use default. | [optional] 
**Port** | Pointer to **int32** | IMAP server port (typically 993 for IMAP over TLS, 143 for STARTTLS) | [optional] 
**ResyncDelay** | Pointer to **int32** | Delay in seconds between full mailbox resynchronizations | [optional] [default to 900]
**Secure** | Pointer to **bool** | Use TLS encryption for the connection (true for port 993, false for STARTTLS on port 143) | [optional] [default to false]
**SentMailPath** | Pointer to **NullableString** | Custom folder path for sent messages. Defaults to auto-detected \&quot;Sent\&quot; folder. Set to null to use default. | [optional] 
**Tls** | Pointer to [**TLS**](TLS.md) |  | [optional] 
**TrashMailPath** | Pointer to **NullableString** | Custom folder path for deleted messages. Defaults to auto-detected \&quot;Trash\&quot; folder. Set to null to use default. | [optional] 
**UseAuthServer** | Pointer to **bool** | Use external authentication server to retrieve credentials dynamically | [optional] 

## Methods

### NewIMAPResponse

`func NewIMAPResponse() *IMAPResponse`

NewIMAPResponse instantiates a new IMAPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIMAPResponseWithDefaults

`func NewIMAPResponseWithDefaults() *IMAPResponse`

NewIMAPResponseWithDefaults instantiates a new IMAPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveMailPath

`func (o *IMAPResponse) GetArchiveMailPath() string`

GetArchiveMailPath returns the ArchiveMailPath field if non-nil, zero value otherwise.

### GetArchiveMailPathOk

`func (o *IMAPResponse) GetArchiveMailPathOk() (*string, bool)`

GetArchiveMailPathOk returns a tuple with the ArchiveMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveMailPath

`func (o *IMAPResponse) SetArchiveMailPath(v string)`

SetArchiveMailPath sets ArchiveMailPath field to given value.

### HasArchiveMailPath

`func (o *IMAPResponse) HasArchiveMailPath() bool`

HasArchiveMailPath returns a boolean if a field has been set.

### SetArchiveMailPathNil

`func (o *IMAPResponse) SetArchiveMailPathNil(b bool)`

 SetArchiveMailPathNil sets the value for ArchiveMailPath to be an explicit nil

### UnsetArchiveMailPath
`func (o *IMAPResponse) UnsetArchiveMailPath()`

UnsetArchiveMailPath ensures that no value is present for ArchiveMailPath, not even an explicit nil
### GetAuth

`func (o *IMAPResponse) GetAuth() Authentication`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IMAPResponse) GetAuthOk() (*Authentication, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IMAPResponse) SetAuth(v Authentication)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *IMAPResponse) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetDisabled

`func (o *IMAPResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *IMAPResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *IMAPResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *IMAPResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDraftsMailPath

`func (o *IMAPResponse) GetDraftsMailPath() string`

GetDraftsMailPath returns the DraftsMailPath field if non-nil, zero value otherwise.

### GetDraftsMailPathOk

`func (o *IMAPResponse) GetDraftsMailPathOk() (*string, bool)`

GetDraftsMailPathOk returns a tuple with the DraftsMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftsMailPath

`func (o *IMAPResponse) SetDraftsMailPath(v string)`

SetDraftsMailPath sets DraftsMailPath field to given value.

### HasDraftsMailPath

`func (o *IMAPResponse) HasDraftsMailPath() bool`

HasDraftsMailPath returns a boolean if a field has been set.

### SetDraftsMailPathNil

`func (o *IMAPResponse) SetDraftsMailPathNil(b bool)`

 SetDraftsMailPathNil sets the value for DraftsMailPath to be an explicit nil

### UnsetDraftsMailPath
`func (o *IMAPResponse) UnsetDraftsMailPath()`

UnsetDraftsMailPath ensures that no value is present for DraftsMailPath, not even an explicit nil
### GetHost

`func (o *IMAPResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IMAPResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IMAPResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IMAPResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJunkMailPath

`func (o *IMAPResponse) GetJunkMailPath() string`

GetJunkMailPath returns the JunkMailPath field if non-nil, zero value otherwise.

### GetJunkMailPathOk

`func (o *IMAPResponse) GetJunkMailPathOk() (*string, bool)`

GetJunkMailPathOk returns a tuple with the JunkMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunkMailPath

`func (o *IMAPResponse) SetJunkMailPath(v string)`

SetJunkMailPath sets JunkMailPath field to given value.

### HasJunkMailPath

`func (o *IMAPResponse) HasJunkMailPath() bool`

HasJunkMailPath returns a boolean if a field has been set.

### SetJunkMailPathNil

`func (o *IMAPResponse) SetJunkMailPathNil(b bool)`

 SetJunkMailPathNil sets the value for JunkMailPath to be an explicit nil

### UnsetJunkMailPath
`func (o *IMAPResponse) UnsetJunkMailPath()`

UnsetJunkMailPath ensures that no value is present for JunkMailPath, not even an explicit nil
### GetPort

`func (o *IMAPResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IMAPResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IMAPResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IMAPResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetResyncDelay

`func (o *IMAPResponse) GetResyncDelay() int32`

GetResyncDelay returns the ResyncDelay field if non-nil, zero value otherwise.

### GetResyncDelayOk

`func (o *IMAPResponse) GetResyncDelayOk() (*int32, bool)`

GetResyncDelayOk returns a tuple with the ResyncDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncDelay

`func (o *IMAPResponse) SetResyncDelay(v int32)`

SetResyncDelay sets ResyncDelay field to given value.

### HasResyncDelay

`func (o *IMAPResponse) HasResyncDelay() bool`

HasResyncDelay returns a boolean if a field has been set.

### GetSecure

`func (o *IMAPResponse) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *IMAPResponse) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *IMAPResponse) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *IMAPResponse) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetSentMailPath

`func (o *IMAPResponse) GetSentMailPath() string`

GetSentMailPath returns the SentMailPath field if non-nil, zero value otherwise.

### GetSentMailPathOk

`func (o *IMAPResponse) GetSentMailPathOk() (*string, bool)`

GetSentMailPathOk returns a tuple with the SentMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentMailPath

`func (o *IMAPResponse) SetSentMailPath(v string)`

SetSentMailPath sets SentMailPath field to given value.

### HasSentMailPath

`func (o *IMAPResponse) HasSentMailPath() bool`

HasSentMailPath returns a boolean if a field has been set.

### SetSentMailPathNil

`func (o *IMAPResponse) SetSentMailPathNil(b bool)`

 SetSentMailPathNil sets the value for SentMailPath to be an explicit nil

### UnsetSentMailPath
`func (o *IMAPResponse) UnsetSentMailPath()`

UnsetSentMailPath ensures that no value is present for SentMailPath, not even an explicit nil
### GetTls

`func (o *IMAPResponse) GetTls() TLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IMAPResponse) GetTlsOk() (*TLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IMAPResponse) SetTls(v TLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IMAPResponse) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTrashMailPath

`func (o *IMAPResponse) GetTrashMailPath() string`

GetTrashMailPath returns the TrashMailPath field if non-nil, zero value otherwise.

### GetTrashMailPathOk

`func (o *IMAPResponse) GetTrashMailPathOk() (*string, bool)`

GetTrashMailPathOk returns a tuple with the TrashMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashMailPath

`func (o *IMAPResponse) SetTrashMailPath(v string)`

SetTrashMailPath sets TrashMailPath field to given value.

### HasTrashMailPath

`func (o *IMAPResponse) HasTrashMailPath() bool`

HasTrashMailPath returns a boolean if a field has been set.

### SetTrashMailPathNil

`func (o *IMAPResponse) SetTrashMailPathNil(b bool)`

 SetTrashMailPathNil sets the value for TrashMailPath to be an explicit nil

### UnsetTrashMailPath
`func (o *IMAPResponse) UnsetTrashMailPath()`

UnsetTrashMailPath ensures that no value is present for TrashMailPath, not even an explicit nil
### GetUseAuthServer

`func (o *IMAPResponse) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *IMAPResponse) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *IMAPResponse) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *IMAPResponse) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


