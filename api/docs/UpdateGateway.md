# UpdateGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Hostname to connect to | [optional] 
**Name** | Pointer to **string** | Account Name | [optional] 
**Pass** | Pointer to **NullableString** | SMTP authentication password | [optional] 
**Port** | Pointer to **int32** | Service port number | [optional] 
**Secure** | Pointer to **bool** | Should connection use TLS. Usually true for port 465 | [optional] 
**User** | Pointer to **NullableString** | SMTP authentication username | [optional] 

## Methods

### NewUpdateGateway

`func NewUpdateGateway() *UpdateGateway`

NewUpdateGateway instantiates a new UpdateGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGatewayWithDefaults

`func NewUpdateGatewayWithDefaults() *UpdateGateway`

NewUpdateGatewayWithDefaults instantiates a new UpdateGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *UpdateGateway) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateGateway) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateGateway) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *UpdateGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPass

`func (o *UpdateGateway) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *UpdateGateway) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *UpdateGateway) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *UpdateGateway) HasPass() bool`

HasPass returns a boolean if a field has been set.

### SetPassNil

`func (o *UpdateGateway) SetPassNil(b bool)`

 SetPassNil sets the value for Pass to be an explicit nil

### UnsetPass
`func (o *UpdateGateway) UnsetPass()`

UnsetPass ensures that no value is present for Pass, not even an explicit nil
### GetPort

`func (o *UpdateGateway) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateGateway) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateGateway) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateGateway) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecure

`func (o *UpdateGateway) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *UpdateGateway) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *UpdateGateway) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *UpdateGateway) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetUser

`func (o *UpdateGateway) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateGateway) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateGateway) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateGateway) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *UpdateGateway) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *UpdateGateway) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


