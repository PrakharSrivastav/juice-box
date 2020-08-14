# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertActionDeleteAlert**](AlertApi.md#AlertActionDeleteAlert) | **Get** /JSON/alert/action/deleteAlert/ | 
[**AlertActionDeleteAllAlerts**](AlertApi.md#AlertActionDeleteAllAlerts) | **Get** /JSON/alert/action/deleteAllAlerts/ | 
[**AlertViewAlert**](AlertApi.md#AlertViewAlert) | **Get** /JSON/alert/view/alert/ | 
[**AlertViewAlertCountsByRisk**](AlertApi.md#AlertViewAlertCountsByRisk) | **Get** /JSON/alert/view/alertCountsByRisk/ | 
[**AlertViewAlerts**](AlertApi.md#AlertViewAlerts) | **Get** /JSON/alert/view/alerts/ | 
[**AlertViewAlertsByRisk**](AlertApi.md#AlertViewAlertsByRisk) | **Get** /JSON/alert/view/alertsByRisk/ | 
[**AlertViewAlertsSummary**](AlertApi.md#AlertViewAlertsSummary) | **Get** /JSON/alert/view/alertsSummary/ | 
[**AlertViewNumberOfAlerts**](AlertApi.md#AlertViewNumberOfAlerts) | **Get** /JSON/alert/view/numberOfAlerts/ | 

# **AlertActionDeleteAlert**
> ModelError AlertActionDeleteAlert(ctx, id)


Deletes the alert with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertActionDeleteAllAlerts**
> ModelError AlertActionDeleteAllAlerts(ctx, )


Deletes all alerts of the current session.

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

# **AlertViewAlert**
> ModelError AlertViewAlert(ctx, optional)


Gets the alert with the given ID, the corresponding HTTP message can be obtained with the 'messageId' field and 'message' API method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertApiAlertViewAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertViewAlertOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertViewAlertCountsByRisk**
> ModelError AlertViewAlertCountsByRisk(ctx, optional)


Gets a count of the alerts, optionally filtered as per alertsPerRisk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertApiAlertViewAlertCountsByRiskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertViewAlertCountsByRiskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **optional.String**|  | 
 **recurse** | **optional.Bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertViewAlerts**
> ModelError AlertViewAlerts(ctx, optional)


Gets the alerts raised by ZAP, optionally filtering by URL or riskId, and paginating with 'start' position and 'count' of alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertApiAlertViewAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertViewAlertsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 
 **riskId** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertViewAlertsByRisk**
> ModelError AlertViewAlertsByRisk(ctx, optional)


Gets a summary of the alerts, optionally filtered by a 'url'. If 'recurse' is true then all alerts that apply to urls that start with the specified 'url' will be returned, otherwise only those on exactly the same 'url' (ignoring url parameters)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertApiAlertViewAlertsByRiskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertViewAlertsByRiskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **optional.String**|  | 
 **recurse** | **optional.Bool**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertViewAlertsSummary**
> ModelError AlertViewAlertsSummary(ctx, optional)


Gets number of alerts grouped by each risk level, optionally filtering by URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertApiAlertViewAlertsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertViewAlertsSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseurl** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertViewNumberOfAlerts**
> ModelError AlertViewNumberOfAlerts(ctx, optional)


Gets the number of alerts, optionally filtering by URL or riskId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertApiAlertViewNumberOfAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertViewNumberOfAlertsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseurl** | **optional.String**|  | 
 **riskId** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

