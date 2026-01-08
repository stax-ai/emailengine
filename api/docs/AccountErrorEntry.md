# AccountErrorEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | Human-readable error message | [optional] 
**ServerResponseCode** | Pointer to **string** | Error code or classification | [optional] 
**TokenRequest** | Pointer to [**TokenRequest**](TokenRequest.md) |  | [optional] 

## Methods

### NewAccountErrorEntry

`func NewAccountErrorEntry() *AccountErrorEntry`

NewAccountErrorEntry instantiates a new AccountErrorEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountErrorEntryWithDefaults

`func NewAccountErrorEntryWithDefaults() *AccountErrorEntry`

NewAccountErrorEntryWithDefaults instantiates a new AccountErrorEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *AccountErrorEntry) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AccountErrorEntry) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AccountErrorEntry) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AccountErrorEntry) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetServerResponseCode

`func (o *AccountErrorEntry) GetServerResponseCode() string`

GetServerResponseCode returns the ServerResponseCode field if non-nil, zero value otherwise.

### GetServerResponseCodeOk

`func (o *AccountErrorEntry) GetServerResponseCodeOk() (*string, bool)`

GetServerResponseCodeOk returns a tuple with the ServerResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerResponseCode

`func (o *AccountErrorEntry) SetServerResponseCode(v string)`

SetServerResponseCode sets ServerResponseCode field to given value.

### HasServerResponseCode

`func (o *AccountErrorEntry) HasServerResponseCode() bool`

HasServerResponseCode returns a boolean if a field has been set.

### GetTokenRequest

`func (o *AccountErrorEntry) GetTokenRequest() TokenRequest`

GetTokenRequest returns the TokenRequest field if non-nil, zero value otherwise.

### GetTokenRequestOk

`func (o *AccountErrorEntry) GetTokenRequestOk() (*TokenRequest, bool)`

GetTokenRequestOk returns a tuple with the TokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequest

`func (o *AccountErrorEntry) SetTokenRequest(v TokenRequest)`

SetTokenRequest sets TokenRequest field to given value.

### HasTokenRequest

`func (o *AccountErrorEntry) HasTokenRequest() bool`

HasTokenRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


