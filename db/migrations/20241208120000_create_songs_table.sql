-- +goose Up
CREATE TABLE IF NOT EXISTS songs (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    group_name VARCHAR(100) NOT NULL,
    song_name VARCHAR(100) NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    release_date VARCHAR(20) DEFAULT '',
    link TEXT DEFAULT ''
    );

-- +goose Down
DROP TABLE IF EXISTS songs;
