# {{classname}}

All URIs are relative to */*

| Method                                                                 | HTTP request                                                        | Description                               |
| ---------------------------------------------------------------------- | ------------------------------------------------------------------- | ----------------------------------------- |
| [**CreateAppToken1**](TokensApiControllerApi.md#CreateAppToken1)       | **Post** /users-web/api/v3/apps/{appID}/tokens                      | Create new app token                      |
| [**DeleteAppToken**](TokensApiControllerApi.md#DeleteAppToken)         | **Delete** /users-web/api/v3/apps/{appID}/tokens/{tokenID}          | Delete app token                          |
| [**GetAppTokens1**](TokensApiControllerApi.md#GetAppTokens1)           | **Get** /users-web/api/v3/apps/{appID}/tokens                       | Get app available tokens                  |
| [**RegenerateAppToken**](TokensApiControllerApi.md#RegenerateAppToken) | **Post** /users-web/api/v3/apps/{appID}/tokens/{tokenID}/regenerate | Regenerate app token)                     |
| [**UpdateAppToken**](TokensApiControllerApi.md#UpdateAppToken)         | **Put** /users-web/api/v3/apps/{appID}/tokens/{tokenID}             | Update app token (enable/disable or name) |

# **CreateAppToken1**
> TokenResponse CreateAppToken1(ctx, body, appID)
Create new app token

### Required Parameters

| Name      | Type                                    | Description                                                                 | Notes |
| --------- | --------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                     | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**CreateTokenDto**](CreateTokenDto.md) | dto                                                                         |
| **appID** | **int64**                               | appID                                                                       |

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAppToken**
> GenericMapBasedApiResponse DeleteAppToken(ctx, appID, tokenID)
Delete app token

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID**   | **int64**           | appID                                                                       |
| **tokenID** | **int64**           | tokenID                                                                     |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppTokens1**
> TokensResponse GetAppTokens1(ctx, appID)
Get app available tokens

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID** | **int64**           | appID                                                                       |

### Return type

[**TokensResponse**](TokensResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateAppToken**
> TokenResponse RegenerateAppToken(ctx, appID, tokenID)
Regenerate app token)

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID**   | **int64**           | appID                                                                       |
| **tokenID** | **int64**           | tokenID                                                                     |

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAppToken**
> TokenResponse UpdateAppToken(ctx, body, appID, tokenID)
Update app token (enable/disable or name)

### Required Parameters

| Name        | Type                                    | Description                                                                 | Notes |
| ----------- | --------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context**                     | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**    | [**UpdateTokenDto**](UpdateTokenDto.md) | dto                                                                         |
| **appID**   | **int64**                               | appID                                                                       |
| **tokenID** | **int64**                               | tokenID                                                                     |

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
