DROP TYPE IF EXISTS show_type;
CREATE TYPE show_type AS ENUM ('movie', 'series');

CREATE TABLE IF NOT EXISTS shows
(
    id                  BIGSERIAL PRIMARY KEY,
    type                show_type    NOT NULL,
    title               VARCHAR(255) NOT NULL,
    original_title      VARCHAR(255) NOT NULL,
    poster_link         VARCHAR(255) NOT NULL,
    description         TEXT         NOT NULL,
    release_year        INT          NOT NULL,
    score               REAL         NOT NULL DEFAULT 0,
    duration_in_minutes INT          NOT NULL DEFAULT 0,

    quantity            VARCHAR(50)  NOT NULL DEFAULT '',

    total_episodes      INT          NOT NULL DEFAULT 0,
    current_episode     INT          NOT NULL DEFAULT 0,

    created_at          timestamptz  NOT NULL DEFAULT current_timestamp,
    updated_at          timestamptz  NOT NULL DEFAULT current_timestamp
);


CREATE TABLE IF NOT EXISTS episodes
(
    id         BIGSERIAL PRIMARY KEY,
    show_id    INT          NOT NULL,
    name       VARCHAR(255) NOT NULL,

    created_at timestamptz  NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz  NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS torrent_links
(
    id          BIGSERIAL PRIMARY KEY,

    quantity    varchar(50)  NOT NULL DEFAULT '',
    link        VARCHAR(255) NOT NULL DEFAULT '',
    magnet_link TEXT         NOT NULL DEFAULT '',

    show_id     INT          NULL,
    foreign key (show_id) references shows (id),
    episode_id  INT          NULL,
    foreign key (episode_id) references episodes (id),

    created_at  timestamptz  NOT NULL DEFAULT current_timestamp,
    updated_at  timestamptz  NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS subtitles
(
    id         BIGSERIAL PRIMARY KEY,
    language   VARCHAR(255) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    owner      VARCHAR(255) NOT NULL,
    link       VARCHAR(255) NOT NULL,

    show_id    INT          NULL,
    foreign key (show_id) references shows (id),
    episode_id INT          NULL,
    foreign key (episode_id) references episodes (id),

    created_at timestamptz  NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz  NOT NULL DEFAULT current_timestamp
);
