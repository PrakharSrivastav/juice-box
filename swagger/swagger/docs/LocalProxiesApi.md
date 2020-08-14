# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalProxiesActionAddAdditionalProxy**](LocalProxiesApi.md#LocalProxiesActionAddAdditionalProxy) | **Get** /JSON/localProxies/action/addAdditionalProxy/ | 
[**LocalProxiesActionRemoveAdditionalProxy**](LocalProxiesApi.md#LocalProxiesActionRemoveAdditionalProxy) | **Get** /JSON/localProxies/action/removeAdditionalProxy/ | 
[**LocalProxiesViewAdditionalProxies**](LocalProxiesApi.md#LocalProxiesViewAdditionalProxies) | **Get** /JSON/localProxies/view/additionalProxies/ | 

# **LocalProxiesActionAddAdditionalProxy**
> ModelError LocalProxiesActionAddAdditionalProxy(ctx, address, port, optional)


Adds an new proxy using the details supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**|  | 
  **port** | **int32**|  | 
 **optional** | ***LocalProxiesApiLocalProxiesActionAddAdditionalProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalProxiesApiLocalProxiesActionAddAdditionalProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **behindNat** | **optional.Bool**|  | 
 **alwaysDecodeZip** | **optional.Bool**|  | 
 **removeUnsupportedEncodings** | **optional.Bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalProxiesActionRemoveAdditionalProxy**
> ModelError LocalProxiesActionRemoveAdditionalProxy(ctx, address, port)


Removes the additional proxy with the specified address and port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**|  | 
  **port** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalProxiesViewAdditionalProxies**
> ModelError LocalProxiesViewAdditionalProxies(ctx, )


Gets all of the additional proxies that have been configured.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

