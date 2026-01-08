# OAuth2FilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]OAuth2ResponseItem**](OAuth2ResponseItem.md) |  | [optional] 
**Page** | Pointer to **int32** | Current page (0-based index) | [optional] 
**Pages** | Pointer to **int32** | Total page count | [optional] 
**Total** | Pointer to **int32** | How many matching entries | [optional] 

## Methods

### NewOAuth2FilterResponse

`func NewOAuth2FilterResponse() *OAuth2FilterResponse`

NewOAuth2FilterResponse instantiates a new OAuth2FilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2FilterResponseWithDefaults

`func NewOAuth2FilterResponseWithDefaults() *OAuth2FilterResponse`

NewOAuth2FilterResponseWithDefaults instantiates a new OAuth2FilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *OAuth2FilterResponse) GetApps() []OAuth2ResponseItem`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OAuth2FilterResponse) GetAppsOk() (*[]OAuth2ResponseItem, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OAuth2FilterResponse) SetApps(v []OAuth2ResponseItem)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OAuth2FilterResponse) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetPage

`func (o *OAuth2FilterResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OAuth2FilterResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OAuth2FilterResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *OAuth2FilterResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *OAuth2FilterResponse) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OAuth2FilterResponse) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OAuth2FilterResponse) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *OAuth2FilterResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *OAuth2FilterResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OAuth2FilterResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OAuth2FilterResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OAuth2FilterResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


