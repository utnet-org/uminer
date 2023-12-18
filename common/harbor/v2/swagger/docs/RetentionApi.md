# \RetentionApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRetention**](RetentionApi.md#CreateRetention) | **Post** /retentions | Create Retention Policy
[**DeleteRetention**](RetentionApi.md#DeleteRetention) | **Delete** /retentions/{id} | Delete Retention Policy
[**GetRentenitionMetadata**](RetentionApi.md#GetRentenitionMetadata) | **Get** /retentions/metadatas | Get Retention Metadatas
[**GetRetention**](RetentionApi.md#GetRetention) | **Get** /retentions/{id} | Get Retention Policy
[**GetRetentionTaskLog**](RetentionApi.md#GetRetentionTaskLog) | **Get** /retentions/{id}/executions/{eid}/tasks/{tid} | Get Retention job task log
[**ListRetentionExecutions**](RetentionApi.md#ListRetentionExecutions) | **Get** /retentions/{id}/executions | Get Retention executions
[**ListRetentionTasks**](RetentionApi.md#ListRetentionTasks) | **Get** /retentions/{id}/executions/{eid}/tasks | Get Retention tasks
[**OperateRetentionExecution**](RetentionApi.md#OperateRetentionExecution) | **Patch** /retentions/{id}/executions/{eid} | Stop a Retention execution
[**TriggerRetentionExecution**](RetentionApi.md#TriggerRetentionExecution) | **Post** /retentions/{id}/executions | Trigger a Retention Execution
[**UpdateRetention**](RetentionApi.md#UpdateRetention) | **Put** /retentions/{id} | Update Retention Policy


# **CreateRetention**
> CreateRetention(ctx, policy, optional)
Create Retention Policy

Create Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when no retention policy binded to project yet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policy** | [**RetentionPolicy**](RetentionPolicy.md)| Create Retention Policy successfully. | 
 **optional** | ***RetentionApiCreateRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiCreateRetentionOpts struct

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

# **DeleteRetention**
> DeleteRetention(ctx, id, optional)
Delete Retention Policy

Delete Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when retention policy has already binded to project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
 **optional** | ***RetentionApiDeleteRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiDeleteRetentionOpts struct

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

# **GetRentenitionMetadata**
> RetentionMetadata GetRentenitionMetadata(ctx, optional)
Get Retention Metadatas

Get Retention Metadatas.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RetentionApiGetRentenitionMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiGetRentenitionMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**RetentionMetadata**](RetentionMetadata.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRetention**
> RetentionPolicy GetRetention(ctx, id, optional)
Get Retention Policy

Get Retention Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
 **optional** | ***RetentionApiGetRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiGetRetentionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**RetentionPolicy**](RetentionPolicy.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRetentionTaskLog**
> string GetRetentionTaskLog(ctx, id, eid, tid, optional)
Get Retention job task log

Get Retention job task log, tags ratain or deletion detail will be shown in a table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
  **eid** | **int64**| Retention execution ID. | 
  **tid** | **int64**| Retention execution ID. | 
 **optional** | ***RetentionApiGetRetentionTaskLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiGetRetentionTaskLogOpts struct

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

# **ListRetentionExecutions**
> []RetentionExecution ListRetentionExecutions(ctx, id, optional)
Get Retention executions

Get Retention executions, execution status may be delayed before job service schedule it up.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
 **optional** | ***RetentionApiListRetentionExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiListRetentionExecutionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **page** | **optional.Int64**| The page number. | 
 **pageSize** | **optional.Int64**| The size of per page. | 

### Return type

[**[]RetentionExecution**](RetentionExecution.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRetentionTasks**
> []RetentionExecutionTask ListRetentionTasks(ctx, id, eid, optional)
Get Retention tasks

Get Retention tasks, each repository as a task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
  **eid** | **int64**| Retention execution ID. | 
 **optional** | ***RetentionApiListRetentionTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiListRetentionTasksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **page** | **optional.Int64**| The page number. | 
 **pageSize** | **optional.Int64**| The size of per page. | 

### Return type

[**[]RetentionExecutionTask**](RetentionExecutionTask.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperateRetentionExecution**
> OperateRetentionExecution(ctx, id, eid, body, optional)
Stop a Retention execution

Stop a Retention execution, only support \"stop\" action now.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
  **eid** | **int64**| Retention execution ID. | 
  **body** | [**Body1**](Body1.md)| The action, only support \&quot;stop\&quot; now. | 
 **optional** | ***RetentionApiOperateRetentionExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiOperateRetentionExecutionOpts struct

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

# **TriggerRetentionExecution**
> TriggerRetentionExecution(ctx, id, body, optional)
Trigger a Retention Execution

Trigger a Retention Execution, if dry_run is True, nothing would be deleted actually.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
  **body** | [**Body**](Body.md)|  | 
 **optional** | ***RetentionApiTriggerRetentionExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiTriggerRetentionExecutionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRetention**
> UpdateRetention(ctx, id, policy, optional)
Update Retention Policy

Update Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when retention policy has already binded to project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Retention ID. | 
  **policy** | [**RetentionPolicy**](RetentionPolicy.md)|  | 
 **optional** | ***RetentionApiUpdateRetentionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetentionApiUpdateRetentionOpts struct

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

