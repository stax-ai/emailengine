# AccountTokensItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Description** | **string** | Token description | 
**Ip** | Pointer to **string** | IP address of the requester | [optional] 
**Metadata** | Pointer to **string** | Related metadata in JSON format | [optional] 
**Restrictions** | Pointer to [**TokenRestrictions**](TokenRestrictions.md) |  | [optional] 

## Methods

### NewAccountTokensItem

`func NewAccountTokensItem(account string, description string, ) *AccountTokensItem`

NewAccountTokensItem instantiates a new AccountTokensItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokensItemWithDefaults

`func NewAccountTokensItemWithDefaults() *AccountTokensItem`

NewAccountTokensItemWithDefaults instantiates a new AccountTokensItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountTokensItem) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountTokensItem) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountTokensItem) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetDescription

`func (o *AccountTokensItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountTokensItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountTokensItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIp

`func (o *AccountTokensItem) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AccountTokensItem) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AccountTokensItem) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AccountTokensItem) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountTokensItem) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountTokensItem) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountTokensItem) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountTokensItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRestrictions

`func (o *AccountTokensItem) GetRestrictions() TokenRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *AccountTokensItem) GetRestrictionsOk() (*TokenRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *AccountTokensItem) SetRestrictions(v TokenRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *AccountTokensItem) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


