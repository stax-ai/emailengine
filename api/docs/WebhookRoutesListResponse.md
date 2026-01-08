# WebhookRoutesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 
**Webhooks** | Pointer to [**[]WebhookRoutesListEntry**](WebhookRoutesListEntry.md) |  | [optional] 

## Methods

### NewWebhookRoutesListResponse

`func NewWebhookRoutesListResponse() *WebhookRoutesListResponse`

NewWebhookRoutesListResponse instantiates a new WebhookRoutesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRoutesListResponseWithDefaults

`func NewWebhookRoutesListResponseWithDefaults() *WebhookRoutesListResponse`

NewWebhookRoutesListResponseWithDefaults instantiates a new WebhookRoutesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *WebhookRoutesListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WebhookRoutesListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WebhookRoutesListResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *WebhookRoutesListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *WebhookRoutesListResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *WebhookRoutesListResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *WebhookRoutesListResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *WebhookRoutesListResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *WebhookRoutesListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WebhookRoutesListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WebhookRoutesListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WebhookRoutesListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWebhooks

`func (o *WebhookRoutesListResponse) GetWebhooks() []WebhookRoutesListEntry`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookRoutesListResponse) GetWebhooksOk() (*[]WebhookRoutesListEntry, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookRoutesListResponse) SetWebhooks(v []WebhookRoutesListEntry)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *WebhookRoutesListResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


