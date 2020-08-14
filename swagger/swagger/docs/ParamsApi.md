# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParamsViewParams**](ParamsApi.md#ParamsViewParams) | **Get** /JSON/params/view/params/ | 

# **ParamsViewParams**
> ModelError ParamsViewParams(ctx, optional)


Shows the parameters for the specified site, or for all sites if the site is not specified

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ParamsApiParamsViewParamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParamsApiParamsViewParamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **site** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

