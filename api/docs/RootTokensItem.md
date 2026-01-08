# RootTokensItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Description** | **string** | Token description | 
**Ip** | Pointer to **string** | IP address of the requester | [optional] 
**Metadata** | Pointer to **string** | Related metadata in JSON format | [optional] 

## Methods

### NewRootTokensItem

`func NewRootTokensItem(account string, description string, ) *RootTokensItem`

NewRootTokensItem instantiates a new RootTokensItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootTokensItemWithDefaults

`func NewRootTokensItemWithDefaults() *RootTokensItem`

NewRootTokensItemWithDefaults instantiates a new RootTokensItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *RootTokensItem) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RootTokensItem) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RootTokensItem) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetDescription

`func (o *RootTokensItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RootTokensItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RootTokensItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIp

`func (o *RootTokensItem) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RootTokensItem) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RootTokensItem) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RootTokensItem) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMetadata

`func (o *RootTokensItem) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RootTokensItem) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RootTokensItem) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RootTokensItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


