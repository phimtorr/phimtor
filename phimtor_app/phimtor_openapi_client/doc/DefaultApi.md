# phimtor_openapi_client.api.DefaultApi

## Load the API package
```dart
import 'package:phimtor_openapi_client/api.dart';
```

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getMovie**](DefaultApi.md#getmovie) | **GET** /movies/{id} | Get movie by id
[**getSeries**](DefaultApi.md#getseries) | **GET** /series/{id} | Get series by id
[**getVersion**](DefaultApi.md#getversion) | **GET** /version | Get version
[**getVideo**](DefaultApi.md#getvideo) | **GET** /videos/{id} | Get video by id
[**listShows**](DefaultApi.md#listshows) | **GET** /shows | List all shows
[**searchShows**](DefaultApi.md#searchshows) | **GET** /shows/search | Search shows


# **getMovie**
> GetMovieResponse getMovie(id)

Get movie by id

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';

final api_instance = DefaultApi();
final id = 789; // int | 

try {
    final result = api_instance.getMovie(id);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->getMovie: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**|  | 

### Return type

[**GetMovieResponse**](GetMovieResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getSeries**
> GetSeriesResponse getSeries(id)

Get series by id

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';

final api_instance = DefaultApi();
final id = 789; // int | 

try {
    final result = api_instance.getSeries(id);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->getSeries: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**|  | 

### Return type

[**GetSeriesResponse**](GetSeriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getVersion**
> GetVersionResponse getVersion()

Get version

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';

final api_instance = DefaultApi();

try {
    final result = api_instance.getVersion();
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->getVersion: $e\n');
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**GetVersionResponse**](GetVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getVideo**
> GetVideoResponse getVideo(id)

Get video by id

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';

final api_instance = DefaultApi();
final id = 789; // int | 

try {
    final result = api_instance.getVideo(id);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->getVideo: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**|  | 

### Return type

[**GetVideoResponse**](GetVideoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listShows**
> ListShowsResponse listShows(page, pageSize, type)

List all shows

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';

final api_instance = DefaultApi();
final page = 56; // int | 
final pageSize = 56; // int | 
final type = ; // ShowType | 

try {
    final result = api_instance.listShows(page, pageSize, type);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->listShows: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**|  | [optional] [default to 1]
 **pageSize** | **int**|  | [optional] [default to 18]
 **type** | [**ShowType**](.md)|  | [optional] 

### Return type

[**ListShowsResponse**](ListShowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **searchShows**
> SearchShowsResponse searchShows(query, page)

Search shows

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';

final api_instance = DefaultApi();
final query = query_example; // String | 
final page = 56; // int | 

try {
    final result = api_instance.searchShows(query, page);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->searchShows: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **String**|  | 
 **page** | **int**|  | [optional] [default to 1]

### Return type

[**SearchShowsResponse**](SearchShowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

