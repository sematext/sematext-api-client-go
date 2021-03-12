# {{classname}}

All URIs are relative to */*

| Method                                                                  | HTTP request                                               | Description                                                        |
| ----------------------------------------------------------------------- | ---------------------------------------------------------- | ------------------------------------------------------------------ |
| [**DeleteUsingDELETE1**](AppsApi.md#DeleteUsingDELETE1)                 | **Delete** /users-web/api/v3/apps/{anyStateAppId}          | delete                                                             |
| [**GetAppTypesUsingGET**](AppsApi.md#GetAppTypesUsingGET)               | **Get** /users-web/api/v3/apps/types                       | Get all App types supported for the account identified with apiKey |
| [**GetUsingGET**](AppsApi.md#GetUsingGET)                               | **Get** /users-web/api/v3/apps/{anyStateAppId}             | Gets defails for one particular App                                |
| [**InviteAppGuestsUsingPOST**](AppsApi.md#InviteAppGuestsUsingPOST)     | **Post** /users-web/api/v3/apps/guests                     | Invite guests to an app                                            |
| [**ListAppsUsersUsingGET1**](AppsApi.md#ListAppsUsersUsingGET1)         | **Get** /users-web/api/v3/apps/users                       | Get all users of apps accessible to this account                   |
| [**ListUsingGET**](AppsApi.md#ListUsingGET)                             | **Get** /users-web/api/v3/apps                             | Get all apps accessible by account identified with apiKey          |
| [**UpdateDescriptionUsingPUT1**](AppsApi.md#UpdateDescriptionUsingPUT1) | **Put** /users-web/api/v3/apps/{anyStateAppId}/description | Update description of the app                                      |
| [**UpdateUsingPUT3**](AppsApi.md#UpdateUsingPUT3)                       | **Put** /users-web/api/v3/apps/{anyStateAppId}             | Update app                                                         |

# **DeleteUsingDELETE1**
> GenericMapBasedApiResponse DeleteUsingDELETE1(ctx, anyStateAppId)
delete

### Required Parameters

| Name              | Type                | Description                                                                 | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **anyStateAppId** | **int64**           | anyStateAppId                                                               |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppTypesUsingGET**
> AppTypesResponse GetAppTypesUsingGET(ctx, )
Get all App types supported for the account identified with apiKey

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AppTypesResponse**](AppTypesResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET**
> AppResponse GetUsingGET(ctx, anyStateAppId)
Gets defails for one particular App

### Required Parameters

| Name              | Type                | Description                                                                 | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **anyStateAppId** | **int64**           | anyStateAppId                                                               |

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteAppGuestsUsingPOST**
> GenericMapBasedApiResponse InviteAppGuestsUsingPOST(ctx, body)
Invite guests to an app

### Required Parameters

| Name     | Type                            | Description                                                                                                                                     | Notes |
| -------- | ------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                     |
| **body** | [**Invitation**](Invitation.md) | For &#x60;app&#x60; and &#x60;apps&#x60; fields only &#x60;id&#x60; needs to be populated.Other fields can be left empty or with default values |

### Return type

[**GenericMapBasedApiResponse**](Generic Map Based Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppsUsersUsingGET1**
> AppsResponse ListAppsUsersUsingGET1(ctx, )
Get all users of apps accessible to this account

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET**
> AppsResponse ListUsingGET(ctx, )
Get all apps accessible by account identified with apiKey

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDescriptionUsingPUT1**
> AppResponse UpdateDescriptionUsingPUT1(ctx, anyStateAppId, optional)
Update description of the app

App can be in any state

### Required Parameters

| Name              | Type                                       | Description                                                                 | Notes                |
| ----------------- | ------------------------------------------ | --------------------------------------------------------------------------- | -------------------- |
| **ctx**           | **context.Context**                        | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **anyStateAppId** | **int64**                                  | App Id                                                                      |
| **optional**      | ***AppsApiUpdateDescriptionUsingPUT1Opts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a AppsApiUpdateDescriptionUsingPUT1Opts struct
| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **body** | [**optional.Interface of AppDescription**](AppDescription.md)| Update Details |

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT3**
> AppResponse UpdateUsingPUT3(ctx, body, anyStateAppId)
Update app

App can be in any state

### Required Parameters

| Name              | Type                                  | Description                                                                 | Notes |
| ----------------- | ------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context**                   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**          | [**UpdateAppInfo**](UpdateAppInfo.md) | dto                                                                         |
| **anyStateAppId** | **int64**                             | App Id                                                                      |

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
