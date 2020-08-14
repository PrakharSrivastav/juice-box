# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatsActionClearStats**](StatsApi.md#StatsActionClearStats) | **Get** /JSON/stats/action/clearStats/ | 
[**StatsActionSetOptionInMemoryEnabled**](StatsApi.md#StatsActionSetOptionInMemoryEnabled) | **Get** /JSON/stats/action/setOptionInMemoryEnabled/ | 
[**StatsActionSetOptionStatsdHost**](StatsApi.md#StatsActionSetOptionStatsdHost) | **Get** /JSON/stats/action/setOptionStatsdHost/ | 
[**StatsActionSetOptionStatsdPort**](StatsApi.md#StatsActionSetOptionStatsdPort) | **Get** /JSON/stats/action/setOptionStatsdPort/ | 
[**StatsActionSetOptionStatsdPrefix**](StatsApi.md#StatsActionSetOptionStatsdPrefix) | **Get** /JSON/stats/action/setOptionStatsdPrefix/ | 
[**StatsViewAllSitesStats**](StatsApi.md#StatsViewAllSitesStats) | **Get** /JSON/stats/view/allSitesStats/ | 
[**StatsViewOptionInMemoryEnabled**](StatsApi.md#StatsViewOptionInMemoryEnabled) | **Get** /JSON/stats/view/optionInMemoryEnabled/ | 
[**StatsViewOptionStatsdEnabled**](StatsApi.md#StatsViewOptionStatsdEnabled) | **Get** /JSON/stats/view/optionStatsdEnabled/ | 
[**StatsViewOptionStatsdHost**](StatsApi.md#StatsViewOptionStatsdHost) | **Get** /JSON/stats/view/optionStatsdHost/ | 
[**StatsViewOptionStatsdPort**](StatsApi.md#StatsViewOptionStatsdPort) | **Get** /JSON/stats/view/optionStatsdPort/ | 
[**StatsViewOptionStatsdPrefix**](StatsApi.md#StatsViewOptionStatsdPrefix) | **Get** /JSON/stats/view/optionStatsdPrefix/ | 
[**StatsViewSiteStats**](StatsApi.md#StatsViewSiteStats) | **Get** /JSON/stats/view/siteStats/ | 
[**StatsViewStats**](StatsApi.md#StatsViewStats) | **Get** /JSON/stats/view/stats/ | 

# **StatsActionClearStats**
> ModelError StatsActionClearStats(ctx, optional)


Clears all of the statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatsApiStatsActionClearStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatsApiStatsActionClearStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyPrefix** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatsActionSetOptionInMemoryEnabled**
> ModelError StatsActionSetOptionInMemoryEnabled(ctx, boolean)


Sets whether in memory statistics are enabled

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **boolean** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatsActionSetOptionStatsdHost**
> ModelError StatsActionSetOptionStatsdHost(ctx, string_)


Sets the Statsd service hostname, supply an empty string to stop using a Statsd service

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

# **StatsActionSetOptionStatsdPort**
> ModelError StatsActionSetOptionStatsdPort(ctx, integer)


Sets the Statsd service port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integer** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatsActionSetOptionStatsdPrefix**
> ModelError StatsActionSetOptionStatsdPrefix(ctx, string_)


Sets the prefix to be applied to all stats sent to the configured Statsd service

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

# **StatsViewAllSitesStats**
> ModelError StatsViewAllSitesStats(ctx, optional)


Gets all of the site based statistics, optionally filtered by a key prefix

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatsApiStatsViewAllSitesStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatsApiStatsViewAllSitesStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyPrefix** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatsViewOptionInMemoryEnabled**
> ModelError StatsViewOptionInMemoryEnabled(ctx, )


Returns 'true' if in memory statistics are enabled, otherwise returns 'false'

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

# **StatsViewOptionStatsdEnabled**
> ModelError StatsViewOptionStatsdEnabled(ctx, )


Returns 'true' if a Statsd server has been correctly configured, otherwise returns 'false'

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

# **StatsViewOptionStatsdHost**
> ModelError StatsViewOptionStatsdHost(ctx, )


Gets the Statsd service hostname

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

# **StatsViewOptionStatsdPort**
> ModelError StatsViewOptionStatsdPort(ctx, )


Gets the Statsd service port

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

# **StatsViewOptionStatsdPrefix**
> ModelError StatsViewOptionStatsdPrefix(ctx, )


Gets the prefix to be applied to all stats sent to the configured Statsd service

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

# **StatsViewSiteStats**
> ModelError StatsViewSiteStats(ctx, site, optional)


Gets all of the global statistics, optionally filtered by a key prefix

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | 
 **optional** | ***StatsApiStatsViewSiteStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatsApiStatsViewSiteStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyPrefix** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatsViewStats**
> ModelError StatsViewStats(ctx, optional)


Statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatsApiStatsViewStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatsApiStatsViewStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyPrefix** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

