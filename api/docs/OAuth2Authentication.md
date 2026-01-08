# OAuth2Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedAccount** | Pointer to **string** | Account ID to use for authenticating the shared mailbox. When provided, EmailEngine uses this account&#39;s credentials instead of creating new ones. | [optional] 
**DelegatedUser** | Pointer to **string** | Shared mailbox username (Microsoft 365 delegation) | [optional] 
**User** | Pointer to **string** | Primary email account username | [optional] 

## Methods

### NewOAuth2Authentication

`func NewOAuth2Authentication() *OAuth2Authentication`

NewOAuth2Authentication instantiates a new OAuth2Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2AuthenticationWithDefaults

`func NewOAuth2AuthenticationWithDefaults() *OAuth2Authentication`

NewOAuth2AuthenticationWithDefaults instantiates a new OAuth2Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedAccount

`func (o *OAuth2Authentication) GetDelegatedAccount() string`

GetDelegatedAccount returns the DelegatedAccount field if non-nil, zero value otherwise.

### GetDelegatedAccountOk

`func (o *OAuth2Authentication) GetDelegatedAccountOk() (*string, bool)`

GetDelegatedAccountOk returns a tuple with the DelegatedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAccount

`func (o *OAuth2Authentication) SetDelegatedAccount(v string)`

SetDelegatedAccount sets DelegatedAccount field to given value.

### HasDelegatedAccount

`func (o *OAuth2Authentication) HasDelegatedAccount() bool`

HasDelegatedAccount returns a boolean if a field has been set.

### GetDelegatedUser

`func (o *OAuth2Authentication) GetDelegatedUser() string`

GetDelegatedUser returns the DelegatedUser field if non-nil, zero value otherwise.

### GetDelegatedUserOk

`func (o *OAuth2Authentication) GetDelegatedUserOk() (*string, bool)`

GetDelegatedUserOk returns a tuple with the DelegatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedUser

`func (o *OAuth2Authentication) SetDelegatedUser(v string)`

SetDelegatedUser sets DelegatedUser field to given value.

### HasDelegatedUser

`func (o *OAuth2Authentication) HasDelegatedUser() bool`

HasDelegatedUser returns a boolean if a field has been set.

### GetUser

`func (o *OAuth2Authentication) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OAuth2Authentication) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OAuth2Authentication) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OAuth2Authentication) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


