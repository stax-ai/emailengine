# Model15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinVersion** | Pointer to **string** | Minimum TLS version to accept | [optional] 
**RejectUnauthorized** | Pointer to **bool** | Reject connections to servers with invalid TLS certificates | [optional] 

## Methods

### NewModel15

`func NewModel15() *Model15`

NewModel15 instantiates a new Model15 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel15WithDefaults

`func NewModel15WithDefaults() *Model15`

NewModel15WithDefaults instantiates a new Model15 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinVersion

`func (o *Model15) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *Model15) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *Model15) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *Model15) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetRejectUnauthorized

`func (o *Model15) GetRejectUnauthorized() bool`

GetRejectUnauthorized returns the RejectUnauthorized field if non-nil, zero value otherwise.

### GetRejectUnauthorizedOk

`func (o *Model15) GetRejectUnauthorizedOk() (*bool, bool)`

GetRejectUnauthorizedOk returns a tuple with the RejectUnauthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUnauthorized

`func (o *Model15) SetRejectUnauthorized(v bool)`

SetRejectUnauthorized sets RejectUnauthorized field to given value.

### HasRejectUnauthorized

`func (o *Model15) HasRejectUnauthorized() bool`

HasRejectUnauthorized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


