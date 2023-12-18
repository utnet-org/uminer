# AuthproxySetting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | The fully qualified URI of login endpoint of authproxy, such as &#39;https://192.168.1.2:8443/login&#39; | [optional] [default to null]
**TokenreivewEndpoint** | **string** | The fully qualified URI of token review endpoint of authproxy, such as &#39;https://192.168.1.2:8443/tokenreview&#39; | [optional] [default to null]
**SkipSearch** | **bool** | The flag to determine whether Harbor can skip search the user/group when adding him as a member. | [optional] [default to null]
**VerifyCert** | **bool** | The flag to determine whether Harbor should verify the certificate when connecting to the auth proxy. | [optional] [default to null]
**ServerCertificate** | **string** | The certificate to be pinned when connecting auth proxy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


