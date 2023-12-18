# \LdapApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportLdapUser**](LdapApi.md#ImportLdapUser) | **Post** /ldap/users/import | Import selected available ldap users.
[**PingLdap**](LdapApi.md#PingLdap) | **Post** /ldap/ping | Ping available ldap service.
[**SearchLdapGroup**](LdapApi.md#SearchLdapGroup) | **Get** /ldap/groups/search | Search available ldap groups.
[**SearchLdapUser**](LdapApi.md#SearchLdapUser) | **Get** /ldap/users/search | Search available ldap users.


# **ImportLdapUser**
> ImportLdapUser(ctx, uidList, optional)
Import selected available ldap users.

This endpoint adds the selected available ldap users to harbor based on related configuration parameters from the system. System will try to guess the user email address and realname, add to harbor user information. If have errors when import user, will return the list of importing failed uid and the failed reason. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uidList** | [**LdapImportUsers**](LdapImportUsers.md)| The uid listed for importing. This list will check users validity of ldap service based on configuration from the system. | 
 **optional** | ***LdapApiImportLdapUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapApiImportLdapUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingLdap**
> LdapPingResult PingLdap(ctx, optional)
Ping available ldap service.

This endpoint ping the available ldap service for test related configuration parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapApiPingLdapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapApiPingLdapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **ldapconf** | [**optional.Interface of LdapConf**](LdapConf.md)| ldap configuration. support input ldap service configuration. If it is a empty request, will load current configuration from the system. | 

### Return type

[**LdapPingResult**](LdapPingResult.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLdapGroup**
> []UserGroup SearchLdapGroup(ctx, optional)
Search available ldap groups.

This endpoint searches the available ldap groups based on related configuration parameters. support to search by groupname or groupdn. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapApiSearchLdapGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapApiSearchLdapGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **groupname** | **optional.String**| Ldap group name | 
 **groupdn** | **optional.String**| The LDAP group DN | 

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLdapUser**
> []LdapUser SearchLdapUser(ctx, optional)
Search available ldap users.

This endpoint searches the available ldap users based on related configuration parameters. Support searched by input ladp configuration, load configuration from the system and specific filter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapApiSearchLdapUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapApiSearchLdapUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **username** | **optional.String**| Registered user ID | 

### Return type

[**[]LdapUser**](LdapUser.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

