# \BillingApi

All URIs are relative to *https://localhost*

| Method                                                                     | HTTP request                                                       | Description            |
| -------------------------------------------------------------------------- | ------------------------------------------------------------------ | ---------------------- |
| [**GetDetailedInvoiceUsingGET**](BillingApi.md#GetDetailedInvoiceUsingGET) | **Get** /users-web/api/v3/billing/invoice/{service}/{year}/{month} | Get invoice details    |
| [**ListAvailablePlansUsingGET**](BillingApi.md#ListAvailablePlansUsingGET) | **Get** /users-web/api/v3/billing/availablePlans                   | Get available plans    |
| [**UpdatePlanUsingPUT**](BillingApi.md#UpdatePlanUsingPUT)                 | **Put** /users-web/api/v3/billing/info/{appId}                     | Update plan for an app |


# **GetDetailedInvoiceUsingGET**
> GenericApiResponse GetDetailedInvoiceUsingGET(ctx, service, year, month)
Get invoice details

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **service** | **string**          | service                                                                     |
| **year**    | **int32**           | year                                                                        |
| **month**   | **int32**           | month                                                                       |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAvailablePlansUsingGET**
> GenericApiResponse ListAvailablePlansUsingGET(ctx, optional)
Get available plans

### Required Parameters

| Name         | Type                                | Description                                                                 | Notes                |
| ------------ | ----------------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**                 | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **optional** | ***ListAvailablePlansUsingGETOpts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a ListAvailablePlansUsingGETOpts struct

| Name              | Type                | Description   | Notes |
| ----------------- | ------------------- | ------------- | ----- |
| **integrationId** | **optional.Int64**  | integrationId |
| **appType**       | **optional.String** | appType       |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePlanUsingPUT**
> GenericApiResponse UpdatePlanUsingPUT(ctx, appId, dto)
Update plan for an app

### Required Parameters

| Name      | Type                              | Description                                                                 | Notes |
| --------- | --------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**               | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appId** | **int64**                         | appId                                                                       |
| **dto**   | [**BillingInfo**](BillingInfo.md) | dto                                                                         |

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
