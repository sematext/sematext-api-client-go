# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForAppUsingPOST**](SubscriptionsApi.md#CreateForAppUsingPOST) | **Post** /users-web/api/v3/apps/{appId}/subscription | Create App subscription
[**CreateForDashUsingPOST1**](SubscriptionsApi.md#CreateForDashUsingPOST1) | **Post** /users-web/api/v3/dashboards/{dashId}/subscription | Create dashboard subscription
[**DeleteUsingDELETE3**](SubscriptionsApi.md#DeleteUsingDELETE3) | **Delete** /users-web/api/v3/subscriptions/{updateableSubscriptionId} | Delete subscription
[**ListUsingGET2**](SubscriptionsApi.md#ListUsingGET2) | **Get** /users-web/api/v3/apps/{appId}/subscriptions | Get subscriptions for an App
[**ListUsingGET5**](SubscriptionsApi.md#ListUsingGET5) | **Get** /users-web/api/v3/subscriptions | Get current account&#x27;s subscriptions
[**SendAppReportUsingPOST1**](SubscriptionsApi.md#SendAppReportUsingPOST1) | **Post** /users-web/api/v3/apps/{appId}/report/send | Email an App report
[**SendDashReportUsingPOST1**](SubscriptionsApi.md#SendDashReportUsingPOST1) | **Post** /users-web/api/v3/dashboards/{dashId}/report/send | Email a dashboard report
[**ToggleEnabledUsingPUT**](SubscriptionsApi.md#ToggleEnabledUsingPUT) | **Put** /users-web/api/v3/subscriptions/{updateableSubscriptionId} | Toggle subscription status
[**UpdateForAppUsingPUT1**](SubscriptionsApi.md#UpdateForAppUsingPUT1) | **Put** /users-web/api/v3/apps/{appId}/subscription | Update App subscription
[**UpdateForDashUsingPUT**](SubscriptionsApi.md#UpdateForDashUsingPUT) | **Put** /users-web/api/v3/dashboards/{dashId}/subscription | Update dashboard subscription

# **CreateForAppUsingPOST**

> SubscriptionResponse CreateForAppUsingPOST(ctx, body, appId)
Create App subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionDto**](SubscriptionDto.md)| subscription |
  **appId** | **int64**| appId |

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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionDashboardDto**](SubscriptionDashboardDto.md)| subscription |
  **dashId** | **int64**| dashId |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE3**

> GenericMapBasedApiResponse DeleteUsingDELETE3(ctx, updateableSubscriptionId)
Delete subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateableSubscriptionId** | **int64**| updateableSubscriptionId |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET2**

> SubscriptionsResponse ListUsingGET2(ctx, appId)
Get subscriptions for an App

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| appId |

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

# **SendAppReportUsingPOST1**

> MailReportResponse SendAppReportUsingPOST1(ctx, body, appId)
Email an App report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReportInfo**](ReportInfo.md)| emailDto |
  **appId** | **int64**| appId |

### Return type

[**MailReportResponse**](MailReportResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendDashReportUsingPOST1**

> MailReportResponse SendDashReportUsingPOST1(ctx, body, dashId)
Email a dashboard report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReportInfo**](ReportInfo.md)| emailDto |
  **dashId** | **int64**| dashId |

### Return type

[**MailReportResponse**](MailReportResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleEnabledUsingPUT**

> SubscriptionResponse ToggleEnabledUsingPUT(ctx, body, updateableSubscriptionId)
Toggle subscription status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSubscriptionDto**](UpdateSubscriptionDto.md)| dto |
  **updateableSubscriptionId** | **int64**| updateableSubscriptionId |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateForAppUsingPUT1**

> SubscriptionResponse UpdateForAppUsingPUT1(ctx, body, appId)
Update App subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionDto**](SubscriptionDto.md)| subscription |
  **appId** | **int64**| appId |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateForDashUsingPUT**

> SubscriptionResponse UpdateForDashUsingPUT(ctx, body, dashId)
Update dashboard subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriptionDashboardDto**](SubscriptionDashboardDto.md)| subscription |
  **dashId** | **int64**| dashId |

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
