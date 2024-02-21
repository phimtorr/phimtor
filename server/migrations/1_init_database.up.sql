CREATE TABLE IF NOT EXISTS videos
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,

    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

CREATE TABLE IF NOT EXISTS `shows`
(
    id                  BIGINT PRIMARY KEY AUTO_INCREMENT,
    type                enum ('movie', 'series') NOT NULL,
    title               VARCHAR(255)             NOT NULL,
    original_title      VARCHAR(255)             NOT NULL,
    poster_link         VARCHAR(255)             NOT NULL,
    description         TEXT                     NOT NULL,
    release_year        INT                      NOT NULL,
    score               REAL                     NOT NULL DEFAULT 0,
    duration_in_minutes INT                      NOT NULL DEFAULT 0,

    quantity            VARCHAR(50)              NOT NULL DEFAULT '',
    video_id            BIGINT,
    FOREIGN KEY (video_id) REFERENCES videos (id),
    UNIQUE (video_id),

    total_episodes      INT                      NOT NULL DEFAULT 0,
    current_episode     INT                      NOT NULL DEFAULT 0,

    created_at          timestamp                NOT NULL DEFAULT current_timestamp,
    updated_at          timestamp                NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);


CREATE TABLE IF NOT EXISTS episodes
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    show_id    BIGINT       NOT NULL,
    FOREIGN KEY (show_id) REFERENCES shows (id),
    name       VARCHAR(255) NOT NULL,

    video_id   BIGINT       NOT NULL,
    FOREIGN KEY (video_id) REFERENCES videos (id),
    UNIQUE (video_id),

    created_at timestamp    NOT NULL DEFAULT current_timestamp,
    updated_at timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);


CREATE TABLE IF NOT EXISTS torrent_links
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    video_id   BIGINT       NOT NULL,
    FOREIGN KEY (video_id) REFERENCES videos (id),

    name       varchar(50)  NOT NULL DEFAULT '',
    link       VARCHAR(255) NOT NULL DEFAULT '',
    file_index INT          NOT NULL DEFAULT 0,
    priority   INT          NOT NULL DEFAULT 0,

    created_at timestamp    NOT NULL DEFAULT current_timestamp,
    updated_at timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

CREATE TABLE IF NOT EXISTS subtitles
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    video_id   BIGINT       NOT NULL,
    FOREIGN KEY (video_id) REFERENCES videos (id),

    language   VARCHAR(255) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    owner      VARCHAR(255) NOT NULL,
    link       VARCHAR(255) NOT NULL,

    created_at timestamp    NOT NULL DEFAULT current_timestamp,
    updated_at timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);
