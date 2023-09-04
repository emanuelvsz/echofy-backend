copy artist_group (artist_id, super_artist_id, joined_at, left_at, is_active)
    from '/sql/sql/migrations/fixtures/000006-artist_group.csv'
    delimiter ';' csv header;

