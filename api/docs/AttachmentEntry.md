# AttachmentEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentId** | Pointer to **string** | Content-ID header value used for embedding images in HTML | [optional] 
**ContentType** | Pointer to **string** | MIME type of the attachment | [optional] 
**Embedded** | Pointer to **bool** | Whether the attachment is embedded in the HTML content | [optional] 
**EncodedSize** | Pointer to **int32** | Size of the attachment as stored in the email (base64 encoded). The actual decoded file size is approximately 75% of this value. | [optional] 
**Filename** | Pointer to **string** | Original filename of the attachment | [optional] 
**Id** | Pointer to **string** | Unique identifier for the attachment | [optional] 
**Inline** | Pointer to **bool** | Whether the attachment should be displayed inline rather than as a download | [optional] 
**Method** | Pointer to **string** | Calendar method (REQUEST, REPLY, CANCEL, etc.) for iCalendar attachments | [optional] 

## Methods

### NewAttachmentEntry

`func NewAttachmentEntry() *AttachmentEntry`

NewAttachmentEntry instantiates a new AttachmentEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentEntryWithDefaults

`func NewAttachmentEntryWithDefaults() *AttachmentEntry`

NewAttachmentEntryWithDefaults instantiates a new AttachmentEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentId

`func (o *AttachmentEntry) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *AttachmentEntry) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *AttachmentEntry) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *AttachmentEntry) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### GetContentType

`func (o *AttachmentEntry) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AttachmentEntry) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AttachmentEntry) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AttachmentEntry) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEmbedded

`func (o *AttachmentEntry) GetEmbedded() bool`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AttachmentEntry) GetEmbeddedOk() (*bool, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AttachmentEntry) SetEmbedded(v bool)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AttachmentEntry) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEncodedSize

`func (o *AttachmentEntry) GetEncodedSize() int32`

GetEncodedSize returns the EncodedSize field if non-nil, zero value otherwise.

### GetEncodedSizeOk

`func (o *AttachmentEntry) GetEncodedSizeOk() (*int32, bool)`

GetEncodedSizeOk returns a tuple with the EncodedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedSize

`func (o *AttachmentEntry) SetEncodedSize(v int32)`

SetEncodedSize sets EncodedSize field to given value.

### HasEncodedSize

`func (o *AttachmentEntry) HasEncodedSize() bool`

HasEncodedSize returns a boolean if a field has been set.

### GetFilename

`func (o *AttachmentEntry) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AttachmentEntry) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AttachmentEntry) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AttachmentEntry) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetId

`func (o *AttachmentEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInline

`func (o *AttachmentEntry) GetInline() bool`

GetInline returns the Inline field if non-nil, zero value otherwise.

### GetInlineOk

`func (o *AttachmentEntry) GetInlineOk() (*bool, bool)`

GetInlineOk returns a tuple with the Inline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInline

`func (o *AttachmentEntry) SetInline(v bool)`

SetInline sets Inline field to given value.

### HasInline

`func (o *AttachmentEntry) HasInline() bool`

HasInline returns a boolean if a field has been set.

### GetMethod

`func (o *AttachmentEntry) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AttachmentEntry) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AttachmentEntry) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AttachmentEntry) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


