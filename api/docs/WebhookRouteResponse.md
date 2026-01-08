# WebhookRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**WebhookRouteContent**](WebhookRouteContent.md) |  | [optional] 
**Created** | Pointer to **time.Time** | The time this route was created | [optional] 
**Description** | Pointer to **string** | Optional description of the webhook route | [optional] 
**Enabled** | Pointer to **bool** | Is the route enabled | [optional] 
**Id** | **string** | Webhook ID | 
**Name** | **string** | Name of the route | 
**TargetUrl** | Pointer to **string** | Target URL that will receive webhook notifications via POST requests | [optional] 
**Tcount** | Pointer to **int32** | How many times this route has been applied | [optional] 
**Updated** | Pointer to **time.Time** | The time this route was last updated | [optional] 

## Methods

### NewWebhookRouteResponse

`func NewWebhookRouteResponse(id string, name string, ) *WebhookRouteResponse`

NewWebhookRouteResponse instantiates a new WebhookRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRouteResponseWithDefaults

`func NewWebhookRouteResponseWithDefaults() *WebhookRouteResponse`

NewWebhookRouteResponseWithDefaults instantiates a new WebhookRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *WebhookRouteResponse) GetContent() WebhookRouteContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebhookRouteResponse) GetContentOk() (*WebhookRouteContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebhookRouteResponse) SetContent(v WebhookRouteContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebhookRouteResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *WebhookRouteResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookRouteResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookRouteResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebhookRouteResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookRouteResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookRouteResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookRouteResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookRouteResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WebhookRouteResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookRouteResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhookRouteResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebhookRouteResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *WebhookRouteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookRouteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookRouteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebhookRouteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookRouteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookRouteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetTargetUrl

`func (o *WebhookRouteResponse) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *WebhookRouteResponse) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *WebhookRouteResponse) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *WebhookRouteResponse) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTcount

`func (o *WebhookRouteResponse) GetTcount() int32`

GetTcount returns the Tcount field if non-nil, zero value otherwise.

### GetTcountOk

`func (o *WebhookRouteResponse) GetTcountOk() (*int32, bool)`

GetTcountOk returns a tuple with the Tcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcount

`func (o *WebhookRouteResponse) SetTcount(v int32)`

SetTcount sets Tcount field to given value.

### HasTcount

`func (o *WebhookRouteResponse) HasTcount() bool`

HasTcount returns a boolean if a field has been set.

### GetUpdated

`func (o *WebhookRouteResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *WebhookRouteResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *WebhookRouteResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *WebhookRouteResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


