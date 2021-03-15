# {{classname}}

All URIs are relative to */*

| Method                                                                   | HTTP request                                      | Description    |
| ------------------------------------------------------------------------ | ------------------------------------------------- | -------------- |
| [**ResetPasswordUsingPOST**](ResetPasswordAPI.md#ResetPasswordUsingPOST) | **Post** /users-web/api/v3/account/password/reset | Reset Password |

# **ResetPasswordUsingPOST**
> GenericMapBasedAPIResponse ResetPasswordUsingPOST(ctx, body)
Reset Password

### Required Parameters

| Name     | Type                        | Description                                                                 | Notes |
| -------- | --------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**         | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**UserInfo**](UserInfo.md) | dto                                                                         |

### Return type

[**GenericMapBasedAPIResponse**](Generic Map Based API Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
