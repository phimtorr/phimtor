import 'package:flutter/material.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/shows/shows_list.dart';
import 'package:phimtor_openapi_client/api.dart';

class HomeView extends StatelessWidget {
  const HomeView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Phim Tor"),
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text(
                "Movies",
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 16),
              SizedBox(
                height: ShowsList.minHeight,
                child: FutureBuilder(
                  future: PhimtorService().defaultApi.listShows(page: 1, pageSize: 10, type: ShowType.movie), 
                  builder: (context, snapshot) {
                    if (snapshot.connectionState == ConnectionState.waiting) {
                      return const Center(child: CircularProgressIndicator());
                    }
                    if (snapshot.hasError) {
                      return Center(child: Text("Error: ${snapshot.error}"));
                    }

                    final response = snapshot.data as ListShowsResponse;
                    return ShowsList(shows: response.shows);
                  },),
              ),
              const Text(
                "TV Series",
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 16),
               SizedBox(
                height: ShowsList.minHeight,
                child: FutureBuilder(
                  future: PhimtorService().defaultApi.listShows(page: 1, pageSize: 10, type: ShowType.series), 
                  builder: (context, snapshot) {
                    if (snapshot.connectionState == ConnectionState.waiting) {
                      return const Center(child: CircularProgressIndicator());
                    }
                    if (snapshot.hasError) {
                      return Center(child: Text("Error: ${snapshot.error}"));
                    }

                    final response = snapshot.data as ListShowsResponse;
                    return ShowsList(shows: response.shows);
                  },),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
