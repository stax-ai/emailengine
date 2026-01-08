# Model2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access Token | 
**Account** | **string** | Unique identifier for the email account | 
**Provider** | [**OAuth2Provider**](OAuth2Provider.md) |  | 
**User** | **string** | Username | 

## Methods

### NewModel2

`func NewModel2(accessToken string, account string, provider OAuth2Provider, user string, ) *Model2`

NewModel2 instantiates a new Model2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel2WithDefaults

`func NewModel2WithDefaults() *Model2`

NewModel2WithDefaults instantiates a new Model2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Model2) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Model2) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Model2) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccount

`func (o *Model2) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Model2) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Model2) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetProvider

`func (o *Model2) GetProvider() OAuth2Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Model2) GetProviderOk() (*OAuth2Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Model2) SetProvider(v OAuth2Provider)`

SetProvider sets Provider field to given value.


### GetUser

`func (o *Model2) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Model2) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Model2) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


