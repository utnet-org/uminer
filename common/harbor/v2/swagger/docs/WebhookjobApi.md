# \WebhookjobApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListWebhookJobs**](WebhookjobApi.md#ListWebhookJobs) | **Get** /projects/{project_name_or_id}/webhook/jobs | List project webhook jobs


# **ListWebhookJobs**
> []WebhookJob ListWebhookJobs(ctx, projectNameOrId, policyId, optional)
List project webhook jobs

This endpoint returns webhook jobs of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **policyId** | **int64**| The policy ID. | 
 **optional** | ***WebhookjobApiListWebhookJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookjobApiListWebhookJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **status** | [**optional.Interface of []string**](string.md)| The status of webhook job. | 

### Return type

[**[]WebhookJob**](WebhookJob.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

