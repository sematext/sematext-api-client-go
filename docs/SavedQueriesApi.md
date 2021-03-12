# \SavedQueriesAPI

All URIs are relative to *https://localhost*

| Method                                                                                | HTTP request                                                  | Description                  |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------- | ---------------------------- |
| [**DeleteSavedQueryUsingDELETE**](SavedQueriesAPI.md#DeleteSavedQueryUsingDELETE)     | **Delete** /users-web/api/v3/savedQueries/{updateableQueryID} | Delete saved query           |
| [**GetSavedQueriesForAppUsingGET**](SavedQueriesAPI.md#GetSavedQueriesForAppUsingGET) | **Get** /users-web/api/v3/apps/{appID}/savedQueries           | Get saved queries for an app |
| [**SaveQueryUsingPOST**](SavedQueriesAPI.md#SaveQueryUsingPOST)                       | **Post** /users-web/api/v3/savedQueries                       | Create saved query           |
| [**SaveQueryUsingPUT**](SavedQueriesAPI.md#SaveQueryUsingPUT)                         | **Put** /users-web/api/v3/savedQueries/{updateableQueryID}    | Update saved query           |


# **DeleteSavedQueryUsingDELETE**
> GenericAPIResponse DeleteSavedQueryUsingDELETE(ctx, updateableQueryID)
Delete saved query

### Required Parameters

| Name                  | Type                | Description                                                                 | Notes |
| --------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **updateableQueryID** | **int64**           | updateableQueryID                                                           |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSavedQueriesForAppUsingGET**
> GenericAPIResponse GetSavedQueriesForAppUsingGET(ctx, appID)
Get saved queries for an app

### Required Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appID** | **int64**           | appID                                                                       |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveQueryUsingPOST**
> GenericAPIResponse SaveQueryUsingPOST(ctx, savedQueryDto)
Create saved query

### Required Parameters

| Name              | Type                            | Description                                                                 | Notes |
| ----------------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **savedQueryDto** | [**SavedQuery**](SavedQuery.md) | savedQueryDto                                                               |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveQueryUsingPUT**
> GenericAPIResponse SaveQueryUsingPUT(ctx, savedQueryDto, updateableQueryID)
Update saved query

### Required Parameters

| Name                  | Type                            | Description                                                                 | Notes |
| --------------------- | ------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**               | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **savedQueryDto**     | [**SavedQuery**](SavedQuery.md) | savedQueryDto                                                               |
| **updateableQueryID** | **int64**                       | updateableQueryID                                                           |

### Return type

[**GenericAPIResponse**](Generic API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
