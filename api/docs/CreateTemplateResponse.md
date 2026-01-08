# CreateTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Created** | Pointer to **bool** | Was the template created or not | [optional] 
**Id** | **string** | Template ID | 

## Methods

### NewCreateTemplateResponse

`func NewCreateTemplateResponse(account string, id string, ) *CreateTemplateResponse`

NewCreateTemplateResponse instantiates a new CreateTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplateResponseWithDefaults

`func NewCreateTemplateResponseWithDefaults() *CreateTemplateResponse`

NewCreateTemplateResponseWithDefaults instantiates a new CreateTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateTemplateResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateTemplateResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateTemplateResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetCreated

`func (o *CreateTemplateResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateTemplateResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateTemplateResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateTemplateResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *CreateTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateTemplateResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


