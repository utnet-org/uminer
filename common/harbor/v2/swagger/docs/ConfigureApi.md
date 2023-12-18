# \ConfigureApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurations**](ConfigureApi.md#GetConfigurations) | **Get** /configurations | Get system configurations.
[**GetInternalconfig**](ConfigureApi.md#GetInternalconfig) | **Get** /internalconfig | Get internal configurations.
[**UpdateConfigurations**](ConfigureApi.md#UpdateConfigurations) | **Put** /configurations | Modify system configurations.


# **GetConfigurations**
> ConfigurationsResponse GetConfigurations(ctx, optional)
Get system configurations.

This endpoint is for retrieving system configurations that only provides for admin user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigureApiGetConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigureApiGetConfigurationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ConfigurationsResponse**](ConfigurationsResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInternalconfig**
> InternalConfigurationsResponse GetInternalconfig(ctx, optional)
Get internal configurations.

This endpoint is for retrieving system configurations that only provides for internal api call. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigureApiGetInternalconfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigureApiGetInternalconfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**InternalConfigurationsResponse**](InternalConfigurationsResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfigurations**
> UpdateConfigurations(ctx, configurations, optional)
Modify system configurations.

This endpoint is for modifying system configurations that only provides for admin user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configurations** | [**Configurations**](Configurations.md)| The configuration map can contain a subset of the attributes of the schema, which are to be updated. | 
 **optional** | ***ConfigureApiUpdateConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigureApiUpdateConfigurationsOpts struct

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

