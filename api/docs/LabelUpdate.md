# LabelUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** | Gmail labels to add (use label ID or path) | [optional] 
**Delete** | Pointer to **[]string** | Gmail labels to remove (use label ID or path) | [optional] 

## Methods

### NewLabelUpdate

`func NewLabelUpdate() *LabelUpdate`

NewLabelUpdate instantiates a new LabelUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelUpdateWithDefaults

`func NewLabelUpdateWithDefaults() *LabelUpdate`

NewLabelUpdateWithDefaults instantiates a new LabelUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *LabelUpdate) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *LabelUpdate) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *LabelUpdate) SetAdd(v []string)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *LabelUpdate) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetDelete

`func (o *LabelUpdate) GetDelete() []string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *LabelUpdate) GetDeleteOk() (*[]string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *LabelUpdate) SetDelete(v []string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *LabelUpdate) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


