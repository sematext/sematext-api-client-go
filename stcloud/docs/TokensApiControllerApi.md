# {{classname}}

All URIs are relative to */*

| Method                                                                 | HTTP request                                                        | Description                               |
| ---------------------------------------------------------------------- | ------------------------------------------------------------------- | ----------------------------------------- |
| [**CreateAppToken1**](TokensApiControllerApi.md#CreateAppToken1)       | **Post** /users-web/api/v3/apps/{appId}/tokens                      | Create new app token                      |
| [**DeleteAppToken1**](TokensApiControllerApi.md#DeleteAppToken1)       | **Delete** /users-web/api/v3/apps/{appId}/tokens/{tokenId}          | Delete app token                          |
| [**GetAppTokens1**](TokensApiControllerApi.md#GetAppTokens1)           | **Get** /users-web/api/v3/apps/{appId}/tokens                       | Get app available tokens                  |
| [**RegenerateAppToken**](TokensApiControllerApi.md#RegenerateAppToken) | **Post** /users-web/api/v3/apps/{appId}/tokens/{tokenId}/regenerate | Regenerate app token)                     |
| [**UpdateAppToken1**](TokensApiControllerApi.md#UpdateAppToken1)       | **Put** /users-web/api/v3/apps/{appId}/tokens/{tokenId}             | Update app token (enable/disable or name) |

# **CreateAppToken1**

> TokenResponse CreateAppToken1(ctx, body, appId)
Create new app token

### Required Parameters

| Name      | Type                                    | Description                                                                 | Notes |
| --------- | --------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                     | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**CreateTokenDto**](CreateTokenDto.md) | dto                                                                         |
| **appId** | **int64**                               | appId                                                                       |

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAppToken1**

> GenericMapBasedApiResponse DeleteAppToken1(ctx, appId, tokenId)
Delete app token

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId**   | **int64**           | appId                                                                       |
| **tokenId** | **int64**           | tokenId                                                                     |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppTokens1**

> TokensResponse GetAppTokens1(ctx, appId)
Get app available tokens

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**           | appId                                                                       |

### Return type

[**TokensResponse**](TokensResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateAppToken**

> TokenResponse RegenerateAppToken(ctx, appId, tokenId)
Regenerate app token)

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId**   | **int64**           | appId                                                                       |
| **tokenId** | **int64**           | tokenId                                                                     |

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAppToken1**

> TokenResponse UpdateAppToken1(ctx, body, appId, tokenId)
Update app token (enable/disable or name)

### Required Parameters

| Name        | Type                                    | Description                                                                 | Notes |
| ----------- | --------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context**                     | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**    | [**UpdateTokenDto**](UpdateTokenDto.md) | dto                                                                         |
| **appId**   | **int64**                               | appId                                                                       |
| **tokenId** | **int64**                               | tokenId                                                                     |

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
