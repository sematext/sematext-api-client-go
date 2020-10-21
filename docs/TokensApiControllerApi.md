# \TokensAPIControllerAPI

All URIs are relative to *https://localhost*

| Method                                                                 | HTTP request                                                        | Description                       |
| ---------------------------------------------------------------------- | ------------------------------------------------------------------- | --------------------------------- |
| [**CreateAppToken**](TokensAPIControllerAPI.md#CreateAppToken)         | **Post** /users-web/api/v3/apps/{appId}/tokens                      | Create new app token              |
| [**DeleteAppToken1**](TokensAPIControllerAPI.md#DeleteAppToken1)       | **Delete** /users-web/api/v3/apps/{appId}/tokens/{tokenId}          | Delete app token                  |
| [**GetAppTokens1**](TokensAPIControllerAPI.md#GetAppTokens1)           | **Get** /users-web/api/v3/apps/{appId}/tokens                       | Get app available tokens          |
| [**RegenerateAppToken**](TokensAPIControllerAPI.md#RegenerateAppToken) | **Post** /users-web/api/v3/apps/{appId}/tokens/{tokenId}/regenerate | Regenerate app token)             |
| [**UpdateAppToken**](TokensAPIControllerAPI.md#UpdateAppToken)         | **Put** /users-web/api/v3/apps/{appId}/tokens/{tokenId}             | Update app token (enable/disable) |


# **CreateAppToken**
> GenericAPIResponse CreateAppToken(ctx, appId, dto)
Create new app token

### Required Parameters

| Name      | Type                                    | Description                                                                 | Notes |
| --------- | --------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**                     | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**                               | appId                                                                       |
| **dto**   | [**CreateTokenDto**](CreateTokenDto.md) | dto                                                                         |

### Return type

[**GenericAPIResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAppToken1**
> GenericAPIResponse DeleteAppToken1(ctx, appId, tokenId)
Delete app token

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId**   | **int64**           | appId                                                                       |
| **tokenId** | **int64**           | tokenId                                                                     |

### Return type

[**GenericAPIResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppTokens1**
> GenericAPIResponse GetAppTokens1(ctx, appId)
Get app available tokens

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**           | appId                                                                       |

### Return type

[**GenericAPIResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateAppToken**
> GenericAPIResponse RegenerateAppToken(ctx, appId, tokenId)
Regenerate app token)

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId**   | **int64**           | appId                                                                       |
| **tokenId** | **int64**           | tokenId                                                                     |

### Return type

[**GenericAPIResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAppToken**
> GenericAPIResponse UpdateAppToken(ctx, appId, tokenId, dto)
Update app token (enable/disable)

### Required Parameters

| Name        | Type                                    | Description                                                                 | Notes |
| ----------- | --------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context**                     | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId**   | **int64**                               | appId                                                                       |
| **tokenId** | **int64**                               | tokenId                                                                     |
| **dto**     | [**UpdateTokenDto**](UpdateTokenDto.md) | dto                                                                         |

### Return type

[**GenericAPIResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
