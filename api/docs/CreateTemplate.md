# CreateTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **NullableString** | Account ID. Use &#x60;null&#x60; for public templates. | 
**Content** | [**CreateTemplateContent**](CreateTemplateContent.md) |  | 
**Description** | Pointer to **string** | Optional description of the template | [optional] 
**Format** | Pointer to [**Format**](Format.md) |  | [optional] [default to FORMAT_HTML]
**Name** | **string** | Name of the template | 

## Methods

### NewCreateTemplate

`func NewCreateTemplate(account NullableString, content CreateTemplateContent, name string, ) *CreateTemplate`

NewCreateTemplate instantiates a new CreateTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplateWithDefaults

`func NewCreateTemplateWithDefaults() *CreateTemplate`

NewCreateTemplateWithDefaults instantiates a new CreateTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateTemplate) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateTemplate) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateTemplate) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *CreateTemplate) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CreateTemplate) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContent

`func (o *CreateTemplate) GetContent() CreateTemplateContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateTemplate) GetContentOk() (*CreateTemplateContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateTemplate) SetContent(v CreateTemplateContent)`

SetContent sets Content field to given value.


### GetDescription

`func (o *CreateTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *CreateTemplate) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateTemplate) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateTemplate) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateTemplate) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *CreateTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTemplate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


