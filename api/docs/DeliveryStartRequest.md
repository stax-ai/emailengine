# DeliveryStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to [**NullableGateway**](Gateway.md) |  | [optional] 

## Methods

### NewDeliveryStartRequest

`func NewDeliveryStartRequest() *DeliveryStartRequest`

NewDeliveryStartRequest instantiates a new DeliveryStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryStartRequestWithDefaults

`func NewDeliveryStartRequestWithDefaults() *DeliveryStartRequest`

NewDeliveryStartRequestWithDefaults instantiates a new DeliveryStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *DeliveryStartRequest) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DeliveryStartRequest) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DeliveryStartRequest) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DeliveryStartRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *DeliveryStartRequest) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *DeliveryStartRequest) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


