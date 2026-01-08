# MessageUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to [**FlagResponse**](FlagResponse.md) |  | [optional] 
**Labels** | Pointer to [**Model24**](Model24.md) |  | [optional] 

## Methods

### NewMessageUpdateResponse

`func NewMessageUpdateResponse() *MessageUpdateResponse`

NewMessageUpdateResponse instantiates a new MessageUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUpdateResponseWithDefaults

`func NewMessageUpdateResponseWithDefaults() *MessageUpdateResponse`

NewMessageUpdateResponseWithDefaults instantiates a new MessageUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *MessageUpdateResponse) GetFlags() FlagResponse`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageUpdateResponse) GetFlagsOk() (*FlagResponse, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageUpdateResponse) SetFlags(v FlagResponse)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageUpdateResponse) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetLabels

`func (o *MessageUpdateResponse) GetLabels() Model24`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MessageUpdateResponse) GetLabelsOk() (*Model24, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MessageUpdateResponse) SetLabels(v Model24)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MessageUpdateResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


