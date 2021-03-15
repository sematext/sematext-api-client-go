# {{classname}}

All URIs are relative to */*

| Method                                                                | HTTP request                                              | Description                                                                                             |
| --------------------------------------------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| [**GetTagNamesUsingGET**](TagAPIControllerAPI.md#GetTagNamesUsingGET) | **Get** /spm-reports/api/v3/apps/{appIDs}/tagNames        | Gets tag names for the given application identifiers appearing in the given time frame.                 |
| [**GetUsingGET2**](TagAPIControllerAPI.md#GetUsingGET2)               | **Get** /spm-reports/api/v3/apps/{appIDs}/metrics/filters | Gets values for specified tags for the given application identifiers appearing in the given time frame. |
| [**GetUsingGET3**](TagAPIControllerAPI.md#GetUsingGET3)               | **Get** /spm-reports/api/v3/apps/{appIDs}/tags            | Gets values for specified tags for the given application identifiers appearing in the given time frame. |

# **GetTagNamesUsingGET**
> TagNamesResponse GetTagNamesUsingGET(ctx, appIDs, optional)
Gets tag names for the given application identifiers appearing in the given time frame.

### Required Parameters

| Name         | Type                                            | Description                                                                 | Notes                |
| ------------ | ----------------------------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**                             | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appIDs**   | **string**                                      | appIDs                                                                      |
| **optional** | ***TagAPIControllerAPIGetTagNamesUsingGETOpts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a TagAPIControllerAPIGetTagNamesUsingGETOpts struct
| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **from** | **optional.Int64**| from |
 **to** | **optional.Int64**| to |
 **metrics** | **optional.Bool**| metrics | [default to true]
 **logs** | **optional.Bool**| logs | [default to true]
 **events** | **optional.Bool**| events | [default to false]
 **rum** | **optional.Bool**| rum | [default to true]

### Return type

[**TagNamesResponse**](TagNamesResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET2**
> map[string]Dimension GetUsingGET2(ctx, appIDs, tag, optional)
Gets values for specified tags for the given application identifiers appearing in the given time frame.

### Required Parameters

| Name         | Type                                     | Description                                                                 | Notes                |
| ------------ | ---------------------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**                      | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appIDs**   | **string**                               | appIDs                                                                      |
| **tag**      | [**[]string**](string.md)                | tag                                                                         |
| **optional** | ***TagAPIControllerAPIGetUsingGET2Opts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a TagAPIControllerAPIGetUsingGET2Opts struct
| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **from** | **optional.Int64**| from |
 **to** | **optional.Int64**| to |
 **metrics** | **optional.Bool**| metrics | [default to true]
 **logs** | **optional.Bool**| logs | [default to true]
 **events** | **optional.Bool**| events | [default to false]
 **rum** | **optional.Bool**| rum | [default to true]

### Return type

[**map[string]Dimension**](Dimension.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET3**
> map[string]Dimension GetUsingGET3(ctx, appIDs, tag, optional)
Gets values for specified tags for the given application identifiers appearing in the given time frame.

### Required Parameters

| Name         | Type                                     | Description                                                                 | Notes                |
| ------------ | ---------------------------------------- | --------------------------------------------------------------------------- | -------------------- |
| **ctx**      | **context.Context**                      | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **appIDs**   | **string**                               | appIDs                                                                      |
| **tag**      | [**[]string**](string.md)                | tag                                                                         |
| **optional** | ***TagAPIControllerAPIGetUsingGET3Opts** | optional parameters                                                         | nil if no parameters |

### Optional Parameters
Optional parameters are passed through a pointer to a TagAPIControllerAPIGetUsingGET3Opts struct
| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **from** | **optional.Int64**| from |
 **to** | **optional.Int64**| to |
 **metrics** | **optional.Bool**| metrics | [default to true]
 **logs** | **optional.Bool**| logs | [default to true]
 **events** | **optional.Bool**| events | [default to false]
 **rum** | **optional.Bool**| rum | [default to true]

### Return type

[**map[string]Dimension**](Dimension.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
