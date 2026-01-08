# EncodedSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Html** | Pointer to **int32** | Size of the HTML part in bytes | [optional] 
**Plain** | Pointer to **int32** | Size of the plain text part in bytes | [optional] 

## Methods

### NewEncodedSize

`func NewEncodedSize() *EncodedSize`

NewEncodedSize instantiates a new EncodedSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncodedSizeWithDefaults

`func NewEncodedSizeWithDefaults() *EncodedSize`

NewEncodedSizeWithDefaults instantiates a new EncodedSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHtml

`func (o *EncodedSize) GetHtml() int32`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *EncodedSize) GetHtmlOk() (*int32, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *EncodedSize) SetHtml(v int32)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *EncodedSize) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetPlain

`func (o *EncodedSize) GetPlain() int32`

GetPlain returns the Plain field if non-nil, zero value otherwise.

### GetPlainOk

`func (o *EncodedSize) GetPlainOk() (*int32, bool)`

GetPlainOk returns a tuple with the Plain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlain

`func (o *EncodedSize) SetPlain(v int32)`

SetPlain sets Plain field to given value.

### HasPlain

`func (o *EncodedSize) HasPlain() bool`

HasPlain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


