# {{classname}}

All URIs are relative to *http://zap*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoupdateActionDownloadLatestRelease**](AutoupdateApi.md#AutoupdateActionDownloadLatestRelease) | **Get** /JSON/autoupdate/action/downloadLatestRelease/ | 
[**AutoupdateActionInstallAddon**](AutoupdateApi.md#AutoupdateActionInstallAddon) | **Get** /JSON/autoupdate/action/installAddon/ | 
[**AutoupdateActionInstallLocalAddon**](AutoupdateApi.md#AutoupdateActionInstallLocalAddon) | **Get** /JSON/autoupdate/action/installLocalAddon/ | 
[**AutoupdateActionSetOptionCheckAddonUpdates**](AutoupdateApi.md#AutoupdateActionSetOptionCheckAddonUpdates) | **Get** /JSON/autoupdate/action/setOptionCheckAddonUpdates/ | 
[**AutoupdateActionSetOptionCheckOnStart**](AutoupdateApi.md#AutoupdateActionSetOptionCheckOnStart) | **Get** /JSON/autoupdate/action/setOptionCheckOnStart/ | 
[**AutoupdateActionSetOptionDownloadNewRelease**](AutoupdateApi.md#AutoupdateActionSetOptionDownloadNewRelease) | **Get** /JSON/autoupdate/action/setOptionDownloadNewRelease/ | 
[**AutoupdateActionSetOptionInstallAddonUpdates**](AutoupdateApi.md#AutoupdateActionSetOptionInstallAddonUpdates) | **Get** /JSON/autoupdate/action/setOptionInstallAddonUpdates/ | 
[**AutoupdateActionSetOptionInstallScannerRules**](AutoupdateApi.md#AutoupdateActionSetOptionInstallScannerRules) | **Get** /JSON/autoupdate/action/setOptionInstallScannerRules/ | 
[**AutoupdateActionSetOptionReportAlphaAddons**](AutoupdateApi.md#AutoupdateActionSetOptionReportAlphaAddons) | **Get** /JSON/autoupdate/action/setOptionReportAlphaAddons/ | 
[**AutoupdateActionSetOptionReportBetaAddons**](AutoupdateApi.md#AutoupdateActionSetOptionReportBetaAddons) | **Get** /JSON/autoupdate/action/setOptionReportBetaAddons/ | 
[**AutoupdateActionSetOptionReportReleaseAddons**](AutoupdateApi.md#AutoupdateActionSetOptionReportReleaseAddons) | **Get** /JSON/autoupdate/action/setOptionReportReleaseAddons/ | 
[**AutoupdateActionUninstallAddon**](AutoupdateApi.md#AutoupdateActionUninstallAddon) | **Get** /JSON/autoupdate/action/uninstallAddon/ | 
[**AutoupdateViewInstalledAddons**](AutoupdateApi.md#AutoupdateViewInstalledAddons) | **Get** /JSON/autoupdate/view/installedAddons/ | 
[**AutoupdateViewIsLatestVersion**](AutoupdateApi.md#AutoupdateViewIsLatestVersion) | **Get** /JSON/autoupdate/view/isLatestVersion/ | 
[**AutoupdateViewLatestVersionNumber**](AutoupdateApi.md#AutoupdateViewLatestVersionNumber) | **Get** /JSON/autoupdate/view/latestVersionNumber/ | 
[**AutoupdateViewLocalAddons**](AutoupdateApi.md#AutoupdateViewLocalAddons) | **Get** /JSON/autoupdate/view/localAddons/ | 
[**AutoupdateViewMarketplaceAddons**](AutoupdateApi.md#AutoupdateViewMarketplaceAddons) | **Get** /JSON/autoupdate/view/marketplaceAddons/ | 
[**AutoupdateViewNewAddons**](AutoupdateApi.md#AutoupdateViewNewAddons) | **Get** /JSON/autoupdate/view/newAddons/ | 
[**AutoupdateViewOptionAddonDirectories**](AutoupdateApi.md#AutoupdateViewOptionAddonDirectories) | **Get** /JSON/autoupdate/view/optionAddonDirectories/ | 
[**AutoupdateViewOptionCheckAddonUpdates**](AutoupdateApi.md#AutoupdateViewOptionCheckAddonUpdates) | **Get** /JSON/autoupdate/view/optionCheckAddonUpdates/ | 
[**AutoupdateViewOptionCheckOnStart**](AutoupdateApi.md#AutoupdateViewOptionCheckOnStart) | **Get** /JSON/autoupdate/view/optionCheckOnStart/ | 
[**AutoupdateViewOptionDayLastChecked**](AutoupdateApi.md#AutoupdateViewOptionDayLastChecked) | **Get** /JSON/autoupdate/view/optionDayLastChecked/ | 
[**AutoupdateViewOptionDayLastInstallWarned**](AutoupdateApi.md#AutoupdateViewOptionDayLastInstallWarned) | **Get** /JSON/autoupdate/view/optionDayLastInstallWarned/ | 
[**AutoupdateViewOptionDayLastUpdateWarned**](AutoupdateApi.md#AutoupdateViewOptionDayLastUpdateWarned) | **Get** /JSON/autoupdate/view/optionDayLastUpdateWarned/ | 
[**AutoupdateViewOptionDownloadDirectory**](AutoupdateApi.md#AutoupdateViewOptionDownloadDirectory) | **Get** /JSON/autoupdate/view/optionDownloadDirectory/ | 
[**AutoupdateViewOptionDownloadNewRelease**](AutoupdateApi.md#AutoupdateViewOptionDownloadNewRelease) | **Get** /JSON/autoupdate/view/optionDownloadNewRelease/ | 
[**AutoupdateViewOptionInstallAddonUpdates**](AutoupdateApi.md#AutoupdateViewOptionInstallAddonUpdates) | **Get** /JSON/autoupdate/view/optionInstallAddonUpdates/ | 
[**AutoupdateViewOptionInstallScannerRules**](AutoupdateApi.md#AutoupdateViewOptionInstallScannerRules) | **Get** /JSON/autoupdate/view/optionInstallScannerRules/ | 
[**AutoupdateViewOptionReportAlphaAddons**](AutoupdateApi.md#AutoupdateViewOptionReportAlphaAddons) | **Get** /JSON/autoupdate/view/optionReportAlphaAddons/ | 
[**AutoupdateViewOptionReportBetaAddons**](AutoupdateApi.md#AutoupdateViewOptionReportBetaAddons) | **Get** /JSON/autoupdate/view/optionReportBetaAddons/ | 
[**AutoupdateViewOptionReportReleaseAddons**](AutoupdateApi.md#AutoupdateViewOptionReportReleaseAddons) | **Get** /JSON/autoupdate/view/optionReportReleaseAddons/ | 
[**AutoupdateViewUpdatedAddons**](AutoupdateApi.md#AutoupdateViewUpdatedAddons) | **Get** /JSON/autoupdate/view/updatedAddons/ | 

# **AutoupdateActionDownloadLatestRelease**
> ModelError AutoupdateActionDownloadLatestRelease(ctx, )


Downloads the latest release, if any

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

# **AutoupdateActionInstallAddon**
> ModelError AutoupdateActionInstallAddon(ctx, id)


Installs or updates the specified add-on, returning when complete (ie not asynchronously)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoupdateActionInstallLocalAddon**
> ModelError AutoupdateActionInstallLocalAddon(ctx, file)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoupdateActionSetOptionCheckAddonUpdates**
> ModelError AutoupdateActionSetOptionCheckAddonUpdates(ctx, boolean)


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

# **AutoupdateActionSetOptionCheckOnStart**
> ModelError AutoupdateActionSetOptionCheckOnStart(ctx, boolean)


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

# **AutoupdateActionSetOptionDownloadNewRelease**
> ModelError AutoupdateActionSetOptionDownloadNewRelease(ctx, boolean)


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

# **AutoupdateActionSetOptionInstallAddonUpdates**
> ModelError AutoupdateActionSetOptionInstallAddonUpdates(ctx, boolean)


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

# **AutoupdateActionSetOptionInstallScannerRules**
> ModelError AutoupdateActionSetOptionInstallScannerRules(ctx, boolean)


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

# **AutoupdateActionSetOptionReportAlphaAddons**
> ModelError AutoupdateActionSetOptionReportAlphaAddons(ctx, boolean)


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

# **AutoupdateActionSetOptionReportBetaAddons**
> ModelError AutoupdateActionSetOptionReportBetaAddons(ctx, boolean)


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

# **AutoupdateActionSetOptionReportReleaseAddons**
> ModelError AutoupdateActionSetOptionReportReleaseAddons(ctx, boolean)


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

# **AutoupdateActionUninstallAddon**
> ModelError AutoupdateActionUninstallAddon(ctx, id)


Uninstalls the specified add-on

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ModelError**](Error.md)

### Authorization

[apiKeyHeader](../README.md#apiKeyHeader), [apiKeyQuery](../README.md#apiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoupdateViewInstalledAddons**
> ModelError AutoupdateViewInstalledAddons(ctx, )


Return a list of all of the installed add-ons

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

# **AutoupdateViewIsLatestVersion**
> ModelError AutoupdateViewIsLatestVersion(ctx, )


Returns 'true' if ZAP is on the latest version

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

# **AutoupdateViewLatestVersionNumber**
> ModelError AutoupdateViewLatestVersionNumber(ctx, )


Returns the latest version number

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

# **AutoupdateViewLocalAddons**
> ModelError AutoupdateViewLocalAddons(ctx, )


Returns a list with all local add-ons, installed or not.

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

# **AutoupdateViewMarketplaceAddons**
> ModelError AutoupdateViewMarketplaceAddons(ctx, )


Return a list of all of the add-ons on the ZAP Marketplace (this information is read once and then cached)

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

# **AutoupdateViewNewAddons**
> ModelError AutoupdateViewNewAddons(ctx, )


Return a list of any add-ons that have been added to the Marketplace since the last check for updates

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

# **AutoupdateViewOptionAddonDirectories**
> ModelError AutoupdateViewOptionAddonDirectories(ctx, )


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

# **AutoupdateViewOptionCheckAddonUpdates**
> ModelError AutoupdateViewOptionCheckAddonUpdates(ctx, )


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

# **AutoupdateViewOptionCheckOnStart**
> ModelError AutoupdateViewOptionCheckOnStart(ctx, )


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

# **AutoupdateViewOptionDayLastChecked**
> ModelError AutoupdateViewOptionDayLastChecked(ctx, )


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

# **AutoupdateViewOptionDayLastInstallWarned**
> ModelError AutoupdateViewOptionDayLastInstallWarned(ctx, )


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

# **AutoupdateViewOptionDayLastUpdateWarned**
> ModelError AutoupdateViewOptionDayLastUpdateWarned(ctx, )


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

# **AutoupdateViewOptionDownloadDirectory**
> ModelError AutoupdateViewOptionDownloadDirectory(ctx, )


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

# **AutoupdateViewOptionDownloadNewRelease**
> ModelError AutoupdateViewOptionDownloadNewRelease(ctx, )


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

# **AutoupdateViewOptionInstallAddonUpdates**
> ModelError AutoupdateViewOptionInstallAddonUpdates(ctx, )


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

# **AutoupdateViewOptionInstallScannerRules**
> ModelError AutoupdateViewOptionInstallScannerRules(ctx, )


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

# **AutoupdateViewOptionReportAlphaAddons**
> ModelError AutoupdateViewOptionReportAlphaAddons(ctx, )


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

# **AutoupdateViewOptionReportBetaAddons**
> ModelError AutoupdateViewOptionReportBetaAddons(ctx, )


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

# **AutoupdateViewOptionReportReleaseAddons**
> ModelError AutoupdateViewOptionReportReleaseAddons(ctx, )


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

# **AutoupdateViewUpdatedAddons**
> ModelError AutoupdateViewUpdatedAddons(ctx, )


Return a list of any add-ons that have been changed in the Marketplace since the last check for updates

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

