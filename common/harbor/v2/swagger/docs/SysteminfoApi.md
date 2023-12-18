# \SysteminfoApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCert**](SysteminfoApi.md#GetCert) | **Get** /systeminfo/getcert | Get default root certificate.
[**GetSystemInfo**](SysteminfoApi.md#GetSystemInfo) | **Get** /systeminfo | Get general system info
[**GetVolumes**](SysteminfoApi.md#GetVolumes) | **Get** /systeminfo/volumes | Get system volume info (total/free size).


# **GetCert**
> *os.File GetCert(ctx, optional)
Get default root certificate.

This endpoint is for downloading a default root certificate. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SysteminfoApiGetCertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SysteminfoApiGetCertOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemInfo**
> GeneralInfo GetSystemInfo(ctx, optional)
Get general system info

This API is for retrieving general system info, this can be called by anonymous request.  Some attributes will be omitted in the response when this API is called by anonymous request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SysteminfoApiGetSystemInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SysteminfoApiGetSystemInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**GeneralInfo**](GeneralInfo.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumes**
> SystemInfo GetVolumes(ctx, optional)
Get system volume info (total/free size).

This endpoint is for retrieving system volume info that only provides for admin user.  Note that the response only reflects the storage status of local disk. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SysteminfoApiGetVolumesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SysteminfoApiGetVolumesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

