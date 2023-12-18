# \OidcApi

All URIs are relative to *http://localhost/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PingOIDC**](OidcApi.md#PingOIDC) | **Post** /system/oidc/ping | Test the OIDC endpoint.


# **PingOIDC**
> PingOIDC(ctx, endpoint, optional)
Test the OIDC endpoint.

Test the OIDC endpoint, the setting of the endpoint is provided in the request.  This API can only be called by system admin. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpoint** | [**Endpoint**](Endpoint.md)| Request body for OIDC endpoint to be tested. | 
 **optional** | ***OidcApiPingOIDCOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OidcApiPingOIDCOpts struct

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

