# UpdateOAuthApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** | Authorization tenant value for Outlook OAuth2 applications | [optional] 
**ClientId** | Pointer to [**NullableClientId**](ClientId.md) |  | [optional] 
**ClientSecret** | Pointer to **string** | Client secret for 3-legged OAuth2 applications | [optional] 
**Cloud** | Pointer to [**AzureCloud**](AzureCloud.md) |  | [optional] 
**Description** | Pointer to **string** | Application description | [optional] 
**Enabled** | Pointer to **bool** | Enable this app | [optional] 
**ExtraScopes** | Pointer to **[]string** | OAuth2 Extra Scopes | [optional] 
**GoogleProjectId** | Pointer to [**NullableGoogleProjectId**](GoogleProjectId.md) |  | [optional] 
**GoogleSubscriptionName** | Pointer to [**NullableGoogleSubscriptionName**](GoogleSubscriptionName.md) |  | [optional] 
**GoogleTopicName** | Pointer to [**NullableGoogleTopicName**](GoogleTopicName.md) |  | [optional] 
**GoogleWorkspaceAccounts** | Pointer to **bool** | Restrict OAuth2 login to Google Workspace accounts only | [optional] 
**Name** | Pointer to **string** | Application name | [optional] 
**PubSubApp** | Pointer to **string** | Cloud Pub/Sub app for Gmail API webhooks | [optional] 
**RedirectUrl** | Pointer to [**NullableRedirectUrl**](RedirectUrl.md) |  | [optional] 
**ServiceClient** | Pointer to [**NullableServiceClient**](ServiceClient.md) |  | [optional] 
**ServiceClientEmail** | Pointer to **string** | Service Client Email for 2-legged OAuth2 applications | [optional] 
**ServiceKey** | Pointer to **string** | PEM formatted service secret for 2-legged OAuth2 applications | [optional] 
**SkipScopes** | Pointer to **[]string** | OAuth2 scopes to skip from the base set | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | Title for the application button | [optional] 

## Methods

### NewUpdateOAuthApp

`func NewUpdateOAuthApp() *UpdateOAuthApp`

NewUpdateOAuthApp instantiates a new UpdateOAuthApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOAuthAppWithDefaults

`func NewUpdateOAuthAppWithDefaults() *UpdateOAuthApp`

NewUpdateOAuthAppWithDefaults instantiates a new UpdateOAuthApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *UpdateOAuthApp) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *UpdateOAuthApp) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *UpdateOAuthApp) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *UpdateOAuthApp) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetClientId

`func (o *UpdateOAuthApp) GetClientId() ClientId`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateOAuthApp) GetClientIdOk() (*ClientId, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateOAuthApp) SetClientId(v ClientId)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UpdateOAuthApp) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *UpdateOAuthApp) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *UpdateOAuthApp) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *UpdateOAuthApp) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *UpdateOAuthApp) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *UpdateOAuthApp) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *UpdateOAuthApp) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCloud

`func (o *UpdateOAuthApp) GetCloud() AzureCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *UpdateOAuthApp) GetCloudOk() (*AzureCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *UpdateOAuthApp) SetCloud(v AzureCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *UpdateOAuthApp) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOAuthApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOAuthApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOAuthApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOAuthApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateOAuthApp) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateOAuthApp) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateOAuthApp) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateOAuthApp) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraScopes

`func (o *UpdateOAuthApp) GetExtraScopes() []string`

GetExtraScopes returns the ExtraScopes field if non-nil, zero value otherwise.

### GetExtraScopesOk

`func (o *UpdateOAuthApp) GetExtraScopesOk() (*[]string, bool)`

GetExtraScopesOk returns a tuple with the ExtraScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraScopes

`func (o *UpdateOAuthApp) SetExtraScopes(v []string)`

SetExtraScopes sets ExtraScopes field to given value.

### HasExtraScopes

`func (o *UpdateOAuthApp) HasExtraScopes() bool`

HasExtraScopes returns a boolean if a field has been set.

### GetGoogleProjectId

`func (o *UpdateOAuthApp) GetGoogleProjectId() GoogleProjectId`

GetGoogleProjectId returns the GoogleProjectId field if non-nil, zero value otherwise.

### GetGoogleProjectIdOk

`func (o *UpdateOAuthApp) GetGoogleProjectIdOk() (*GoogleProjectId, bool)`

GetGoogleProjectIdOk returns a tuple with the GoogleProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProjectId

`func (o *UpdateOAuthApp) SetGoogleProjectId(v GoogleProjectId)`

SetGoogleProjectId sets GoogleProjectId field to given value.

### HasGoogleProjectId

`func (o *UpdateOAuthApp) HasGoogleProjectId() bool`

HasGoogleProjectId returns a boolean if a field has been set.

### SetGoogleProjectIdNil

`func (o *UpdateOAuthApp) SetGoogleProjectIdNil(b bool)`

 SetGoogleProjectIdNil sets the value for GoogleProjectId to be an explicit nil

### UnsetGoogleProjectId
`func (o *UpdateOAuthApp) UnsetGoogleProjectId()`

UnsetGoogleProjectId ensures that no value is present for GoogleProjectId, not even an explicit nil
### GetGoogleSubscriptionName

`func (o *UpdateOAuthApp) GetGoogleSubscriptionName() GoogleSubscriptionName`

GetGoogleSubscriptionName returns the GoogleSubscriptionName field if non-nil, zero value otherwise.

### GetGoogleSubscriptionNameOk

`func (o *UpdateOAuthApp) GetGoogleSubscriptionNameOk() (*GoogleSubscriptionName, bool)`

GetGoogleSubscriptionNameOk returns a tuple with the GoogleSubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleSubscriptionName

`func (o *UpdateOAuthApp) SetGoogleSubscriptionName(v GoogleSubscriptionName)`

SetGoogleSubscriptionName sets GoogleSubscriptionName field to given value.

### HasGoogleSubscriptionName

`func (o *UpdateOAuthApp) HasGoogleSubscriptionName() bool`

HasGoogleSubscriptionName returns a boolean if a field has been set.

### SetGoogleSubscriptionNameNil

`func (o *UpdateOAuthApp) SetGoogleSubscriptionNameNil(b bool)`

 SetGoogleSubscriptionNameNil sets the value for GoogleSubscriptionName to be an explicit nil

### UnsetGoogleSubscriptionName
`func (o *UpdateOAuthApp) UnsetGoogleSubscriptionName()`

UnsetGoogleSubscriptionName ensures that no value is present for GoogleSubscriptionName, not even an explicit nil
### GetGoogleTopicName

`func (o *UpdateOAuthApp) GetGoogleTopicName() GoogleTopicName`

GetGoogleTopicName returns the GoogleTopicName field if non-nil, zero value otherwise.

### GetGoogleTopicNameOk

`func (o *UpdateOAuthApp) GetGoogleTopicNameOk() (*GoogleTopicName, bool)`

GetGoogleTopicNameOk returns a tuple with the GoogleTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTopicName

`func (o *UpdateOAuthApp) SetGoogleTopicName(v GoogleTopicName)`

SetGoogleTopicName sets GoogleTopicName field to given value.

### HasGoogleTopicName

`func (o *UpdateOAuthApp) HasGoogleTopicName() bool`

HasGoogleTopicName returns a boolean if a field has been set.

### SetGoogleTopicNameNil

`func (o *UpdateOAuthApp) SetGoogleTopicNameNil(b bool)`

 SetGoogleTopicNameNil sets the value for GoogleTopicName to be an explicit nil

### UnsetGoogleTopicName
`func (o *UpdateOAuthApp) UnsetGoogleTopicName()`

UnsetGoogleTopicName ensures that no value is present for GoogleTopicName, not even an explicit nil
### GetGoogleWorkspaceAccounts

`func (o *UpdateOAuthApp) GetGoogleWorkspaceAccounts() bool`

GetGoogleWorkspaceAccounts returns the GoogleWorkspaceAccounts field if non-nil, zero value otherwise.

### GetGoogleWorkspaceAccountsOk

`func (o *UpdateOAuthApp) GetGoogleWorkspaceAccountsOk() (*bool, bool)`

GetGoogleWorkspaceAccountsOk returns a tuple with the GoogleWorkspaceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleWorkspaceAccounts

`func (o *UpdateOAuthApp) SetGoogleWorkspaceAccounts(v bool)`

SetGoogleWorkspaceAccounts sets GoogleWorkspaceAccounts field to given value.

### HasGoogleWorkspaceAccounts

`func (o *UpdateOAuthApp) HasGoogleWorkspaceAccounts() bool`

HasGoogleWorkspaceAccounts returns a boolean if a field has been set.

### GetName

`func (o *UpdateOAuthApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOAuthApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOAuthApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOAuthApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPubSubApp

`func (o *UpdateOAuthApp) GetPubSubApp() string`

GetPubSubApp returns the PubSubApp field if non-nil, zero value otherwise.

### GetPubSubAppOk

`func (o *UpdateOAuthApp) GetPubSubAppOk() (*string, bool)`

GetPubSubAppOk returns a tuple with the PubSubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubSubApp

`func (o *UpdateOAuthApp) SetPubSubApp(v string)`

SetPubSubApp sets PubSubApp field to given value.

### HasPubSubApp

`func (o *UpdateOAuthApp) HasPubSubApp() bool`

HasPubSubApp returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UpdateOAuthApp) GetRedirectUrl() RedirectUrl`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateOAuthApp) GetRedirectUrlOk() (*RedirectUrl, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateOAuthApp) SetRedirectUrl(v RedirectUrl)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateOAuthApp) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *UpdateOAuthApp) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *UpdateOAuthApp) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetServiceClient

`func (o *UpdateOAuthApp) GetServiceClient() ServiceClient`

GetServiceClient returns the ServiceClient field if non-nil, zero value otherwise.

### GetServiceClientOk

`func (o *UpdateOAuthApp) GetServiceClientOk() (*ServiceClient, bool)`

GetServiceClientOk returns a tuple with the ServiceClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClient

`func (o *UpdateOAuthApp) SetServiceClient(v ServiceClient)`

SetServiceClient sets ServiceClient field to given value.

### HasServiceClient

`func (o *UpdateOAuthApp) HasServiceClient() bool`

HasServiceClient returns a boolean if a field has been set.

### SetServiceClientNil

`func (o *UpdateOAuthApp) SetServiceClientNil(b bool)`

 SetServiceClientNil sets the value for ServiceClient to be an explicit nil

### UnsetServiceClient
`func (o *UpdateOAuthApp) UnsetServiceClient()`

UnsetServiceClient ensures that no value is present for ServiceClient, not even an explicit nil
### GetServiceClientEmail

`func (o *UpdateOAuthApp) GetServiceClientEmail() string`

GetServiceClientEmail returns the ServiceClientEmail field if non-nil, zero value otherwise.

### GetServiceClientEmailOk

`func (o *UpdateOAuthApp) GetServiceClientEmailOk() (*string, bool)`

GetServiceClientEmailOk returns a tuple with the ServiceClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClientEmail

`func (o *UpdateOAuthApp) SetServiceClientEmail(v string)`

SetServiceClientEmail sets ServiceClientEmail field to given value.

### HasServiceClientEmail

`func (o *UpdateOAuthApp) HasServiceClientEmail() bool`

HasServiceClientEmail returns a boolean if a field has been set.

### GetServiceKey

`func (o *UpdateOAuthApp) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *UpdateOAuthApp) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *UpdateOAuthApp) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *UpdateOAuthApp) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetSkipScopes

`func (o *UpdateOAuthApp) GetSkipScopes() []string`

GetSkipScopes returns the SkipScopes field if non-nil, zero value otherwise.

### GetSkipScopesOk

`func (o *UpdateOAuthApp) GetSkipScopesOk() (*[]string, bool)`

GetSkipScopesOk returns a tuple with the SkipScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipScopes

`func (o *UpdateOAuthApp) SetSkipScopes(v []string)`

SetSkipScopes sets SkipScopes field to given value.

### HasSkipScopes

`func (o *UpdateOAuthApp) HasSkipScopes() bool`

HasSkipScopes returns a boolean if a field has been set.

### GetTenant

`func (o *UpdateOAuthApp) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *UpdateOAuthApp) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *UpdateOAuthApp) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *UpdateOAuthApp) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateOAuthApp) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateOAuthApp) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateOAuthApp) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateOAuthApp) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


