# TemplateRender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**Format**](Format.md) |  | [optional] [default to FORMAT_HTML]
**Params** | Pointer to **map[string]interface{}** | An object of variables for the template renderer | [optional] 

## Methods

### NewTemplateRender

`func NewTemplateRender() *TemplateRender`

NewTemplateRender instantiates a new TemplateRender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateRenderWithDefaults

`func NewTemplateRenderWithDefaults() *TemplateRender`

NewTemplateRenderWithDefaults instantiates a new TemplateRender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *TemplateRender) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TemplateRender) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TemplateRender) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TemplateRender) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetParams

`func (o *TemplateRender) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TemplateRender) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TemplateRender) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *TemplateRender) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


