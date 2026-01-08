# SMTPUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**Model16**](Model16.md) |  | [optional] 
**Host** | Pointer to **string** | SMTP server hostname | [optional] 
**Partial** | Pointer to **bool** | Update only the provided fields instead of replacing the entire configuration | [optional] [default to false]
**Port** | Pointer to **int32** | SMTP server port | [optional] 
**Secure** | Pointer to **bool** | Use TLS encryption from the start | [optional] 
**Tls** | Pointer to [**Model17**](Model17.md) |  | [optional] 
**UseAuthServer** | Pointer to **bool** | Use external authentication server to retrieve credentials dynamically | [optional] 

## Methods

### NewSMTPUpdate

`func NewSMTPUpdate() *SMTPUpdate`

NewSMTPUpdate instantiates a new SMTPUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPUpdateWithDefaults

`func NewSMTPUpdateWithDefaults() *SMTPUpdate`

NewSMTPUpdateWithDefaults instantiates a new SMTPUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *SMTPUpdate) GetAuth() Model16`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *SMTPUpdate) GetAuthOk() (*Model16, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *SMTPUpdate) SetAuth(v Model16)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *SMTPUpdate) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetHost

`func (o *SMTPUpdate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPUpdate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPUpdate) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SMTPUpdate) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPartial

`func (o *SMTPUpdate) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *SMTPUpdate) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *SMTPUpdate) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *SMTPUpdate) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetPort

`func (o *SMTPUpdate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPUpdate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPUpdate) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPUpdate) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecure

`func (o *SMTPUpdate) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *SMTPUpdate) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *SMTPUpdate) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *SMTPUpdate) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTls

`func (o *SMTPUpdate) GetTls() Model17`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *SMTPUpdate) GetTlsOk() (*Model17, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *SMTPUpdate) SetTls(v Model17)`

SetTls sets Tls field to given value.

### HasTls

`func (o *SMTPUpdate) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetUseAuthServer

`func (o *SMTPUpdate) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *SMTPUpdate) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *SMTPUpdate) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *SMTPUpdate) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


