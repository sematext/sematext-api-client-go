# \SavedQueriesApi

All URIs are relative to *https://localhost*

| Method                                                                                | HTTP request                                                  | Description                  |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------- | ---------------------------- |
| [**DeleteSavedQueryUsingDELETE**](SavedQueriesApi.md#DeleteSavedQueryUsingDELETE)     | **Delete** /users-web/api/v3/savedQueries/{updateableQueryId} | Delete saved query           |
| [**GetSavedQueriesForAppUsingGET**](SavedQueriesApi.md#GetSavedQueriesForAppUsingGET) | **Get** /users-web/api/v3/apps/{appId}/savedQueries           | Get saved queries for an app |
| [**SaveQueryUsingPOST**](SavedQueriesApi.md#SaveQueryUsingPOST)                       | **Post** /users-web/api/v3/savedQueries                       | Create saved query           |
| [**SaveQueryUsingPUT**](SavedQueriesApi.md#SaveQueryUsingPUT)                         | **Put** /users-web/api/v3/savedQueries/{updateableQueryId}    | Update saved query           |


# **DeleteSavedQueryUsingDELETE**
> GenericApiResponse DeleteSavedQueryUsingDELETE(ctx, updateableQueryId)
Delete saved query

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableQueryId** | **int64**           | updateableQueryId                                                           |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSavedQueriesForAppUsingGET**
> GenericApiResponse GetSavedQueriesForAppUsingGET(ctx, appId)
Get saved queries for an app

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**           | appId                                                                       |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveQueryUsingPOST**
> GenericApiResponse SaveQueryUsingPOST(ctx, savedQueryDto)
Create saved query

### Required Parameters

| Name              | Type                            | Description                                                                 | Notes |
| ----------------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **savedQueryDto** | [**SavedQuery**](SavedQuery.md) | savedQueryDto                                                               |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveQueryUsingPUT**
> GenericApiResponse SaveQueryUsingPUT(ctx, savedQueryDto, updateableQueryId)
Update saved query

### Required Parameters

| Name                  | Type                            | Description                                                                 | Notes |
| --------------------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **savedQueryDto**     | [**SavedQuery**](SavedQuery.md) | savedQueryDto                                                               |
| **updateableQueryId** | **int64**                       | updateableQueryId                                                           |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
