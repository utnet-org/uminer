# UserResp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | [optional] [default to null]
**Realname** | **string** |  | [optional] [default to null]
**Comment** | **string** |  | [optional] [default to null]
**UserId** | **int32** |  | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]
**SysadminFlag** | **bool** |  | [optional] [default to null]
**AdminRoleInAuth** | **bool** | indicate the admin privilege is grant by authenticator (LDAP), is always false unless it is the current login user | [optional] [default to null]
**OidcUserMeta** | [***OidcUserInfo**](OIDCUserInfo.md) |  | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The creation time of the user. | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | The update time of the user. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


