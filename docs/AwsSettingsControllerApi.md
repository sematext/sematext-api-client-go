# \AwsSettingsControllerApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateUsingPUT**](AwsSettingsControllerApi.md#UpdateUsingPUT) | **Put** /users-web/api/v3/apps/{appId}/aws | Update App&#39;s AWS CloudWatch settings


# **UpdateUsingPUT**
> GenericApiResponse UpdateUsingPUT(ctx, appId, dto)
Update App's AWS CloudWatch settings

Applicable only for AWS Apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **int64**| appId | 
  **dto** | [**CloudWatchSettings**](CloudWatchSettings.md)| dto | 

### Return type

[**GenericApiResponse**](Generic Api Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

