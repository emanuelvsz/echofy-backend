copy country (id, name)
    from '/sql/sql/migrations/fixtures/000004-country.csv'
    delimiter ';' csv header;

copy record_company (id, name, founded_at, website_url, country_id)
    from '/sql/sql/migrations/fixtures/000004-record_company.csv'
    delimiter ';' csv header;

copy artist (id, name, super_artist_id, description, founded_at, terminated_at, image_url, record_company_id, country_id, spotify_url)
    from '/sql/sql/migrations/fixtures/000002-artist.csv'
    delimiter ';' csv header;

copy album (id, name, artist_id, release_date, description, image_url)
    from '/sql/sql/migrations/fixtures/000002-album.csv'
    delimiter ';' csv header;

copy song (id, name, album_id, release_date, duration)
    from '/sql/sql/migrations/fixtures/000002-song.csv'
    delimiter ';' csv header;

