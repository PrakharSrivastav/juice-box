# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchOtherHarByHeaderRegex**](SearchApi.md#SearchOtherHarByHeaderRegex) | **Get** /OTHER/search/other/harByHeaderRegex/ | 
[**SearchOtherHarByRequestRegex**](SearchApi.md#SearchOtherHarByRequestRegex) | **Get** /OTHER/search/other/harByRequestRegex/ | 
[**SearchOtherHarByResponseRegex**](SearchApi.md#SearchOtherHarByResponseRegex) | **Get** /OTHER/search/other/harByResponseRegex/ | 
[**SearchOtherHarByUrlRegex**](SearchApi.md#SearchOtherHarByUrlRegex) | **Get** /OTHER/search/other/harByUrlRegex/ | 
[**SearchViewMessagesByHeaderRegex**](SearchApi.md#SearchViewMessagesByHeaderRegex) | **Get** /JSON/search/view/messagesByHeaderRegex/ | 
[**SearchViewMessagesByRequestRegex**](SearchApi.md#SearchViewMessagesByRequestRegex) | **Get** /JSON/search/view/messagesByRequestRegex/ | 
[**SearchViewMessagesByResponseRegex**](SearchApi.md#SearchViewMessagesByResponseRegex) | **Get** /JSON/search/view/messagesByResponseRegex/ | 
[**SearchViewMessagesByUrlRegex**](SearchApi.md#SearchViewMessagesByUrlRegex) | **Get** /JSON/search/view/messagesByUrlRegex/ | 
[**SearchViewUrlsByHeaderRegex**](SearchApi.md#SearchViewUrlsByHeaderRegex) | **Get** /JSON/search/view/urlsByHeaderRegex/ | 
[**SearchViewUrlsByRequestRegex**](SearchApi.md#SearchViewUrlsByRequestRegex) | **Get** /JSON/search/view/urlsByRequestRegex/ | 
[**SearchViewUrlsByResponseRegex**](SearchApi.md#SearchViewUrlsByResponseRegex) | **Get** /JSON/search/view/urlsByResponseRegex/ | 
[**SearchViewUrlsByUrlRegex**](SearchApi.md#SearchViewUrlsByUrlRegex) | **Get** /JSON/search/view/urlsByUrlRegex/ | 

# **SearchOtherHarByHeaderRegex**
> ModelError SearchOtherHarByHeaderRegex(ctx, regex, optional)


Returns the HTTP messages, in HAR format, that match the given regular expression in the header(s) optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchOtherHarByHeaderRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOtherHarByHeaderRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOtherHarByRequestRegex**
> ModelError SearchOtherHarByRequestRegex(ctx, regex, optional)


Returns the HTTP messages, in HAR format, that match the given regular expression in the request optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchOtherHarByRequestRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOtherHarByRequestRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOtherHarByResponseRegex**
> ModelError SearchOtherHarByResponseRegex(ctx, regex, optional)


Returns the HTTP messages, in HAR format, that match the given regular expression in the response optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchOtherHarByResponseRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOtherHarByResponseRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOtherHarByUrlRegex**
> ModelError SearchOtherHarByUrlRegex(ctx, regex, optional)


Returns the HTTP messages, in HAR format, that match the given regular expression in the URL optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchOtherHarByUrlRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOtherHarByUrlRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewMessagesByHeaderRegex**
> ModelError SearchViewMessagesByHeaderRegex(ctx, regex, optional)


Returns the HTTP messages that match the given regular expression in the header(s) optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewMessagesByHeaderRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewMessagesByHeaderRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewMessagesByRequestRegex**
> ModelError SearchViewMessagesByRequestRegex(ctx, regex, optional)


Returns the HTTP messages that match the given regular expression in the request optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewMessagesByRequestRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewMessagesByRequestRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewMessagesByResponseRegex**
> ModelError SearchViewMessagesByResponseRegex(ctx, regex, optional)


Returns the HTTP messages that match the given regular expression in the response optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewMessagesByResponseRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewMessagesByResponseRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewMessagesByUrlRegex**
> ModelError SearchViewMessagesByUrlRegex(ctx, regex, optional)


Returns the HTTP messages that match the given regular expression in the URL optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewMessagesByUrlRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewMessagesByUrlRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewUrlsByHeaderRegex**
> ModelError SearchViewUrlsByHeaderRegex(ctx, regex, optional)


Returns the URLs of the HTTP messages that match the given regular expression in the header(s) optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewUrlsByHeaderRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewUrlsByHeaderRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewUrlsByRequestRegex**
> ModelError SearchViewUrlsByRequestRegex(ctx, regex, optional)


Returns the URLs of the HTTP messages that match the given regular expression in the request optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewUrlsByRequestRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewUrlsByRequestRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewUrlsByResponseRegex**
> ModelError SearchViewUrlsByResponseRegex(ctx, regex, optional)


Returns the URLs of the HTTP messages that match the given regular expression in the response optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewUrlsByResponseRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewUrlsByResponseRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchViewUrlsByUrlRegex**
> ModelError SearchViewUrlsByUrlRegex(ctx, regex, optional)


Returns the URLs of the HTTP messages that match the given regular expression in the URL optionally filtered by URL and paginated with 'start' position and 'count' of messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regex** | **string**|  | 
 **optional** | ***SearchApiSearchViewUrlsByUrlRegexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchViewUrlsByUrlRegexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseurl** | **optional.String**|  | 
 **start** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

