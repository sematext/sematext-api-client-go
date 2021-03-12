# {{classname}}

All URIs are relative to */*

| Method                                                                       | HTTP request                                                       | Description            |
| ---------------------------------------------------------------------------- | ------------------------------------------------------------------ | ---------------------- |
| [**GetDetailedInvoiceUsingGET**](BillingApi.md#GetDetailedInvoiceUsingGET)   | **Get** /users-web/api/v3/billing/invoice/{service}/{year}/{month} | Get invoice details    |
| [**ListAvailablePlansUsingGET1**](BillingApi.md#ListAvailablePlansUsingGET1) | **Get** /users-web/api/v3/billing/availablePlans                   | Get available plans    |
| [**UpdatePlanUsingPUT1**](BillingApi.md#UpdatePlanUsingPUT1)                 | **Put** /users-web/api/v3/billing/info/{appId}                     | Update plan for an app |

# **GetDetailedInvoiceUsingGET**
> InvoiceResponse GetDetailedInvoiceUsingGET(ctx, service, year, month)
Get invoice details

### Required Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **service** | **string**          | service                                                                     |
| **year**    | **int32**           | year                                                                        |
| **month**   | **int32**           | month                                                                       |

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAvailablePlansUsingGET1**
> PlansResponse ListAvailablePlansUsingGET1(ctx, optional)
Get available plans

### Required Parameters

| Name         | Type                                           | Description                                                                 | Notes                |
| ------------ | ---------------------------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**                            | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **optional** | ***BillingApiListAvailablePlansUsingGET1Opts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a BillingApiListAvailablePlansUsingGET1Opts struct
| Name              | Type                | Description   | Notes |
| ----------------- | ------------------- | ------------- | ----- |
| **integrationId** | **optional.Int64**  | integrationId |
| **appType**       | **optional.String** | appType       |

### Return type

[**PlansResponse**](PlansResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePlanUsingPUT1**
> UpdatePlanResponse UpdatePlanUsingPUT1(ctx, body, appId)
Update plan for an app

### Required Parameters

| Name      | Type                              | Description                                                                 | Notes |
| --------- | --------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context**               | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body**  | [**BillingInfo**](BillingInfo.md) | dto                                                                         |
| **appId** | **int64**                         | appId                                                                       |

### Return type

[**UpdatePlanResponse**](UpdatePlanResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
