# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PscanActionDisableAllScanners**](PscanApi.md#PscanActionDisableAllScanners) | **Get** /JSON/pscan/action/disableAllScanners/ | 
[**PscanActionDisableScanners**](PscanApi.md#PscanActionDisableScanners) | **Get** /JSON/pscan/action/disableScanners/ | 
[**PscanActionEnableAllScanners**](PscanApi.md#PscanActionEnableAllScanners) | **Get** /JSON/pscan/action/enableAllScanners/ | 
[**PscanActionEnableScanners**](PscanApi.md#PscanActionEnableScanners) | **Get** /JSON/pscan/action/enableScanners/ | 
[**PscanActionSetEnabled**](PscanApi.md#PscanActionSetEnabled) | **Get** /JSON/pscan/action/setEnabled/ | 
[**PscanActionSetMaxAlertsPerRule**](PscanApi.md#PscanActionSetMaxAlertsPerRule) | **Get** /JSON/pscan/action/setMaxAlertsPerRule/ | 
[**PscanActionSetScanOnlyInScope**](PscanApi.md#PscanActionSetScanOnlyInScope) | **Get** /JSON/pscan/action/setScanOnlyInScope/ | 
[**PscanActionSetScannerAlertThreshold**](PscanApi.md#PscanActionSetScannerAlertThreshold) | **Get** /JSON/pscan/action/setScannerAlertThreshold/ | 
[**PscanViewCurrentRule**](PscanApi.md#PscanViewCurrentRule) | **Get** /JSON/pscan/view/currentRule/ | 
[**PscanViewMaxAlertsPerRule**](PscanApi.md#PscanViewMaxAlertsPerRule) | **Get** /JSON/pscan/view/maxAlertsPerRule/ | 
[**PscanViewRecordsToScan**](PscanApi.md#PscanViewRecordsToScan) | **Get** /JSON/pscan/view/recordsToScan/ | 
[**PscanViewScanOnlyInScope**](PscanApi.md#PscanViewScanOnlyInScope) | **Get** /JSON/pscan/view/scanOnlyInScope/ | 
[**PscanViewScanners**](PscanApi.md#PscanViewScanners) | **Get** /JSON/pscan/view/scanners/ | 

# **PscanActionDisableAllScanners**
> ModelError PscanActionDisableAllScanners(ctx, )


Disables all passive scanners

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

# **PscanActionDisableScanners**
> ModelError PscanActionDisableScanners(ctx, ids)


Disables all passive scanners with the given IDs (comma separated list of IDs)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PscanActionEnableAllScanners**
> ModelError PscanActionEnableAllScanners(ctx, )


Enables all passive scanners

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

# **PscanActionEnableScanners**
> ModelError PscanActionEnableScanners(ctx, ids)


Enables all passive scanners with the given IDs (comma separated list of IDs)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PscanActionSetEnabled**
> ModelError PscanActionSetEnabled(ctx, enabled)


Sets whether or not the passive scanning is enabled (Note: the enabled state is not persisted).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enabled** | **bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PscanActionSetMaxAlertsPerRule**
> ModelError PscanActionSetMaxAlertsPerRule(ctx, maxAlerts)


Sets the maximum number of alerts a passive scan rule should raise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maxAlerts** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PscanActionSetScanOnlyInScope**
> ModelError PscanActionSetScanOnlyInScope(ctx, onlyInScope)


Sets whether or not the passive scan should be performed only on messages that are in scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **onlyInScope** | **bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PscanActionSetScannerAlertThreshold**
> ModelError PscanActionSetScannerAlertThreshold(ctx, id, alertThreshold)


Sets the alert threshold of the passive scanner with the given ID, accepted values for alert threshold: OFF, DEFAULT, LOW, MEDIUM and HIGH

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **alertThreshold** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PscanViewCurrentRule**
> ModelError PscanViewCurrentRule(ctx, )


Show information about the passive scan rule currently being run (if any).

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

# **PscanViewMaxAlertsPerRule**
> ModelError PscanViewMaxAlertsPerRule(ctx, )


Gets the maximum number of alerts a passive scan rule should raise.

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

# **PscanViewRecordsToScan**
> ModelError PscanViewRecordsToScan(ctx, )


The number of records the passive scanner still has to scan

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

# **PscanViewScanOnlyInScope**
> ModelError PscanViewScanOnlyInScope(ctx, )


Tells whether or not the passive scan should be performed only on messages that are in scope.

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

# **PscanViewScanners**
> ModelError PscanViewScanners(ctx, )


Lists all passive scanners with its ID, name, enabled state and alert threshold.

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

