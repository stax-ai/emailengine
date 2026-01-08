# GatewaysFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateways** | Pointer to [**[]GatewayResponseItem**](GatewayResponseItem.md) |  | [optional] 
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 

## Methods

### NewGatewaysFilterResponse

`func NewGatewaysFilterResponse() *GatewaysFilterResponse`

NewGatewaysFilterResponse instantiates a new GatewaysFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaysFilterResponseWithDefaults

`func NewGatewaysFilterResponseWithDefaults() *GatewaysFilterResponse`

NewGatewaysFilterResponseWithDefaults instantiates a new GatewaysFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateways

`func (o *GatewaysFilterResponse) GetGateways() []GatewayResponseItem`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *GatewaysFilterResponse) GetGatewaysOk() (*[]GatewayResponseItem, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *GatewaysFilterResponse) SetGateways(v []GatewayResponseItem)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *GatewaysFilterResponse) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetPage

`func (o *GatewaysFilterResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GatewaysFilterResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GatewaysFilterResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GatewaysFilterResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *GatewaysFilterResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GatewaysFilterResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GatewaysFilterResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GatewaysFilterResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *GatewaysFilterResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GatewaysFilterResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GatewaysFilterResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GatewaysFilterResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


