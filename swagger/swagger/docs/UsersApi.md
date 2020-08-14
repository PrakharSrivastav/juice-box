# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersActionNewUser**](UsersApi.md#UsersActionNewUser) | **Get** /JSON/users/action/newUser/ | 
[**UsersActionRemoveUser**](UsersApi.md#UsersActionRemoveUser) | **Get** /JSON/users/action/removeUser/ | 
[**UsersActionSetAuthenticationCredentials**](UsersApi.md#UsersActionSetAuthenticationCredentials) | **Get** /JSON/users/action/setAuthenticationCredentials/ | 
[**UsersActionSetUserEnabled**](UsersApi.md#UsersActionSetUserEnabled) | **Get** /JSON/users/action/setUserEnabled/ | 
[**UsersActionSetUserName**](UsersApi.md#UsersActionSetUserName) | **Get** /JSON/users/action/setUserName/ | 
[**UsersViewGetAuthenticationCredentials**](UsersApi.md#UsersViewGetAuthenticationCredentials) | **Get** /JSON/users/view/getAuthenticationCredentials/ | 
[**UsersViewGetAuthenticationCredentialsConfigParams**](UsersApi.md#UsersViewGetAuthenticationCredentialsConfigParams) | **Get** /JSON/users/view/getAuthenticationCredentialsConfigParams/ | 
[**UsersViewGetUserById**](UsersApi.md#UsersViewGetUserById) | **Get** /JSON/users/view/getUserById/ | 
[**UsersViewUsersList**](UsersApi.md#UsersViewUsersList) | **Get** /JSON/users/view/usersList/ | 

# **UsersActionNewUser**
> ModelError UsersActionNewUser(ctx, contextId, name)


Creates a new user with the given name for the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **name** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersActionRemoveUser**
> ModelError UsersActionRemoveUser(ctx, contextId, userId)


Removes the user with the given ID that belongs to the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **userId** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersActionSetAuthenticationCredentials**
> ModelError UsersActionSetAuthenticationCredentials(ctx, contextId, userId, optional)


Sets the authentication credentials for the user with the given ID that belongs to the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **userId** | **int32**|  | 
 **optional** | ***UsersApiUsersActionSetAuthenticationCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersActionSetAuthenticationCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authCredentialsConfigParams** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersActionSetUserEnabled**
> ModelError UsersActionSetUserEnabled(ctx, contextId, userId, enabled)


Sets whether or not the user, with the given ID that belongs to the context with the given ID, should be enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **userId** | **int32**|  | 
  **enabled** | **bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersActionSetUserName**
> ModelError UsersActionSetUserName(ctx, contextId, userId, name)


Renames the user with the given ID that belongs to the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **userId** | **int32**|  | 
  **name** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersViewGetAuthenticationCredentials**
> ModelError UsersViewGetAuthenticationCredentials(ctx, contextId, userId)


Gets the authentication credentials of the user with given ID that belongs to the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **userId** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersViewGetAuthenticationCredentialsConfigParams**
> ModelError UsersViewGetAuthenticationCredentialsConfigParams(ctx, contextId)


Gets the configuration parameters for the credentials of the context with the given ID.

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

# **UsersViewGetUserById**
> ModelError UsersViewGetUserById(ctx, contextId, userId)


Gets the data of the user with the given ID that belongs to the context with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **int32**|  | 
  **userId** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersViewUsersList**
> ModelError UsersViewUsersList(ctx, optional)


Gets a list of users that belong to the context with the given ID, or all users if none provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersViewUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersViewUsersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextId** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

