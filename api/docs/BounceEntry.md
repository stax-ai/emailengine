# BounceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Bounce action (failed, delayed, etc.) | [optional] 
**Date** | Pointer to **time.Time** | When the bounce was detected | [optional] 
**Message** | **string** | EmailEngine identifier of the bounce notification | 
**Recipient** | Pointer to **string** | Email address that bounced | [optional] 
**Response** | Pointer to [**BounceResponse**](BounceResponse.md) |  | [optional] 

## Methods

### NewBounceEntry

`func NewBounceEntry(message string, ) *BounceEntry`

NewBounceEntry instantiates a new BounceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceEntryWithDefaults

`func NewBounceEntryWithDefaults() *BounceEntry`

NewBounceEntryWithDefaults instantiates a new BounceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BounceEntry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BounceEntry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BounceEntry) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BounceEntry) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDate

`func (o *BounceEntry) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BounceEntry) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BounceEntry) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BounceEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMessage

`func (o *BounceEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BounceEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BounceEntry) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRecipient

`func (o *BounceEntry) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *BounceEntry) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *BounceEntry) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *BounceEntry) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetResponse

`func (o *BounceEntry) GetResponse() BounceResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *BounceEntry) GetResponseOk() (*BounceResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *BounceEntry) SetResponse(v BounceResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *BounceEntry) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


