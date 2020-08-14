# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcsrfActionAddOptionToken**](AcsrfApi.md#AcsrfActionAddOptionToken) | **Get** /JSON/acsrf/action/addOptionToken/ | 
[**AcsrfActionRemoveOptionToken**](AcsrfApi.md#AcsrfActionRemoveOptionToken) | **Get** /JSON/acsrf/action/removeOptionToken/ | 
[**AcsrfOtherGenForm**](AcsrfApi.md#AcsrfOtherGenForm) | **Get** /OTHER/acsrf/other/genForm/ | 
[**AcsrfViewOptionTokensNames**](AcsrfApi.md#AcsrfViewOptionTokensNames) | **Get** /JSON/acsrf/view/optionTokensNames/ | 

# **AcsrfActionAddOptionToken**
> ModelError AcsrfActionAddOptionToken(ctx, string_)


Adds an anti-CSRF token with the given name, enabled by default

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **string_** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AcsrfActionRemoveOptionToken**
> ModelError AcsrfActionRemoveOptionToken(ctx, string_)


Removes the anti-CSRF token with the given name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **string_** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AcsrfOtherGenForm**
> ModelError AcsrfOtherGenForm(ctx, hrefId)


Generate a form for testing lack of anti-CSRF tokens - typically invoked via ZAP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hrefId** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AcsrfViewOptionTokensNames**
> ModelError AcsrfViewOptionTokensNames(ctx, )


Lists the names of all anti-CSRF tokens

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

