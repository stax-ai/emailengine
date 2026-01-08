# OAuth2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | OAuth2 access token for the email account | [optional] 
**Auth** | [**OAuth2Authentication**](OAuth2Authentication.md) |  | 
**Authorize** | Pointer to **bool** | Request an OAuth2 authorization URL instead of directly configuring credentials | [optional] 
**Expires** | Pointer to **time.Time** | Access token expiration timestamp | [optional] 
**Provider** | Pointer to **string** | OAuth2 Application ID configured in EmailEngine | [optional] 
**RedirectUrl** | Pointer to **string** | URL to redirect to after OAuth2 authorization completes (only used when authorize&#x3D;true) | [optional] 
**RefreshToken** | Pointer to **string** | OAuth2 refresh token for obtaining new access tokens | [optional] 
**UseAuthServer** | Pointer to **bool** | Use external authentication server for token management instead of EmailEngine | [optional] 

## Methods

### NewOAuth2

`func NewOAuth2(auth OAuth2Authentication, ) *OAuth2`

NewOAuth2 instantiates a new OAuth2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2WithDefaults

`func NewOAuth2WithDefaults() *OAuth2`

NewOAuth2WithDefaults instantiates a new OAuth2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuth2) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuth2) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuth2) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OAuth2) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAuth

`func (o *OAuth2) GetAuth() OAuth2Authentication`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *OAuth2) GetAuthOk() (*OAuth2Authentication, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *OAuth2) SetAuth(v OAuth2Authentication)`

SetAuth sets Auth field to given value.


### GetAuthorize

`func (o *OAuth2) GetAuthorize() bool`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *OAuth2) GetAuthorizeOk() (*bool, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *OAuth2) SetAuthorize(v bool)`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *OAuth2) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### GetExpires

`func (o *OAuth2) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OAuth2) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OAuth2) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *OAuth2) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetProvider

`func (o *OAuth2) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OAuth2) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OAuth2) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *OAuth2) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *OAuth2) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *OAuth2) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *OAuth2) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *OAuth2) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OAuth2) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuth2) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuth2) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuth2) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUseAuthServer

`func (o *OAuth2) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *OAuth2) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *OAuth2) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *OAuth2) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


