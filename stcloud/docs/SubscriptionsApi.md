# {{classname}}

All URIs are relative to */*

| Method                                                                     | HTTP request                                                          | Description                              |
| -------------------------------------------------------------------------- | --------------------------------------------------------------------- | ---------------------------------------- |
| [**CreateForAppUsingPOST**](SubscriptionsApi.md#CreateForAppUsingPOST)     | **Post** /users-web/api/v3/apps/{appId}/subscription                  | Create App subscription                  |
| [**CreateForDashUsingPOST1**](SubscriptionsApi.md#CreateForDashUsingPOST1) | **Post** /users-web/api/v3/dashboards/{dashId}/subscription           | Create dashboard subscription            |
| [**DeleteUsingDELETE2**](SubscriptionsApi.md#DeleteUsingDELETE2)           | **Delete** /users-web/api/v3/subscriptions/{updateableSubscriptionId} | Delete subscription                      |
| [**ListUsingGET3**](SubscriptionsApi.md#ListUsingGET3)                     | **Get** /users-web/api/v3/apps/{appId}/subscriptions                  | Get subscriptions for an App             |
| [**ListUsingGET5**](SubscriptionsApi.md#ListUsingGET5)                     | **Get** /users-web/api/v3/subscriptions                               | Get current account&#x27;s subscriptions |
| [**SendAppReportUsingPOST**](SubscriptionsApi.md#SendAppReportUsingPOST)   | **Post** /users-web/api/v3/apps/{appId}/report/send                   | Email an App report                      |
| [**SendDashReportUsingPOST**](SubscriptionsApi.md#SendDashReportUsingPOST) | **Post** /users-web/api/v3/dashboards/{dashId}/report/send            | Email a dashboard report                 |
| [**ToggleEnabledUsingPUT1**](SubscriptionsApi.md#ToggleEnabledUsingPUT1)   | **Put** /users-web/api/v3/subscriptions/{updateableSubscriptionId}    | Toggle subscription status               |
| [**UpdateForAppUsingPUT**](SubscriptionsApi.md#UpdateForAppUsingPUT)       | **Put** /users-web/api/v3/apps/{appId}/subscription                   | Update App subscription                  |
| [**UpdateForDashUsingPUT1**](SubscriptionsApi.md#UpdateForDashUsingPUT1)   | **Put** /users-web/api/v3/dashboards/{dashId}/subscription            | Update dashboard subscription            |

# **CreateForAppUsingPOST**

> SubscriptionResponse CreateForAppUsingPOST(ctx, body, appId)
Create App subscription

### Required Parameters

| Name      | Type                                      | Description                                                                 | Notes |
| --------- | ----------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                       | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**SubscriptionDto**](SubscriptionDto.md) | subscription                                                                |
| **appId** | **int64**                                 | appId                                                                       |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateForDashUsingPOST1**

> SubscriptionResponse CreateForDashUsingPOST1(ctx, body, dashId)
Create dashboard subscription

### Required Parameters

| Name       | Type                                                        | Description                                                                 | Notes |
| ---------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context**                                         | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**   | [**SubscriptionDashboardDto**](SubscriptionDashboardDto.md) | subscription                                                                |
| **dashId** | **int64**                                                   | dashId                                                                      |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE2**

> GenericMapBasedApiResponse DeleteUsingDELETE2(ctx, updateableSubscriptionId)
Delete subscription

### Required Parameters

| Name                         | Type                | Description                                                                 | Notes |
| ---------------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**                      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableSubscriptionId** | **int64**           | updateableSubscriptionId                                                    |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET3**

> SubscriptionsResponse ListUsingGET3(ctx, appId)
Get subscriptions for an App

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**           | appId                                                                       |

### Return type

[**SubscriptionsResponse**](SubscriptionsResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET5**

> SubscriptionsResponse ListUsingGET5(ctx, )
Get current account's subscriptions

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SubscriptionsResponse**](SubscriptionsResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendAppReportUsingPOST**

> MailReportResponse SendAppReportUsingPOST(ctx, body, appId)
Email an App report

### Required Parameters

| Name      | Type                            | Description                                                                 | Notes |
| --------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**ReportInfo**](ReportInfo.md) | emailDto                                                                    |
| **appId** | **int64**                       | appId                                                                       |

### Return type

[**MailReportResponse**](MailReportResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendDashReportUsingPOST**

> MailReportResponse SendDashReportUsingPOST(ctx, body, dashId)
Email a dashboard report

### Required Parameters

| Name       | Type                            | Description                                                                 | Notes |
| ---------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**   | [**ReportInfo**](ReportInfo.md) | emailDto                                                                    |
| **dashId** | **int64**                       | dashId                                                                      |

### Return type

[**MailReportResponse**](MailReportResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleEnabledUsingPUT1**

> SubscriptionResponse ToggleEnabledUsingPUT1(ctx, body, updateableSubscriptionId)
Toggle subscription status

### Required Parameters

| Name                         | Type                                                  | Description                                                                 | Notes |
| ---------------------------- | ----------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**                      | **context.Context**                                   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**                     | [**UpdateSubscriptionDto**](UpdateSubscriptionDto.md) | dto                                                                         |
| **updateableSubscriptionId** | **int64**                                             | updateableSubscriptionId                                                    |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateForAppUsingPUT**

> SubscriptionResponse UpdateForAppUsingPUT(ctx, body, appId)
Update App subscription

### Required Parameters

| Name      | Type                                      | Description                                                                 | Notes |
| --------- | ----------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                       | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**SubscriptionDto**](SubscriptionDto.md) | subscription                                                                |
| **appId** | **int64**                                 | appId                                                                       |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateForDashUsingPUT1**

> SubscriptionResponse UpdateForDashUsingPUT1(ctx, body, dashId)
Update dashboard subscription

### Required Parameters

| Name       | Type                                                        | Description                                                                 | Notes |
| ---------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context**                                         | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**   | [**SubscriptionDashboardDto**](SubscriptionDashboardDto.md) | subscription                                                                |
| **dashId** | **int64**                                                   | dashId                                                                      |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
