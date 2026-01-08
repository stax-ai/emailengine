# CreateGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | **string** | Gateway ID | 
**Host** | **string** | Hostname to connect to | 
**Name** | **string** | Account Name | 
**Pass** | Pointer to **string** | SMTP authentication password | [optional] 
**Port** | **int32** | Service port number | 
**Secure** | Pointer to **bool** | Should connection use TLS. Usually true for port 465 | [optional] [default to false]
**User** | Pointer to **string** | SMTP authentication username | [optional] 

## Methods

### NewCreateGateway

`func NewCreateGateway(gateway string, host string, name string, port int32, ) *CreateGateway`

NewCreateGateway instantiates a new CreateGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGatewayWithDefaults

`func NewCreateGatewayWithDefaults() *CreateGateway`

NewCreateGatewayWithDefaults instantiates a new CreateGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *CreateGateway) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateGateway) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateGateway) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetHost

`func (o *CreateGateway) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateGateway) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateGateway) SetHost(v string)`

SetHost sets Host field to given value.


### GetName

`func (o *CreateGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGateway) SetName(v string)`

SetName sets Name field to given value.


### GetPass

`func (o *CreateGateway) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *CreateGateway) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *CreateGateway) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *CreateGateway) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetPort

`func (o *CreateGateway) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateGateway) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateGateway) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecure

`func (o *CreateGateway) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *CreateGateway) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *CreateGateway) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *CreateGateway) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetUser

`func (o *CreateGateway) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateGateway) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateGateway) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateGateway) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


