# LicenseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | Licensed application identifier | [optional] 
**Created** | Pointer to **string** | License creation date | [optional] 
**Hostname** | Pointer to **string** | Licensed hostname or environment | [optional] 
**Key** | Pointer to **string** | License key | [optional] 
**LicensedTo** | Pointer to **string** | Organization or individual the license is issued to | [optional] 

## Methods

### NewLicenseDetails

`func NewLicenseDetails() *LicenseDetails`

NewLicenseDetails instantiates a new LicenseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseDetailsWithDefaults

`func NewLicenseDetailsWithDefaults() *LicenseDetails`

NewLicenseDetailsWithDefaults instantiates a new LicenseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *LicenseDetails) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LicenseDetails) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LicenseDetails) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LicenseDetails) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreated

`func (o *LicenseDetails) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LicenseDetails) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LicenseDetails) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LicenseDetails) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetHostname

`func (o *LicenseDetails) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LicenseDetails) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LicenseDetails) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LicenseDetails) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetKey

`func (o *LicenseDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LicenseDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LicenseDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LicenseDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLicensedTo

`func (o *LicenseDetails) GetLicensedTo() string`

GetLicensedTo returns the LicensedTo field if non-nil, zero value otherwise.

### GetLicensedToOk

`func (o *LicenseDetails) GetLicensedToOk() (*string, bool)`

GetLicensedToOk returns a tuple with the LicensedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedTo

`func (o *LicenseDetails) SetLicensedTo(v string)`

SetLicensedTo sets LicensedTo field to given value.

### HasLicensedTo

`func (o *LicenseDetails) HasLicensedTo() bool`

HasLicensedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


