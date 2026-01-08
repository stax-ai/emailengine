# RcptAddressEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Email address | 
**Name** | Pointer to **string** | Display name for the email address | [optional] 

## Methods

### NewRcptAddressEntry

`func NewRcptAddressEntry(address string, ) *RcptAddressEntry`

NewRcptAddressEntry instantiates a new RcptAddressEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcptAddressEntryWithDefaults

`func NewRcptAddressEntryWithDefaults() *RcptAddressEntry`

NewRcptAddressEntryWithDefaults instantiates a new RcptAddressEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RcptAddressEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RcptAddressEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RcptAddressEntry) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *RcptAddressEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RcptAddressEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RcptAddressEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RcptAddressEntry) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


