# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScriptActionClearGlobalVar**](ScriptApi.md#ScriptActionClearGlobalVar) | **Get** /JSON/script/action/clearGlobalVar/ | 
[**ScriptActionClearGlobalVars**](ScriptApi.md#ScriptActionClearGlobalVars) | **Get** /JSON/script/action/clearGlobalVars/ | 
[**ScriptActionClearScriptVar**](ScriptApi.md#ScriptActionClearScriptVar) | **Get** /JSON/script/action/clearScriptVar/ | 
[**ScriptActionClearScriptVars**](ScriptApi.md#ScriptActionClearScriptVars) | **Get** /JSON/script/action/clearScriptVars/ | 
[**ScriptActionDisable**](ScriptApi.md#ScriptActionDisable) | **Get** /JSON/script/action/disable/ | 
[**ScriptActionEnable**](ScriptApi.md#ScriptActionEnable) | **Get** /JSON/script/action/enable/ | 
[**ScriptActionLoad**](ScriptApi.md#ScriptActionLoad) | **Get** /JSON/script/action/load/ | 
[**ScriptActionRemove**](ScriptApi.md#ScriptActionRemove) | **Get** /JSON/script/action/remove/ | 
[**ScriptActionRunStandAloneScript**](ScriptApi.md#ScriptActionRunStandAloneScript) | **Get** /JSON/script/action/runStandAloneScript/ | 
[**ScriptActionSetGlobalVar**](ScriptApi.md#ScriptActionSetGlobalVar) | **Get** /JSON/script/action/setGlobalVar/ | 
[**ScriptActionSetScriptVar**](ScriptApi.md#ScriptActionSetScriptVar) | **Get** /JSON/script/action/setScriptVar/ | 
[**ScriptViewGlobalVar**](ScriptApi.md#ScriptViewGlobalVar) | **Get** /JSON/script/view/globalVar/ | 
[**ScriptViewGlobalVars**](ScriptApi.md#ScriptViewGlobalVars) | **Get** /JSON/script/view/globalVars/ | 
[**ScriptViewListEngines**](ScriptApi.md#ScriptViewListEngines) | **Get** /JSON/script/view/listEngines/ | 
[**ScriptViewListScripts**](ScriptApi.md#ScriptViewListScripts) | **Get** /JSON/script/view/listScripts/ | 
[**ScriptViewListTypes**](ScriptApi.md#ScriptViewListTypes) | **Get** /JSON/script/view/listTypes/ | 
[**ScriptViewScriptVar**](ScriptApi.md#ScriptViewScriptVar) | **Get** /JSON/script/view/scriptVar/ | 
[**ScriptViewScriptVars**](ScriptApi.md#ScriptViewScriptVars) | **Get** /JSON/script/view/scriptVars/ | 

# **ScriptActionClearGlobalVar**
> ModelError ScriptActionClearGlobalVar(ctx, varKey)


Clears the global variable with the given key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **varKey** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionClearGlobalVars**
> ModelError ScriptActionClearGlobalVars(ctx, )


Clears the global variables.

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

# **ScriptActionClearScriptVar**
> ModelError ScriptActionClearScriptVar(ctx, scriptName, varKey)


Clears the variable with the given key of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 
  **varKey** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionClearScriptVars**
> ModelError ScriptActionClearScriptVars(ctx, scriptName)


Clears the variables of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionDisable**
> ModelError ScriptActionDisable(ctx, scriptName)


Disables the script with the given name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionEnable**
> ModelError ScriptActionEnable(ctx, scriptName)


Enables the script with the given name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionLoad**
> ModelError ScriptActionLoad(ctx, scriptName, scriptType, scriptEngine, fileName, optional)


Loads a script into ZAP from the given local file, with the given name, type and engine, optionally with a description, and a charset name to read the script (the charset name is required if the script is not in UTF-8, for example, in ISO-8859-1).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 
  **scriptType** | **string**|  | 
  **scriptEngine** | **string**|  | 
  **fileName** | **string**|  | 
 **optional** | ***ScriptApiScriptActionLoadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScriptApiScriptActionLoadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **scriptDescription** | **optional.String**|  | 
 **charset** | **optional.String**|  | [default to UTF-8]

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionRemove**
> ModelError ScriptActionRemove(ctx, scriptName)


Removes the script with the given name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionRunStandAloneScript**
> ModelError ScriptActionRunStandAloneScript(ctx, scriptName)


Runs the stand alone script with the given name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionSetGlobalVar**
> ModelError ScriptActionSetGlobalVar(ctx, varKey, optional)


Sets the value of the global variable with the given key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **varKey** | **string**|  | 
 **optional** | ***ScriptApiScriptActionSetGlobalVarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScriptApiScriptActionSetGlobalVarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **varValue** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptActionSetScriptVar**
> ModelError ScriptActionSetScriptVar(ctx, scriptName, varKey, optional)


Sets the value of the variable with the given key of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 
  **varKey** | **string**|  | 
 **optional** | ***ScriptApiScriptActionSetScriptVarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScriptApiScriptActionSetScriptVarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **varValue** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptViewGlobalVar**
> ModelError ScriptViewGlobalVar(ctx, varKey)


Gets the value of the global variable with the given key. Returns an API error (DOES_NOT_EXIST) if no value was previously set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **varKey** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptViewGlobalVars**
> ModelError ScriptViewGlobalVars(ctx, )


Gets all the global variables (key/value pairs).

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

# **ScriptViewListEngines**
> ModelError ScriptViewListEngines(ctx, )


Lists the script engines available

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

# **ScriptViewListScripts**
> ModelError ScriptViewListScripts(ctx, )


Lists the scripts available, with its engine, name, description, type and error state.

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

# **ScriptViewListTypes**
> ModelError ScriptViewListTypes(ctx, )


Lists the script types available.

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

# **ScriptViewScriptVar**
> ModelError ScriptViewScriptVar(ctx, scriptName, varKey)


Gets the value of the variable with the given key for the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists or if no value was previously set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 
  **varKey** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScriptViewScriptVars**
> ModelError ScriptViewScriptVars(ctx, scriptName)


Gets all the variables (key/value pairs) of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scriptName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

