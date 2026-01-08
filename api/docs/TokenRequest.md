# TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | OAuth2 client ID used for authentication | [optional] 
**Grant** | Pointer to [**Grant**](Grant.md) |  | [optional] 
**Provider** | Pointer to **string** | OAuth2 provider name | [optional] 
**Response** | Pointer to **map[string]interface{}** | Raw error response from the OAuth2 server | [optional] 
**Scopes** | Pointer to **[]string** | Requested OAuth2 permission scopes | [optional] 
**Status** | Pointer to **int32** | HTTP status code from the OAuth2 server | [optional] 

## Methods

### NewTokenRequest

`func NewTokenRequest() *TokenRequest`

NewTokenRequest instantiates a new TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestWithDefaults

`func NewTokenRequestWithDefaults() *TokenRequest`

NewTokenRequestWithDefaults instantiates a new TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetGrant

`func (o *TokenRequest) GetGrant() Grant`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *TokenRequest) GetGrantOk() (*Grant, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *TokenRequest) SetGrant(v Grant)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *TokenRequest) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetProvider

`func (o *TokenRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TokenRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TokenRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TokenRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetResponse

`func (o *TokenRequest) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *TokenRequest) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *TokenRequest) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *TokenRequest) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetScopes

`func (o *TokenRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *TokenRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TokenRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


