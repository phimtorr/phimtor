ALTER TABLE shows
    ADD INDEX idx_shows_updated_at (updated_at);

ALTER TABLE shows
    ADD INDEX idx_shows_type_updated_at (type, updated_at);
