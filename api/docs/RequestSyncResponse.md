# RequestSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sync** | Pointer to **bool** | Sync status | [optional] [default to false]

## Methods

### NewRequestSyncResponse

`func NewRequestSyncResponse() *RequestSyncResponse`

NewRequestSyncResponse instantiates a new RequestSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSyncResponseWithDefaults

`func NewRequestSyncResponseWithDefaults() *RequestSyncResponse`

NewRequestSyncResponseWithDefaults instantiates a new RequestSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSync

`func (o *RequestSyncResponse) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *RequestSyncResponse) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *RequestSyncResponse) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *RequestSyncResponse) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


