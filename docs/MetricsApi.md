# \MetricsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDataSeriesUsingPOST**](MetricsApi.md#ListDataSeriesUsingPOST) | **Post** /spm-reports/api/v3/apps/{appId}/metrics/data | Get metrics data points for an app
[**ListFiltersUsingPOST1**](MetricsApi.md#ListFiltersUsingPOST1) | **Post** /spm-reports/api/v3/apps/{appId}/metrics/filters | Get metrics filters and their values for an app
[**ListMetricsKeysUsingGET**](MetricsApi.md#ListMetricsKeysUsingGET) | **Get** /spm-reports/api/v3/apps/{appId}/metrics/keys | Get metrics keys for an app
[**ListMetricsUsingGET**](MetricsApi.md#ListMetricsUsingGET) | **Get** /spm-reports/api/v3/apps/{appId}/metrics | Get metrics info for an app


# **ListDataSeriesUsingPOST**
> GenericApiResponse ListDataSeriesUsingPOST(ctx, appId, requestBody)
Get metrics data points for an app

Default value of interval is 5m

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| appId | 
  **requestBody** | [**DataSeriesRequest**](DataSeriesRequest.md)| Metric data points request | 

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFiltersUsingPOST1**
> GenericApiResponse ListFiltersUsingPOST1(ctx, appId, requestBody)
Get metrics filters and their values for an app

Default value of interval is 5m

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| appId | 
  **requestBody** | [**DataSeriesRequest**](DataSeriesRequest.md)| Metric filters request | 

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMetricsKeysUsingGET**
> GenericApiResponse ListMetricsKeysUsingGET(ctx, appId)
Get metrics keys for an app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| appId | 

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMetricsUsingGET**
> GenericApiResponse ListMetricsUsingGET(ctx, appId)
Get metrics info for an app

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| appId | 

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

