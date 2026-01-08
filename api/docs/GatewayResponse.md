# GatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deliveries** | Pointer to **int32** | Count of email deliveries using this gateway | [optional] 
**Gateway** | **string** | Gateway ID | 
**Host** | Pointer to **string** | Hostname to connect to | [optional] 
**LastError** | Pointer to [**NullableAccountErrorEntry**](AccountErrorEntry.md) |  | [optional] 
**LastUse** | Pointer to **time.Time** | Last delivery time | [optional] 
**Name** | **string** | Display name for the gateway | 
**Pass** | Pointer to **string** | SMTP authentication password | [optional] 
**Port** | Pointer to **int32** | Service port number | [optional] 
**Secure** | Pointer to **bool** | Should connection use TLS. Usually true for port 465 | [optional] [default to false]
**User** | Pointer to **string** | SMTP authentication username | [optional] 

## Methods

### NewGatewayResponse

`func NewGatewayResponse(gateway string, name string, ) *GatewayResponse`

NewGatewayResponse instantiates a new GatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayResponseWithDefaults

`func NewGatewayResponseWithDefaults() *GatewayResponse`

NewGatewayResponseWithDefaults instantiates a new GatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveries

`func (o *GatewayResponse) GetDeliveries() int32`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *GatewayResponse) GetDeliveriesOk() (*int32, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *GatewayResponse) SetDeliveries(v int32)`

SetDeliveries sets Deliveries field to given value.

### HasDeliveries

`func (o *GatewayResponse) HasDeliveries() bool`

HasDeliveries returns a boolean if a field has been set.

### GetGateway

`func (o *GatewayResponse) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GatewayResponse) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GatewayResponse) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetHost

`func (o *GatewayResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GatewayResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GatewayResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GatewayResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLastError

`func (o *GatewayResponse) GetLastError() AccountErrorEntry`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GatewayResponse) GetLastErrorOk() (*AccountErrorEntry, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GatewayResponse) SetLastError(v AccountErrorEntry)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GatewayResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GatewayResponse) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GatewayResponse) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastUse

`func (o *GatewayResponse) GetLastUse() time.Time`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *GatewayResponse) GetLastUseOk() (*time.Time, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *GatewayResponse) SetLastUse(v time.Time)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *GatewayResponse) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.

### GetName

`func (o *GatewayResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPass

`func (o *GatewayResponse) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *GatewayResponse) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *GatewayResponse) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *GatewayResponse) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetPort

`func (o *GatewayResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GatewayResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecure

`func (o *GatewayResponse) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *GatewayResponse) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *GatewayResponse) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *GatewayResponse) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetUser

`func (o *GatewayResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GatewayResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GatewayResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GatewayResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


