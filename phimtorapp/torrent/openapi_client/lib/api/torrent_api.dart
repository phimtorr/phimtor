//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;


class TorrentApi {
  TorrentApi([ApiClient? apiClient]) : apiClient = apiClient ?? defaultApiClient;

  final ApiClient apiClient;

  /// Add torrent
  ///
  /// Add torrent
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [AddTorrentRequest] addTorrentRequest (required):
  ///
  /// * [bool] dropOthers:
  ///   Drop other torrents
  ///
  /// * [bool] deleteOthers:
  ///   Delete other torrents
  Future<Response> addTorrentWithHttpInfo(AddTorrentRequest addTorrentRequest, { bool? dropOthers, bool? deleteOthers, }) async {
    // ignore: prefer_const_declarations
    final path = r'/torrents';

    // ignore: prefer_final_locals
    Object? postBody = addTorrentRequest;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

    if (dropOthers != null) {
      queryParams.addAll(_queryParams('', 'dropOthers', dropOthers));
    }
    if (deleteOthers != null) {
      queryParams.addAll(_queryParams('', 'deleteOthers', deleteOthers));
    }

    const contentTypes = <String>['application/json'];


    return apiClient.invokeAPI(
      path,
      'POST',
      queryParams,
      postBody,
      headerParams,
      formParams,
      contentTypes.isEmpty ? null : contentTypes.first,
    );
  }

  /// Add torrent
  ///
  /// Add torrent
  ///
  /// Parameters:
  ///
  /// * [AddTorrentRequest] addTorrentRequest (required):
  ///
  /// * [bool] dropOthers:
  ///   Drop other torrents
  ///
  /// * [bool] deleteOthers:
  ///   Delete other torrents
  Future<AddTorrent200Response?> addTorrent(AddTorrentRequest addTorrentRequest, { bool? dropOthers, bool? deleteOthers, }) async {
    final response = await addTorrentWithHttpInfo(addTorrentRequest,  dropOthers: dropOthers, deleteOthers: deleteOthers, );
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'AddTorrent200Response',) as AddTorrent200Response;
    
    }
    return null;
  }

  /// Delete torrent
  ///
  /// Delete torrent
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [String] infoHash (required):
  ///   Torrent info hash
  Future<Response> deleteTorrentWithHttpInfo(String infoHash,) async {
    // ignore: prefer_const_declarations
    final path = r'/torrents/{infoHash}'
      .replaceAll('{infoHash}', infoHash);

    // ignore: prefer_final_locals
    Object? postBody;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

    const contentTypes = <String>[];


    return apiClient.invokeAPI(
      path,
      'DELETE',
      queryParams,
      postBody,
      headerParams,
      formParams,
      contentTypes.isEmpty ? null : contentTypes.first,
    );
  }

  /// Delete torrent
  ///
  /// Delete torrent
  ///
  /// Parameters:
  ///
  /// * [String] infoHash (required):
  ///   Torrent info hash
  Future<void> deleteTorrent(String infoHash,) async {
    final response = await deleteTorrentWithHttpInfo(infoHash,);
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
  }

  /// List torrents
  ///
  /// List of torrents
  ///
  /// Note: This method returns the HTTP [Response].
  Future<Response> listTorrentsWithHttpInfo() async {
    // ignore: prefer_const_declarations
    final path = r'/torrents';

    // ignore: prefer_final_locals
    Object? postBody;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

    const contentTypes = <String>[];


    return apiClient.invokeAPI(
      path,
      'GET',
      queryParams,
      postBody,
      headerParams,
      formParams,
      contentTypes.isEmpty ? null : contentTypes.first,
    );
  }

  /// List torrents
  ///
  /// List of torrents
  Future<ListTorrents200Response?> listTorrents() async {
    final response = await listTorrentsWithHttpInfo();
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'ListTorrents200Response',) as ListTorrents200Response;
    
    }
    return null;
  }
}
