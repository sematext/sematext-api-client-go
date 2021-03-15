# {{classname}}

All URIs are relative to */*

| Method                                                                     | HTTP request                                                          | Description                              |
| -------------------------------------------------------------------------- | --------------------------------------------------------------------- | ---------------------------------------- |
| [**CreateForAppUsingPOST1**](SubscriptionsAPI.md#CreateForAppUsingPOST1)   | **Post** /users-web/api/v3/apps/{appID}/subscription                  | Create App subscription                  |
| [**CreateForDashUsingPOST1**](SubscriptionsAPI.md#CreateForDashUsingPOST1) | **Post** /users-web/api/v3/dashboards/{dashID}/subscription           | Create dashboard subscription            |
| [**DeleteUsingDELETE3**](SubscriptionsAPI.md#DeleteUsingDELETE3)           | **Delete** /users-web/api/v3/subscriptions/{updateableSubscriptionID} | Delete subscription                      |
| [**ListUsingGET3**](SubscriptionsAPI.md#ListUsingGET3)                     | **Get** /users-web/api/v3/apps/{appID}/subscriptions                  | Get subscriptions for an App             |
| [**ListUsingGET5**](SubscriptionsAPI.md#ListUsingGET5)                     | **Get** /users-web/api/v3/subscriptions                               | Get current account&#x27;s subscriptions |
| [**SendAppReportUsingPOST**](SubscriptionsAPI.md#SendAppReportUsingPOST)   | **Post** /users-web/api/v3/apps/{appID}/report/send                   | Email an App report                      |
| [**SendDashReportUsingPOST**](SubscriptionsAPI.md#SendDashReportUsingPOST) | **Post** /users-web/api/v3/dashboards/{dashID}/report/send            | Email a dashboard report                 |
| [**ToggleEnabledUsingPUT1**](SubscriptionsAPI.md#ToggleEnabledUsingPUT1)   | **Put** /users-web/api/v3/subscriptions/{updateableSubscriptionID}    | Toggle subscription status               |
| [**UpdateForAppUsingPUT1**](SubscriptionsAPI.md#UpdateForAppUsingPUT1)     | **Put** /users-web/api/v3/apps/{appID}/subscription                   | Update App subscription                  |
| [**UpdateForDashUsingPUT1**](SubscriptionsAPI.md#UpdateForDashUsingPUT1)   | **Put** /users-web/api/v3/dashboards/{dashID}/subscription            | Update dashboard subscription            |

# **CreateForAppUsingPOST1**
> SubscriptionResponse CreateForAppUsingPOST1(ctx, body, appID)
Create App subscription

### Required Parameters

| Name      | Type                                      | Description                                                                 | Notes |
| --------- | ----------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                       | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**SubscriptionDto**](SubscriptionDto.md) | subscription                                                                |
| **appID** | **int64**                                 | appID                                                                       |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateForDashUsingPOST1**
> SubscriptionResponse CreateForDashUsingPOST1(ctx, body, dashID)
Create dashboard subscription

### Required Parameters

| Name       | Type                                                        | Description                                                                 | Notes |
| ---------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context**                                         | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**   | [**SubscriptionDashboardDto**](SubscriptionDashboardDto.md) | subscription                                                                |
| **dashID** | **int64**                                                   | dashID                                                                      |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE3**
> GenericMapBasedAPIResponse DeleteUsingDELETE3(ctx, updateableSubscriptionID)
Delete subscription

### Required Parameters

| Name                         | Type                | Description                                                                 | Notes |
| ---------------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**                      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableSubscriptionID** | **int64**           | updateableSubscriptionID                                                    |

### Return type

[**GenericMapBasedAPIResponse**](Generic Map Based API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET3**
> SubscriptionsResponse ListUsingGET3(ctx, appID)
Get subscriptions for an App

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID** | **int64**           | appID                                                                       |

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
> MailReportResponse SendAppReportUsingPOST(ctx, body, appID)
Email an App report

### Required Parameters

| Name      | Type                            | Description                                                                 | Notes |
| --------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**ReportInfo**](ReportInfo.md) | emailDto                                                                    |
| **appID** | **int64**                       | appID                                                                       |

### Return type

[**MailReportResponse**](MailReportResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendDashReportUsingPOST**
> MailReportResponse SendDashReportUsingPOST(ctx, body, dashID)
Email a dashboard report

### Required Parameters

| Name       | Type                            | Description                                                                 | Notes |
| ---------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**   | [**ReportInfo**](ReportInfo.md) | emailDto                                                                    |
| **dashID** | **int64**                       | dashID                                                                      |

### Return type

[**MailReportResponse**](MailReportResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleEnabledUsingPUT1**
> SubscriptionResponse ToggleEnabledUsingPUT1(ctx, body, updateableSubscriptionID)
Toggle subscription status

### Required Parameters

| Name                         | Type                                                  | Description                                                                 | Notes |
| ---------------------------- | ----------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**                      | **context.Context**                                   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**                     | [**UpdateSubscriptionDto**](UpdateSubscriptionDto.md) | dto                                                                         |
| **updateableSubscriptionID** | **int64**                                             | updateableSubscriptionID                                                    |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateForAppUsingPUT1**
> SubscriptionResponse UpdateForAppUsingPUT1(ctx, body, appID)
Update App subscription

### Required Parameters

| Name      | Type                                      | Description                                                                 | Notes |
| --------- | ----------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                       | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**SubscriptionDto**](SubscriptionDto.md) | subscription                                                                |
| **appID** | **int64**                                 | appID                                                                       |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateForDashUsingPUT1**
> SubscriptionResponse UpdateForDashUsingPUT1(ctx, body, dashID)
Update dashboard subscription

### Required Parameters

| Name       | Type                                                        | Description                                                                 | Notes |
| ---------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context**                                         | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**   | [**SubscriptionDashboardDto**](SubscriptionDashboardDto.md) | subscription                                                                |
| **dashID** | **int64**                                                   | dashID                                                                      |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
