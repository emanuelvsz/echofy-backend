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
where ag.super_artist_id = 'a6d6488b-6ed0-4d41-9c55-4e899fdd7e47';