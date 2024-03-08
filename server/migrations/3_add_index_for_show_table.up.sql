ALTER TABLE shows
ADD INDEX idx_shows_created_at (created_at);

ALTER TABLE shows
ADD INDEX idx_shows_type_created_at (type, created_at);
