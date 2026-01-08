# ImapConfiguration

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

### NewImapConfiguration

`func NewImapConfiguration() *ImapConfiguration`

NewImapConfiguration instantiates a new ImapConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImapConfigurationWithDefaults

`func NewImapConfigurationWithDefaults() *ImapConfiguration`

NewImapConfigurationWithDefaults instantiates a new ImapConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveMailPath

`func (o *ImapConfiguration) GetArchiveMailPath() string`

GetArchiveMailPath returns the ArchiveMailPath field if non-nil, zero value otherwise.

### GetArchiveMailPathOk

`func (o *ImapConfiguration) GetArchiveMailPathOk() (*string, bool)`

GetArchiveMailPathOk returns a tuple with the ArchiveMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveMailPath

`func (o *ImapConfiguration) SetArchiveMailPath(v string)`

SetArchiveMailPath sets ArchiveMailPath field to given value.

### HasArchiveMailPath

`func (o *ImapConfiguration) HasArchiveMailPath() bool`

HasArchiveMailPath returns a boolean if a field has been set.

### SetArchiveMailPathNil

`func (o *ImapConfiguration) SetArchiveMailPathNil(b bool)`

 SetArchiveMailPathNil sets the value for ArchiveMailPath to be an explicit nil

### UnsetArchiveMailPath
`func (o *ImapConfiguration) UnsetArchiveMailPath()`

UnsetArchiveMailPath ensures that no value is present for ArchiveMailPath, not even an explicit nil
### GetAuth

`func (o *ImapConfiguration) GetAuth() Authentication`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ImapConfiguration) GetAuthOk() (*Authentication, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ImapConfiguration) SetAuth(v Authentication)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ImapConfiguration) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetDisabled

`func (o *ImapConfiguration) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ImapConfiguration) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ImapConfiguration) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ImapConfiguration) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDraftsMailPath

`func (o *ImapConfiguration) GetDraftsMailPath() string`

GetDraftsMailPath returns the DraftsMailPath field if non-nil, zero value otherwise.

### GetDraftsMailPathOk

`func (o *ImapConfiguration) GetDraftsMailPathOk() (*string, bool)`

GetDraftsMailPathOk returns a tuple with the DraftsMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftsMailPath

`func (o *ImapConfiguration) SetDraftsMailPath(v string)`

SetDraftsMailPath sets DraftsMailPath field to given value.

### HasDraftsMailPath

`func (o *ImapConfiguration) HasDraftsMailPath() bool`

HasDraftsMailPath returns a boolean if a field has been set.

### SetDraftsMailPathNil

`func (o *ImapConfiguration) SetDraftsMailPathNil(b bool)`

 SetDraftsMailPathNil sets the value for DraftsMailPath to be an explicit nil

### UnsetDraftsMailPath
`func (o *ImapConfiguration) UnsetDraftsMailPath()`

UnsetDraftsMailPath ensures that no value is present for DraftsMailPath, not even an explicit nil
### GetHost

`func (o *ImapConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ImapConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ImapConfiguration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ImapConfiguration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJunkMailPath

`func (o *ImapConfiguration) GetJunkMailPath() string`

GetJunkMailPath returns the JunkMailPath field if non-nil, zero value otherwise.

### GetJunkMailPathOk

`func (o *ImapConfiguration) GetJunkMailPathOk() (*string, bool)`

GetJunkMailPathOk returns a tuple with the JunkMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunkMailPath

`func (o *ImapConfiguration) SetJunkMailPath(v string)`

SetJunkMailPath sets JunkMailPath field to given value.

### HasJunkMailPath

`func (o *ImapConfiguration) HasJunkMailPath() bool`

HasJunkMailPath returns a boolean if a field has been set.

### SetJunkMailPathNil

`func (o *ImapConfiguration) SetJunkMailPathNil(b bool)`

 SetJunkMailPathNil sets the value for JunkMailPath to be an explicit nil

### UnsetJunkMailPath
`func (o *ImapConfiguration) UnsetJunkMailPath()`

UnsetJunkMailPath ensures that no value is present for JunkMailPath, not even an explicit nil
### GetPort

`func (o *ImapConfiguration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ImapConfiguration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ImapConfiguration) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ImapConfiguration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetResyncDelay

`func (o *ImapConfiguration) GetResyncDelay() int32`

GetResyncDelay returns the ResyncDelay field if non-nil, zero value otherwise.

### GetResyncDelayOk

`func (o *ImapConfiguration) GetResyncDelayOk() (*int32, bool)`

GetResyncDelayOk returns a tuple with the ResyncDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncDelay

`func (o *ImapConfiguration) SetResyncDelay(v int32)`

SetResyncDelay sets ResyncDelay field to given value.

### HasResyncDelay

`func (o *ImapConfiguration) HasResyncDelay() bool`

HasResyncDelay returns a boolean if a field has been set.

### GetSecure

`func (o *ImapConfiguration) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *ImapConfiguration) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *ImapConfiguration) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *ImapConfiguration) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetSentMailPath

`func (o *ImapConfiguration) GetSentMailPath() string`

GetSentMailPath returns the SentMailPath field if non-nil, zero value otherwise.

### GetSentMailPathOk

`func (o *ImapConfiguration) GetSentMailPathOk() (*string, bool)`

GetSentMailPathOk returns a tuple with the SentMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentMailPath

`func (o *ImapConfiguration) SetSentMailPath(v string)`

SetSentMailPath sets SentMailPath field to given value.

### HasSentMailPath

`func (o *ImapConfiguration) HasSentMailPath() bool`

HasSentMailPath returns a boolean if a field has been set.

### SetSentMailPathNil

`func (o *ImapConfiguration) SetSentMailPathNil(b bool)`

 SetSentMailPathNil sets the value for SentMailPath to be an explicit nil

### UnsetSentMailPath
`func (o *ImapConfiguration) UnsetSentMailPath()`

UnsetSentMailPath ensures that no value is present for SentMailPath, not even an explicit nil
### GetTls

`func (o *ImapConfiguration) GetTls() TLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *ImapConfiguration) GetTlsOk() (*TLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *ImapConfiguration) SetTls(v TLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *ImapConfiguration) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTrashMailPath

`func (o *ImapConfiguration) GetTrashMailPath() string`

GetTrashMailPath returns the TrashMailPath field if non-nil, zero value otherwise.

### GetTrashMailPathOk

`func (o *ImapConfiguration) GetTrashMailPathOk() (*string, bool)`

GetTrashMailPathOk returns a tuple with the TrashMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashMailPath

`func (o *ImapConfiguration) SetTrashMailPath(v string)`

SetTrashMailPath sets TrashMailPath field to given value.

### HasTrashMailPath

`func (o *ImapConfiguration) HasTrashMailPath() bool`

HasTrashMailPath returns a boolean if a field has been set.

### SetTrashMailPathNil

`func (o *ImapConfiguration) SetTrashMailPathNil(b bool)`

 SetTrashMailPathNil sets the value for TrashMailPath to be an explicit nil

### UnsetTrashMailPath
`func (o *ImapConfiguration) UnsetTrashMailPath()`

UnsetTrashMailPath ensures that no value is present for TrashMailPath, not even an explicit nil
### GetUseAuthServer

`func (o *ImapConfiguration) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *ImapConfiguration) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *ImapConfiguration) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *ImapConfiguration) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


