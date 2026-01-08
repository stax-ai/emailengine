# LicenseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether a valid license is currently active | [optional] 
**Details** | Pointer to [**LicenseDetails**](LicenseDetails.md) |  | [optional] 
**Suspended** | Pointer to **bool** | Whether email operations are suspended due to license issues | [optional] 
**Type** | Pointer to **string** | License type/product name | [optional] 

## Methods

### NewLicenseResponse

`func NewLicenseResponse() *LicenseResponse`

NewLicenseResponse instantiates a new LicenseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseResponseWithDefaults

`func NewLicenseResponseWithDefaults() *LicenseResponse`

NewLicenseResponseWithDefaults instantiates a new LicenseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *LicenseResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *LicenseResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *LicenseResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *LicenseResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDetails

`func (o *LicenseResponse) GetDetails() LicenseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LicenseResponse) GetDetailsOk() (*LicenseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *LicenseResponse) SetDetails(v LicenseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *LicenseResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetSuspended

`func (o *LicenseResponse) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *LicenseResponse) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *LicenseResponse) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *LicenseResponse) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetType

`func (o *LicenseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LicenseResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


