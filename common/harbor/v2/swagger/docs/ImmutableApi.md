# \ImmutableApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImmuRule**](ImmutableApi.md#CreateImmuRule) | **Post** /projects/{project_name_or_id}/immutabletagrules | Add an immutable tag rule to current project
[**DeleteImmuRule**](ImmutableApi.md#DeleteImmuRule) | **Delete** /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id} | Delete the immutable tag rule.
[**ListImmuRules**](ImmutableApi.md#ListImmuRules) | **Get** /projects/{project_name_or_id}/immutabletagrules | List all immutable tag rules of current project
[**UpdateImmuRule**](ImmutableApi.md#UpdateImmuRule) | **Put** /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id} | Update the immutable tag rule or enable or disable the rule


# **CreateImmuRule**
> CreateImmuRule(ctx, projectNameOrId, immutableRule, optional)
Add an immutable tag rule to current project

This endpoint add an immutable tag rule to the project 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **immutableRule** | [**ImmutableRule**](ImmutableRule.md)|  | 
 **optional** | ***ImmutableApiCreateImmuRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImmutableApiCreateImmuRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImmuRule**
> DeleteImmuRule(ctx, projectNameOrId, immutableRuleId, optional)
Delete the immutable tag rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **immutableRuleId** | **int64**| The ID of the immutable rule | 
 **optional** | ***ImmutableApiDeleteImmuRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImmutableApiDeleteImmuRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListImmuRules**
> []ImmutableRule ListImmuRules(ctx, projectNameOrId, optional)
List all immutable tag rules of current project

This endpoint returns the immutable tag rules of a project 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***ImmutableApiListImmuRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImmutableApiListImmuRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 

### Return type

[**[]ImmutableRule**](ImmutableRule.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateImmuRule**
> UpdateImmuRule(ctx, projectNameOrId, immutableRuleId, immutableRule, optional)
Update the immutable tag rule or enable or disable the rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **immutableRuleId** | **int64**| The ID of the immutable rule | 
  **immutableRule** | [**ImmutableRule**](ImmutableRule.md)|  | 
 **optional** | ***ImmutableApiUpdateImmuRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImmutableApiUpdateImmuRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

