ALTER TABLE torrent_links
    ADD COLUMN required_premium BOOLEAN NOT NULL DEFAULT FALSE AFTER priority;