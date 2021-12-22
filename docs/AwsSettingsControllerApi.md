# {{classname}}

All URIs are relative to */*

| Method                                                             | HTTP request                               | Description                               |
| ------------------------------------------------------------------ | ------------------------------------------ | ----------------------------------------- |
| [**UpdateUsingPUT1**](AwsSettingsControllerApi.md#UpdateUsingPUT1) | **Put** /users-web/api/v3/apps/{appId}/aws | Update App&#x27;s AWS CloudWatch settings |

# **UpdateUsingPUT1**

> CloudWatchSettingsResponse UpdateUsingPUT1(ctx, body, appId)
Update App's AWS CloudWatch settings

Applicable only for AWS Apps

### Required Parameters

| Name      | Type                                            | Description                                                                 | Notes |
| --------- | ----------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**CloudWatchSettings**](CloudWatchSettings.md) | dto                                                                         |
| **appId** | **int64**                                       | appId                                                                       |

### Return type

[**CloudWatchSettingsResponse**](CloudWatchSettingsResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
