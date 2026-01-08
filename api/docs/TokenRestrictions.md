# TokenRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**AddressAllowlist**](AddressAllowlist.md) |  | [optional] 
**RateLimit** | Pointer to [**AddressRateLimit**](AddressRateLimit.md) |  | [optional] 
**Referrers** | Pointer to [**ReferrerAllowlist**](ReferrerAllowlist.md) |  | [optional] 

## Methods

### NewTokenRestrictions

`func NewTokenRestrictions() *TokenRestrictions`

NewTokenRestrictions instantiates a new TokenRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRestrictionsWithDefaults

`func NewTokenRestrictionsWithDefaults() *TokenRestrictions`

NewTokenRestrictionsWithDefaults instantiates a new TokenRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *TokenRestrictions) GetAddresses() AddressAllowlist`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenRestrictions) GetAddressesOk() (*AddressAllowlist, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenRestrictions) SetAddresses(v AddressAllowlist)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *TokenRestrictions) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetRateLimit

`func (o *TokenRestrictions) GetRateLimit() AddressRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *TokenRestrictions) GetRateLimitOk() (*AddressRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *TokenRestrictions) SetRateLimit(v AddressRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *TokenRestrictions) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetReferrers

`func (o *TokenRestrictions) GetReferrers() ReferrerAllowlist`

GetReferrers returns the Referrers field if non-nil, zero value otherwise.

### GetReferrersOk

`func (o *TokenRestrictions) GetReferrersOk() (*ReferrerAllowlist, bool)`

GetReferrersOk returns a tuple with the Referrers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrers

`func (o *TokenRestrictions) SetReferrers(v ReferrerAllowlist)`

SetReferrers sets Referrers field to given value.

### HasReferrers

`func (o *TokenRestrictions) HasReferrers() bool`

HasReferrers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


