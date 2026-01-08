# ApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int32** | The number of accounts registered with this application. Not available for legacy apps. | [optional] 
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

### NewApplicationResponse

`func NewApplicationResponse(created time.Time, id string, provider OAuth2Provider, ) *ApplicationResponse`

NewApplicationResponse instantiates a new ApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseWithDefaults

`func NewApplicationResponseWithDefaults() *ApplicationResponse`

NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ApplicationResponse) GetAccounts() int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ApplicationResponse) GetAccountsOk() (*int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ApplicationResponse) SetAccounts(v int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ApplicationResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAuthority

`func (o *ApplicationResponse) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *ApplicationResponse) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *ApplicationResponse) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *ApplicationResponse) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetClientId

`func (o *ApplicationResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApplicationResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApplicationResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ApplicationResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ApplicationResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApplicationResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApplicationResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ApplicationResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCreated

`func (o *ApplicationResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDescription

`func (o *ApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGoogleProjectId

`func (o *ApplicationResponse) GetGoogleProjectId() GoogleProjectId`

GetGoogleProjectId returns the GoogleProjectId field if non-nil, zero value otherwise.

### GetGoogleProjectIdOk

`func (o *ApplicationResponse) GetGoogleProjectIdOk() (*GoogleProjectId, bool)`

GetGoogleProjectIdOk returns a tuple with the GoogleProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProjectId

`func (o *ApplicationResponse) SetGoogleProjectId(v GoogleProjectId)`

SetGoogleProjectId sets GoogleProjectId field to given value.

### HasGoogleProjectId

`func (o *ApplicationResponse) HasGoogleProjectId() bool`

HasGoogleProjectId returns a boolean if a field has been set.

### SetGoogleProjectIdNil

`func (o *ApplicationResponse) SetGoogleProjectIdNil(b bool)`

 SetGoogleProjectIdNil sets the value for GoogleProjectId to be an explicit nil

### UnsetGoogleProjectId
`func (o *ApplicationResponse) UnsetGoogleProjectId()`

UnsetGoogleProjectId ensures that no value is present for GoogleProjectId, not even an explicit nil
### GetGoogleSubscriptionName

`func (o *ApplicationResponse) GetGoogleSubscriptionName() GoogleSubscriptionName`

GetGoogleSubscriptionName returns the GoogleSubscriptionName field if non-nil, zero value otherwise.

### GetGoogleSubscriptionNameOk

`func (o *ApplicationResponse) GetGoogleSubscriptionNameOk() (*GoogleSubscriptionName, bool)`

GetGoogleSubscriptionNameOk returns a tuple with the GoogleSubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleSubscriptionName

`func (o *ApplicationResponse) SetGoogleSubscriptionName(v GoogleSubscriptionName)`

SetGoogleSubscriptionName sets GoogleSubscriptionName field to given value.

### HasGoogleSubscriptionName

`func (o *ApplicationResponse) HasGoogleSubscriptionName() bool`

HasGoogleSubscriptionName returns a boolean if a field has been set.

### SetGoogleSubscriptionNameNil

`func (o *ApplicationResponse) SetGoogleSubscriptionNameNil(b bool)`

 SetGoogleSubscriptionNameNil sets the value for GoogleSubscriptionName to be an explicit nil

### UnsetGoogleSubscriptionName
`func (o *ApplicationResponse) UnsetGoogleSubscriptionName()`

UnsetGoogleSubscriptionName ensures that no value is present for GoogleSubscriptionName, not even an explicit nil
### GetGoogleTopicName

`func (o *ApplicationResponse) GetGoogleTopicName() GoogleTopicName`

GetGoogleTopicName returns the GoogleTopicName field if non-nil, zero value otherwise.

### GetGoogleTopicNameOk

`func (o *ApplicationResponse) GetGoogleTopicNameOk() (*GoogleTopicName, bool)`

GetGoogleTopicNameOk returns a tuple with the GoogleTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTopicName

`func (o *ApplicationResponse) SetGoogleTopicName(v GoogleTopicName)`

SetGoogleTopicName sets GoogleTopicName field to given value.

### HasGoogleTopicName

`func (o *ApplicationResponse) HasGoogleTopicName() bool`

HasGoogleTopicName returns a boolean if a field has been set.

### SetGoogleTopicNameNil

`func (o *ApplicationResponse) SetGoogleTopicNameNil(b bool)`

 SetGoogleTopicNameNil sets the value for GoogleTopicName to be an explicit nil

### UnsetGoogleTopicName
`func (o *ApplicationResponse) UnsetGoogleTopicName()`

UnsetGoogleTopicName ensures that no value is present for GoogleTopicName, not even an explicit nil
### GetGoogleWorkspaceAccounts

`func (o *ApplicationResponse) GetGoogleWorkspaceAccounts() bool`

GetGoogleWorkspaceAccounts returns the GoogleWorkspaceAccounts field if non-nil, zero value otherwise.

### GetGoogleWorkspaceAccountsOk

`func (o *ApplicationResponse) GetGoogleWorkspaceAccountsOk() (*bool, bool)`

GetGoogleWorkspaceAccountsOk returns a tuple with the GoogleWorkspaceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleWorkspaceAccounts

`func (o *ApplicationResponse) SetGoogleWorkspaceAccounts(v bool)`

SetGoogleWorkspaceAccounts sets GoogleWorkspaceAccounts field to given value.

### HasGoogleWorkspaceAccounts

`func (o *ApplicationResponse) HasGoogleWorkspaceAccounts() bool`

HasGoogleWorkspaceAccounts returns a boolean if a field has been set.

### GetId

`func (o *ApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeInListing

`func (o *ApplicationResponse) GetIncludeInListing() bool`

GetIncludeInListing returns the IncludeInListing field if non-nil, zero value otherwise.

### GetIncludeInListingOk

`func (o *ApplicationResponse) GetIncludeInListingOk() (*bool, bool)`

GetIncludeInListingOk returns a tuple with the IncludeInListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInListing

`func (o *ApplicationResponse) SetIncludeInListing(v bool)`

SetIncludeInListing sets IncludeInListing field to given value.

### HasIncludeInListing

`func (o *ApplicationResponse) HasIncludeInListing() bool`

HasIncludeInListing returns a boolean if a field has been set.

### GetLastError

`func (o *ApplicationResponse) GetLastError() AccountErrorEntry`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ApplicationResponse) GetLastErrorOk() (*AccountErrorEntry, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ApplicationResponse) SetLastError(v AccountErrorEntry)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ApplicationResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ApplicationResponse) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ApplicationResponse) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLegacy

`func (o *ApplicationResponse) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *ApplicationResponse) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *ApplicationResponse) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *ApplicationResponse) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetName

`func (o *ApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *ApplicationResponse) GetProvider() OAuth2Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApplicationResponse) GetProviderOk() (*OAuth2Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApplicationResponse) SetProvider(v OAuth2Provider)`

SetProvider sets Provider field to given value.


### GetRedirectUrl

`func (o *ApplicationResponse) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ApplicationResponse) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ApplicationResponse) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ApplicationResponse) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetServiceClient

`func (o *ApplicationResponse) GetServiceClient() string`

GetServiceClient returns the ServiceClient field if non-nil, zero value otherwise.

### GetServiceClientOk

`func (o *ApplicationResponse) GetServiceClientOk() (*string, bool)`

GetServiceClientOk returns a tuple with the ServiceClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClient

`func (o *ApplicationResponse) SetServiceClient(v string)`

SetServiceClient sets ServiceClient field to given value.

### HasServiceClient

`func (o *ApplicationResponse) HasServiceClient() bool`

HasServiceClient returns a boolean if a field has been set.

### GetServiceClientEmail

`func (o *ApplicationResponse) GetServiceClientEmail() string`

GetServiceClientEmail returns the ServiceClientEmail field if non-nil, zero value otherwise.

### GetServiceClientEmailOk

`func (o *ApplicationResponse) GetServiceClientEmailOk() (*string, bool)`

GetServiceClientEmailOk returns a tuple with the ServiceClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClientEmail

`func (o *ApplicationResponse) SetServiceClientEmail(v string)`

SetServiceClientEmail sets ServiceClientEmail field to given value.

### HasServiceClientEmail

`func (o *ApplicationResponse) HasServiceClientEmail() bool`

HasServiceClientEmail returns a boolean if a field has been set.

### GetServiceKey

`func (o *ApplicationResponse) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *ApplicationResponse) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *ApplicationResponse) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *ApplicationResponse) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetTitle

`func (o *ApplicationResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApplicationResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApplicationResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApplicationResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *ApplicationResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ApplicationResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ApplicationResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ApplicationResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


