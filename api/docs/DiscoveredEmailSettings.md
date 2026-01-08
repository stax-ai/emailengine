# DiscoveredEmailSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Source for the detected info | [optional] 
**Imap** | Pointer to [**ResolvedServerSettings**](ResolvedServerSettings.md) |  | [optional] 
**Smtp** | Pointer to [**DiscoveredServerSettings**](DiscoveredServerSettings.md) |  | [optional] 

## Methods

### NewDiscoveredEmailSettings

`func NewDiscoveredEmailSettings() *DiscoveredEmailSettings`

NewDiscoveredEmailSettings instantiates a new DiscoveredEmailSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredEmailSettingsWithDefaults

`func NewDiscoveredEmailSettingsWithDefaults() *DiscoveredEmailSettings`

NewDiscoveredEmailSettingsWithDefaults instantiates a new DiscoveredEmailSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DiscoveredEmailSettings) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DiscoveredEmailSettings) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DiscoveredEmailSettings) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DiscoveredEmailSettings) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetImap

`func (o *DiscoveredEmailSettings) GetImap() ResolvedServerSettings`

GetImap returns the Imap field if non-nil, zero value otherwise.

### GetImapOk

`func (o *DiscoveredEmailSettings) GetImapOk() (*ResolvedServerSettings, bool)`

GetImapOk returns a tuple with the Imap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImap

`func (o *DiscoveredEmailSettings) SetImap(v ResolvedServerSettings)`

SetImap sets Imap field to given value.

### HasImap

`func (o *DiscoveredEmailSettings) HasImap() bool`

HasImap returns a boolean if a field has been set.

### GetSmtp

`func (o *DiscoveredEmailSettings) GetSmtp() DiscoveredServerSettings`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *DiscoveredEmailSettings) GetSmtpOk() (*DiscoveredServerSettings, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *DiscoveredEmailSettings) SetSmtp(v DiscoveredServerSettings)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *DiscoveredEmailSettings) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


