# UpdateTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Id** | **string** | Template ID | 
**Updated** | Pointer to **bool** | Was the template updated or not | [optional] 

## Methods

### NewUpdateTemplateResponse

`func NewUpdateTemplateResponse(account string, id string, ) *UpdateTemplateResponse`

NewUpdateTemplateResponse instantiates a new UpdateTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateResponseWithDefaults

`func NewUpdateTemplateResponseWithDefaults() *UpdateTemplateResponse`

NewUpdateTemplateResponseWithDefaults instantiates a new UpdateTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UpdateTemplateResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateTemplateResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateTemplateResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetId

`func (o *UpdateTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUpdated

`func (o *UpdateTemplateResponse) GetUpdated() bool`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpdateTemplateResponse) GetUpdatedOk() (*bool, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpdateTemplateResponse) SetUpdated(v bool)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UpdateTemplateResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


