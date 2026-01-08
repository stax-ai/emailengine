# MessageUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to [**FlagUpdate**](FlagUpdate.md) |  | [optional] 
**Labels** | Pointer to [**LabelUpdate**](LabelUpdate.md) |  | [optional] 

## Methods

### NewMessageUpdate

`func NewMessageUpdate() *MessageUpdate`

NewMessageUpdate instantiates a new MessageUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUpdateWithDefaults

`func NewMessageUpdateWithDefaults() *MessageUpdate`

NewMessageUpdateWithDefaults instantiates a new MessageUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *MessageUpdate) GetFlags() FlagUpdate`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *MessageUpdate) GetFlagsOk() (*FlagUpdate, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *MessageUpdate) SetFlags(v FlagUpdate)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *MessageUpdate) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetLabels

`func (o *MessageUpdate) GetLabels() LabelUpdate`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MessageUpdate) GetLabelsOk() (*LabelUpdate, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MessageUpdate) SetLabels(v LabelUpdate)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MessageUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


