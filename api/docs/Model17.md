# Model17

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinVersion** | Pointer to **string** | Minimum TLS version to accept | [optional] 
**RejectUnauthorized** | Pointer to **bool** | Reject connections to servers with invalid TLS certificates | [optional] 

## Methods

### NewModel17

`func NewModel17() *Model17`

NewModel17 instantiates a new Model17 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel17WithDefaults

`func NewModel17WithDefaults() *Model17`

NewModel17WithDefaults instantiates a new Model17 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinVersion

`func (o *Model17) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *Model17) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *Model17) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *Model17) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetRejectUnauthorized

`func (o *Model17) GetRejectUnauthorized() bool`

GetRejectUnauthorized returns the RejectUnauthorized field if non-nil, zero value otherwise.

### GetRejectUnauthorizedOk

`func (o *Model17) GetRejectUnauthorizedOk() (*bool, bool)`

GetRejectUnauthorizedOk returns a tuple with the RejectUnauthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUnauthorized

`func (o *Model17) SetRejectUnauthorized(v bool)`

SetRejectUnauthorized sets RejectUnauthorized field to given value.

### HasRejectUnauthorized

`func (o *Model17) HasRejectUnauthorized() bool`

HasRejectUnauthorized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


