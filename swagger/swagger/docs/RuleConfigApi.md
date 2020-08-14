# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RuleConfigActionResetAllRuleConfigValues**](RuleConfigApi.md#RuleConfigActionResetAllRuleConfigValues) | **Get** /JSON/ruleConfig/action/resetAllRuleConfigValues/ | 
[**RuleConfigActionResetRuleConfigValue**](RuleConfigApi.md#RuleConfigActionResetRuleConfigValue) | **Get** /JSON/ruleConfig/action/resetRuleConfigValue/ | 
[**RuleConfigActionSetRuleConfigValue**](RuleConfigApi.md#RuleConfigActionSetRuleConfigValue) | **Get** /JSON/ruleConfig/action/setRuleConfigValue/ | 
[**RuleConfigViewAllRuleConfigs**](RuleConfigApi.md#RuleConfigViewAllRuleConfigs) | **Get** /JSON/ruleConfig/view/allRuleConfigs/ | 
[**RuleConfigViewRuleConfigValue**](RuleConfigApi.md#RuleConfigViewRuleConfigValue) | **Get** /JSON/ruleConfig/view/ruleConfigValue/ | 

# **RuleConfigActionResetAllRuleConfigValues**
> ModelError RuleConfigActionResetAllRuleConfigValues(ctx, )


Reset all of the rule configurations

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

# **RuleConfigActionResetRuleConfigValue**
> ModelError RuleConfigActionResetRuleConfigValue(ctx, key)


Reset the specified rule configuration, which must already exist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleConfigActionSetRuleConfigValue**
> ModelError RuleConfigActionSetRuleConfigValue(ctx, key, optional)


Set the specified rule configuration, which must already exist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
 **optional** | ***RuleConfigApiRuleConfigActionSetRuleConfigValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RuleConfigApiRuleConfigActionSetRuleConfigValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **value** | **optional.String**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleConfigViewAllRuleConfigs**
> ModelError RuleConfigViewAllRuleConfigs(ctx, )


Show all of the rule configurations

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

# **RuleConfigViewRuleConfigValue**
> ModelError RuleConfigViewRuleConfigValue(ctx, key)


Show the specified rule configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

