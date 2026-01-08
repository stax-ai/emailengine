# AddressRateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxRequests** | Pointer to **int32** | Maximum requests allowed in the time window | [optional] 
**TimeWindow** | Pointer to **int32** | Time window duration in seconds | [optional] 

## Methods

### NewAddressRateLimit

`func NewAddressRateLimit() *AddressRateLimit`

NewAddressRateLimit instantiates a new AddressRateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressRateLimitWithDefaults

`func NewAddressRateLimitWithDefaults() *AddressRateLimit`

NewAddressRateLimitWithDefaults instantiates a new AddressRateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxRequests

`func (o *AddressRateLimit) GetMaxRequests() int32`

GetMaxRequests returns the MaxRequests field if non-nil, zero value otherwise.

### GetMaxRequestsOk

`func (o *AddressRateLimit) GetMaxRequestsOk() (*int32, bool)`

GetMaxRequestsOk returns a tuple with the MaxRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequests

`func (o *AddressRateLimit) SetMaxRequests(v int32)`

SetMaxRequests sets MaxRequests field to given value.

### HasMaxRequests

`func (o *AddressRateLimit) HasMaxRequests() bool`

HasMaxRequests returns a boolean if a field has been set.

### GetTimeWindow

`func (o *AddressRateLimit) GetTimeWindow() int32`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *AddressRateLimit) GetTimeWindowOk() (*int32, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *AddressRateLimit) SetTimeWindow(v int32)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *AddressRateLimit) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


