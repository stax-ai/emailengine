# Envelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEnvelope

`func NewEnvelope() *Envelope`

NewEnvelope instantiates a new Envelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopeWithDefaults

`func NewEnvelopeWithDefaults() *Envelope`

NewEnvelopeWithDefaults instantiates a new Envelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *Envelope) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Envelope) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Envelope) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Envelope) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Envelope) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Envelope) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Envelope) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *Envelope) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


