# \WeightApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWeight**](WeightApi.md#GetWeight) | **Get** /weight | Returns the total size of the files under a directory.


# **GetWeight**
> WeightResponse GetWeight(ctx, ignorePaths)
Returns the total size of the files under a directory.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ignorePaths** | [**ExcludePath**](ExcludePath.md)| Specifies which directories to search and which to ignore. | 

### Return type

[**WeightResponse**](WeightResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

