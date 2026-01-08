# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to [**AccessToken**](AccessToken.md) |  | [optional] 
**Pass** | Pointer to **string** | Password for IMAP authentication | [optional] 
**User** | **string** | Username or email address for IMAP authentication | 

## Methods

### NewAuthentication

`func NewAuthentication(user string, ) *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Authentication) GetAccessToken() AccessToken`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Authentication) GetAccessTokenOk() (*AccessToken, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Authentication) SetAccessToken(v AccessToken)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Authentication) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetPass

`func (o *Authentication) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *Authentication) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *Authentication) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *Authentication) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetUser

`func (o *Authentication) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Authentication) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Authentication) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


