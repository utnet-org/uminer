# \DefaultApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesScanAllPost**](DefaultApi.md#RepositoriesScanAllPost) | **Post** /repositories/scanAll | Scan all images of the registry.


# **RepositoriesScanAllPost**
> RepositoriesScanAllPost(ctx, optional)
Scan all images of the registry.

The server will launch different jobs to scan each image on the regsitry, so this is equivalent to calling  the API to scan the image one by one in background, so there's no way to track the overall status of the \"scan all\" action.  Only system adim has permission to call this API.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiRepositoriesScanAllPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRepositoriesScanAllPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **optional.Int32**| When this parm is set only the images under the project identified by the project_id will be scanned. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

