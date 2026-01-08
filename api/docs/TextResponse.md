# TextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | Is the current text output capped or not | [optional] 
**Html** | Pointer to **string** | HTML content | [optional] 
**Plain** | Pointer to **string** | Plaintext content | [optional] 

## Methods

### NewTextResponse

`func NewTextResponse() *TextResponse`

NewTextResponse instantiates a new TextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextResponseWithDefaults

`func NewTextResponseWithDefaults() *TextResponse`

NewTextResponseWithDefaults instantiates a new TextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *TextResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TextResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TextResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *TextResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetHtml

`func (o *TextResponse) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TextResponse) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TextResponse) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *TextResponse) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetPlain

`func (o *TextResponse) GetPlain() string`

GetPlain returns the Plain field if non-nil, zero value otherwise.

### GetPlainOk

`func (o *TextResponse) GetPlainOk() (*string, bool)`

GetPlainOk returns a tuple with the Plain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlain

`func (o *TextResponse) SetPlain(v string)`

SetPlain sets Plain field to given value.

### HasPlain

`func (o *TextResponse) HasPlain() bool`

HasPlain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


