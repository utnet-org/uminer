# OidcUserInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | the ID of the OIDC info record | [optional] [default to null]
**UserId** | **int32** | the ID of the user | [optional] [default to null]
**Subiss** | **string** | the concatenation of sub and issuer in the ID token | [optional] [default to null]
**Secret** | **string** | the secret of the OIDC user that can be used for CLI to push/pull artifacts | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The creation time of the OIDC user info record. | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | The update time of the OIDC user info record. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


