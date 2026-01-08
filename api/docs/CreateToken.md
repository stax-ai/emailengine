# CreateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Description** | **string** | Token description | 
**Ip** | Pointer to **string** | IP address of the requester | [optional] 
**Metadata** | Pointer to **string** | Related metadata in JSON format | [optional] 
**Restrictions** | Pointer to [**TokenRestrictions**](TokenRestrictions.md) |  | [optional] 
**Scopes** | [**[]TokenScope**](TokenScope.md) | Token permission scopes: \&quot;api\&quot; for REST API access, \&quot;smtp\&quot; for SMTP submission, \&quot;imap-proxy\&quot; for IMAP proxy authentication | [default to ["api"]]

## Methods

### NewCreateToken

`func NewCreateToken(account string, description string, scopes []TokenScope, ) *CreateToken`

NewCreateToken instantiates a new CreateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenWithDefaults

`func NewCreateTokenWithDefaults() *CreateToken`

NewCreateTokenWithDefaults instantiates a new CreateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateToken) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateToken) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateToken) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetDescription

`func (o *CreateToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateToken) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIp

`func (o *CreateToken) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateToken) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateToken) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateToken) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateToken) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateToken) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateToken) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateToken) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRestrictions

`func (o *CreateToken) GetRestrictions() TokenRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *CreateToken) GetRestrictionsOk() (*TokenRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *CreateToken) SetRestrictions(v TokenRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *CreateToken) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetScopes

`func (o *CreateToken) GetScopes() []TokenScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateToken) GetScopesOk() (*[]TokenScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateToken) SetScopes(v []TokenScope)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


