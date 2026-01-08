# DSN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The envelope identifier that would be included in the response (ENVID) | [optional] 
**Notify** | Pointer to [**[]NotifyEntry**](NotifyEntry.md) | Defines the conditions under which a DSN response should be sent | [optional] 
**Recipient** | Pointer to **string** | The email address the DSN should be sent (ORCPT) | [optional] 
**Return** | [**ModelReturn**](ModelReturn.md) |  | 

## Methods

### NewDSN

`func NewDSN(return_ ModelReturn, ) *DSN`

NewDSN instantiates a new DSN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSNWithDefaults

`func NewDSNWithDefaults() *DSN`

NewDSNWithDefaults instantiates a new DSN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DSN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DSN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DSN) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DSN) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotify

`func (o *DSN) GetNotify() []NotifyEntry`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *DSN) GetNotifyOk() (*[]NotifyEntry, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *DSN) SetNotify(v []NotifyEntry)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *DSN) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetRecipient

`func (o *DSN) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *DSN) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *DSN) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *DSN) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetReturn

`func (o *DSN) GetReturn() ModelReturn`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *DSN) GetReturnOk() (*ModelReturn, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *DSN) SetReturn(v ModelReturn)`

SetReturn sets Return field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


