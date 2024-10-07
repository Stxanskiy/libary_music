-- Добавление музыкальных групп
INSERT INTO music_band (name) VALUES
    ('Muse'),
    ('Queen'),
    ('The Beatles');

-- Добавление песен
INSERT INTO song (music_band_id, title, release_date, lyrics, link) VALUES
    (1, 'Supermassive Black Hole', '2006-07-16', E'Ooh baby, don\'t you know I suffer?\\nOoh baby, can you hear me moan?', 'https://www.youtube.com/watch?v=Xsp3_a-PMTw'),
    (2, 'Bohemian Rhapsody', '1975-10-31', E'Is this the real life? Is this just fantasy?\\nCaught in a landslide, no escape from reality', 'https://www.youtube.com/watch?v=fJ9rUzIMcZQ'),
    (3, 'Hey Jude', '1968-08-26', E'Hey Jude, don\'t make it bad\\nTake a sad song and make it better', 'https://www.youtube.com/watch?v=A_MjCqQoLLA');

-- Добавление куплетов для каждой песни
INSERT INTO verse (song_id, content, position) VALUES
    (1, E'Ooh baby, don\'t you know I suffer?', 1),
    (1, E'Ooh baby, can you hear me moan?', 2),
    (2, E'Is this the real life?', 1),
    (2, E'Is this just fantasy?', 2),
    (3, E'Hey Jude, don\'t make it bad', 1),
    (3, E'Take a sad song and make it better', 2);
