# FlagUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** | Flags to add to the message | [optional] 
**Delete** | Pointer to **[]string** | Flags to remove from the message | [optional] 
**Set** | Pointer to **[]string** | Replace all flags with this list | [optional] 

## Methods

### NewFlagUpdate

`func NewFlagUpdate() *FlagUpdate`

NewFlagUpdate instantiates a new FlagUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagUpdateWithDefaults

`func NewFlagUpdateWithDefaults() *FlagUpdate`

NewFlagUpdateWithDefaults instantiates a new FlagUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *FlagUpdate) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *FlagUpdate) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *FlagUpdate) SetAdd(v []string)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *FlagUpdate) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetDelete

`func (o *FlagUpdate) GetDelete() []string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *FlagUpdate) GetDeleteOk() (*[]string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *FlagUpdate) SetDelete(v []string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *FlagUpdate) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetSet

`func (o *FlagUpdate) GetSet() []string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *FlagUpdate) GetSetOk() (*[]string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *FlagUpdate) SetSet(v []string)`

SetSet sets Set field to given value.

### HasSet

`func (o *FlagUpdate) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


