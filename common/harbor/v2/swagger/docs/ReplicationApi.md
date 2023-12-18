# \ReplicationApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplicationPolicy**](ReplicationApi.md#CreateReplicationPolicy) | **Post** /replication/policies | Create a replication policy
[**DeleteReplicationPolicy**](ReplicationApi.md#DeleteReplicationPolicy) | **Delete** /replication/policies/{id} | Delete the specific replication policy
[**GetReplicationExecution**](ReplicationApi.md#GetReplicationExecution) | **Get** /replication/executions/{id} | Get the specific replication execution
[**GetReplicationLog**](ReplicationApi.md#GetReplicationLog) | **Get** /replication/executions/{id}/tasks/{task_id}/log | Get the log of the specific replication task
[**GetReplicationPolicy**](ReplicationApi.md#GetReplicationPolicy) | **Get** /replication/policies/{id} | Get the specific replication policy
[**ListReplicationExecutions**](ReplicationApi.md#ListReplicationExecutions) | **Get** /replication/executions | List replication executions
[**ListReplicationPolicies**](ReplicationApi.md#ListReplicationPolicies) | **Get** /replication/policies | List replication policies
[**ListReplicationTasks**](ReplicationApi.md#ListReplicationTasks) | **Get** /replication/executions/{id}/tasks | List replication tasks for a specific execution
[**StartReplication**](ReplicationApi.md#StartReplication) | **Post** /replication/executions | Start one replication execution
[**StopReplication**](ReplicationApi.md#StopReplication) | **Put** /replication/executions/{id} | Stop the specific replication execution
[**UpdateReplicationPolicy**](ReplicationApi.md#UpdateReplicationPolicy) | **Put** /replication/policies/{id} | Update the replication policy


# **CreateReplicationPolicy**
> CreateReplicationPolicy(ctx, policy, optional)
Create a replication policy

Create a replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policy** | [**ReplicationPolicy**](ReplicationPolicy.md)| The replication policy | 
 **optional** | ***ReplicationApiCreateReplicationPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiCreateReplicationPolicyOpts struct

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

# **DeleteReplicationPolicy**
> DeleteReplicationPolicy(ctx, id, optional)
Delete the specific replication policy

Delete the specific replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Replication policy ID | 
 **optional** | ***ReplicationApiDeleteReplicationPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiDeleteReplicationPolicyOpts struct

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

# **GetReplicationExecution**
> ReplicationExecution GetReplicationExecution(ctx, id, optional)
Get the specific replication execution

Get the replication execution specified by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the execution. | 
 **optional** | ***ReplicationApiGetReplicationExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiGetReplicationExecutionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ReplicationExecution**](ReplicationExecution.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplicationLog**
> string GetReplicationLog(ctx, id, taskId, optional)
Get the log of the specific replication task

Get the log of the specific replication task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the execution that the tasks belongs to. | 
  **taskId** | **int64**| The ID of the task. | 
 **optional** | ***ReplicationApiGetReplicationLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiGetReplicationLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplicationPolicy**
> ReplicationPolicy GetReplicationPolicy(ctx, id, optional)
Get the specific replication policy

Get the specific replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Policy ID | 
 **optional** | ***ReplicationApiGetReplicationPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiGetReplicationPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReplicationExecutions**
> []ReplicationExecution ListReplicationExecutions(ctx, optional)
List replication executions

List replication executions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationApiListReplicationExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiListReplicationExecutionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **policyId** | **optional.Int32**| The ID of the policy that the executions belong to. | 
 **status** | **optional.String**| The execution status. | 
 **trigger** | **optional.String**| The trigger mode. | 

### Return type

[**[]ReplicationExecution**](ReplicationExecution.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReplicationPolicies**
> []ReplicationPolicy ListReplicationPolicies(ctx, optional)
List replication policies

List replication policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationApiListReplicationPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiListReplicationPoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **name** | **optional.String**| Deprecated, use \&quot;query\&quot; instead. The policy name. | 

### Return type

[**[]ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReplicationTasks**
> []ReplicationTask ListReplicationTasks(ctx, id, optional)
List replication tasks for a specific execution

List replication tasks for a specific execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the execution that the tasks belongs to. | 
 **optional** | ***ReplicationApiListReplicationTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiListReplicationTasksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **status** | **optional.String**| The task status. | 
 **resourceType** | **optional.String**| The resource type. | 

### Return type

[**[]ReplicationTask**](ReplicationTask.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartReplication**
> StartReplication(ctx, execution, optional)
Start one replication execution

Start one replication execution according to the policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **execution** | [**StartReplicationExecution**](StartReplicationExecution.md)| The ID of policy that the execution belongs to | 
 **optional** | ***ReplicationApiStartReplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiStartReplicationOpts struct

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

# **StopReplication**
> StopReplication(ctx, id, optional)
Stop the specific replication execution

Stop the replication execution specified by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the execution. | 
 **optional** | ***ReplicationApiStopReplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiStopReplicationOpts struct

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

# **UpdateReplicationPolicy**
> UpdateReplicationPolicy(ctx, id, policy, optional)
Update the replication policy

Update the replication policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The policy ID | 
  **policy** | [**ReplicationPolicy**](ReplicationPolicy.md)| The replication policy | 
 **optional** | ***ReplicationApiUpdateReplicationPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationApiUpdateReplicationPolicyOpts struct

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

