# AccountTemplatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Templates** | Pointer to [**[]AccountTemplate**](AccountTemplate.md) |  | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 

## Methods

### NewAccountTemplatesResponse

`func NewAccountTemplatesResponse(account string, ) *AccountTemplatesResponse`

NewAccountTemplatesResponse instantiates a new AccountTemplatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTemplatesResponseWithDefaults

`func NewAccountTemplatesResponseWithDefaults() *AccountTemplatesResponse`

NewAccountTemplatesResponseWithDefaults instantiates a new AccountTemplatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountTemplatesResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountTemplatesResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountTemplatesResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetPage

`func (o *AccountTemplatesResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AccountTemplatesResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AccountTemplatesResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AccountTemplatesResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *AccountTemplatesResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AccountTemplatesResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AccountTemplatesResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *AccountTemplatesResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTemplates

`func (o *AccountTemplatesResponse) GetTemplates() []AccountTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *AccountTemplatesResponse) GetTemplatesOk() (*[]AccountTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *AccountTemplatesResponse) SetTemplates(v []AccountTemplate)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *AccountTemplatesResponse) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTotal

`func (o *AccountTemplatesResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountTemplatesResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountTemplatesResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AccountTemplatesResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


