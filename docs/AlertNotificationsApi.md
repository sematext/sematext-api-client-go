# \AlertNotificationsApi

All URIs are relative to *https://localhost*

| Method                                                                                                      | HTTP request                                                 | Description                        |
| ----------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | ---------------------------------- |
| [**GetAlertNotificationsForAppUsingPOST**](AlertNotificationsApi.md#GetAlertNotificationsForAppUsingPOST)   | **Post** /users-web/api/v3/apps/{appId}/notifications/alerts | Get alert notifications for an app |
| [**GetAlertNotificationsForUserUsingPOST**](AlertNotificationsApi.md#GetAlertNotificationsForUserUsingPOST) | **Post** /users-web/api/v3/notifications/alerts              | Get alert notifications for a user |


# **GetAlertNotificationsForAppUsingPOST**
> GenericApiResponse GetAlertNotificationsForAppUsingPOST(ctx, appId, timeInterval)
Get alert notifications for an app

Default value of interval is 1d

### Required Parameters

| Name             | Type                                                        | Description                                                                 | Notes |
| ---------------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**          | **context.Context**                                         | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId**        | **int64**                                                   | appId                                                                       |
| **timeInterval** | [**AlertNotificationRequest**](AlertNotificationRequest.md) | Time Interval                                                               |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertNotificationsForUserUsingPOST**
> GenericApiResponse GetAlertNotificationsForUserUsingPOST(ctx, timeInterval)
Get alert notifications for a user

Default value of interval is 1d

### Required Parameters

| Name             | Type                                                        | Description                                                                 | Notes |
| ---------------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**          | **context.Context**                                         | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **timeInterval** | [**AlertNotificationRequest**](AlertNotificationRequest.md) | Time Interval                                                               |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
