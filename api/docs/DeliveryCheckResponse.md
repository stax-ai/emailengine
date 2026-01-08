# DeliveryCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arc** | Pointer to **map[string]interface{}** | ARC results | [optional] 
**Bimi** | Pointer to **map[string]interface{}** | BIMI results | [optional] 
**Dkim** | Pointer to **map[string]interface{}** | DKIM results | [optional] 
**Dmarc** | Pointer to **map[string]interface{}** | DMARC results | [optional] 
**MainSig** | Pointer to **map[string]interface{}** | Primary DKIM signature. &#x60;status.aligned&#x60; should be set, otherwise DKIM check should not be considered as passed. | [optional] 
**Spf** | Pointer to **map[string]interface{}** | SPF results | [optional] 
**Success** | Pointer to **bool** | Was the test completed | [optional] 

## Methods

### NewDeliveryCheckResponse

`func NewDeliveryCheckResponse() *DeliveryCheckResponse`

NewDeliveryCheckResponse instantiates a new DeliveryCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryCheckResponseWithDefaults

`func NewDeliveryCheckResponseWithDefaults() *DeliveryCheckResponse`

NewDeliveryCheckResponseWithDefaults instantiates a new DeliveryCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArc

`func (o *DeliveryCheckResponse) GetArc() map[string]interface{}`

GetArc returns the Arc field if non-nil, zero value otherwise.

### GetArcOk

`func (o *DeliveryCheckResponse) GetArcOk() (*map[string]interface{}, bool)`

GetArcOk returns a tuple with the Arc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArc

`func (o *DeliveryCheckResponse) SetArc(v map[string]interface{})`

SetArc sets Arc field to given value.

### HasArc

`func (o *DeliveryCheckResponse) HasArc() bool`

HasArc returns a boolean if a field has been set.

### GetBimi

`func (o *DeliveryCheckResponse) GetBimi() map[string]interface{}`

GetBimi returns the Bimi field if non-nil, zero value otherwise.

### GetBimiOk

`func (o *DeliveryCheckResponse) GetBimiOk() (*map[string]interface{}, bool)`

GetBimiOk returns a tuple with the Bimi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBimi

`func (o *DeliveryCheckResponse) SetBimi(v map[string]interface{})`

SetBimi sets Bimi field to given value.

### HasBimi

`func (o *DeliveryCheckResponse) HasBimi() bool`

HasBimi returns a boolean if a field has been set.

### GetDkim

`func (o *DeliveryCheckResponse) GetDkim() map[string]interface{}`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *DeliveryCheckResponse) GetDkimOk() (*map[string]interface{}, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *DeliveryCheckResponse) SetDkim(v map[string]interface{})`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *DeliveryCheckResponse) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetDmarc

`func (o *DeliveryCheckResponse) GetDmarc() map[string]interface{}`

GetDmarc returns the Dmarc field if non-nil, zero value otherwise.

### GetDmarcOk

`func (o *DeliveryCheckResponse) GetDmarcOk() (*map[string]interface{}, bool)`

GetDmarcOk returns a tuple with the Dmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarc

`func (o *DeliveryCheckResponse) SetDmarc(v map[string]interface{})`

SetDmarc sets Dmarc field to given value.

### HasDmarc

`func (o *DeliveryCheckResponse) HasDmarc() bool`

HasDmarc returns a boolean if a field has been set.

### GetMainSig

`func (o *DeliveryCheckResponse) GetMainSig() map[string]interface{}`

GetMainSig returns the MainSig field if non-nil, zero value otherwise.

### GetMainSigOk

`func (o *DeliveryCheckResponse) GetMainSigOk() (*map[string]interface{}, bool)`

GetMainSigOk returns a tuple with the MainSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainSig

`func (o *DeliveryCheckResponse) SetMainSig(v map[string]interface{})`

SetMainSig sets MainSig field to given value.

### HasMainSig

`func (o *DeliveryCheckResponse) HasMainSig() bool`

HasMainSig returns a boolean if a field has been set.

### GetSpf

`func (o *DeliveryCheckResponse) GetSpf() map[string]interface{}`

GetSpf returns the Spf field if non-nil, zero value otherwise.

### GetSpfOk

`func (o *DeliveryCheckResponse) GetSpfOk() (*map[string]interface{}, bool)`

GetSpfOk returns a tuple with the Spf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpf

`func (o *DeliveryCheckResponse) SetSpf(v map[string]interface{})`

SetSpf sets Spf field to given value.

### HasSpf

`func (o *DeliveryCheckResponse) HasSpf() bool`

HasSpf returns a boolean if a field has been set.

### GetSuccess

`func (o *DeliveryCheckResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeliveryCheckResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeliveryCheckResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DeliveryCheckResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


