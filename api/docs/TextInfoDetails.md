# TextInfoDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedSize** | Pointer to [**EncodedSize**](EncodedSize.md) |  | [optional] 
**HasMore** | Pointer to **bool** | Whether the message content was truncated (true if more content is available via separate API call) | [optional] 
**Html** | Pointer to **string** | HTML version of the message | [optional] 
**Id** | Pointer to **string** | Identifier for fetching additional text content | [optional] 
**Plain** | Pointer to **string** | Plain text version of the message | [optional] 

## Methods

### NewTextInfoDetails

`func NewTextInfoDetails() *TextInfoDetails`

NewTextInfoDetails instantiates a new TextInfoDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextInfoDetailsWithDefaults

`func NewTextInfoDetailsWithDefaults() *TextInfoDetails`

NewTextInfoDetailsWithDefaults instantiates a new TextInfoDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedSize

`func (o *TextInfoDetails) GetEncodedSize() EncodedSize`

GetEncodedSize returns the EncodedSize field if non-nil, zero value otherwise.

### GetEncodedSizeOk

`func (o *TextInfoDetails) GetEncodedSizeOk() (*EncodedSize, bool)`

GetEncodedSizeOk returns a tuple with the EncodedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedSize

`func (o *TextInfoDetails) SetEncodedSize(v EncodedSize)`

SetEncodedSize sets EncodedSize field to given value.

### HasEncodedSize

`func (o *TextInfoDetails) HasEncodedSize() bool`

HasEncodedSize returns a boolean if a field has been set.

### GetHasMore

`func (o *TextInfoDetails) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TextInfoDetails) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TextInfoDetails) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *TextInfoDetails) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetHtml

`func (o *TextInfoDetails) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TextInfoDetails) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TextInfoDetails) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *TextInfoDetails) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetId

`func (o *TextInfoDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextInfoDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextInfoDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TextInfoDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlain

`func (o *TextInfoDetails) GetPlain() string`

GetPlain returns the Plain field if non-nil, zero value otherwise.

### GetPlainOk

`func (o *TextInfoDetails) GetPlainOk() (*string, bool)`

GetPlainOk returns a tuple with the Plain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlain

`func (o *TextInfoDetails) SetPlain(v string)`

SetPlain sets Plain field to given value.

### HasPlain

`func (o *TextInfoDetails) HasPlain() bool`

HasPlain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


