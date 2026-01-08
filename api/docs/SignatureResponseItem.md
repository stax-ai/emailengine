# SignatureResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Email address associated with the signature | 
**Signature** | **string** | Signature HTML code | 

## Methods

### NewSignatureResponseItem

`func NewSignatureResponseItem(address string, signature string, ) *SignatureResponseItem`

NewSignatureResponseItem instantiates a new SignatureResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureResponseItemWithDefaults

`func NewSignatureResponseItemWithDefaults() *SignatureResponseItem`

NewSignatureResponseItemWithDefaults instantiates a new SignatureResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SignatureResponseItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SignatureResponseItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SignatureResponseItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSignature

`func (o *SignatureResponseItem) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SignatureResponseItem) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SignatureResponseItem) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


