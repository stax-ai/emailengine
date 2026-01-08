# SettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int32** | Number of registered accounts | [optional] 
**Connections** | Pointer to [**ConnectionsStats**](ConnectionsStats.md) |  | [optional] 
**Counters** | Pointer to **map[string]interface{}** |  | [optional] 
**License** | Pointer to **string** | EmailEngine license | [optional] 
**Node** | Pointer to **string** | Node.js Version | [optional] 
**Redis** | Pointer to **string** | Redis Version | [optional] 
**Version** | Pointer to **string** | EmailEngine version number | [optional] 

## Methods

### NewSettingsResponse

`func NewSettingsResponse() *SettingsResponse`

NewSettingsResponse instantiates a new SettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsResponseWithDefaults

`func NewSettingsResponseWithDefaults() *SettingsResponse`

NewSettingsResponseWithDefaults instantiates a new SettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *SettingsResponse) GetAccounts() int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SettingsResponse) GetAccountsOk() (*int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SettingsResponse) SetAccounts(v int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SettingsResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetConnections

`func (o *SettingsResponse) GetConnections() ConnectionsStats`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *SettingsResponse) GetConnectionsOk() (*ConnectionsStats, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *SettingsResponse) SetConnections(v ConnectionsStats)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *SettingsResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetCounters

`func (o *SettingsResponse) GetCounters() map[string]interface{}`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *SettingsResponse) GetCountersOk() (*map[string]interface{}, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *SettingsResponse) SetCounters(v map[string]interface{})`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *SettingsResponse) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetLicense

`func (o *SettingsResponse) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SettingsResponse) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SettingsResponse) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SettingsResponse) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetNode

`func (o *SettingsResponse) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *SettingsResponse) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *SettingsResponse) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *SettingsResponse) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetRedis

`func (o *SettingsResponse) GetRedis() string`

GetRedis returns the Redis field if non-nil, zero value otherwise.

### GetRedisOk

`func (o *SettingsResponse) GetRedisOk() (*string, bool)`

GetRedisOk returns a tuple with the Redis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis

`func (o *SettingsResponse) SetRedis(v string)`

SetRedis sets Redis field to given value.

### HasRedis

`func (o *SettingsResponse) HasRedis() bool`

HasRedis returns a boolean if a field has been set.

### GetVersion

`func (o *SettingsResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SettingsResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SettingsResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SettingsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


