# \ScannerApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScanner**](ScannerApi.md#CreateScanner) | **Post** /scanners | Create a scanner registration
[**DeleteScanner**](ScannerApi.md#DeleteScanner) | **Delete** /scanners/{registration_id} | Delete a scanner registration
[**GetScanner**](ScannerApi.md#GetScanner) | **Get** /scanners/{registration_id} | Get a scanner registration details
[**GetScannerMetadata**](ScannerApi.md#GetScannerMetadata) | **Get** /scanners/{registration_id}/metadata | Get the metadata of the specified scanner registration
[**ListScanners**](ScannerApi.md#ListScanners) | **Get** /scanners | List scanner registrations
[**PingScanner**](ScannerApi.md#PingScanner) | **Post** /scanners/ping | Tests scanner registration settings
[**SetScannerAsDefault**](ScannerApi.md#SetScannerAsDefault) | **Patch** /scanners/{registration_id} | Set system default scanner registration
[**UpdateScanner**](ScannerApi.md#UpdateScanner) | **Put** /scanners/{registration_id} | Update a scanner registration


# **CreateScanner**
> CreateScanner(ctx, registration, optional)
Create a scanner registration

Creats a new scanner registration with the given data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registration** | [**ScannerRegistrationReq**](ScannerRegistrationReq.md)| A scanner registration to be created. | 
 **optional** | ***ScannerApiCreateScannerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiCreateScannerOpts struct

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

# **DeleteScanner**
> ScannerRegistration DeleteScanner(ctx, registrationId, optional)
Delete a scanner registration

Deletes the specified scanner registration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifier. | 
 **optional** | ***ScannerApiDeleteScannerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiDeleteScannerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScanner**
> ScannerRegistration GetScanner(ctx, registrationId, optional)
Get a scanner registration details

Retruns the details of the specified scanner registration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifer. | 
 **optional** | ***ScannerApiGetScannerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiGetScannerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScannerMetadata**
> ScannerAdapterMetadata GetScannerMetadata(ctx, registrationId, optional)
Get the metadata of the specified scanner registration

Get the metadata of the specified scanner registration, including the capabilities and customized properties. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifier. | 
 **optional** | ***ScannerApiGetScannerMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiGetScannerMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| An unique ID for the request | 

### Return type

[**ScannerAdapterMetadata**](ScannerAdapterMetadata.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListScanners**
> []ScannerRegistration ListScanners(ctx, optional)
List scanner registrations

Returns a list of currently configured scanner registrations. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScannerApiListScannersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiListScannersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| An unique ID for the request | 
 **q** | **optional.String**| Query string to query resources. Supported query patterns are \&quot;exact match(k&#x3D;v)\&quot;, \&quot;fuzzy match(k&#x3D;~v)\&quot;, \&quot;range(k&#x3D;[min~max])\&quot;, \&quot;list with union releationship(k&#x3D;{v1 v2 v3})\&quot; and \&quot;list with intersetion relationship(k&#x3D;(v1 v2 v3))\&quot;. The value of range and list can be string(enclosed by \&quot; or &#39;), integer or time(in format \&quot;2020-04-09 02:36:00\&quot;). All of these query patterns should be put in the query string \&quot;q&#x3D;xxx\&quot; and splitted by \&quot;,\&quot;. e.g. q&#x3D;k1&#x3D;v1,k2&#x3D;~v2,k3&#x3D;[min~max] | 
 **sort** | **optional.String**| Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with \&quot;sort&#x3D;field1,-field2\&quot; | 
 **page** | **optional.Int64**| The page number | [default to 1]
 **pageSize** | **optional.Int64**| The size of per page | [default to 10]

### Return type

[**[]ScannerRegistration**](ScannerRegistration.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingScanner**
> PingScanner(ctx, settings, optional)
Tests scanner registration settings

Pings scanner adapter to test endpoint URL and authorization settings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settings** | [**ScannerRegistrationSettings**](ScannerRegistrationSettings.md)| A scanner registration settings to be tested. | 
 **optional** | ***ScannerApiPingScannerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiPingScannerOpts struct

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

# **SetScannerAsDefault**
> SetScannerAsDefault(ctx, registrationId, payload, optional)
Set system default scanner registration

Set the specified scanner registration as the system default one. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifier. | 
  **payload** | [**IsDefault**](IsDefault.md)|  | 
 **optional** | ***ScannerApiSetScannerAsDefaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiSetScannerAsDefaultOpts struct

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

# **UpdateScanner**
> UpdateScanner(ctx, registrationId, registration, optional)
Update a scanner registration

Updates the specified scanner registration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registrationId** | **string**| The scanner registration identifier. | 
  **registration** | [**ScannerRegistrationReq**](ScannerRegistrationReq.md)| A scanner registraiton to be updated. | 
 **optional** | ***ScannerApiUpdateScannerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScannerApiUpdateScannerOpts struct

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

