# AccountsFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountResponseItem**](AccountResponseItem.md) |  | [optional] 
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 

## Methods

### NewAccountsFilterResponse

`func NewAccountsFilterResponse() *AccountsFilterResponse`

NewAccountsFilterResponse instantiates a new AccountsFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsFilterResponseWithDefaults

`func NewAccountsFilterResponseWithDefaults() *AccountsFilterResponse`

NewAccountsFilterResponseWithDefaults instantiates a new AccountsFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountsFilterResponse) GetAccounts() []AccountResponseItem`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountsFilterResponse) GetAccountsOk() (*[]AccountResponseItem, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountsFilterResponse) SetAccounts(v []AccountResponseItem)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AccountsFilterResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetPage

`func (o *AccountsFilterResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AccountsFilterResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AccountsFilterResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AccountsFilterResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *AccountsFilterResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AccountsFilterResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AccountsFilterResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *AccountsFilterResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *AccountsFilterResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountsFilterResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountsFilterResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AccountsFilterResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


