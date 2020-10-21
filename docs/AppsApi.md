# \AppsApi

All URIs are relative to *https://localhost*

| Method                                                                | HTTP request                                               | Description                                                        |
| --------------------------------------------------------------------- | ---------------------------------------------------------- | ------------------------------------------------------------------ |
| [**GetAppTypesUsingGET**](AppsApi.md#GetAppTypesUsingGET)             | **Get** /users-web/api/v3/apps/types                       | Get all App types supported for the account identified with apiKey |
| [**GetUsingGET**](AppsApi.md#GetUsingGET)                             | **Get** /users-web/api/v3/apps/{anyStateAppId}             | Gets defails for one particular App                                |
| [**InviteAppGuestsUsingPOST**](AppsApi.md#InviteAppGuestsUsingPOST)   | **Post** /users-web/api/v3/apps/guests                     | Invite guests to an app                                            |
| [**ListAppsUsersUsingGET**](AppsApi.md#ListAppsUsersUsingGET)         | **Get** /users-web/api/v3/apps/users                       | Get all users of apps accessible to this account                   |
| [**ListUsingGET**](AppsApi.md#ListUsingGET)                           | **Get** /users-web/api/v3/apps                             | Get all apps accessible by account identified with apiKey          |
| [**UpdateDescriptionUsingPUT**](AppsApi.md#UpdateDescriptionUsingPUT) | **Put** /users-web/api/v3/apps/{anyStateAppId}/description | Update description of the app                                      |
| [**UpdateUsingPUT1**](AppsApi.md#UpdateUsingPUT1)                     | **Put** /users-web/api/v3/apps/{anyStateAppId}             | Update app                                                         |


# **GetAppTypesUsingGET**
> GenericApiResponse GetAppTypesUsingGET(ctx, )
Get all App types supported for the account identified with apiKey

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET**
> GenericApiResponse GetUsingGET(ctx, anyStateAppId)
Gets defails for one particular App

### Required Parameters

| Name              | Type                | Description                                                                 | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **anyStateAppId** | **int64**           | anyStateAppId                                                               |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteAppGuestsUsingPOST**
> GenericApiResponse InviteAppGuestsUsingPOST(ctx, invitation)
Invite guests to an app

### Required Parameters

| Name           | Type                            | Description                                                                                                                                     | Notes |
| -------------- | ------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**             | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                     |
| **invitation** | [**Invitation**](Invitation.md) | For &#x60;app&#x60; and &#x60;apps&#x60; fields only &#x60;id&#x60; needs to be populated.Other fields can be left empty or with default values |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppsUsersUsingGET**
> GenericApiResponse ListAppsUsersUsingGET(ctx, )
Get all users of apps accessible to this account

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET**
> GenericApiResponse ListUsingGET(ctx, )
Get all apps accessible by account identified with apiKey

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDescriptionUsingPUT**
> GenericApiResponse UpdateDescriptionUsingPUT(ctx, anyStateAppId, optional)
Update description of the app

App can be in any state

### Required Parameters

| Name              | Type                               | Description                                                                 | Notes                |
| ----------------- | ---------------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**           | **context.Context**                | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **anyStateAppId** | **int64**                          | App Id                                                                      |
| **optional**      | ***UpdateDescriptionUsingPUTOpts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDescriptionUsingPUTOpts struct

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **updateDetails** | [**optional.Interface of AppDescription**](AppDescription.md)| Update Details |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT1**
> GenericApiResponse UpdateUsingPUT1(ctx, dto, anyStateAppId)
Update app

App can be in any state

### Required Parameters

| Name              | Type                                  | Description                                                                 | Notes |
| ----------------- | ------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context**                   | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **dto**           | [**UpdateAppInfo**](UpdateAppInfo.md) | dto                                                                         |
| **anyStateAppId** | **int64**                             | App Id                                                                      |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
