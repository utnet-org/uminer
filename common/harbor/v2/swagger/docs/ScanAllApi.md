# \ScanAllApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScanAllSchedule**](ScanAllApi.md#CreateScanAllSchedule) | **Post** /system/scanAll/schedule | Create a schedule or a manual trigger for the scan all job.
[**GetLatestScanAllMetrics**](ScanAllApi.md#GetLatestScanAllMetrics) | **Get** /scans/all/metrics | Get the metrics of the latest scan all process
[**GetLatestScheduledScanAllMetrics**](ScanAllApi.md#GetLatestScheduledScanAllMetrics) | **Get** /scans/schedule/metrics | Get the metrics of the latest scheduled scan all process
[**GetScanAllSchedule**](ScanAllApi.md#GetScanAllSchedule) | **Get** /system/scanAll/schedule | Get scan all&#39;s schedule.
[**UpdateScanAllSchedule**](ScanAllApi.md#UpdateScanAllSchedule) | **Put** /system/scanAll/schedule | Update scan all&#39;s schedule.


# **CreateScanAllSchedule**
> CreateScanAllSchedule(ctx, schedule, optional)
Create a schedule or a manual trigger for the scan all job.

This endpoint is for creating a schedule or a manual trigger for the scan all job, which scans all of images in Harbor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**Schedule**](Schedule.md)| Create a schedule or a manual trigger for the scan all job. | 
 **optional** | ***ScanAllApiCreateScanAllScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanAllApiCreateScanAllScheduleOpts struct

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

# **GetLatestScanAllMetrics**
> Stats GetLatestScanAllMetrics(ctx, optional)
Get the metrics of the latest scan all process

Get the metrics of the latest scan all process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScanAllApiGetLatestScanAllMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanAllApiGetLatestScanAllMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**Stats**](Stats.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestScheduledScanAllMetrics**
> Stats GetLatestScheduledScanAllMetrics(ctx, optional)
Get the metrics of the latest scheduled scan all process

Get the metrics of the latest scheduled scan all process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScanAllApiGetLatestScheduledScanAllMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanAllApiGetLatestScheduledScanAllMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**Stats**](Stats.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScanAllSchedule**
> Schedule GetScanAllSchedule(ctx, optional)
Get scan all's schedule.

This endpoint is for getting a schedule for the scan all job, which scans all of images in Harbor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScanAllApiGetScanAllScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanAllApiGetScanAllScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateScanAllSchedule**
> UpdateScanAllSchedule(ctx, schedule, optional)
Update scan all's schedule.

This endpoint is for updating the schedule of scan all job, which scans all of images in Harbor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**Schedule**](Schedule.md)| Updates the schedule of scan all job, which scans all of images in Harbor. | 
 **optional** | ***ScanAllApiUpdateScanAllScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScanAllApiUpdateScanAllScheduleOpts struct

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

