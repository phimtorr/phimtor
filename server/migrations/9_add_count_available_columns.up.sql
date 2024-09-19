RENAME TABLE tv_series_show TO tv_series_shows;

ALTER TABLE tv_seasons ADD COLUMN count_available_episodes INT NOT NULL DEFAULT 0 AFTER total_episodes;