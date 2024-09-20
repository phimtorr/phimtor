# phimtor_openapi_client.api.DefaultApi

## Load the API package
```dart
import 'package:phimtor_openapi_client/api.dart';
```

All URIs are relative to *http://localhost:8080/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getMovie**](DefaultApi.md#getmovie) | **GET** /movies/{movieId} | Get movie by id
[**getTvSeason**](DefaultApi.md#gettvseason) | **GET** /tv-series/{tvSeriesId}/seasons/{seasonNumber} | Get tv season by tv series id and season number
[**getTvSeries**](DefaultApi.md#gettvseries) | **GET** /tv-series/{tvSeriesId} | Get tv series by id
[**getVersion**](DefaultApi.md#getversion) | **GET** /version | Get version
[**getVideo**](DefaultApi.md#getvideo) | **GET** /videos/{id} | Get video by id
[**listLatestEpisodes**](DefaultApi.md#listlatestepisodes) | **GET** /shows/latest-episodes | List latest episodes
[**listLatestMovies**](DefaultApi.md#listlatestmovies) | **GET** /shows/latest-movies | List latest movies
[**listLatestTvSeries**](DefaultApi.md#listlatesttvseries) | **GET** /shows/latest-tv-series | List latest tv series
[**listRecentlyAddedMovies**](DefaultApi.md#listrecentlyaddedmovies) | **GET** /shows/recently-added-movies | List recently added movies
[**searchShows**](DefaultApi.md#searchshows) | **GET** /shows/search | Search shows


# **getMovie**
> GetMovieResponse getMovie(movieId)

Get movie by id

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final movieId = 789; // int | 

try {
    final result = api_instance.getMovie(movieId);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->getMovie: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int**|  | 

### Return type

[**GetMovieResponse**](GetMovieResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getTvSeason**
> GetTvSeasonResponse getTvSeason(tvSeriesId, seasonNumber)

Get tv season by tv series id and season number

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final tvSeriesId = 789; // int | 
final seasonNumber = 56; // int | 

try {
    final result = api_instance.getTvSeason(tvSeriesId, seasonNumber);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->getTvSeason: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvSeriesId** | **int**|  | 
 **seasonNumber** | **int**|  | 

### Return type

[**GetTvSeasonResponse**](GetTvSeasonResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getTvSeries**
> GetTvSeriesResponse getTvSeries(tvSeriesId)

Get tv series by id

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final tvSeriesId = 789; // int | 

try {
    final result = api_instance.getTvSeries(tvSeriesId);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->getTvSeries: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvSeriesId** | **int**|  | 

### Return type

[**GetTvSeriesResponse**](GetTvSeriesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

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

[bearerAuth](../README.md#bearerAuth)

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
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listLatestEpisodes**
> GetLatestEpisodesResponse listLatestEpisodes(page, pageSize)

List latest episodes

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final page = 56; // int | 
final pageSize = 56; // int | 

try {
    final result = api_instance.listLatestEpisodes(page, pageSize);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->listLatestEpisodes: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**|  | [optional] [default to 1]
 **pageSize** | **int**|  | [optional] [default to 18]

### Return type

[**GetLatestEpisodesResponse**](GetLatestEpisodesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listLatestMovies**
> GetLatestMoviesResponse listLatestMovies(page, pageSize)

List latest movies

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final page = 56; // int | 
final pageSize = 56; // int | 

try {
    final result = api_instance.listLatestMovies(page, pageSize);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->listLatestMovies: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**|  | [optional] [default to 1]
 **pageSize** | **int**|  | [optional] [default to 18]

### Return type

[**GetLatestMoviesResponse**](GetLatestMoviesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listLatestTvSeries**
> GetLatestTvSeriesResponse listLatestTvSeries(page, pageSize)

List latest tv series

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final page = 56; // int | 
final pageSize = 56; // int | 

try {
    final result = api_instance.listLatestTvSeries(page, pageSize);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->listLatestTvSeries: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**|  | [optional] [default to 1]
 **pageSize** | **int**|  | [optional] [default to 18]

### Return type

[**GetLatestTvSeriesResponse**](GetLatestTvSeriesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listRecentlyAddedMovies**
> GetLatestMoviesResponse listRecentlyAddedMovies(page, pageSize)

List recently added movies

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final page = 56; // int | 
final pageSize = 56; // int | 

try {
    final result = api_instance.listRecentlyAddedMovies(page, pageSize);
    print(result);
} catch (e) {
    print('Exception when calling DefaultApi->listRecentlyAddedMovies: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**|  | [optional] [default to 1]
 **pageSize** | **int**|  | [optional] [default to 18]

### Return type

[**GetLatestMoviesResponse**](GetLatestMoviesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **searchShows**
> SearchShowsResponse searchShows(query, page, pageSize)

Search shows

### Example
```dart
import 'package:phimtor_openapi_client/api.dart';
// TODO Configure HTTP Bearer authorization: bearerAuth
// Case 1. Use String Token
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken('YOUR_ACCESS_TOKEN');
// Case 2. Use Function which generate token.
// String yourTokenGeneratorFunction() { ... }
//defaultApiClient.getAuthentication<HttpBearerAuth>('bearerAuth').setAccessToken(yourTokenGeneratorFunction);

final api_instance = DefaultApi();
final query = query_example; // String | 
final page = 56; // int | 
final pageSize = 56; // int | 

try {
    final result = api_instance.searchShows(query, page, pageSize);
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
 **pageSize** | **int**|  | [optional] [default to 18]

### Return type

[**SearchShowsResponse**](SearchShowsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

