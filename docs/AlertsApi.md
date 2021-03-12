# {{classname}}

All URIs are relative to */*

| Method                                                                      | HTTP request                                                 | Description                |
| --------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------- |
| [**CreateAlertUsingPOST**](AlertsAPI.md#CreateAlertUsingPOST)               | **Post** /users-web/api/v3/alerts                            | Create alert rule          |
| [**DeleteAlertRuleUsingDELETE**](AlertsAPI.md#DeleteAlertRuleUsingDELETE)   | **Delete** /users-web/api/v3/alerts/{updateableAlertID}      | Delete alert rule          |
| [**DisableAlertRuleUsingPUT**](AlertsAPI.md#DisableAlertRuleUsingPUT)       | **Put** /users-web/api/v3/alerts/{updateableAlertID}/disable | Disable alert rule         |
| [**EnableAlertRuleUsingPUT**](AlertsAPI.md#EnableAlertRuleUsingPUT)         | **Put** /users-web/api/v3/alerts/{updateableAlertID}/enable  | Enable alert rule          |
| [**GetAlertRulesForAppUsingGET**](AlertsAPI.md#GetAlertRulesForAppUsingGET) | **Get** /users-web/api/v3/apps/{appID}/alerts                | Get alert rules for an app |

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
> GenericMapBasedAPIResponse DeleteAlertRuleUsingDELETE(ctx, updateableAlertID)
Delete alert rule

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableAlertID** | **int64**           | updateableAlertID                                                           |

### Return type

[**GenericMapBasedAPIResponse**](Generic Map Based API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableAlertRuleUsingPUT**
> GenericMapBasedAPIResponse DisableAlertRuleUsingPUT(ctx, updateableAlertID)
Disable alert rule

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableAlertID** | **int64**           | updateableAlertID                                                           |

### Return type

[**GenericMapBasedAPIResponse**](Generic Map Based API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAlertRuleUsingPUT**
> GenericMapBasedAPIResponse EnableAlertRuleUsingPUT(ctx, updateableAlertID)
Enable alert rule

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableAlertID** | **int64**           | updateableAlertID                                                           |

### Return type

[**GenericMapBasedAPIResponse**](Generic Map Based API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertRulesForAppUsingGET**
> AlertRulesResponse GetAlertRulesForAppUsingGET(ctx, appID)
Get alert rules for an app

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID** | **int64**           | appID                                                                       |

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
