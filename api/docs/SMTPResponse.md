# SMTPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**Model1**](Model1.md) |  | [optional] 
**Host** | **string** | SMTP server hostname | 
**Port** | **int32** | SMTP server port (typically 587 for STARTTLS, 465 for SMTP over TLS, 25 for unencrypted) | 
**Secure** | Pointer to **bool** | Use TLS encryption from the start (true for port 465, false for STARTTLS on ports 587/25) | [optional] [default to false]
**Tls** | Pointer to [**TLS**](TLS.md) |  | [optional] 
**UseAuthServer** | Pointer to **bool** | Use external authentication server to retrieve credentials dynamically | [optional] 

## Methods

### NewSMTPResponse

`func NewSMTPResponse(host string, port int32, ) *SMTPResponse`

NewSMTPResponse instantiates a new SMTPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPResponseWithDefaults

`func NewSMTPResponseWithDefaults() *SMTPResponse`

NewSMTPResponseWithDefaults instantiates a new SMTPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *SMTPResponse) GetAuth() Model1`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *SMTPResponse) GetAuthOk() (*Model1, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *SMTPResponse) SetAuth(v Model1)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *SMTPResponse) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetHost

`func (o *SMTPResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPResponse) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SMTPResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPResponse) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecure

`func (o *SMTPResponse) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *SMTPResponse) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *SMTPResponse) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *SMTPResponse) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTls

`func (o *SMTPResponse) GetTls() TLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *SMTPResponse) GetTlsOk() (*TLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *SMTPResponse) SetTls(v TLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *SMTPResponse) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetUseAuthServer

`func (o *SMTPResponse) GetUseAuthServer() bool`

GetUseAuthServer returns the UseAuthServer field if non-nil, zero value otherwise.

### GetUseAuthServerOk

`func (o *SMTPResponse) GetUseAuthServerOk() (*bool, bool)`

GetUseAuthServerOk returns a tuple with the UseAuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAuthServer

`func (o *SMTPResponse) SetUseAuthServer(v bool)`

SetUseAuthServer sets UseAuthServer field to given value.

### HasUseAuthServer

`func (o *SMTPResponse) HasUseAuthServer() bool`

HasUseAuthServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


