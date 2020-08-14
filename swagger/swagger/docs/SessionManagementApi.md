# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionManagementActionSetSessionManagementMethod**](SessionManagementApi.md#SessionManagementActionSetSessionManagementMethod) | **Get** /JSON/sessionManagement/action/setSessionManagementMethod/ | 
[**SessionManagementViewGetSessionManagementMethod**](SessionManagementApi.md#SessionManagementViewGetSessionManagementMethod) | **Get** /JSON/sessionManagement/view/getSessionManagementMethod/ | 
[**SessionManagementViewGetSessionManagementMethodConfigParams**](SessionManagementApi.md#SessionManagementViewGetSessionManagementMethodConfigParams) | **Get** /JSON/sessionManagement/view/getSessionManagementMethodConfigParams/ | 
[**SessionManagementViewGetSupportedSessionManagementMethods**](SessionManagementApi.md#SessionManagementViewGetSupportedSessionManagementMethods) | **Get** /JSON/sessionManagement/view/getSupportedSessionManagementMethods/ | 

# **SessionManagementActionSetSessionManagementMethod**
> ModelError SessionManagementActionSetSessionManagementMethod(ctx, contextId, methodName, optional)


Sets the session management method for the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **methodName** | **string**|  | 
 **optional** | ***SessionManagementApiSessionManagementActionSetSessionManagementMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SessionManagementApiSessionManagementActionSetSessionManagementMethodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **methodConfigParams** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionManagementViewGetSessionManagementMethod**
> ModelError SessionManagementViewGetSessionManagementMethod(ctx, contextId)


Gets the name of the session management method for the context with the given ID.

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

# **SessionManagementViewGetSessionManagementMethodConfigParams**
> ModelError SessionManagementViewGetSessionManagementMethodConfigParams(ctx, methodName)


Gets the configuration parameters for the session management method with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **methodName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionManagementViewGetSupportedSessionManagementMethods**
> ModelError SessionManagementViewGetSupportedSessionManagementMethods(ctx, )


Gets the name of the session management methods.

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

