# ConnectionsStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationError** | Pointer to **int32** | Authentication failed | [optional] 
**ConnectError** | Pointer to **int32** | Connection failed due to technical error | [optional] 
**Connected** | Pointer to **int32** | Successfully connected accounts | [optional] 
**Connecting** | Pointer to **int32** | Connection is being established | [optional] 
**Disconnected** | Pointer to **int32** | IMAP connection was closed | [optional] 
**Init** | Pointer to **int32** | Accounts not yet initialized | [optional] 
**Unset** | Pointer to **int32** | Accounts without valid IMAP settings | [optional] 

## Methods

### NewConnectionsStats

`func NewConnectionsStats() *ConnectionsStats`

NewConnectionsStats instantiates a new ConnectionsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionsStatsWithDefaults

`func NewConnectionsStatsWithDefaults() *ConnectionsStats`

NewConnectionsStatsWithDefaults instantiates a new ConnectionsStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationError

`func (o *ConnectionsStats) GetAuthenticationError() int32`

GetAuthenticationError returns the AuthenticationError field if non-nil, zero value otherwise.

### GetAuthenticationErrorOk

`func (o *ConnectionsStats) GetAuthenticationErrorOk() (*int32, bool)`

GetAuthenticationErrorOk returns a tuple with the AuthenticationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationError

`func (o *ConnectionsStats) SetAuthenticationError(v int32)`

SetAuthenticationError sets AuthenticationError field to given value.

### HasAuthenticationError

`func (o *ConnectionsStats) HasAuthenticationError() bool`

HasAuthenticationError returns a boolean if a field has been set.

### GetConnectError

`func (o *ConnectionsStats) GetConnectError() int32`

GetConnectError returns the ConnectError field if non-nil, zero value otherwise.

### GetConnectErrorOk

`func (o *ConnectionsStats) GetConnectErrorOk() (*int32, bool)`

GetConnectErrorOk returns a tuple with the ConnectError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectError

`func (o *ConnectionsStats) SetConnectError(v int32)`

SetConnectError sets ConnectError field to given value.

### HasConnectError

`func (o *ConnectionsStats) HasConnectError() bool`

HasConnectError returns a boolean if a field has been set.

### GetConnected

`func (o *ConnectionsStats) GetConnected() int32`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ConnectionsStats) GetConnectedOk() (*int32, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ConnectionsStats) SetConnected(v int32)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ConnectionsStats) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnecting

`func (o *ConnectionsStats) GetConnecting() int32`

GetConnecting returns the Connecting field if non-nil, zero value otherwise.

### GetConnectingOk

`func (o *ConnectionsStats) GetConnectingOk() (*int32, bool)`

GetConnectingOk returns a tuple with the Connecting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnecting

`func (o *ConnectionsStats) SetConnecting(v int32)`

SetConnecting sets Connecting field to given value.

### HasConnecting

`func (o *ConnectionsStats) HasConnecting() bool`

HasConnecting returns a boolean if a field has been set.

### GetDisconnected

`func (o *ConnectionsStats) GetDisconnected() int32`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ConnectionsStats) GetDisconnectedOk() (*int32, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ConnectionsStats) SetDisconnected(v int32)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ConnectionsStats) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetInit

`func (o *ConnectionsStats) GetInit() int32`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *ConnectionsStats) GetInitOk() (*int32, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *ConnectionsStats) SetInit(v int32)`

SetInit sets Init field to given value.

### HasInit

`func (o *ConnectionsStats) HasInit() bool`

HasInit returns a boolean if a field has been set.

### GetUnset

`func (o *ConnectionsStats) GetUnset() int32`

GetUnset returns the Unset field if non-nil, zero value otherwise.

### GetUnsetOk

`func (o *ConnectionsStats) GetUnsetOk() (*int32, bool)`

GetUnsetOk returns a tuple with the Unset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnset

`func (o *ConnectionsStats) SetUnset(v int32)`

SetUnset sets Unset field to given value.

### HasUnset

`func (o *ConnectionsStats) HasUnset() bool`

HasUnset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


