# DiscoveredServerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**DetectedAuthenticationInfo**](DetectedAuthenticationInfo.md) |  | [optional] 
**Host** | **string** | Hostname to connect to | 
**Port** | **int32** | Service port number | 
**Secure** | Pointer to **bool** | Should connection use TLS. Usually true for port 993 | [optional] [default to false]

## Methods

### NewDiscoveredServerSettings

`func NewDiscoveredServerSettings(host string, port int32, ) *DiscoveredServerSettings`

NewDiscoveredServerSettings instantiates a new DiscoveredServerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredServerSettingsWithDefaults

`func NewDiscoveredServerSettingsWithDefaults() *DiscoveredServerSettings`

NewDiscoveredServerSettingsWithDefaults instantiates a new DiscoveredServerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *DiscoveredServerSettings) GetAuth() DetectedAuthenticationInfo`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *DiscoveredServerSettings) GetAuthOk() (*DetectedAuthenticationInfo, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *DiscoveredServerSettings) SetAuth(v DetectedAuthenticationInfo)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *DiscoveredServerSettings) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetHost

`func (o *DiscoveredServerSettings) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DiscoveredServerSettings) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DiscoveredServerSettings) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DiscoveredServerSettings) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DiscoveredServerSettings) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DiscoveredServerSettings) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecure

`func (o *DiscoveredServerSettings) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *DiscoveredServerSettings) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *DiscoveredServerSettings) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *DiscoveredServerSettings) HasSecure() bool`

HasSecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


