# OutboxEntryProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**OutboxListProgressError**](OutboxListProgressError.md) |  | [optional] 
**Response** | Pointer to **string** | SMTP server response (when status is \&quot;processing\&quot;) | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewOutboxEntryProgress

`func NewOutboxEntryProgress() *OutboxEntryProgress`

NewOutboxEntryProgress instantiates a new OutboxEntryProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxEntryProgressWithDefaults

`func NewOutboxEntryProgressWithDefaults() *OutboxEntryProgress`

NewOutboxEntryProgressWithDefaults instantiates a new OutboxEntryProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OutboxEntryProgress) GetError() OutboxListProgressError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OutboxEntryProgress) GetErrorOk() (*OutboxListProgressError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OutboxEntryProgress) SetError(v OutboxListProgressError)`

SetError sets Error field to given value.

### HasError

`func (o *OutboxEntryProgress) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResponse

`func (o *OutboxEntryProgress) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OutboxEntryProgress) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OutboxEntryProgress) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OutboxEntryProgress) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *OutboxEntryProgress) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutboxEntryProgress) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutboxEntryProgress) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OutboxEntryProgress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


