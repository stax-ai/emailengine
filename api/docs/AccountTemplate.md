# AccountTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The time this template was created | [optional] 
**Description** | Pointer to **string** | Optional description of the template | [optional] 
**Format** | Pointer to [**Format**](Format.md) |  | [optional] [default to FORMAT_HTML]
**Id** | **string** | Template ID | 
**Name** | **string** | Name of the template | 
**Updated** | Pointer to **time.Time** | The time this template was last updated | [optional] 

## Methods

### NewAccountTemplate

`func NewAccountTemplate(id string, name string, ) *AccountTemplate`

NewAccountTemplate instantiates a new AccountTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTemplateWithDefaults

`func NewAccountTemplateWithDefaults() *AccountTemplate`

NewAccountTemplateWithDefaults instantiates a new AccountTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AccountTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *AccountTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *AccountTemplate) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AccountTemplate) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AccountTemplate) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AccountTemplate) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetId

`func (o *AccountTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccountTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetUpdated

`func (o *AccountTemplate) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AccountTemplate) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AccountTemplate) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AccountTemplate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


