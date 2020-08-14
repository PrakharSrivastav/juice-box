# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationActionSetBasicAuthorizationDetectionMethod**](AuthorizationApi.md#AuthorizationActionSetBasicAuthorizationDetectionMethod) | **Get** /JSON/authorization/action/setBasicAuthorizationDetectionMethod/ | 
[**AuthorizationViewGetAuthorizationDetectionMethod**](AuthorizationApi.md#AuthorizationViewGetAuthorizationDetectionMethod) | **Get** /JSON/authorization/view/getAuthorizationDetectionMethod/ | 

# **AuthorizationActionSetBasicAuthorizationDetectionMethod**
> ModelError AuthorizationActionSetBasicAuthorizationDetectionMethod(ctx, contextId, optional)


Sets the authorization detection method for a context as one that identifies un-authorized messages based on: the message's status code or a regex pattern in the response's header or body. Also, whether all conditions must match or just some can be specified via the logicalOperator parameter, which accepts two values: \"AND\" (default), \"OR\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
 **optional** | ***AuthorizationApiAuthorizationActionSetBasicAuthorizationDetectionMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationActionSetBasicAuthorizationDetectionMethodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **headerRegex** | **optional.String**|  | 
 **bodyRegex** | **optional.String**|  | 
 **statusCode** | **optional.String**|  | 
 **logicalOperator** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationViewGetAuthorizationDetectionMethod**
> ModelError AuthorizationViewGetAuthorizationDetectionMethod(ctx, contextId)


Obtains all the configuration of the authorization detection method that is currently set for a context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

