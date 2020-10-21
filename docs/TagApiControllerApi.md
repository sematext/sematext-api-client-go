# \TagApiControllerApi

All URIs are relative to *https://localhost*

| Method                                                                  | HTTP request                                              | Description                                                                                             |
| ----------------------------------------------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| [**GetTagNamesUsingGET1**](TagApiControllerApi.md#GetTagNamesUsingGET1) | **Get** /spm-reports/api/v3/apps/{appIds}/tagNames        | Gets tag names for the given application identifiers appearing in the given time frame.                 |
| [**GetUsingGET1**](TagApiControllerApi.md#GetUsingGET1)                 | **Get** /spm-reports/api/v3/apps/{appIds}/tags            | Gets values for specified tags for the given application identifiers appearing in the given time frame. |
| [**GetUsingGET2**](TagApiControllerApi.md#GetUsingGET2)                 | **Get** /spm-reports/api/v3/apps/{appIds}/metrics/filters | Gets values for specified tags for the given application identifiers appearing in the given time frame. |


# **GetTagNamesUsingGET1**
> interface{} GetTagNamesUsingGET1(ctx, appIds, optional)
Gets tag names for the given application identifiers appearing in the given time frame.

### Required Parameters

| Name         | Type                          | Description                                                                 | Notes                |
| ------------ | ----------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**           | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appIds**   | **string**                    | appIds                                                                      |
| **optional** | ***GetTagNamesUsingGET1Opts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a GetTagNamesUsingGET1Opts struct

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **from** | **optional.Int64**| from |
 **to** | **optional.Int64**| to |
 **metrics** | **optional.Bool**| metrics | [default to true]
 **logs** | **optional.Bool**| logs | [default to true]
 **events** | **optional.Bool**| events | [default to false]
 **rum** | **optional.Bool**| rum | [default to true]

### Return type

[**interface{}**](interface{}.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET1**
> interface{} GetUsingGET1(ctx, appIds, tag, optional)
Gets values for specified tags for the given application identifiers appearing in the given time frame.

### Required Parameters

| Name         | Type                      | Description                                                                 | Notes                |
| ------------ | ------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**       | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appIds**   | **string**                | appIds                                                                      |
| **tag**      | [**[]string**](string.md) | tag                                                                         |
| **optional** | ***GetUsingGET1Opts**     | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsingGET1Opts struct

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **from** | **optional.Int64**| from |
 **to** | **optional.Int64**| to |
 **metrics** | **optional.Bool**| metrics | [default to true]
 **logs** | **optional.Bool**| logs | [default to true]
 **events** | **optional.Bool**| events | [default to false]
 **rum** | **optional.Bool**| rum | [default to true]

### Return type

[**interface{}**](interface{}.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET2**
> interface{} GetUsingGET2(ctx, appIds, tag, optional)
Gets values for specified tags for the given application identifiers appearing in the given time frame.

### Required Parameters

| Name         | Type                      | Description                                                                 | Notes                |
| ------------ | ------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**       | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appIds**   | **string**                | appIds                                                                      |
| **tag**      | [**[]string**](string.md) | tag                                                                         |
| **optional** | ***GetUsingGET2Opts**     | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsingGET2Opts struct

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **from** | **optional.Int64**| from |
 **to** | **optional.Int64**| to |
 **metrics** | **optional.Bool**| metrics | [default to true]
 **logs** | **optional.Bool**| logs | [default to true]
 **events** | **optional.Bool**| events | [default to false]
 **rum** | **optional.Bool**| rum | [default to true]

### Return type

[**interface{}**](interface{}.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
