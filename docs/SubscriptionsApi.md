# \SubscriptionsApi

All URIs are relative to *https://localhost*

| Method                                                             | HTTP request                                         | Description                           |
| ------------------------------------------------------------------ | ---------------------------------------------------- | ------------------------------------- |
| [**ListUsingGET1**](SubscriptionsApi.md#ListUsingGET1)             | **Get** /users-web/api/v3/apps/{appId}/subscriptions | Get subscriptions for an app          |
| [**SendReportUsingPOST**](SubscriptionsApi.md#SendReportUsingPOST) | **Post** /users-web/api/v3/apps/{appId}/report/send  | Trigger emailing of report for an app |


# **ListUsingGET1**
> GenericApiResponse ListUsingGET1(ctx, appId)
Get subscriptions for an app

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**           | appId                                                                       |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendReportUsingPOST**
> GenericApiResponse SendReportUsingPOST(ctx, appId, emailDto)
Trigger emailing of report for an app

### Required Parameters

| Name         | Type                            | Description                                                                 | Notes |
| ------------ | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId**    | **int64**                       | appId                                                                       |
| **emailDto** | [**ReportInfo**](ReportInfo.md) | emailDto                                                                    |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
