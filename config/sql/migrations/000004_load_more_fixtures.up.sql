
copy genre (id, name, description, created_at)
    from '/sql/sql/migrations/fixtures/000002-genre.csv'
    delimiter ';' csv header;
    
copy song_artist (song_id, artist_id)
    from '/sql/sql/migrations/fixtures/000002-artist_song.csv'
    delimiter ';' csv header;

copy artist_genre (genre_id, artist_id)
    from '/sql/sql/migrations/fixtures/000002-artist_genre.csv'
    delimiter ';' csv header;