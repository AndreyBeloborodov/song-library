-- +goose Up
INSERT INTO songs (group_name, song_name, text)
VALUES
    ('Muse', 'Supermassive Black Hole', 'Ooh baby, don\t you know I suffer?\nOoh baby, can you hear me moan?'),
    ('Coldplay', 'Fix You', 'When you try your best but you don\t succeed\nWhen you get what you want but not what you need'),
    ('Radiohead', 'Creep', 'When you were here before\nCouldn\t look you in the eye');

-- +goose Down
DELETE FROM songs WHERE song_name IN ('Supermassive Black Hole', 'Fix You', 'Creep');
