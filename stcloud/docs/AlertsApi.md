# {{classname}}

All URIs are relative to */*

| Method                                                                      | HTTP request                                                 | Description                |
| --------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------- |
| [**CreateAlertUsingPOST**](AlertsApi.md#CreateAlertUsingPOST)               | **Post** /users-web/api/v3/alerts                            | Create alert rule          |
| [**DeleteAlertRuleUsingDELETE**](AlertsApi.md#DeleteAlertRuleUsingDELETE)   | **Delete** /users-web/api/v3/alerts/{updateableAlertId}      | Delete alert rule          |
| [**DisableAlertRuleUsingPUT**](AlertsApi.md#DisableAlertRuleUsingPUT)       | **Put** /users-web/api/v3/alerts/{updateableAlertId}/disable | Disable alert rule         |
| [**EnableAlertRuleUsingPUT**](AlertsApi.md#EnableAlertRuleUsingPUT)         | **Put** /users-web/api/v3/alerts/{updateableAlertId}/enable  | Enable alert rule          |
| [**GetAlertRulesForAppUsingGET**](AlertsApi.md#GetAlertRulesForAppUsingGET) | **Get** /users-web/api/v3/apps/{appId}/alerts                | Get alert rules for an app |

# **CreateAlertUsingPOST**

> AlertRuleResponse CreateAlertUsingPOST(ctx, body)
Create alert rule

### Required Parameters

| Name     | Type                          | Description                                                                 | Notes |
| -------- | ----------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**           | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**AlertRule**](AlertRule.md) | dto                                                                         |

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertRuleUsingDELETE**

> GenericMapBasedApiResponse DeleteAlertRuleUsingDELETE(ctx, updateableAlertId)
Delete alert rule

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableAlertId** | **int64**           | updateableAlertId                                                           |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableAlertRuleUsingPUT**

> GenericMapBasedApiResponse DisableAlertRuleUsingPUT(ctx, updateableAlertId)
Disable alert rule

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableAlertId** | **int64**           | updateableAlertId                                                           |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAlertRuleUsingPUT**

> GenericMapBasedApiResponse EnableAlertRuleUsingPUT(ctx, updateableAlertId)
Enable alert rule

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableAlertId** | **int64**           | updateableAlertId                                                           |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertRulesForAppUsingGET**

> AlertRulesResponse GetAlertRulesForAppUsingGET(ctx, appId)
Get alert rules for an app

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**           | appId                                                                       |

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
