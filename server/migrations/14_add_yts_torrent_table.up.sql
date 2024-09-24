ALTER TABLE videos
    ADD COLUMN yts_movie_id BIGINT AFTER id,
    ADD UNIQUE INDEX yts_movie_id_UNIQUE (yts_movie_id);

CREATE TABLE yts_torrents
(
    hash           CHAR(40) PRIMARY KEY,
    movie_id       BIGINT      NOT NULL,
    quality        VARCHAR(255) NOT NULL,
    resolution     INT          NOT NULL DEFAULT 0,
    type           VARCHAR(20)  NOT NULL,
    is_repack      BOOLEAN      NOT NULL DEFAULT FALSE,
    video_codec    VARCHAR(20)  NOT NULL,
    bit_depth      INT          NOT NULL DEFAULT 8,
    audio_channels VARCHAR(20)  NOT NULL,
    seeds          INTEGER      NOT NULL,
    peers          INTEGER      NOT NULL,
    size_bytes     BIGINT       NOT NULL,
    date_uploaded  TIMESTAMP    NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (movie_id) REFERENCES videos (yts_movie_id),
    INDEX movie_id_INDEX (movie_id)
);