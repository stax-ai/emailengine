# GatewayResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deliveries** | Pointer to **int32** | Count of email deliveries using this gateway | [optional] 
**Gateway** | **string** | Gateway ID | 
**LastError** | Pointer to [**NullableAccountErrorEntry**](AccountErrorEntry.md) |  | [optional] 
**LastUse** | Pointer to **time.Time** | Last delivery time | [optional] 
**Name** | Pointer to **string** | Display name for the gateway | [optional] 

## Methods

### NewGatewayResponseItem

`func NewGatewayResponseItem(gateway string, ) *GatewayResponseItem`

NewGatewayResponseItem instantiates a new GatewayResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayResponseItemWithDefaults

`func NewGatewayResponseItemWithDefaults() *GatewayResponseItem`

NewGatewayResponseItemWithDefaults instantiates a new GatewayResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveries

`func (o *GatewayResponseItem) GetDeliveries() int32`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *GatewayResponseItem) GetDeliveriesOk() (*int32, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *GatewayResponseItem) SetDeliveries(v int32)`

SetDeliveries sets Deliveries field to given value.

### HasDeliveries

`func (o *GatewayResponseItem) HasDeliveries() bool`

HasDeliveries returns a boolean if a field has been set.

### GetGateway

`func (o *GatewayResponseItem) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GatewayResponseItem) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GatewayResponseItem) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetLastError

`func (o *GatewayResponseItem) GetLastError() AccountErrorEntry`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GatewayResponseItem) GetLastErrorOk() (*AccountErrorEntry, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GatewayResponseItem) SetLastError(v AccountErrorEntry)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GatewayResponseItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GatewayResponseItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GatewayResponseItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastUse

`func (o *GatewayResponseItem) GetLastUse() time.Time`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *GatewayResponseItem) GetLastUseOk() (*time.Time, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *GatewayResponseItem) SetLastUse(v time.Time)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *GatewayResponseItem) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.

### GetName

`func (o *GatewayResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


