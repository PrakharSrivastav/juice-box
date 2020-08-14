# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BreakActionAddHttpBreakpoint**](BreakApi.md#BreakActionAddHttpBreakpoint) | **Get** /JSON/break/action/addHttpBreakpoint/ | 
[**BreakActionBreak**](BreakApi.md#BreakActionBreak) | **Get** /JSON/break/action/break/ | 
[**BreakActionContinue**](BreakApi.md#BreakActionContinue) | **Get** /JSON/break/action/continue/ | 
[**BreakActionDrop**](BreakApi.md#BreakActionDrop) | **Get** /JSON/break/action/drop/ | 
[**BreakActionRemoveHttpBreakpoint**](BreakApi.md#BreakActionRemoveHttpBreakpoint) | **Get** /JSON/break/action/removeHttpBreakpoint/ | 
[**BreakActionSetHttpMessage**](BreakApi.md#BreakActionSetHttpMessage) | **Get** /JSON/break/action/setHttpMessage/ | 
[**BreakActionStep**](BreakApi.md#BreakActionStep) | **Get** /JSON/break/action/step/ | 
[**BreakViewHttpMessage**](BreakApi.md#BreakViewHttpMessage) | **Get** /JSON/break/view/httpMessage/ | 
[**BreakViewIsBreakAll**](BreakApi.md#BreakViewIsBreakAll) | **Get** /JSON/break/view/isBreakAll/ | 
[**BreakViewIsBreakRequest**](BreakApi.md#BreakViewIsBreakRequest) | **Get** /JSON/break/view/isBreakRequest/ | 
[**BreakViewIsBreakResponse**](BreakApi.md#BreakViewIsBreakResponse) | **Get** /JSON/break/view/isBreakResponse/ | 

# **BreakActionAddHttpBreakpoint**
> ModelError BreakActionAddHttpBreakpoint(ctx, string_, location, match, inverse, ignorecase)


Adds a custom HTTP breakpoint. The string is the string to match. Location may be one of: url, request_header, request_body, response_header or response_body. Match may be: contains or regex. Inverse (match) may be true or false. Lastly, ignorecase (when matching the string) may be true or false.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **string_** | **string**|  | 
  **location** | **string**|  | 
  **match** | **string**|  | 
  **inverse** | **bool**|  | 
  **ignorecase** | **bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BreakActionBreak**
> ModelError BreakActionBreak(ctx, type_, state, optional)


Controls the global break functionality. The type may be one of: http-all, http-request or http-response. The state may be true (for turning break on for the specified type) or false (for turning break off). Scope is not currently used.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **state** | **string**|  | 
 **optional** | ***BreakApiBreakActionBreakOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BreakApiBreakActionBreakOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BreakActionContinue**
> ModelError BreakActionContinue(ctx, )


Submits the currently intercepted message and unsets the global request/response break points

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

# **BreakActionDrop**
> ModelError BreakActionDrop(ctx, )


Drops the currently intercepted message

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

# **BreakActionRemoveHttpBreakpoint**
> ModelError BreakActionRemoveHttpBreakpoint(ctx, string_, location, match, inverse, ignorecase)


Removes the specified break point

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **string_** | **string**|  | 
  **location** | **string**|  | 
  **match** | **string**|  | 
  **inverse** | **bool**|  | 
  **ignorecase** | **bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BreakActionSetHttpMessage**
> ModelError BreakActionSetHttpMessage(ctx, httpHeader, optional)


Overwrites the currently intercepted message with the data provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **httpHeader** | **string**|  | 
 **optional** | ***BreakApiBreakActionSetHttpMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BreakApiBreakActionSetHttpMessageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **httpBody** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BreakActionStep**
> ModelError BreakActionStep(ctx, )


Submits the currently intercepted message, the next request or response will automatically be intercepted

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

# **BreakViewHttpMessage**
> ModelError BreakViewHttpMessage(ctx, )


Returns the HTTP message currently intercepted (if any)

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

# **BreakViewIsBreakAll**
> ModelError BreakViewIsBreakAll(ctx, )


Returns True if ZAP will break on both requests and responses

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

# **BreakViewIsBreakRequest**
> ModelError BreakViewIsBreakRequest(ctx, )


Returns True if ZAP will break on requests

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

# **BreakViewIsBreakResponse**
> ModelError BreakViewIsBreakResponse(ctx, )


Returns True if ZAP will break on responses

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

