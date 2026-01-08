# OAuth2ResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** | Authorization tenant value for Outlook OAuth2 applications | [optional] 
**ClientId** | Pointer to **string** | Client or Application ID for 3-legged OAuth2 applications | [optional] 
**ClientSecret** | Pointer to **string** | Client secret for 3-legged OAuth2 applications. Actual value is not revealed. | [optional] 
**Created** | **time.Time** | The time this entry was added | 
**Description** | Pointer to **string** | OAuth2 application description | [optional] 
**Enabled** | Pointer to **bool** | Is the application enabled | [optional] 
**GoogleProjectId** | Pointer to [**NullableGoogleProjectId**](GoogleProjectId.md) |  | [optional] 
**GoogleSubscriptionName** | Pointer to [**NullableGoogleSubscriptionName**](GoogleSubscriptionName.md) |  | [optional] 
**GoogleTopicName** | Pointer to [**NullableGoogleTopicName**](GoogleTopicName.md) |  | [optional] 
**GoogleWorkspaceAccounts** | Pointer to **bool** | Restrict OAuth2 login to Google Workspace accounts only | [optional] 
**Id** | **string** | OAuth2 application ID | 
**IncludeInListing** | Pointer to **bool** | Is the application listed in the hosted authentication form | [optional] 
**LastError** | Pointer to [**NullableAccountErrorEntry**](AccountErrorEntry.md) |  | [optional] 
**Legacy** | Pointer to **bool** | &#x60;true&#x60; for older OAuth2 apps set via the settings endpoint | [optional] 
**Name** | Pointer to **string** | Display name for the app | [optional] 
**Provider** | [**OAuth2Provider**](OAuth2Provider.md) |  | 
**RedirectUrl** | Pointer to **string** | Redirect URL for 3-legged OAuth2 applications | [optional] 
**ServiceClient** | Pointer to **string** | Service client ID for 2-legged OAuth2 applications | [optional] 
**ServiceClientEmail** | Pointer to **string** | Service Client Email for 2-legged OAuth2 applications | [optional] 
**ServiceKey** | Pointer to **string** | PEM formatted service secret for 2-legged OAuth2 applications. Actual value is not revealed. | [optional] 
**Title** | Pointer to **string** | Title for the application button | [optional] 
**Updated** | Pointer to **time.Time** | The time this entry was updated | [optional] 

## Methods

### NewOAuth2ResponseItem

`func NewOAuth2ResponseItem(created time.Time, id string, provider OAuth2Provider, ) *OAuth2ResponseItem`

NewOAuth2ResponseItem instantiates a new OAuth2ResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ResponseItemWithDefaults

`func NewOAuth2ResponseItemWithDefaults() *OAuth2ResponseItem`

NewOAuth2ResponseItemWithDefaults instantiates a new OAuth2ResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *OAuth2ResponseItem) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *OAuth2ResponseItem) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *OAuth2ResponseItem) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *OAuth2ResponseItem) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetClientId

`func (o *OAuth2ResponseItem) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2ResponseItem) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2ResponseItem) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2ResponseItem) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuth2ResponseItem) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2ResponseItem) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2ResponseItem) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2ResponseItem) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCreated

`func (o *OAuth2ResponseItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OAuth2ResponseItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OAuth2ResponseItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDescription

`func (o *OAuth2ResponseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuth2ResponseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuth2ResponseItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OAuth2ResponseItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OAuth2ResponseItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OAuth2ResponseItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OAuth2ResponseItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OAuth2ResponseItem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGoogleProjectId

`func (o *OAuth2ResponseItem) GetGoogleProjectId() GoogleProjectId`

GetGoogleProjectId returns the GoogleProjectId field if non-nil, zero value otherwise.

### GetGoogleProjectIdOk

`func (o *OAuth2ResponseItem) GetGoogleProjectIdOk() (*GoogleProjectId, bool)`

GetGoogleProjectIdOk returns a tuple with the GoogleProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProjectId

`func (o *OAuth2ResponseItem) SetGoogleProjectId(v GoogleProjectId)`

SetGoogleProjectId sets GoogleProjectId field to given value.

### HasGoogleProjectId

`func (o *OAuth2ResponseItem) HasGoogleProjectId() bool`

HasGoogleProjectId returns a boolean if a field has been set.

### SetGoogleProjectIdNil

`func (o *OAuth2ResponseItem) SetGoogleProjectIdNil(b bool)`

 SetGoogleProjectIdNil sets the value for GoogleProjectId to be an explicit nil

### UnsetGoogleProjectId
`func (o *OAuth2ResponseItem) UnsetGoogleProjectId()`

UnsetGoogleProjectId ensures that no value is present for GoogleProjectId, not even an explicit nil
### GetGoogleSubscriptionName

`func (o *OAuth2ResponseItem) GetGoogleSubscriptionName() GoogleSubscriptionName`

GetGoogleSubscriptionName returns the GoogleSubscriptionName field if non-nil, zero value otherwise.

### GetGoogleSubscriptionNameOk

`func (o *OAuth2ResponseItem) GetGoogleSubscriptionNameOk() (*GoogleSubscriptionName, bool)`

GetGoogleSubscriptionNameOk returns a tuple with the GoogleSubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleSubscriptionName

`func (o *OAuth2ResponseItem) SetGoogleSubscriptionName(v GoogleSubscriptionName)`

SetGoogleSubscriptionName sets GoogleSubscriptionName field to given value.

### HasGoogleSubscriptionName

`func (o *OAuth2ResponseItem) HasGoogleSubscriptionName() bool`

HasGoogleSubscriptionName returns a boolean if a field has been set.

### SetGoogleSubscriptionNameNil

`func (o *OAuth2ResponseItem) SetGoogleSubscriptionNameNil(b bool)`

 SetGoogleSubscriptionNameNil sets the value for GoogleSubscriptionName to be an explicit nil

### UnsetGoogleSubscriptionName
`func (o *OAuth2ResponseItem) UnsetGoogleSubscriptionName()`

UnsetGoogleSubscriptionName ensures that no value is present for GoogleSubscriptionName, not even an explicit nil
### GetGoogleTopicName

`func (o *OAuth2ResponseItem) GetGoogleTopicName() GoogleTopicName`

GetGoogleTopicName returns the GoogleTopicName field if non-nil, zero value otherwise.

### GetGoogleTopicNameOk

`func (o *OAuth2ResponseItem) GetGoogleTopicNameOk() (*GoogleTopicName, bool)`

GetGoogleTopicNameOk returns a tuple with the GoogleTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTopicName

`func (o *OAuth2ResponseItem) SetGoogleTopicName(v GoogleTopicName)`

SetGoogleTopicName sets GoogleTopicName field to given value.

### HasGoogleTopicName

`func (o *OAuth2ResponseItem) HasGoogleTopicName() bool`

HasGoogleTopicName returns a boolean if a field has been set.

### SetGoogleTopicNameNil

`func (o *OAuth2ResponseItem) SetGoogleTopicNameNil(b bool)`

 SetGoogleTopicNameNil sets the value for GoogleTopicName to be an explicit nil

### UnsetGoogleTopicName
`func (o *OAuth2ResponseItem) UnsetGoogleTopicName()`

UnsetGoogleTopicName ensures that no value is present for GoogleTopicName, not even an explicit nil
### GetGoogleWorkspaceAccounts

`func (o *OAuth2ResponseItem) GetGoogleWorkspaceAccounts() bool`

GetGoogleWorkspaceAccounts returns the GoogleWorkspaceAccounts field if non-nil, zero value otherwise.

### GetGoogleWorkspaceAccountsOk

`func (o *OAuth2ResponseItem) GetGoogleWorkspaceAccountsOk() (*bool, bool)`

GetGoogleWorkspaceAccountsOk returns a tuple with the GoogleWorkspaceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleWorkspaceAccounts

`func (o *OAuth2ResponseItem) SetGoogleWorkspaceAccounts(v bool)`

SetGoogleWorkspaceAccounts sets GoogleWorkspaceAccounts field to given value.

### HasGoogleWorkspaceAccounts

`func (o *OAuth2ResponseItem) HasGoogleWorkspaceAccounts() bool`

HasGoogleWorkspaceAccounts returns a boolean if a field has been set.

### GetId

`func (o *OAuth2ResponseItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2ResponseItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2ResponseItem) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeInListing

`func (o *OAuth2ResponseItem) GetIncludeInListing() bool`

GetIncludeInListing returns the IncludeInListing field if non-nil, zero value otherwise.

### GetIncludeInListingOk

`func (o *OAuth2ResponseItem) GetIncludeInListingOk() (*bool, bool)`

GetIncludeInListingOk returns a tuple with the IncludeInListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInListing

`func (o *OAuth2ResponseItem) SetIncludeInListing(v bool)`

SetIncludeInListing sets IncludeInListing field to given value.

### HasIncludeInListing

`func (o *OAuth2ResponseItem) HasIncludeInListing() bool`

HasIncludeInListing returns a boolean if a field has been set.

### GetLastError

`func (o *OAuth2ResponseItem) GetLastError() AccountErrorEntry`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *OAuth2ResponseItem) GetLastErrorOk() (*AccountErrorEntry, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *OAuth2ResponseItem) SetLastError(v AccountErrorEntry)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *OAuth2ResponseItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *OAuth2ResponseItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *OAuth2ResponseItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLegacy

`func (o *OAuth2ResponseItem) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *OAuth2ResponseItem) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *OAuth2ResponseItem) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *OAuth2ResponseItem) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetName

`func (o *OAuth2ResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2ResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2ResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OAuth2ResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *OAuth2ResponseItem) GetProvider() OAuth2Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OAuth2ResponseItem) GetProviderOk() (*OAuth2Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OAuth2ResponseItem) SetProvider(v OAuth2Provider)`

SetProvider sets Provider field to given value.


### GetRedirectUrl

`func (o *OAuth2ResponseItem) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *OAuth2ResponseItem) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *OAuth2ResponseItem) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *OAuth2ResponseItem) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetServiceClient

`func (o *OAuth2ResponseItem) GetServiceClient() string`

GetServiceClient returns the ServiceClient field if non-nil, zero value otherwise.

### GetServiceClientOk

`func (o *OAuth2ResponseItem) GetServiceClientOk() (*string, bool)`

GetServiceClientOk returns a tuple with the ServiceClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClient

`func (o *OAuth2ResponseItem) SetServiceClient(v string)`

SetServiceClient sets ServiceClient field to given value.

### HasServiceClient

`func (o *OAuth2ResponseItem) HasServiceClient() bool`

HasServiceClient returns a boolean if a field has been set.

### GetServiceClientEmail

`func (o *OAuth2ResponseItem) GetServiceClientEmail() string`

GetServiceClientEmail returns the ServiceClientEmail field if non-nil, zero value otherwise.

### GetServiceClientEmailOk

`func (o *OAuth2ResponseItem) GetServiceClientEmailOk() (*string, bool)`

GetServiceClientEmailOk returns a tuple with the ServiceClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClientEmail

`func (o *OAuth2ResponseItem) SetServiceClientEmail(v string)`

SetServiceClientEmail sets ServiceClientEmail field to given value.

### HasServiceClientEmail

`func (o *OAuth2ResponseItem) HasServiceClientEmail() bool`

HasServiceClientEmail returns a boolean if a field has been set.

### GetServiceKey

`func (o *OAuth2ResponseItem) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *OAuth2ResponseItem) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *OAuth2ResponseItem) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *OAuth2ResponseItem) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetTitle

`func (o *OAuth2ResponseItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OAuth2ResponseItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OAuth2ResponseItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OAuth2ResponseItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *OAuth2ResponseItem) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *OAuth2ResponseItem) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *OAuth2ResponseItem) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *OAuth2ResponseItem) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


