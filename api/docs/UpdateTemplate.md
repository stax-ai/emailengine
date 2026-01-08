# UpdateTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**UpdateTemplateContent**](UpdateTemplateContent.md) |  | [optional] 
**Description** | Pointer to **string** | Optional description of the template | [optional] 
**Format** | Pointer to [**Format**](Format.md) |  | [optional] [default to FORMAT_HTML]
**Name** | Pointer to **string** | Name of the template | [optional] 

## Methods

### NewUpdateTemplate

`func NewUpdateTemplate() *UpdateTemplate`

NewUpdateTemplate instantiates a new UpdateTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateWithDefaults

`func NewUpdateTemplateWithDefaults() *UpdateTemplate`

NewUpdateTemplateWithDefaults instantiates a new UpdateTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *UpdateTemplate) GetContent() UpdateTemplateContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateTemplate) GetContentOk() (*UpdateTemplateContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateTemplate) SetContent(v UpdateTemplateContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateTemplate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *UpdateTemplate) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UpdateTemplate) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UpdateTemplate) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UpdateTemplate) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *UpdateTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTemplate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


