# AccountQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | How many bytes can the account store emails | [optional] 
**Status** | Pointer to **string** | Textual information about the usage | [optional] 
**Usage** | Pointer to **int32** | How many bytes has the account stored in emails | [optional] 

## Methods

### NewAccountQuota

`func NewAccountQuota() *AccountQuota`

NewAccountQuota instantiates a new AccountQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountQuotaWithDefaults

`func NewAccountQuotaWithDefaults() *AccountQuota`

NewAccountQuotaWithDefaults instantiates a new AccountQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AccountQuota) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccountQuota) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccountQuota) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccountQuota) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetStatus

`func (o *AccountQuota) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountQuota) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountQuota) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountQuota) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsage

`func (o *AccountQuota) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AccountQuota) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AccountQuota) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AccountQuota) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


