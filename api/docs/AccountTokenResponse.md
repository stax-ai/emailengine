# AccountTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signatures** | Pointer to [**[]SignatureResponseItem**](SignatureResponseItem.md) |  | [optional] 

## Methods

### NewAccountTokenResponse

`func NewAccountTokenResponse() *AccountTokenResponse`

NewAccountTokenResponse instantiates a new AccountTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokenResponseWithDefaults

`func NewAccountTokenResponseWithDefaults() *AccountTokenResponse`

NewAccountTokenResponseWithDefaults instantiates a new AccountTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatures

`func (o *AccountTokenResponse) GetSignatures() []SignatureResponseItem`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *AccountTokenResponse) GetSignaturesOk() (*[]SignatureResponseItem, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *AccountTokenResponse) SetSignatures(v []SignatureResponseItem)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *AccountTokenResponse) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


