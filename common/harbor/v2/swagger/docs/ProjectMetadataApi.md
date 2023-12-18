# \ProjectMetadataApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProjectMetadatas**](ProjectMetadataApi.md#AddProjectMetadatas) | **Post** /projects/{project_name_or_id}/metadatas/ | Add metadata for the specific project
[**DeleteProjectMetadata**](ProjectMetadataApi.md#DeleteProjectMetadata) | **Delete** /projects/{project_name_or_id}/metadatas/{meta_name} | Delete the specific metadata for the specific project
[**GetProjectMetadata**](ProjectMetadataApi.md#GetProjectMetadata) | **Get** /projects/{project_name_or_id}/metadatas/{meta_name} | Get the specific metadata of the specific project
[**ListProjectMetadatas**](ProjectMetadataApi.md#ListProjectMetadatas) | **Get** /projects/{project_name_or_id}/metadatas/ | Get the metadata of the specific project
[**UpdateProjectMetadata**](ProjectMetadataApi.md#UpdateProjectMetadata) | **Put** /projects/{project_name_or_id}/metadatas/{meta_name} | Update the specific metadata for the specific project


# **AddProjectMetadatas**
> AddProjectMetadatas(ctx, projectNameOrId, optional)
Add metadata for the specific project

Add metadata for the specific project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***ProjectMetadataApiAddProjectMetadatasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMetadataApiAddProjectMetadatasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]
 **metadata** | **optional.Interface{}**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectMetadata**
> DeleteProjectMetadata(ctx, projectNameOrId, metaName, optional)
Delete the specific metadata for the specific project

Delete the specific metadata for the specific project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **metaName** | **string**| The name of metadata. | 
 **optional** | ***ProjectMetadataApiDeleteProjectMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMetadataApiDeleteProjectMetadataOpts struct

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

# **GetProjectMetadata**
> map[string]string GetProjectMetadata(ctx, projectNameOrId, metaName, optional)
Get the specific metadata of the specific project

Get the specific metadata of the specific project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **metaName** | **string**| The name of metadata. | 
 **optional** | ***ProjectMetadataApiGetProjectMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMetadataApiGetProjectMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]

### Return type

**map[string]string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjectMetadatas**
> map[string]string ListProjectMetadatas(ctx, projectNameOrId, optional)
Get the metadata of the specific project

Get the metadata of the specific project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
 **optional** | ***ProjectMetadataApiListProjectMetadatasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMetadataApiListProjectMetadatasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]

### Return type

**map[string]string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectMetadata**
> UpdateProjectMetadata(ctx, projectNameOrId, metaName, optional)
Update the specific metadata for the specific project

Update the specific metadata for the specific project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectNameOrId** | **string**| The name or id of the project | 
  **metaName** | **string**| The name of metadata. | 
 **optional** | ***ProjectMetadataApiUpdateProjectMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectMetadataApiUpdateProjectMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| An unique ID for the request | 
 **xIsResourceName** | **optional.Bool**| The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name. | [default to false]
 **metadata** | **optional.Interface{}**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

