# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContextActionExcludeAllContextTechnologies**](ContextApi.md#ContextActionExcludeAllContextTechnologies) | **Get** /JSON/context/action/excludeAllContextTechnologies/ | 
[**ContextActionExcludeContextTechnologies**](ContextApi.md#ContextActionExcludeContextTechnologies) | **Get** /JSON/context/action/excludeContextTechnologies/ | 
[**ContextActionExcludeFromContext**](ContextApi.md#ContextActionExcludeFromContext) | **Get** /JSON/context/action/excludeFromContext/ | 
[**ContextActionExportContext**](ContextApi.md#ContextActionExportContext) | **Get** /JSON/context/action/exportContext/ | 
[**ContextActionImportContext**](ContextApi.md#ContextActionImportContext) | **Get** /JSON/context/action/importContext/ | 
[**ContextActionIncludeAllContextTechnologies**](ContextApi.md#ContextActionIncludeAllContextTechnologies) | **Get** /JSON/context/action/includeAllContextTechnologies/ | 
[**ContextActionIncludeContextTechnologies**](ContextApi.md#ContextActionIncludeContextTechnologies) | **Get** /JSON/context/action/includeContextTechnologies/ | 
[**ContextActionIncludeInContext**](ContextApi.md#ContextActionIncludeInContext) | **Get** /JSON/context/action/includeInContext/ | 
[**ContextActionNewContext**](ContextApi.md#ContextActionNewContext) | **Get** /JSON/context/action/newContext/ | 
[**ContextActionRemoveContext**](ContextApi.md#ContextActionRemoveContext) | **Get** /JSON/context/action/removeContext/ | 
[**ContextActionSetContextInScope**](ContextApi.md#ContextActionSetContextInScope) | **Get** /JSON/context/action/setContextInScope/ | 
[**ContextActionSetContextRegexs**](ContextApi.md#ContextActionSetContextRegexs) | **Get** /JSON/context/action/setContextRegexs/ | 
[**ContextViewContext**](ContextApi.md#ContextViewContext) | **Get** /JSON/context/view/context/ | 
[**ContextViewContextList**](ContextApi.md#ContextViewContextList) | **Get** /JSON/context/view/contextList/ | 
[**ContextViewExcludeRegexs**](ContextApi.md#ContextViewExcludeRegexs) | **Get** /JSON/context/view/excludeRegexs/ | 
[**ContextViewExcludedTechnologyList**](ContextApi.md#ContextViewExcludedTechnologyList) | **Get** /JSON/context/view/excludedTechnologyList/ | 
[**ContextViewIncludeRegexs**](ContextApi.md#ContextViewIncludeRegexs) | **Get** /JSON/context/view/includeRegexs/ | 
[**ContextViewIncludedTechnologyList**](ContextApi.md#ContextViewIncludedTechnologyList) | **Get** /JSON/context/view/includedTechnologyList/ | 
[**ContextViewTechnologyList**](ContextApi.md#ContextViewTechnologyList) | **Get** /JSON/context/view/technologyList/ | 
[**ContextViewUrls**](ContextApi.md#ContextViewUrls) | **Get** /JSON/context/view/urls/ | 

# **ContextActionExcludeAllContextTechnologies**
> ModelError ContextActionExcludeAllContextTechnologies(ctx, contextName)


Excludes all built in technologies from a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionExcludeContextTechnologies**
> ModelError ContextActionExcludeContextTechnologies(ctx, contextName, technologyNames)


Excludes technologies with the given names, separated by a comma, from a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 
  **technologyNames** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionExcludeFromContext**
> ModelError ContextActionExcludeFromContext(ctx, contextName, regex)


Add exclude regex to context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 
  **regex** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionExportContext**
> ModelError ContextActionExportContext(ctx, contextName, contextFile)


Exports the context with the given name to a file. If a relative file path is specified it will be resolved against the \"contexts\" directory in ZAP \"home\" dir.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 
  **contextFile** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionImportContext**
> ModelError ContextActionImportContext(ctx, contextFile)


Imports a context from a file. If a relative file path is specified it will be resolved against the \"contexts\" directory in ZAP \"home\" dir.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextFile** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionIncludeAllContextTechnologies**
> ModelError ContextActionIncludeAllContextTechnologies(ctx, contextName)


Includes all built in technologies in to a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionIncludeContextTechnologies**
> ModelError ContextActionIncludeContextTechnologies(ctx, contextName, technologyNames)


Includes technologies with the given names, separated by a comma, to a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 
  **technologyNames** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionIncludeInContext**
> ModelError ContextActionIncludeInContext(ctx, contextName, regex)


Add include regex to context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 
  **regex** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionNewContext**
> ModelError ContextActionNewContext(ctx, contextName)


Creates a new context with the given name in the current session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionRemoveContext**
> ModelError ContextActionRemoveContext(ctx, contextName)


Removes a context in the current session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionSetContextInScope**
> ModelError ContextActionSetContextInScope(ctx, contextName, booleanInScope)


Sets a context to in scope (contexts are in scope by default)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 
  **booleanInScope** | **bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextActionSetContextRegexs**
> ModelError ContextActionSetContextRegexs(ctx, contextName, incRegexs, excRegexs)


Set the regexs to include and exclude for a context, both supplied as JSON string arrays

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 
  **incRegexs** | **string**|  | 
  **excRegexs** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextViewContext**
> ModelError ContextViewContext(ctx, contextName)


List the information about the named context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextViewContextList**
> ModelError ContextViewContextList(ctx, )


List context names of current session

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

# **ContextViewExcludeRegexs**
> ModelError ContextViewExcludeRegexs(ctx, contextName)


List excluded regexs for context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextViewExcludedTechnologyList**
> ModelError ContextViewExcludedTechnologyList(ctx, contextName)


Lists the names of all technologies excluded from a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextViewIncludeRegexs**
> ModelError ContextViewIncludeRegexs(ctx, contextName)


List included regexs for context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextViewIncludedTechnologyList**
> ModelError ContextViewIncludedTechnologyList(ctx, contextName)


Lists the names of all technologies included in a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContextViewTechnologyList**
> ModelError ContextViewTechnologyList(ctx, )


Lists the names of all built in technologies

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

# **ContextViewUrls**
> ModelError ContextViewUrls(ctx, contextName)


Lists the URLs accessed through/by ZAP, that belong to the context with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

