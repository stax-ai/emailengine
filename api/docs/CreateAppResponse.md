# CreateAppResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Was the app created | [optional] [default to true]
**Id** | **string** | OAuth2 application ID | 

## Methods

### NewCreateAppResponse

`func NewCreateAppResponse(id string, ) *CreateAppResponse`

NewCreateAppResponse instantiates a new CreateAppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppResponseWithDefaults

`func NewCreateAppResponseWithDefaults() *CreateAppResponse`

NewCreateAppResponseWithDefaults instantiates a new CreateAppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CreateAppResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateAppResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateAppResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateAppResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *CreateAppResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAppResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAppResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


