# LogSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Enable detailed logging for all email accounts | [optional] [default to false]
**MaxLogLines** | Pointer to **int32** | Maximum number of log entries to retain per account | [optional] [default to 10000]

## Methods

### NewLogSettings

`func NewLogSettings() *LogSettings`

NewLogSettings instantiates a new LogSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSettingsWithDefaults

`func NewLogSettingsWithDefaults() *LogSettings`

NewLogSettingsWithDefaults instantiates a new LogSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *LogSettings) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *LogSettings) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *LogSettings) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *LogSettings) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetMaxLogLines

`func (o *LogSettings) GetMaxLogLines() int32`

GetMaxLogLines returns the MaxLogLines field if non-nil, zero value otherwise.

### GetMaxLogLinesOk

`func (o *LogSettings) GetMaxLogLinesOk() (*int32, bool)`

GetMaxLogLinesOk returns a tuple with the MaxLogLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLogLines

`func (o *LogSettings) SetMaxLogLines(v int32)`

SetMaxLogLines sets MaxLogLines field to given value.

### HasMaxLogLines

`func (o *LogSettings) HasMaxLogLines() bool`

HasMaxLogLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


