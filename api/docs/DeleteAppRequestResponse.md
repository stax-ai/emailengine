# DeleteAppRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int32** | The number of accounts registered with this application. Not available for legacy apps. | [optional] 
**Deleted** | Pointer to **bool** | Was the OAuth2 application deleted | [optional] [default to true]
**Id** | **string** | OAuth2 application ID | 

## Methods

### NewDeleteAppRequestResponse

`func NewDeleteAppRequestResponse(id string, ) *DeleteAppRequestResponse`

NewDeleteAppRequestResponse instantiates a new DeleteAppRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAppRequestResponseWithDefaults

`func NewDeleteAppRequestResponseWithDefaults() *DeleteAppRequestResponse`

NewDeleteAppRequestResponseWithDefaults instantiates a new DeleteAppRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *DeleteAppRequestResponse) GetAccounts() int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *DeleteAppRequestResponse) GetAccountsOk() (*int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *DeleteAppRequestResponse) SetAccounts(v int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *DeleteAppRequestResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetDeleted

`func (o *DeleteAppRequestResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteAppRequestResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteAppRequestResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteAppRequestResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *DeleteAppRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteAppRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteAppRequestResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


