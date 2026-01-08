# CreateGatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | **string** | Gateway ID | 
**State** | [**CreateGatewayState**](CreateGatewayState.md) |  | 

## Methods

### NewCreateGatewayResponse

`func NewCreateGatewayResponse(gateway string, state CreateGatewayState, ) *CreateGatewayResponse`

NewCreateGatewayResponse instantiates a new CreateGatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGatewayResponseWithDefaults

`func NewCreateGatewayResponseWithDefaults() *CreateGatewayResponse`

NewCreateGatewayResponseWithDefaults instantiates a new CreateGatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *CreateGatewayResponse) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateGatewayResponse) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateGatewayResponse) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetState

`func (o *CreateGatewayResponse) GetState() CreateGatewayState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateGatewayResponse) GetStateOk() (*CreateGatewayState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateGatewayResponse) SetState(v CreateGatewayState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


