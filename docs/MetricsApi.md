# \MetricsAPI

All URIs are relative to *https://localhost*

| Method                                                               | HTTP request                                              | Description                                     |
| -------------------------------------------------------------------- | --------------------------------------------------------- | ----------------------------------------------- |
| [**ListDataSeriesUsingPOST**](MetricsAPI.md#ListDataSeriesUsingPOST) | **Post** /spm-reports/api/v3/apps/{appID}/metrics/data    | Get metrics data points for an app              |
| [**ListFiltersUsingPOST1**](MetricsAPI.md#ListFiltersUsingPOST1)     | **Post** /spm-reports/api/v3/apps/{appID}/metrics/filters | Get metrics filters and their values for an app |
| [**ListMetricsKeysUsingGET**](MetricsAPI.md#ListMetricsKeysUsingGET) | **Get** /spm-reports/api/v3/apps/{appID}/metrics/keys     | Get metrics keys for an app                     |
| [**ListMetricsUsingGET**](MetricsAPI.md#ListMetricsUsingGET)         | **Get** /spm-reports/api/v3/apps/{appID}/metrics          | Get metrics info for an app                     |


# **ListDataSeriesUsingPOST**
> GenericAPIResponse ListDataSeriesUsingPOST(ctx, appID, requestBody)
Get metrics data points for an app

Default value of interval is 5m

### Required Parameters

| Name            | Type                                          | Description                                                                 | Notes |
| --------------- | --------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context**                           | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID**       | **int64**                                     | appID                                                                       |
| **requestBody** | [**DataSeriesRequest**](DataSeriesRequest.md) | Metric data points request                                                  |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFiltersUsingPOST1**
> GenericAPIResponse ListFiltersUsingPOST1(ctx, appID, requestBody)
Get metrics filters and their values for an app

Default value of interval is 5m

### Required Parameters

| Name            | Type                                          | Description                                                                 | Notes |
| --------------- | --------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context**                           | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID**       | **int64**                                     | appID                                                                       |
| **requestBody** | [**DataSeriesRequest**](DataSeriesRequest.md) | Metric filters request                                                      |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMetricsKeysUsingGET**
> GenericAPIResponse ListMetricsKeysUsingGET(ctx, appID)
Get metrics keys for an app

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID** | **int64**           | appID                                                                       |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMetricsUsingGET**
> GenericAPIResponse ListMetricsUsingGET(ctx, appID)
Get metrics info for an app

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID** | **int64**           | appID                                                                       |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
