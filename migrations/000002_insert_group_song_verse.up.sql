
INSERT INTO music_band(name) VALUES
    ('Queen'),
    ('Metallica'),
    ('Linkin Park'),
    ('Muse'),
    ('Nirvana');


INSERT INTO song(music_band_id, title, release_date, lyrics, link) VALUES
    ((SELECT 1 FROM music_band WHERE name = 'Queen'), 'Bohemian Rhapsody', '1975-10-31', 'Mama, just killed a man...', 'https://www.youtube.com/watch?v=fJ9rUzIMcZQ'),
    ((SELECT 2 FROM music_band WHERE name = 'Metallica'), 'Enter Sandman', '1991-07-29', 'Say your prayers, little one...', 'https://www.youtube.com/watch?v=CD-E-LDc384'),
    ((SELECT 3 FROM music_band WHERE name = 'Linkin Park'), 'Numb', '2003-03-25', 'I`ve become so numb...', 'https://www.youtube.com/watch?v=kXYiU_JCYtU'),
    ((SELECT 4 FROM music_band WHERE name = 'Muse'), 'Supermassive Black Hole', '2006-07-16', 'Ooh baby, don`t you know I suffer?', 'https://www.youtube.com/watch?v=Xsp3_a-PMTw'),
    ((SELECT 5 FROM music_band WHERE name = 'Nirvana'), 'Smells Like Teen Spirit', '1991-09-10', 'Load up on guns, bring your friends...', 'https://www.youtube.com/watch?v=hTWKbfoikeg');


INSERT INTO verse (song_id, content, position)VALUES
    ((SELECT 1 FROM song WHERE title = 'Bohemian Rhapsody'), 'Mama, just killed a man\nPut a gun against his head\nPulled my trigger, now he`s dead...', 1),
    ((SELECT 2 FROM song WHERE title = 'Bohemian Rhapsody'), 'Mama, life had just begun\nBut now I`ve gone and thrown it all away...', 2),
    ((SELECT 3 FROM song WHERE title = 'Enter Sandman'), 'Say your prayers, little one\nDon`t forget, my son\nTo include everyone...', 1),
    ((SELECT 4 FROM song WHERE title = 'Enter Sandman'), 'Sleep with one eye open\nGripping your pillow tight...', 2);
