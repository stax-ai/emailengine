# Oauth2Response

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

### NewOauth2Response

`func NewOauth2Response(auth OAuth2Authentication, ) *Oauth2Response`

NewOauth2Response instantiates a new Oauth2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2ResponseWithDefaults

`func NewOauth2ResponseWithDefaults() *Oauth2Response`

NewOauth2ResponseWithDefaults instantiates a new Oauth2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Oauth2Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Oauth2Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Oauth2Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Oauth2Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAuth

`func (o *Oauth2Response) GetAuth() OAuth2Authentication`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Oauth2Response) GetAuthOk() (*OAuth2Authentication, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Oauth2Response) SetAuth(v OAuth2Authentication)`

SetAuth sets Auth field to given value.


### GetAuthorize

`func (o *Oauth2Response) GetAuthorize() bool`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *Oauth2Response) GetAuthorizeOk() (*bool, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *Oauth2Response) SetAuthorize(v bool)`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *Oauth2Response) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### GetExpires

`func (o *Oauth2Response) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Oauth2Response) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Oauth2Response) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Oauth2Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetProvider

`func (o *Oauth2Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Oauth2Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Oauth2Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Oauth2Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *Oauth2Response) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *Oauth2Response) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *Oauth2Response) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *Oauth2Response) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetRefreshToken

`func (o *Oauth2Response) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Oauth2Response) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Oauth2Response) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Oauth2Response) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUseAuthServer

`func (o *Oauth2Response) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *Oauth2Response) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *Oauth2Response) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *Oauth2Response) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


