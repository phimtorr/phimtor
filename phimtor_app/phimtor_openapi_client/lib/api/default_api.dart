//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;


class DefaultApi {
  DefaultApi([ApiClient? apiClient]) : apiClient = apiClient ?? defaultApiClient;

  final ApiClient apiClient;

  /// Get movie by id
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] movieId (required):
  Future<Response> getMovieWithHttpInfo(int movieId,) async {
    // ignore: prefer_const_declarations
    final path = r'/movies/{movieId}'
      .replaceAll('{movieId}', movieId.toString());

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

  /// Get movie by id
  ///
  /// Parameters:
  ///
  /// * [int] movieId (required):
  Future<GetMovieResponse?> getMovie(int movieId,) async {
    final response = await getMovieWithHttpInfo(movieId,);
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetMovieResponse',) as GetMovieResponse;
    
    }
    return null;
  }

  /// Get tv season by tv series id and season number
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] tvSeriesId (required):
  ///
  /// * [int] seasonNumber (required):
  Future<Response> getTvSeasonWithHttpInfo(int tvSeriesId, int seasonNumber,) async {
    // ignore: prefer_const_declarations
    final path = r'/tv-series/{tvSeriesId}/seasons/{seasonNumber}'
      .replaceAll('{tvSeriesId}', tvSeriesId.toString())
      .replaceAll('{seasonNumber}', seasonNumber.toString());

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

  /// Get tv season by tv series id and season number
  ///
  /// Parameters:
  ///
  /// * [int] tvSeriesId (required):
  ///
  /// * [int] seasonNumber (required):
  Future<GetTvSeasonResponse?> getTvSeason(int tvSeriesId, int seasonNumber,) async {
    final response = await getTvSeasonWithHttpInfo(tvSeriesId, seasonNumber,);
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetTvSeasonResponse',) as GetTvSeasonResponse;
    
    }
    return null;
  }

  /// Get tv series by id
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] tvSeriesId (required):
  Future<Response> getTvSeriesWithHttpInfo(int tvSeriesId,) async {
    // ignore: prefer_const_declarations
    final path = r'/tv-series/{tvSeriesId}'
      .replaceAll('{tvSeriesId}', tvSeriesId.toString());

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

  /// Get tv series by id
  ///
  /// Parameters:
  ///
  /// * [int] tvSeriesId (required):
  Future<GetTvSeriesResponse?> getTvSeries(int tvSeriesId,) async {
    final response = await getTvSeriesWithHttpInfo(tvSeriesId,);
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetTvSeriesResponse',) as GetTvSeriesResponse;
    
    }
    return null;
  }

  /// Get version
  ///
  /// Note: This method returns the HTTP [Response].
  Future<Response> getVersionWithHttpInfo() async {
    // ignore: prefer_const_declarations
    final path = r'/version';

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

  /// Get version
  Future<GetVersionResponse?> getVersion() async {
    final response = await getVersionWithHttpInfo();
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetVersionResponse',) as GetVersionResponse;
    
    }
    return null;
  }

  /// Get video by id
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] id (required):
  Future<Response> getVideoWithHttpInfo(int id,) async {
    // ignore: prefer_const_declarations
    final path = r'/videos/{id}'
      .replaceAll('{id}', id.toString());

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

  /// Get video by id
  ///
  /// Parameters:
  ///
  /// * [int] id (required):
  Future<GetVideoResponse?> getVideo(int id,) async {
    final response = await getVideoWithHttpInfo(id,);
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetVideoResponse',) as GetVideoResponse;
    
    }
    return null;
  }

  /// List latest episodes
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<Response> listLatestEpisodesWithHttpInfo({ int? page, int? pageSize, }) async {
    // ignore: prefer_const_declarations
    final path = r'/shows/latest-episodes';

    // ignore: prefer_final_locals
    Object? postBody;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

    if (page != null) {
      queryParams.addAll(_queryParams('', 'page', page));
    }
    if (pageSize != null) {
      queryParams.addAll(_queryParams('', 'pageSize', pageSize));
    }

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

  /// List latest episodes
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<GetLatestEpisodesResponse?> listLatestEpisodes({ int? page, int? pageSize, }) async {
    final response = await listLatestEpisodesWithHttpInfo( page: page, pageSize: pageSize, );
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetLatestEpisodesResponse',) as GetLatestEpisodesResponse;
    
    }
    return null;
  }

  /// List latest movies
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<Response> listLatestMoviesWithHttpInfo({ int? page, int? pageSize, }) async {
    // ignore: prefer_const_declarations
    final path = r'/shows/latest-movies';

    // ignore: prefer_final_locals
    Object? postBody;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

    if (page != null) {
      queryParams.addAll(_queryParams('', 'page', page));
    }
    if (pageSize != null) {
      queryParams.addAll(_queryParams('', 'pageSize', pageSize));
    }

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

  /// List latest movies
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<GetLatestMoviesResponse?> listLatestMovies({ int? page, int? pageSize, }) async {
    final response = await listLatestMoviesWithHttpInfo( page: page, pageSize: pageSize, );
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetLatestMoviesResponse',) as GetLatestMoviesResponse;
    
    }
    return null;
  }

  /// List latest tv series
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<Response> listLatestTvSeriesWithHttpInfo({ int? page, int? pageSize, }) async {
    // ignore: prefer_const_declarations
    final path = r'/shows/latest-tv-series';

    // ignore: prefer_final_locals
    Object? postBody;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

    if (page != null) {
      queryParams.addAll(_queryParams('', 'page', page));
    }
    if (pageSize != null) {
      queryParams.addAll(_queryParams('', 'pageSize', pageSize));
    }

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

  /// List latest tv series
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<GetLatestTvSeriesResponse?> listLatestTvSeries({ int? page, int? pageSize, }) async {
    final response = await listLatestTvSeriesWithHttpInfo( page: page, pageSize: pageSize, );
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetLatestTvSeriesResponse',) as GetLatestTvSeriesResponse;
    
    }
    return null;
  }

  /// List recently added movies
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<Response> listRecentlyAddedMoviesWithHttpInfo({ int? page, int? pageSize, }) async {
    // ignore: prefer_const_declarations
    final path = r'/shows/recently-added-movies';

    // ignore: prefer_final_locals
    Object? postBody;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

    if (page != null) {
      queryParams.addAll(_queryParams('', 'page', page));
    }
    if (pageSize != null) {
      queryParams.addAll(_queryParams('', 'pageSize', pageSize));
    }

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

  /// List recently added movies
  ///
  /// Parameters:
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<GetLatestMoviesResponse?> listRecentlyAddedMovies({ int? page, int? pageSize, }) async {
    final response = await listRecentlyAddedMoviesWithHttpInfo( page: page, pageSize: pageSize, );
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'GetLatestMoviesResponse',) as GetLatestMoviesResponse;
    
    }
    return null;
  }

  /// Search shows
  ///
  /// Note: This method returns the HTTP [Response].
  ///
  /// Parameters:
  ///
  /// * [String] query (required):
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<Response> searchShowsWithHttpInfo(String query, { int? page, int? pageSize, }) async {
    // ignore: prefer_const_declarations
    final path = r'/shows/search';

    // ignore: prefer_final_locals
    Object? postBody;

    final queryParams = <QueryParam>[];
    final headerParams = <String, String>{};
    final formParams = <String, String>{};

      queryParams.addAll(_queryParams('', 'query', query));
    if (page != null) {
      queryParams.addAll(_queryParams('', 'page', page));
    }
    if (pageSize != null) {
      queryParams.addAll(_queryParams('', 'pageSize', pageSize));
    }

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

  /// Search shows
  ///
  /// Parameters:
  ///
  /// * [String] query (required):
  ///
  /// * [int] page:
  ///
  /// * [int] pageSize:
  Future<SearchShowsResponse?> searchShows(String query, { int? page, int? pageSize, }) async {
    final response = await searchShowsWithHttpInfo(query,  page: page, pageSize: pageSize, );
    if (response.statusCode >= HttpStatus.badRequest) {
      throw ApiException(response.statusCode, await _decodeBodyBytes(response));
    }
    // When a remote server returns no body with a status of 204, we shall not decode it.
    // At the time of writing this, `dart:convert` will throw an "Unexpected end of input"
    // FormatException when trying to decode an empty string.
    if (response.body.isNotEmpty && response.statusCode != HttpStatus.noContent) {
      return await apiClient.deserializeAsync(await _decodeBodyBytes(response), 'SearchShowsResponse',) as SearchShowsResponse;
    
    }
    return null;
  }
}
