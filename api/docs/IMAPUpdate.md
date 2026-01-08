# IMAPUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveMailPath** | Pointer to **NullableString** | Custom folder path for archived messages. Set to null to use default. | [optional] 
**Auth** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**Disabled** | Pointer to **bool** | Temporarily disable IMAP operations for this account | [optional] 
**DraftsMailPath** | Pointer to **NullableString** | Custom folder path for draft messages. Set to null to use default. | [optional] 
**Host** | Pointer to **string** | IMAP server hostname | [optional] 
**JunkMailPath** | Pointer to **NullableString** | Custom folder path for spam/junk messages. Set to null to use default. | [optional] 
**Partial** | Pointer to **bool** | Update only the provided fields instead of replacing the entire configuration | [optional] [default to false]
**Port** | Pointer to **int32** | IMAP server port | [optional] 
**ResyncDelay** | Pointer to **int32** | Delay in seconds between full mailbox resynchronizations | [optional] 
**Secure** | Pointer to **bool** | Use TLS encryption for the connection | [optional] 
**SentMailPath** | Pointer to **NullableString** | Custom folder path for sent messages. Set to null to use default. | [optional] 
**Tls** | Pointer to [**Model15**](Model15.md) |  | [optional] 
**TrashMailPath** | Pointer to **NullableString** | Custom folder path for deleted messages. Set to null to use default. | [optional] 
**UseAuthServer** | Pointer to **bool** | Use external authentication server to retrieve credentials dynamically | [optional] 

## Methods

### NewIMAPUpdate

`func NewIMAPUpdate() *IMAPUpdate`

NewIMAPUpdate instantiates a new IMAPUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIMAPUpdateWithDefaults

`func NewIMAPUpdateWithDefaults() *IMAPUpdate`

NewIMAPUpdateWithDefaults instantiates a new IMAPUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveMailPath

`func (o *IMAPUpdate) GetArchiveMailPath() string`

GetArchiveMailPath returns the ArchiveMailPath field if non-nil, zero value otherwise.

### GetArchiveMailPathOk

`func (o *IMAPUpdate) GetArchiveMailPathOk() (*string, bool)`

GetArchiveMailPathOk returns a tuple with the ArchiveMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveMailPath

`func (o *IMAPUpdate) SetArchiveMailPath(v string)`

SetArchiveMailPath sets ArchiveMailPath field to given value.

### HasArchiveMailPath

`func (o *IMAPUpdate) HasArchiveMailPath() bool`

HasArchiveMailPath returns a boolean if a field has been set.

### SetArchiveMailPathNil

`func (o *IMAPUpdate) SetArchiveMailPathNil(b bool)`

 SetArchiveMailPathNil sets the value for ArchiveMailPath to be an explicit nil

### UnsetArchiveMailPath
`func (o *IMAPUpdate) UnsetArchiveMailPath()`

UnsetArchiveMailPath ensures that no value is present for ArchiveMailPath, not even an explicit nil
### GetAuth

`func (o *IMAPUpdate) GetAuth() Authentication`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IMAPUpdate) GetAuthOk() (*Authentication, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IMAPUpdate) SetAuth(v Authentication)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *IMAPUpdate) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetDisabled

`func (o *IMAPUpdate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *IMAPUpdate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *IMAPUpdate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *IMAPUpdate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDraftsMailPath

`func (o *IMAPUpdate) GetDraftsMailPath() string`

GetDraftsMailPath returns the DraftsMailPath field if non-nil, zero value otherwise.

### GetDraftsMailPathOk

`func (o *IMAPUpdate) GetDraftsMailPathOk() (*string, bool)`

GetDraftsMailPathOk returns a tuple with the DraftsMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftsMailPath

`func (o *IMAPUpdate) SetDraftsMailPath(v string)`

SetDraftsMailPath sets DraftsMailPath field to given value.

### HasDraftsMailPath

`func (o *IMAPUpdate) HasDraftsMailPath() bool`

HasDraftsMailPath returns a boolean if a field has been set.

### SetDraftsMailPathNil

`func (o *IMAPUpdate) SetDraftsMailPathNil(b bool)`

 SetDraftsMailPathNil sets the value for DraftsMailPath to be an explicit nil

### UnsetDraftsMailPath
`func (o *IMAPUpdate) UnsetDraftsMailPath()`

UnsetDraftsMailPath ensures that no value is present for DraftsMailPath, not even an explicit nil
### GetHost

`func (o *IMAPUpdate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IMAPUpdate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IMAPUpdate) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IMAPUpdate) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetJunkMailPath

`func (o *IMAPUpdate) GetJunkMailPath() string`

GetJunkMailPath returns the JunkMailPath field if non-nil, zero value otherwise.

### GetJunkMailPathOk

`func (o *IMAPUpdate) GetJunkMailPathOk() (*string, bool)`

GetJunkMailPathOk returns a tuple with the JunkMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunkMailPath

`func (o *IMAPUpdate) SetJunkMailPath(v string)`

SetJunkMailPath sets JunkMailPath field to given value.

### HasJunkMailPath

`func (o *IMAPUpdate) HasJunkMailPath() bool`

HasJunkMailPath returns a boolean if a field has been set.

### SetJunkMailPathNil

`func (o *IMAPUpdate) SetJunkMailPathNil(b bool)`

 SetJunkMailPathNil sets the value for JunkMailPath to be an explicit nil

### UnsetJunkMailPath
`func (o *IMAPUpdate) UnsetJunkMailPath()`

UnsetJunkMailPath ensures that no value is present for JunkMailPath, not even an explicit nil
### GetPartial

`func (o *IMAPUpdate) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *IMAPUpdate) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *IMAPUpdate) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *IMAPUpdate) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetPort

`func (o *IMAPUpdate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IMAPUpdate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IMAPUpdate) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IMAPUpdate) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetResyncDelay

`func (o *IMAPUpdate) GetResyncDelay() int32`

GetResyncDelay returns the ResyncDelay field if non-nil, zero value otherwise.

### GetResyncDelayOk

`func (o *IMAPUpdate) GetResyncDelayOk() (*int32, bool)`

GetResyncDelayOk returns a tuple with the ResyncDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncDelay

`func (o *IMAPUpdate) SetResyncDelay(v int32)`

SetResyncDelay sets ResyncDelay field to given value.

### HasResyncDelay

`func (o *IMAPUpdate) HasResyncDelay() bool`

HasResyncDelay returns a boolean if a field has been set.

### GetSecure

`func (o *IMAPUpdate) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *IMAPUpdate) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *IMAPUpdate) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *IMAPUpdate) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetSentMailPath

`func (o *IMAPUpdate) GetSentMailPath() string`

GetSentMailPath returns the SentMailPath field if non-nil, zero value otherwise.

### GetSentMailPathOk

`func (o *IMAPUpdate) GetSentMailPathOk() (*string, bool)`

GetSentMailPathOk returns a tuple with the SentMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentMailPath

`func (o *IMAPUpdate) SetSentMailPath(v string)`

SetSentMailPath sets SentMailPath field to given value.

### HasSentMailPath

`func (o *IMAPUpdate) HasSentMailPath() bool`

HasSentMailPath returns a boolean if a field has been set.

### SetSentMailPathNil

`func (o *IMAPUpdate) SetSentMailPathNil(b bool)`

 SetSentMailPathNil sets the value for SentMailPath to be an explicit nil

### UnsetSentMailPath
`func (o *IMAPUpdate) UnsetSentMailPath()`

UnsetSentMailPath ensures that no value is present for SentMailPath, not even an explicit nil
### GetTls

`func (o *IMAPUpdate) GetTls() Model15`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IMAPUpdate) GetTlsOk() (*Model15, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IMAPUpdate) SetTls(v Model15)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IMAPUpdate) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTrashMailPath

`func (o *IMAPUpdate) GetTrashMailPath() string`

GetTrashMailPath returns the TrashMailPath field if non-nil, zero value otherwise.

### GetTrashMailPathOk

`func (o *IMAPUpdate) GetTrashMailPathOk() (*string, bool)`

GetTrashMailPathOk returns a tuple with the TrashMailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashMailPath

`func (o *IMAPUpdate) SetTrashMailPath(v string)`

SetTrashMailPath sets TrashMailPath field to given value.

### HasTrashMailPath

`func (o *IMAPUpdate) HasTrashMailPath() bool`

HasTrashMailPath returns a boolean if a field has been set.

### SetTrashMailPathNil

`func (o *IMAPUpdate) SetTrashMailPathNil(b bool)`

 SetTrashMailPathNil sets the value for TrashMailPath to be an explicit nil

### UnsetTrashMailPath
`func (o *IMAPUpdate) UnsetTrashMailPath()`

UnsetTrashMailPath ensures that no value is present for TrashMailPath, not even an explicit nil
### GetUseAuthServer

`func (o *IMAPUpdate) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *IMAPUpdate) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *IMAPUpdate) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *IMAPUpdate) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


