-- name: SelectAlbumSongs :many
select s.id as id,
    s.name as name,
    s.album_id as album_id,
    s.release_date as release_date,
    s.duration as duration,
    s.lyrics as lyrics,
    s.track_number as track_number,
    s.spotify_url as spotify_url
    from song s where s.album_id = @album_id
    order by s.track_number;

-- name: SelectAlbums :many
select a.id as id,
    a.name as name,
    a.artist_id as artist_id,
    a.release_date as release_date,
    a.description as description,
    a.image_url as image_url
    from album a;

-- name: SelectArtists :many
select a.id as id,
    a.name as name,
    a.super_artist_id as super_artist_id,
    a.description as description,
    a.founded_at as founded_at,
    a.terminated_at as terminated_at 
        from artist a
    order by a.name;

-- name: SelectArtistSongs :many
select s.id as id,
    s.name as name,
    s.album_id as album_id,
    s.release_date as release_date,
    s.duration as duration,
    sa.artist_id as artist_id
        from song s
    inner join song_artist sa on sa.song_id = s.id
    where sa.artist_id = @artist_id;

-- name: SelectArtistByID :one
select a.id as id,
    a.name as name,
    a.super_artist_id as super_artist_id,
    a.description as description,
    a.founded_at as founded_at,
    a.terminated_at as terminated_at,
    a.image_url as image_url,
    a.record_company_id as record_company_id,
    a.country_id as country_id,
    a.spotify_url as spotify_url,
    g.name as genre_name,
    g.description as genre_description,
    g.created_at as genre_created_at
        from artist a
    inner join artist_genre ag on a.id = ag.artist_id
    inner join genre g on g.id = ag.genre_id
    where a.id = @artist_id;

-- name: SelectSubArtists :many
select a.id as id,
    a.name as name,
    a.description as description,
    a.founded_at as founded_at,
    a.terminated_at as terminated_at,
    a.image_url as image_url,
    a.record_company_id as record_company_id,
    a.country_id as country_id,
    a.spotify_url as spotify_url
from artist as a
join artist_group as ag on a.id = ag.artist_id
join artist as g on g.id = ag.super_artist_id
where ag.super_artist_id = @super_artist_id;

