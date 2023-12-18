# \SystemCVEAllowlistApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemCVEAllowlist**](SystemCVEAllowlistApi.md#GetSystemCVEAllowlist) | **Get** /system/CVEAllowlist | Get the system level allowlist of CVE.
[**PutSystemCVEAllowlist**](SystemCVEAllowlistApi.md#PutSystemCVEAllowlist) | **Put** /system/CVEAllowlist | Update the system level allowlist of CVE.


# **GetSystemCVEAllowlist**
> CveAllowlist GetSystemCVEAllowlist(ctx, optional)
Get the system level allowlist of CVE.

Get the system level allowlist of CVE.  This API can be called by all authenticated users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemCVEAllowlistApiGetSystemCVEAllowlistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemCVEAllowlistApiGetSystemCVEAllowlistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**CveAllowlist**](CVEAllowlist.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSystemCVEAllowlist**
> PutSystemCVEAllowlist(ctx, optional)
Update the system level allowlist of CVE.

This API overwrites the system level allowlist of CVE with the list in request body.  Only system Admin has permission to call this API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemCVEAllowlistApiPutSystemCVEAllowlistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemCVEAllowlistApiPutSystemCVEAllowlistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **allowlist** | [**optional.Interface of CveAllowlist**](CveAllowlist.md)| The allowlist with new content | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

