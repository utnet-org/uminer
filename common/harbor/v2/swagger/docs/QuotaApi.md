# \QuotaApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuota**](QuotaApi.md#GetQuota) | **Get** /quotas/{id} | Get the specified quota
[**ListQuotas**](QuotaApi.md#ListQuotas) | **Get** /quotas | List quotas
[**UpdateQuota**](QuotaApi.md#UpdateQuota) | **Put** /quotas/{id} | Update the specified quota


# **GetQuota**
> Quota GetQuota(ctx, id, optional)
Get the specified quota

Get the specified quota

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Quota ID | 
 **optional** | ***QuotaApiGetQuotaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotaApiGetQuotaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**Quota**](Quota.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuotas**
> []Quota ListQuotas(ctx, optional)
List quotas

List quotas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotaApiListQuotasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotaApiListQuotasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]
 **reference** | **optional.String**| The reference type of quota. | 
 **referenceId** | **optional.String**| The reference id of quota. | 
 **sort** | **optional.String**| Sort method, valid values include: &#39;hard.resource_name&#39;, &#39;-hard.resource_name&#39;, &#39;used.resource_name&#39;, &#39;-used.resource_name&#39;. Here &#39;-&#39; stands for descending order, resource_name should be the real resource name of the quota.  | 

### Return type

[**[]Quota**](Quota.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuota**
> UpdateQuota(ctx, id, hard, optional)
Update the specified quota

Update hard limits of the specified quota

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Quota ID | 
  **hard** | [**QuotaUpdateReq**](QuotaUpdateReq.md)| The new hard limits for the quota | 
 **optional** | ***QuotaApiUpdateQuotaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotaApiUpdateQuotaOpts struct

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

