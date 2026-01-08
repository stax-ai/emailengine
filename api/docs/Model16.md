# Model16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to [**AccessToken**](AccessToken.md) |  | [optional] 
**Pass** | Pointer to **string** | Password for SMTP authentication | [optional] 
**User** | **string** | Username or email address for SMTP authentication | 

## Methods

### NewModel16

`func NewModel16(user string, ) *Model16`

NewModel16 instantiates a new Model16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel16WithDefaults

`func NewModel16WithDefaults() *Model16`

NewModel16WithDefaults instantiates a new Model16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Model16) GetAccessToken() AccessToken`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Model16) GetAccessTokenOk() (*AccessToken, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Model16) SetAccessToken(v AccessToken)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Model16) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetPass

`func (o *Model16) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *Model16) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *Model16) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *Model16) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetUser

`func (o *Model16) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Model16) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Model16) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


