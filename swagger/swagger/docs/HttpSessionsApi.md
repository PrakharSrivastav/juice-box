# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HttpSessionsActionAddDefaultSessionToken**](HttpSessionsApi.md#HttpSessionsActionAddDefaultSessionToken) | **Get** /JSON/httpSessions/action/addDefaultSessionToken/ | 
[**HttpSessionsActionAddSessionToken**](HttpSessionsApi.md#HttpSessionsActionAddSessionToken) | **Get** /JSON/httpSessions/action/addSessionToken/ | 
[**HttpSessionsActionCreateEmptySession**](HttpSessionsApi.md#HttpSessionsActionCreateEmptySession) | **Get** /JSON/httpSessions/action/createEmptySession/ | 
[**HttpSessionsActionRemoveDefaultSessionToken**](HttpSessionsApi.md#HttpSessionsActionRemoveDefaultSessionToken) | **Get** /JSON/httpSessions/action/removeDefaultSessionToken/ | 
[**HttpSessionsActionRemoveSession**](HttpSessionsApi.md#HttpSessionsActionRemoveSession) | **Get** /JSON/httpSessions/action/removeSession/ | 
[**HttpSessionsActionRemoveSessionToken**](HttpSessionsApi.md#HttpSessionsActionRemoveSessionToken) | **Get** /JSON/httpSessions/action/removeSessionToken/ | 
[**HttpSessionsActionRenameSession**](HttpSessionsApi.md#HttpSessionsActionRenameSession) | **Get** /JSON/httpSessions/action/renameSession/ | 
[**HttpSessionsActionSetActiveSession**](HttpSessionsApi.md#HttpSessionsActionSetActiveSession) | **Get** /JSON/httpSessions/action/setActiveSession/ | 
[**HttpSessionsActionSetDefaultSessionTokenEnabled**](HttpSessionsApi.md#HttpSessionsActionSetDefaultSessionTokenEnabled) | **Get** /JSON/httpSessions/action/setDefaultSessionTokenEnabled/ | 
[**HttpSessionsActionSetSessionTokenValue**](HttpSessionsApi.md#HttpSessionsActionSetSessionTokenValue) | **Get** /JSON/httpSessions/action/setSessionTokenValue/ | 
[**HttpSessionsActionUnsetActiveSession**](HttpSessionsApi.md#HttpSessionsActionUnsetActiveSession) | **Get** /JSON/httpSessions/action/unsetActiveSession/ | 
[**HttpSessionsViewActiveSession**](HttpSessionsApi.md#HttpSessionsViewActiveSession) | **Get** /JSON/httpSessions/view/activeSession/ | 
[**HttpSessionsViewDefaultSessionTokens**](HttpSessionsApi.md#HttpSessionsViewDefaultSessionTokens) | **Get** /JSON/httpSessions/view/defaultSessionTokens/ | 
[**HttpSessionsViewSessionTokens**](HttpSessionsApi.md#HttpSessionsViewSessionTokens) | **Get** /JSON/httpSessions/view/sessionTokens/ | 
[**HttpSessionsViewSessions**](HttpSessionsApi.md#HttpSessionsViewSessions) | **Get** /JSON/httpSessions/view/sessions/ | 
[**HttpSessionsViewSites**](HttpSessionsApi.md#HttpSessionsViewSites) | **Get** /JSON/httpSessions/view/sites/ | 

# **HttpSessionsActionAddDefaultSessionToken**
> ModelError HttpSessionsActionAddDefaultSessionToken(ctx, sessionToken, optional)


Adds a default session token with the given name and enabled state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionToken** | **string**|  | 
 **optional** | ***HttpSessionsApiHttpSessionsActionAddDefaultSessionTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HttpSessionsApiHttpSessionsActionAddDefaultSessionTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenEnabled** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionAddSessionToken**
> ModelError HttpSessionsActionAddSessionToken(ctx, site, sessionToken)


Adds the session token to the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionCreateEmptySession**
> ModelError HttpSessionsActionCreateEmptySession(ctx, site, optional)


Creates an empty session for the given site. Optionally with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
 **optional** | ***HttpSessionsApiHttpSessionsActionCreateEmptySessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HttpSessionsApiHttpSessionsActionCreateEmptySessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **session** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionRemoveDefaultSessionToken**
> ModelError HttpSessionsActionRemoveDefaultSessionToken(ctx, sessionToken)


Removes the default session token with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionToken** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionRemoveSession**
> ModelError HttpSessionsActionRemoveSession(ctx, site, session)


Removes the session from the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
  **session** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionRemoveSessionToken**
> ModelError HttpSessionsActionRemoveSessionToken(ctx, site, sessionToken)


Removes the session token from the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionRenameSession**
> ModelError HttpSessionsActionRenameSession(ctx, site, oldSessionName, newSessionName)


Renames the session of the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
  **oldSessionName** | **string**|  | 
  **newSessionName** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionSetActiveSession**
> ModelError HttpSessionsActionSetActiveSession(ctx, site, session)


Sets the given session as active for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
  **session** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionSetDefaultSessionTokenEnabled**
> ModelError HttpSessionsActionSetDefaultSessionTokenEnabled(ctx, sessionToken, tokenEnabled)


Sets whether or not the default session token with the given name is enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionToken** | **string**|  | 
  **tokenEnabled** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionSetSessionTokenValue**
> ModelError HttpSessionsActionSetSessionTokenValue(ctx, site, session, sessionToken, tokenValue)


Sets the value of the session token of the given session for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
  **session** | **string**|  | 
  **sessionToken** | **string**|  | 
  **tokenValue** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsActionUnsetActiveSession**
> ModelError HttpSessionsActionUnsetActiveSession(ctx, site)


Unsets the active session of the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsViewActiveSession**
> ModelError HttpSessionsViewActiveSession(ctx, site)


Gets the name of the active session for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsViewDefaultSessionTokens**
> ModelError HttpSessionsViewDefaultSessionTokens(ctx, )


Gets the default session tokens.

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

# **HttpSessionsViewSessionTokens**
> ModelError HttpSessionsViewSessionTokens(ctx, site)


Gets the names of the session tokens for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsViewSessions**
> ModelError HttpSessionsViewSessions(ctx, site, session)


Gets the sessions for the given site. Optionally returning just the session with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
  **session** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HttpSessionsViewSites**
> ModelError HttpSessionsViewSites(ctx, )


Gets all of the sites that have sessions.

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

