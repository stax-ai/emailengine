# OAuth2Update

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | OAuth2 access token for the email account | [optional] 
**Auth** | [**OAuth2Authentication**](OAuth2Authentication.md) |  | 
**Authorize** | Pointer to **bool** | Request an OAuth2 authorization URL instead of directly configuring credentials | [optional] 
**Expires** | Pointer to **time.Time** | Access token expiration timestamp | [optional] 
**Partial** | Pointer to **bool** | Update only the provided fields instead of replacing the entire configuration | [optional] [default to false]
**Provider** | Pointer to **string** | OAuth2 Application ID configured in EmailEngine | [optional] 
**RefreshToken** | Pointer to **string** | OAuth2 refresh token for obtaining new access tokens | [optional] 
**UseAuthServer** | Pointer to **bool** | Use external authentication server for token management instead of EmailEngine | [optional] 

## Methods

### NewOAuth2Update

`func NewOAuth2Update(auth OAuth2Authentication, ) *OAuth2Update`

NewOAuth2Update instantiates a new OAuth2Update object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2UpdateWithDefaults

`func NewOAuth2UpdateWithDefaults() *OAuth2Update`

NewOAuth2UpdateWithDefaults instantiates a new OAuth2Update object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuth2Update) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuth2Update) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuth2Update) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OAuth2Update) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAuth

`func (o *OAuth2Update) GetAuth() OAuth2Authentication`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *OAuth2Update) GetAuthOk() (*OAuth2Authentication, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *OAuth2Update) SetAuth(v OAuth2Authentication)`

SetAuth sets Auth field to given value.


### GetAuthorize

`func (o *OAuth2Update) GetAuthorize() bool`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *OAuth2Update) GetAuthorizeOk() (*bool, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *OAuth2Update) SetAuthorize(v bool)`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *OAuth2Update) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### GetExpires

`func (o *OAuth2Update) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OAuth2Update) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OAuth2Update) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *OAuth2Update) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetPartial

`func (o *OAuth2Update) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *OAuth2Update) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *OAuth2Update) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *OAuth2Update) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetProvider

`func (o *OAuth2Update) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OAuth2Update) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OAuth2Update) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *OAuth2Update) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OAuth2Update) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuth2Update) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuth2Update) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuth2Update) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUseAuthServer

`func (o *OAuth2Update) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *OAuth2Update) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *OAuth2Update) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *OAuth2Update) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


