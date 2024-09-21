ALTER TABLE latest_shows
    ADD COLUMN video_id BIGINT NOT NULL DEFAULT 0 AFTER episode_number,
    ADD COLUMN has_vi_sub BOOLEAN NOT NULL DEFAULT FALSE AFTER quality;

ALTER TABLE torrent_links
    ADD COLUMN resolution INT NOT NULL AFTER video_id,
    ADD COLUMN type VARCHAR(20) NOT NULL AFTER resolution,
    ADD COLUMN source VARCHAR(20) NOT NULL AFTER type;

ALTER TABLE videos
    ADD COLUMN max_resolution INT NOT NULL AFTER id,
    ADD COLUMN has_vi_sub BOOLEAN NOT NULL AFTER max_resolution,
    ADD COLUMN has_en_sub BOOLEAN NOT NULL AFTER has_vi_sub;

ALTER TABLE torrent_links
    ADD COLUMN codec VARCHAR(20) NOT NULL AFTER type;
