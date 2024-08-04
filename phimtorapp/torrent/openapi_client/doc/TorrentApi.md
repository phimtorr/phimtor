# openapi_client.api.TorrentApi

## Load the API package
```dart
import 'package:openapi_client/api.dart';
```

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addTorrent**](TorrentApi.md#addtorrent) | **POST** /torrents | Add torrent
[**deleteTorrent**](TorrentApi.md#deletetorrent) | **DELETE** /torrents/{infoHash} | Delete torrent
[**listTorrents**](TorrentApi.md#listtorrents) | **GET** /torrents | List torrents


# **addTorrent**
> AddTorrent200Response addTorrent(addTorrentRequest, dropOthers, deleteOthers)

Add torrent

Add torrent

### Example
```dart
import 'package:openapi_client/api.dart';

final api_instance = TorrentApi();
final addTorrentRequest = AddTorrentRequest(); // AddTorrentRequest | 
final dropOthers = true; // bool | Drop other torrents
final deleteOthers = true; // bool | Delete other torrents

try {
    final result = api_instance.addTorrent(addTorrentRequest, dropOthers, deleteOthers);
    print(result);
} catch (e) {
    print('Exception when calling TorrentApi->addTorrent: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTorrentRequest** | [**AddTorrentRequest**](AddTorrentRequest.md)|  | 
 **dropOthers** | **bool**| Drop other torrents | [optional] 
 **deleteOthers** | **bool**| Delete other torrents | [optional] 

### Return type

[**AddTorrent200Response**](AddTorrent200Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteTorrent**
> deleteTorrent(infoHash)

Delete torrent

Delete torrent

### Example
```dart
import 'package:openapi_client/api.dart';

final api_instance = TorrentApi();
final infoHash = infoHash_example; // String | Torrent info hash

try {
    api_instance.deleteTorrent(infoHash);
} catch (e) {
    print('Exception when calling TorrentApi->deleteTorrent: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **infoHash** | **String**| Torrent info hash | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listTorrents**
> ListTorrents200Response listTorrents()

List torrents

List of torrents

### Example
```dart
import 'package:openapi_client/api.dart';

final api_instance = TorrentApi();

try {
    final result = api_instance.listTorrents();
    print(result);
} catch (e) {
    print('Exception when calling TorrentApi->listTorrents: $e\n');
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ListTorrents200Response**](ListTorrents200Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

