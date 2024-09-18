CREATE TABLE movies
(
    id             BIGINT PRIMARY KEY,
    imdb_id        VARCHAR(255) NOT NULL,
    title          VARCHAR(255) NOT NULL,
    original_title VARCHAR(255) NOT NULL,
    status         VARCHAR(255) NOT NULL,
    tagline        VARCHAR(255) NOT NULL,
    genres         JSON         NOT NULL,
    overview       TEXT         NOT NULL,
    poster_path    VARCHAR(255) NOT NULL,
    backdrop_path  VARCHAR(255) NOT NULL,
    release_date   DATE         NOT NULL,
    runtime        INT          NOT NULL,
    vote_average   FLOAT        NOT NULL,
    vote_count     INT          NOT NULL,
    video_id       BIGINT       NOT NULL DEFAULT 0,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE tv_series_show
(
    id                 BIGINT PRIMARY KEY,
    name               VARCHAR(255) NOT NULL,
    original_name      VARCHAR(255) NOT NULL,
    status             VARCHAR(255) NOT NULL,
    tagline            VARCHAR(255) NOT NULL,
    genres             JSON         NOT NULL,
    overview           TEXT         NOT NULL,
    poster_path        VARCHAR(255) NOT NULL,
    backdrop_path      VARCHAR(255) NOT NULL,
    first_air_date     DATE         NULL,
    last_air_date      DATE         NULL,
    vote_average       FLOAT        NOT NULL,
    vote_count         INT          NOT NULL,
    number_of_episodes INT          NOT NULL,
    number_of_seasons  INT          NOT NULL,
    created_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE tv_seasons
(
    id             BIGINT PRIMARY KEY,
    show_id        BIGINT       NOT NULL,
    season_number  INT          NOT NULL,
    name           VARCHAR(255) NOT NULL,
    poster_path    VARCHAR(255) NOT NULL,
    overview       TEXT         NOT NULL,
    air_date       DATE         NULL,
    vote_average   FLOAT        NOT NULL,
    total_episodes INT          NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY idx_show_id_season_number (show_id, season_number),
    INDEX idx_show_id (show_id),
    FOREIGN KEY (show_id) REFERENCES tv_series_show (id)
);

CREATE TABLE tv_episodes
(
    id             BIGINT PRIMARY KEY,
    show_id        BIGINT       NOT NULL,
    season_number  INT          NOT NULL,
    episode_number INT          NOT NULL,
    name           VARCHAR(255) NOT NULL,
    overview       TEXT         NOT NULL,
    air_date       DATE         NULL,
    runtime        INT          NOT NULL,
    still_path     VARCHAR(255) NOT NULL,
    vote_average   FLOAT        NOT NULL,
    vote_count     INT          NOT NULL,
    video_id       BIGINT       NOT NULL DEFAULT 0,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY idx_show_id_season_number_episode_number (show_id, season_number, episode_number),
    INDEX idx_show_id_season_number (show_id, season_number),
    INDEX idx_show_id (show_id),
    FOREIGN KEY (show_id, season_number) REFERENCES tv_seasons (show_id, season_number)
);

CREATE TABLE latest_shows
(
    id             BIGINT PRIMARY KEY AUTO_INCREMENT,
    type           enum ('movie', 'tv-series', 'episode') NOT NULL,
    show_id        BIGINT                                 NOT NULL, -- it can be movie_id or tv_series_id
    title          VARCHAR(255)                           NOT NULL,
    original_title VARCHAR(255)                           NOT NULL,
    poster_path    VARCHAR(255)                           NOT NULL,
    air_date       DATE                                   NULL,
    runtime        INT                                    NULL,
    vote_average   FLOAT                                  NOT NULL,
    quality        VARCHAR(255)                           NOT NULL,
    season_number  INT,                                             -- in case of episode
    episode_number INT,                                             -- in case of episode
    created_at     TIMESTAMP                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP                              NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FULLTEXT INDEX search_index (title, original_title),
    UNIQUE KEY idx_type_show_id (type, show_id)
);

