# TextInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedSize** | Pointer to [**EncodedSize**](EncodedSize.md) |  | [optional] 
**Id** | Pointer to **string** | Identifier for fetching the full message text | [optional] 

## Methods

### NewTextInfo

`func NewTextInfo() *TextInfo`

NewTextInfo instantiates a new TextInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextInfoWithDefaults

`func NewTextInfoWithDefaults() *TextInfo`

NewTextInfoWithDefaults instantiates a new TextInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedSize

`func (o *TextInfo) GetEncodedSize() EncodedSize`

GetEncodedSize returns the EncodedSize field if non-nil, zero value otherwise.

### GetEncodedSizeOk

`func (o *TextInfo) GetEncodedSizeOk() (*EncodedSize, bool)`

GetEncodedSizeOk returns a tuple with the EncodedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedSize

`func (o *TextInfo) SetEncodedSize(v EncodedSize)`

SetEncodedSize sets EncodedSize field to given value.

### HasEncodedSize

`func (o *TextInfo) HasEncodedSize() bool`

HasEncodedSize returns a boolean if a field has been set.

### GetId

`func (o *TextInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TextInfo) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


