# {{classname}}

All URIs are relative to */*

| Method                                                                      | HTTP request                                                   | Description |
| --------------------------------------------------------------------------- | -------------------------------------------------------------- | ----------- |
| [**GetForRangeUsingGET**](LogsUsageAPIControllerAPI.md#GetForRangeUsingGET) | **Get** /logsene-reports/api/v3/apps/{appID}/usage/{from}/{to} | getForRange |

# **GetForRangeUsingGET**
> UsageResponse GetForRangeUsingGET(ctx, appID, from, to)
getForRange

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID** | **int64**           | appID                                                                       |
| **from**  | **int64**           | from                                                                        |
| **to**    | **int64**           | to                                                                          |

### Return type

[**UsageResponse**](UsageResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
