import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/routes/route_names.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart'
    as phimtor_service;
import 'package:phimtor_app/views/search_section.dart';
import 'package:phimtor_app/views/shows/shows_list.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class HomeView extends StatelessWidget {
  const HomeView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.title),
      ),
      body: const SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              SearchSection(),
              SizedBox(height: 32),
              MoviesSection(),
              SizedBox(height: 32),
              TVSeriesSection(),
              SizedBox(height: 32),
              TVEpisodesSection(),
            ],
          ),
        ),
      ),
    );
  }
}

class MoviesSection extends StatelessWidget {
  const MoviesSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Text(
              context.loc.latest_movies,
              style: Theme.of(context).textTheme.headlineLarge,
            ),
            const SizedBox(width: 8),
            ElevatedButton.icon(
              onPressed: () {
                context.goNamed(routeNameMovies);
              },
              label: Text(context.loc.load_more),
              icon: const Icon(Icons.arrow_forward),
              iconAlignment: IconAlignment.end,
            ),
          ],
        ),
        const SizedBox(height: 16),
        SizedBox(
          height: 320.0,
          child: FutureBuilder(
            future: phimtor_service.PhimtorService()
                .defaultApi
                .getLatestMovies(page: 1, pageSize: 10),
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const Center(child: CircularProgressIndicator());
              }
              if (snapshot.hasError) {
                return Center(
                    child: Text(context.loc.error(snapshot.error.toString())));
              }

              final response = snapshot.data as phimtor_api.GetLatestMoviesResponse;
              return ShowsList(shows: response.movies);
            },
          ),
        ),
      ],
    );
  }
}

class TVSeriesSection extends StatelessWidget {
  const TVSeriesSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          children: [
            Text(
              context.loc.latest_tv_series,
              style: Theme.of(context).textTheme.headlineLarge,
            ),
            const SizedBox(width: 8),
            ElevatedButton.icon(
              onPressed: () {
                context.goNamed(routeNameTVSeries);
              },
              label: Text(context.loc.load_more),
              icon: const Icon(Icons.arrow_forward),
              iconAlignment: IconAlignment.end,
            ),
          ],
        ),
        const SizedBox(height: 16),
        SizedBox(
          height: 320.0,
          child: FutureBuilder(
            future: phimtor_service.PhimtorService()
                .defaultApi
                .getLatestTvSeries(page: 1, pageSize: 10,),
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const Center(child: CircularProgressIndicator());
              }
              if (snapshot.hasError) {
                return Center(
                    child: Text(context.loc.error(snapshot.error.toString())));
              }

              final response = snapshot.data as phimtor_api.GetLatestTvSeriesResponse;
              return ShowsList(shows: response.tvSeries);
            },
          ),
        ),
      ],
    );
  }

}

class TVEpisodesSection extends StatelessWidget {
  const TVEpisodesSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          children: [
            Text(
              context.loc.latest_episodes,
              style: Theme.of(context).textTheme.headlineLarge,
            ),
            const SizedBox(width: 8),
            ElevatedButton.icon(
              onPressed: () {
                context.goNamed(routeNameTVSeries);
              },
              label: Text(context.loc.load_more),
              icon: const Icon(Icons.arrow_forward),
              iconAlignment: IconAlignment.end,
            ),
          ],
        ),
        const SizedBox(height: 16),
        SizedBox(
          height: 320.0,
          child: FutureBuilder(
            future: phimtor_service.PhimtorService()
                .defaultApi
                .getLatestEpisodes(page: 1, pageSize: 10,),
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const Center(child: CircularProgressIndicator());
              }
              if (snapshot.hasError) {
                return Center(
                    child: Text(context.loc.error(snapshot.error.toString())));
              }

              final response = snapshot.data as phimtor_api.GetLatestEpisodesResponse;
              return ShowsList(shows: response.episodes);
            },
          ),
        ),
      ],
    );
  }
}