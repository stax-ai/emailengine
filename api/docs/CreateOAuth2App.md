# CreateOAuth2App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | **string** | Microsoft tenant configuration (common, organizations, consumers, or tenant ID) | 
**BaseScopes** | Pointer to [**Model4**](Model4.md) |  | [optional] 
**ClientId** | Pointer to **string** | OAuth2 client ID from the provider | [optional] 
**ClientSecret** | Pointer to **string** | OAuth2 client secret from the provider | [optional] 
**Cloud** | Pointer to **string** | Microsoft Azure cloud environment | [optional] 
**Description** | Pointer to **string** | Detailed description of the application | [optional] 
**Enabled** | Pointer to **bool** | Whether this OAuth2 app is active | [optional] [default to false]
**ExtraScopes** | Pointer to **[]string** |  | [optional] 
**GoogleProjectId** | Pointer to [**NullableGoogleProjectId**](GoogleProjectId.md) |  | [optional] 
**GoogleSubscriptionName** | Pointer to [**NullableGoogleSubscriptionName**](GoogleSubscriptionName.md) |  | [optional] 
**GoogleTopicName** | Pointer to [**NullableGoogleTopicName**](GoogleTopicName.md) |  | [optional] 
**GoogleWorkspaceAccounts** | Pointer to **bool** | Restrict OAuth2 login to Google Workspace accounts only | [optional] 
**Name** | **string** | Display name for the OAuth2 application | 
**Provider** | [**Provider**](Provider.md) |  | 
**PubSubApp** | Pointer to [**NullablePubSubApp**](PubSubApp.md) |  | [optional] 
**RedirectUrl** | Pointer to **string** | OAuth2 redirect URI configured in the provider | [optional] 
**ServiceClient** | **string** | Service account unique ID (for 2-legged OAuth2) | 
**ServiceClientEmail** | **string** | Service account email address | 
**ServiceKey** | **string** | Service account private key in PEM format | 
**SkipScopes** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** | Short title shown on the OAuth2 button | [optional] 

## Methods

### NewCreateOAuth2App

`func NewCreateOAuth2App(authority string, name string, provider Provider, serviceClient string, serviceClientEmail string, serviceKey string, ) *CreateOAuth2App`

NewCreateOAuth2App instantiates a new CreateOAuth2App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOAuth2AppWithDefaults

`func NewCreateOAuth2AppWithDefaults() *CreateOAuth2App`

NewCreateOAuth2AppWithDefaults instantiates a new CreateOAuth2App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *CreateOAuth2App) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *CreateOAuth2App) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *CreateOAuth2App) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetBaseScopes

`func (o *CreateOAuth2App) GetBaseScopes() Model4`

GetBaseScopes returns the BaseScopes field if non-nil, zero value otherwise.

### GetBaseScopesOk

`func (o *CreateOAuth2App) GetBaseScopesOk() (*Model4, bool)`

GetBaseScopesOk returns a tuple with the BaseScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScopes

`func (o *CreateOAuth2App) SetBaseScopes(v Model4)`

SetBaseScopes sets BaseScopes field to given value.

### HasBaseScopes

`func (o *CreateOAuth2App) HasBaseScopes() bool`

HasBaseScopes returns a boolean if a field has been set.

### GetClientId

`func (o *CreateOAuth2App) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateOAuth2App) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateOAuth2App) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateOAuth2App) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CreateOAuth2App) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateOAuth2App) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateOAuth2App) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CreateOAuth2App) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCloud

`func (o *CreateOAuth2App) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CreateOAuth2App) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CreateOAuth2App) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CreateOAuth2App) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOAuth2App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOAuth2App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOAuth2App) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOAuth2App) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateOAuth2App) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOAuth2App) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOAuth2App) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateOAuth2App) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraScopes

`func (o *CreateOAuth2App) GetExtraScopes() []string`

GetExtraScopes returns the ExtraScopes field if non-nil, zero value otherwise.

### GetExtraScopesOk

`func (o *CreateOAuth2App) GetExtraScopesOk() (*[]string, bool)`

GetExtraScopesOk returns a tuple with the ExtraScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraScopes

`func (o *CreateOAuth2App) SetExtraScopes(v []string)`

SetExtraScopes sets ExtraScopes field to given value.

### HasExtraScopes

`func (o *CreateOAuth2App) HasExtraScopes() bool`

HasExtraScopes returns a boolean if a field has been set.

### GetGoogleProjectId

`func (o *CreateOAuth2App) GetGoogleProjectId() GoogleProjectId`

GetGoogleProjectId returns the GoogleProjectId field if non-nil, zero value otherwise.

### GetGoogleProjectIdOk

`func (o *CreateOAuth2App) GetGoogleProjectIdOk() (*GoogleProjectId, bool)`

GetGoogleProjectIdOk returns a tuple with the GoogleProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProjectId

`func (o *CreateOAuth2App) SetGoogleProjectId(v GoogleProjectId)`

SetGoogleProjectId sets GoogleProjectId field to given value.

### HasGoogleProjectId

`func (o *CreateOAuth2App) HasGoogleProjectId() bool`

HasGoogleProjectId returns a boolean if a field has been set.

### SetGoogleProjectIdNil

`func (o *CreateOAuth2App) SetGoogleProjectIdNil(b bool)`

 SetGoogleProjectIdNil sets the value for GoogleProjectId to be an explicit nil

### UnsetGoogleProjectId
`func (o *CreateOAuth2App) UnsetGoogleProjectId()`

UnsetGoogleProjectId ensures that no value is present for GoogleProjectId, not even an explicit nil
### GetGoogleSubscriptionName

`func (o *CreateOAuth2App) GetGoogleSubscriptionName() GoogleSubscriptionName`

GetGoogleSubscriptionName returns the GoogleSubscriptionName field if non-nil, zero value otherwise.

### GetGoogleSubscriptionNameOk

`func (o *CreateOAuth2App) GetGoogleSubscriptionNameOk() (*GoogleSubscriptionName, bool)`

GetGoogleSubscriptionNameOk returns a tuple with the GoogleSubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleSubscriptionName

`func (o *CreateOAuth2App) SetGoogleSubscriptionName(v GoogleSubscriptionName)`

SetGoogleSubscriptionName sets GoogleSubscriptionName field to given value.

### HasGoogleSubscriptionName

`func (o *CreateOAuth2App) HasGoogleSubscriptionName() bool`

HasGoogleSubscriptionName returns a boolean if a field has been set.

### SetGoogleSubscriptionNameNil

`func (o *CreateOAuth2App) SetGoogleSubscriptionNameNil(b bool)`

 SetGoogleSubscriptionNameNil sets the value for GoogleSubscriptionName to be an explicit nil

### UnsetGoogleSubscriptionName
`func (o *CreateOAuth2App) UnsetGoogleSubscriptionName()`

UnsetGoogleSubscriptionName ensures that no value is present for GoogleSubscriptionName, not even an explicit nil
### GetGoogleTopicName

`func (o *CreateOAuth2App) GetGoogleTopicName() GoogleTopicName`

GetGoogleTopicName returns the GoogleTopicName field if non-nil, zero value otherwise.

### GetGoogleTopicNameOk

`func (o *CreateOAuth2App) GetGoogleTopicNameOk() (*GoogleTopicName, bool)`

GetGoogleTopicNameOk returns a tuple with the GoogleTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTopicName

`func (o *CreateOAuth2App) SetGoogleTopicName(v GoogleTopicName)`

SetGoogleTopicName sets GoogleTopicName field to given value.

### HasGoogleTopicName

`func (o *CreateOAuth2App) HasGoogleTopicName() bool`

HasGoogleTopicName returns a boolean if a field has been set.

### SetGoogleTopicNameNil

`func (o *CreateOAuth2App) SetGoogleTopicNameNil(b bool)`

 SetGoogleTopicNameNil sets the value for GoogleTopicName to be an explicit nil

### UnsetGoogleTopicName
`func (o *CreateOAuth2App) UnsetGoogleTopicName()`

UnsetGoogleTopicName ensures that no value is present for GoogleTopicName, not even an explicit nil
### GetGoogleWorkspaceAccounts

`func (o *CreateOAuth2App) GetGoogleWorkspaceAccounts() bool`

GetGoogleWorkspaceAccounts returns the GoogleWorkspaceAccounts field if non-nil, zero value otherwise.

### GetGoogleWorkspaceAccountsOk

`func (o *CreateOAuth2App) GetGoogleWorkspaceAccountsOk() (*bool, bool)`

GetGoogleWorkspaceAccountsOk returns a tuple with the GoogleWorkspaceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleWorkspaceAccounts

`func (o *CreateOAuth2App) SetGoogleWorkspaceAccounts(v bool)`

SetGoogleWorkspaceAccounts sets GoogleWorkspaceAccounts field to given value.

### HasGoogleWorkspaceAccounts

`func (o *CreateOAuth2App) HasGoogleWorkspaceAccounts() bool`

HasGoogleWorkspaceAccounts returns a boolean if a field has been set.

### GetName

`func (o *CreateOAuth2App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOAuth2App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOAuth2App) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *CreateOAuth2App) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateOAuth2App) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateOAuth2App) SetProvider(v Provider)`

SetProvider sets Provider field to given value.


### GetPubSubApp

`func (o *CreateOAuth2App) GetPubSubApp() PubSubApp`

GetPubSubApp returns the PubSubApp field if non-nil, zero value otherwise.

### GetPubSubAppOk

`func (o *CreateOAuth2App) GetPubSubAppOk() (*PubSubApp, bool)`

GetPubSubAppOk returns a tuple with the PubSubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubSubApp

`func (o *CreateOAuth2App) SetPubSubApp(v PubSubApp)`

SetPubSubApp sets PubSubApp field to given value.

### HasPubSubApp

`func (o *CreateOAuth2App) HasPubSubApp() bool`

HasPubSubApp returns a boolean if a field has been set.

### SetPubSubAppNil

`func (o *CreateOAuth2App) SetPubSubAppNil(b bool)`

 SetPubSubAppNil sets the value for PubSubApp to be an explicit nil

### UnsetPubSubApp
`func (o *CreateOAuth2App) UnsetPubSubApp()`

UnsetPubSubApp ensures that no value is present for PubSubApp, not even an explicit nil
### GetRedirectUrl

`func (o *CreateOAuth2App) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateOAuth2App) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateOAuth2App) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateOAuth2App) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetServiceClient

`func (o *CreateOAuth2App) GetServiceClient() string`

GetServiceClient returns the ServiceClient field if non-nil, zero value otherwise.

### GetServiceClientOk

`func (o *CreateOAuth2App) GetServiceClientOk() (*string, bool)`

GetServiceClientOk returns a tuple with the ServiceClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClient

`func (o *CreateOAuth2App) SetServiceClient(v string)`

SetServiceClient sets ServiceClient field to given value.


### GetServiceClientEmail

`func (o *CreateOAuth2App) GetServiceClientEmail() string`

GetServiceClientEmail returns the ServiceClientEmail field if non-nil, zero value otherwise.

### GetServiceClientEmailOk

`func (o *CreateOAuth2App) GetServiceClientEmailOk() (*string, bool)`

GetServiceClientEmailOk returns a tuple with the ServiceClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClientEmail

`func (o *CreateOAuth2App) SetServiceClientEmail(v string)`

SetServiceClientEmail sets ServiceClientEmail field to given value.


### GetServiceKey

`func (o *CreateOAuth2App) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *CreateOAuth2App) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *CreateOAuth2App) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetSkipScopes

`func (o *CreateOAuth2App) GetSkipScopes() []string`

GetSkipScopes returns the SkipScopes field if non-nil, zero value otherwise.

### GetSkipScopesOk

`func (o *CreateOAuth2App) GetSkipScopesOk() (*[]string, bool)`

GetSkipScopesOk returns a tuple with the SkipScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipScopes

`func (o *CreateOAuth2App) SetSkipScopes(v []string)`

SetSkipScopes sets SkipScopes field to given value.

### HasSkipScopes

`func (o *CreateOAuth2App) HasSkipScopes() bool`

HasSkipScopes returns a boolean if a field has been set.

### GetTitle

`func (o *CreateOAuth2App) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateOAuth2App) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateOAuth2App) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateOAuth2App) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


