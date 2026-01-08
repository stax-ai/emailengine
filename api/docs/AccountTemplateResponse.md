# AccountTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Unique identifier for the email account | 
**Content** | Pointer to [**RequestTemplateContent**](RequestTemplateContent.md) |  | [optional] 
**Created** | Pointer to **time.Time** | The time this template was created | [optional] 
**Description** | Pointer to **string** | Optional description of the template | [optional] 
**Format** | Pointer to [**Format**](Format.md) |  | [optional] [default to FORMAT_HTML]
**Id** | **string** | Template ID | 
**Name** | **string** | Name of the template | 
**Updated** | Pointer to **time.Time** | The time this template was last updated | [optional] 

## Methods

### NewAccountTemplateResponse

`func NewAccountTemplateResponse(account string, id string, name string, ) *AccountTemplateResponse`

NewAccountTemplateResponse instantiates a new AccountTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTemplateResponseWithDefaults

`func NewAccountTemplateResponseWithDefaults() *AccountTemplateResponse`

NewAccountTemplateResponseWithDefaults instantiates a new AccountTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountTemplateResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountTemplateResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountTemplateResponse) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetContent

`func (o *AccountTemplateResponse) GetContent() RequestTemplateContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AccountTemplateResponse) GetContentOk() (*RequestTemplateContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AccountTemplateResponse) SetContent(v RequestTemplateContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *AccountTemplateResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *AccountTemplateResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountTemplateResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountTemplateResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountTemplateResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *AccountTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *AccountTemplateResponse) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AccountTemplateResponse) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AccountTemplateResponse) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AccountTemplateResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *AccountTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccountTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUpdated

`func (o *AccountTemplateResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AccountTemplateResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AccountTemplateResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AccountTemplateResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


