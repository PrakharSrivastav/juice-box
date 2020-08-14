# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationActionSetAuthenticationMethod**](AuthenticationApi.md#AuthenticationActionSetAuthenticationMethod) | **Get** /JSON/authentication/action/setAuthenticationMethod/ | 
[**AuthenticationActionSetLoggedInIndicator**](AuthenticationApi.md#AuthenticationActionSetLoggedInIndicator) | **Get** /JSON/authentication/action/setLoggedInIndicator/ | 
[**AuthenticationActionSetLoggedOutIndicator**](AuthenticationApi.md#AuthenticationActionSetLoggedOutIndicator) | **Get** /JSON/authentication/action/setLoggedOutIndicator/ | 
[**AuthenticationViewGetAuthenticationMethod**](AuthenticationApi.md#AuthenticationViewGetAuthenticationMethod) | **Get** /JSON/authentication/view/getAuthenticationMethod/ | 
[**AuthenticationViewGetAuthenticationMethodConfigParams**](AuthenticationApi.md#AuthenticationViewGetAuthenticationMethodConfigParams) | **Get** /JSON/authentication/view/getAuthenticationMethodConfigParams/ | 
[**AuthenticationViewGetLoggedInIndicator**](AuthenticationApi.md#AuthenticationViewGetLoggedInIndicator) | **Get** /JSON/authentication/view/getLoggedInIndicator/ | 
[**AuthenticationViewGetLoggedOutIndicator**](AuthenticationApi.md#AuthenticationViewGetLoggedOutIndicator) | **Get** /JSON/authentication/view/getLoggedOutIndicator/ | 
[**AuthenticationViewGetSupportedAuthenticationMethods**](AuthenticationApi.md#AuthenticationViewGetSupportedAuthenticationMethods) | **Get** /JSON/authentication/view/getSupportedAuthenticationMethods/ | 

# **AuthenticationActionSetAuthenticationMethod**
> ModelError AuthenticationActionSetAuthenticationMethod(ctx, contextId, authMethodName, optional)


Sets the authentication method for the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **authMethodName** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationActionSetAuthenticationMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationActionSetAuthenticationMethodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authMethodConfigParams** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationActionSetLoggedInIndicator**
> ModelError AuthenticationActionSetLoggedInIndicator(ctx, contextId, loggedInIndicatorRegex)


Sets the logged in indicator for the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **loggedInIndicatorRegex** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationActionSetLoggedOutIndicator**
> ModelError AuthenticationActionSetLoggedOutIndicator(ctx, contextId, loggedOutIndicatorRegex)


Sets the logged out indicator for the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **loggedOutIndicatorRegex** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationViewGetAuthenticationMethod**
> ModelError AuthenticationViewGetAuthenticationMethod(ctx, contextId)


Gets the name of the authentication method for the context with the given ID.

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

# **AuthenticationViewGetAuthenticationMethodConfigParams**
> ModelError AuthenticationViewGetAuthenticationMethodConfigParams(ctx, authMethodName)


Gets the configuration parameters for the authentication method with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authMethodName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationViewGetLoggedInIndicator**
> ModelError AuthenticationViewGetLoggedInIndicator(ctx, contextId)


Gets the logged in indicator for the context with the given ID.

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

# **AuthenticationViewGetLoggedOutIndicator**
> ModelError AuthenticationViewGetLoggedOutIndicator(ctx, contextId)


Gets the logged out indicator for the context with the given ID.

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

# **AuthenticationViewGetSupportedAuthenticationMethods**
> ModelError AuthenticationViewGetSupportedAuthenticationMethods(ctx, )


Gets the name of the authentication methods.

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

