# DeliveryStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryTest** | Pointer to **string** | Test ID | [optional] 
**Success** | Pointer to **bool** | Was the test started | [optional] 

## Methods

### NewDeliveryStartResponse

`func NewDeliveryStartResponse() *DeliveryStartResponse`

NewDeliveryStartResponse instantiates a new DeliveryStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryStartResponseWithDefaults

`func NewDeliveryStartResponseWithDefaults() *DeliveryStartResponse`

NewDeliveryStartResponseWithDefaults instantiates a new DeliveryStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryTest

`func (o *DeliveryStartResponse) GetDeliveryTest() string`

GetDeliveryTest returns the DeliveryTest field if non-nil, zero value otherwise.

### GetDeliveryTestOk

`func (o *DeliveryStartResponse) GetDeliveryTestOk() (*string, bool)`

GetDeliveryTestOk returns a tuple with the DeliveryTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTest

`func (o *DeliveryStartResponse) SetDeliveryTest(v string)`

SetDeliveryTest sets DeliveryTest field to given value.

### HasDeliveryTest

`func (o *DeliveryStartResponse) HasDeliveryTest() bool`

HasDeliveryTest returns a boolean if a field has been set.

### GetSuccess

`func (o *DeliveryStartResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeliveryStartResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeliveryStartResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DeliveryStartResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


