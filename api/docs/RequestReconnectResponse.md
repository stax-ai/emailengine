# RequestReconnectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reconnect** | Pointer to **bool** | Reconnection status | [optional] [default to false]

## Methods

### NewRequestReconnectResponse

`func NewRequestReconnectResponse() *RequestReconnectResponse`

NewRequestReconnectResponse instantiates a new RequestReconnectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestReconnectResponseWithDefaults

`func NewRequestReconnectResponseWithDefaults() *RequestReconnectResponse`

NewRequestReconnectResponseWithDefaults instantiates a new RequestReconnectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReconnect

`func (o *RequestReconnectResponse) GetReconnect() bool`

GetReconnect returns the Reconnect field if non-nil, zero value otherwise.

### GetReconnectOk

`func (o *RequestReconnectResponse) GetReconnectOk() (*bool, bool)`

GetReconnectOk returns a tuple with the Reconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconnect

`func (o *RequestReconnectResponse) SetReconnect(v bool)`

SetReconnect sets Reconnect field to given value.

### HasReconnect

`func (o *RequestReconnectResponse) HasReconnect() bool`

HasReconnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


