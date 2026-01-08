# AccountCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **map[string]interface{}** | Cumulative event counters for the account lifetime | [optional] 

## Methods

### NewAccountCounters

`func NewAccountCounters() *AccountCounters`

NewAccountCounters instantiates a new AccountCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCountersWithDefaults

`func NewAccountCountersWithDefaults() *AccountCounters`

NewAccountCountersWithDefaults instantiates a new AccountCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AccountCounters) GetEvents() map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AccountCounters) GetEventsOk() (*map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AccountCounters) SetEvents(v map[string]interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AccountCounters) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


