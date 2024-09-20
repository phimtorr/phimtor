ALTER TABLE latest_shows
    ADD INDEX idx_type_air_date (type, air_date);

ALTER TABLE latest_shows
    ADD INDEX idx_type_created_at (type, created_at);