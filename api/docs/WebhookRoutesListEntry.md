# WebhookRoutesListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The time this route was created | [optional] 
**Description** | Pointer to **string** | Optional description of the webhook route | [optional] 
**Enabled** | Pointer to **bool** | Is the route enabled | [optional] 
**Id** | **string** | Webhook ID | 
**Name** | **string** | Name of the route | 
**TargetUrl** | Pointer to **string** | Target URL that will receive webhook notifications via POST requests | [optional] 
**Tcount** | Pointer to **int32** | How many times this route has been applied | [optional] 
**Updated** | Pointer to **time.Time** | The time this route was last updated | [optional] 

## Methods

### NewWebhookRoutesListEntry

`func NewWebhookRoutesListEntry(id string, name string, ) *WebhookRoutesListEntry`

NewWebhookRoutesListEntry instantiates a new WebhookRoutesListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRoutesListEntryWithDefaults

`func NewWebhookRoutesListEntryWithDefaults() *WebhookRoutesListEntry`

NewWebhookRoutesListEntryWithDefaults instantiates a new WebhookRoutesListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *WebhookRoutesListEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookRoutesListEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookRoutesListEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebhookRoutesListEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookRoutesListEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookRoutesListEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookRoutesListEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookRoutesListEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WebhookRoutesListEntry) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookRoutesListEntry) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhookRoutesListEntry) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebhookRoutesListEntry) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *WebhookRoutesListEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookRoutesListEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookRoutesListEntry) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebhookRoutesListEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookRoutesListEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookRoutesListEntry) SetName(v string)`

SetName sets Name field to given value.


### GetTargetUrl

`func (o *WebhookRoutesListEntry) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *WebhookRoutesListEntry) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *WebhookRoutesListEntry) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *WebhookRoutesListEntry) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetTcount

`func (o *WebhookRoutesListEntry) GetTcount() int32`

GetTcount returns the Tcount field if non-nil, zero value otherwise.

### GetTcountOk

`func (o *WebhookRoutesListEntry) GetTcountOk() (*int32, bool)`

GetTcountOk returns a tuple with the Tcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcount

`func (o *WebhookRoutesListEntry) SetTcount(v int32)`

SetTcount sets Tcount field to given value.

### HasTcount

`func (o *WebhookRoutesListEntry) HasTcount() bool`

HasTcount returns a boolean if a field has been set.

### GetUpdated

`func (o *WebhookRoutesListEntry) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *WebhookRoutesListEntry) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *WebhookRoutesListEntry) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *WebhookRoutesListEntry) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


