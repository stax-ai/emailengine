# RequestTemplateContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**Format**](Format.md) |  | [optional] [default to FORMAT_HTML]
**Html** | Pointer to **string** | HTML message content | [optional] 
**PreviewText** | Pointer to **string** | Preview text shown in email clients after the subject line | [optional] 
**Subject** | Pointer to **string** | Email subject line | [optional] 
**Text** | Pointer to **string** | Plain text message content | [optional] 

## Methods

### NewRequestTemplateContent

`func NewRequestTemplateContent() *RequestTemplateContent`

NewRequestTemplateContent instantiates a new RequestTemplateContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTemplateContentWithDefaults

`func NewRequestTemplateContentWithDefaults() *RequestTemplateContent`

NewRequestTemplateContentWithDefaults instantiates a new RequestTemplateContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RequestTemplateContent) GetFormat() Format`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RequestTemplateContent) GetFormatOk() (*Format, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RequestTemplateContent) SetFormat(v Format)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RequestTemplateContent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHtml

`func (o *RequestTemplateContent) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *RequestTemplateContent) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *RequestTemplateContent) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *RequestTemplateContent) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetPreviewText

`func (o *RequestTemplateContent) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *RequestTemplateContent) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *RequestTemplateContent) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *RequestTemplateContent) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### GetSubject

`func (o *RequestTemplateContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestTemplateContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestTemplateContent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RequestTemplateContent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *RequestTemplateContent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RequestTemplateContent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RequestTemplateContent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RequestTemplateContent) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


