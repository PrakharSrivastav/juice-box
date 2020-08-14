# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForcedUserActionSetForcedUser**](ForcedUserApi.md#ForcedUserActionSetForcedUser) | **Get** /JSON/forcedUser/action/setForcedUser/ | 
[**ForcedUserActionSetForcedUserModeEnabled**](ForcedUserApi.md#ForcedUserActionSetForcedUserModeEnabled) | **Get** /JSON/forcedUser/action/setForcedUserModeEnabled/ | 
[**ForcedUserViewGetForcedUser**](ForcedUserApi.md#ForcedUserViewGetForcedUser) | **Get** /JSON/forcedUser/view/getForcedUser/ | 
[**ForcedUserViewIsForcedUserModeEnabled**](ForcedUserApi.md#ForcedUserViewIsForcedUserModeEnabled) | **Get** /JSON/forcedUser/view/isForcedUserModeEnabled/ | 

# **ForcedUserActionSetForcedUser**
> ModelError ForcedUserActionSetForcedUser(ctx, contextId, userId)


Sets the user (ID) that should be used in 'forced user' mode for the given context (ID)

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

# **ForcedUserActionSetForcedUserModeEnabled**
> ModelError ForcedUserActionSetForcedUserModeEnabled(ctx, boolean)


Sets if 'forced user' mode should be enabled or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **boolean** | **bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForcedUserViewGetForcedUser**
> ModelError ForcedUserViewGetForcedUser(ctx, contextId)


Gets the user (ID) set as 'forced user' for the given context (ID)

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

# **ForcedUserViewIsForcedUserModeEnabled**
> ModelError ForcedUserViewIsForcedUserModeEnabled(ctx, )


Returns 'true' if 'forced user' mode is enabled, 'false' otherwise

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

