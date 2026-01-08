# CreateAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**State** | [**CreateAccountState**](CreateAccountState.md) |  | 

## Methods

### NewCreateAccountResponse

`func NewCreateAccountResponse(account string, state CreateAccountState, ) *CreateAccountResponse`

NewCreateAccountResponse instantiates a new CreateAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountResponseWithDefaults

`func NewCreateAccountResponseWithDefaults() *CreateAccountResponse`

NewCreateAccountResponseWithDefaults instantiates a new CreateAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateAccountResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateAccountResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateAccountResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetState

`func (o *CreateAccountResponse) GetState() CreateAccountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateAccountResponse) GetStateOk() (*CreateAccountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateAccountResponse) SetState(v CreateAccountState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


