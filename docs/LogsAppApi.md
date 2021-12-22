# {{classname}}

All URIs are relative to */*

| Method                                                                 | HTTP request                          | Description     |
| ---------------------------------------------------------------------- | ------------------------------------- | --------------- |
| [**CreateLogseneApplication**](LogsAppApi.md#CreateLogseneApplication) | **Post** /logsene-reports/api/v3/apps | Create Logs App |

# **CreateLogseneApplication**

> AppsResponse CreateLogseneApplication(ctx, body)
Create Logs App

### Required Parameters

| Name     | Type                                  | Description                                                                 | Notes |
| -------- | ------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**CreateAppInfo**](CreateAppInfo.md) | Details of the application to be created                                    |

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
