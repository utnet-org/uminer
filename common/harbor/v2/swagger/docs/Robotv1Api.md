# \Robotv1Api

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRobotV1**](Robotv1Api.md#CreateRobotV1) | **Post** /projects/{project_name_or_id}/robots | Create a robot account
[**DeleteRobotV1**](Robotv1Api.md#DeleteRobotV1) | **Delete** /projects/{project_name_or_id}/robots/{robot_id} | Delete a robot account
[**GetRobotByIDV1**](Robotv1Api.md#GetRobotByIDV1) | **Get** /projects/{project_name_or_id}/robots/{robot_id} | Get a robot account
[**ListRobotV1**](Robotv1Api.md#ListRobotV1) | **Get** /projects/{project_name_or_id}/robots | Get all robot accounts of specified project
[**UpdateRobotV1**](Robotv1Api.md#UpdateRobotV1) | **Put** /projects/{project_name_or_id}/robots/{robot_id} | Update status of robot account.


# **CreateRobotV1**
> RobotCreated CreateRobotV1(ctx, projectNameOrId, robot, optional)
Create a robot account

Create a robot account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **robot** | [**RobotCreateV1**](RobotCreateV1.md)| The JSON object of a robot account. | 
 **optional** | ***Robotv1ApiCreateRobotV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Robotv1ApiCreateRobotV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]

### Return type

[**RobotCreated**](RobotCreated.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRobotV1**
> DeleteRobotV1(ctx, projectNameOrId, robotId, optional)
Delete a robot account

This endpoint deletes specific robot account information by robot ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **robotId** | **int32**| Robot ID | 
 **optional** | ***Robotv1ApiDeleteRobotV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Robotv1ApiDeleteRobotV1Opts struct

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

# **GetRobotByIDV1**
> Robot GetRobotByIDV1(ctx, projectNameOrId, robotId, optional)
Get a robot account

This endpoint returns specific robot account information by robot ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **robotId** | **int32**| Robot ID | 
 **optional** | ***Robotv1ApiGetRobotByIDV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Robotv1ApiGetRobotByIDV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]

### Return type

[**Robot**](Robot.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRobotV1**
> []Robot ListRobotV1(ctx, projectNameOrId, optional)
Get all robot accounts of specified project

Get all robot accounts of specified project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***Robotv1ApiListRobotV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Robotv1ApiListRobotV1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 

### Return type

[**[]Robot**](Robot.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRobotV1**
> UpdateRobotV1(ctx, projectNameOrId, robotId, robot, optional)
Update status of robot account.

Used to disable/enable a specified robot account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **robotId** | **int32**| Robot ID | 
  **robot** | [**Robot**](Robot.md)| The JSON object of a robot account. | 
 **optional** | ***Robotv1ApiUpdateRobotV1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Robotv1ApiUpdateRobotV1Opts struct

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

