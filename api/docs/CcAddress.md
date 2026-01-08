# CcAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Email address | 
**Name** | Pointer to **string** | Display name for the email address | [optional] 

## Methods

### NewCcAddress

`func NewCcAddress(address string, ) *CcAddress`

NewCcAddress instantiates a new CcAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCcAddressWithDefaults

`func NewCcAddressWithDefaults() *CcAddress`

NewCcAddressWithDefaults instantiates a new CcAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CcAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CcAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CcAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *CcAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CcAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CcAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CcAddress) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


