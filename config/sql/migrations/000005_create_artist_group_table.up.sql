create extension if not exists "uuid-ossp";

create table artist_group (
  id uuid not null
    constraint pk_artist_group_id primary key
    constraint df_artist_group_id default uuid_generate_v4(),
  artist_id uuid not null,
  super_artist_id uuid not null,
  joined_at timestamp not null,
  left_at timestamp,
  is_active boolean not null,
  foreign key (artist_id) references artist (id),
  foreign key (super_artist_id) references artist (id),

  constraint unique_artist_group unique (artist_id, super_artist_id),
  constraint valid_dates check (joined_at <= coalesce(left_at, current_timestamp)),
  constraint valid_active_status check (is_active = true or left_at is not null)
);

create or replace function check_super_artist()
returns trigger as $$
begin
  if new.artist_id = new.super_artist_id then
    raise exception 'Um artista não pode ser seu próprio super artista';
  end if;
  return new;
end;
$$ language plpgsql;

create trigger prevent_self_super_artist
before insert or update on artist_group
for each row
execute function check_super_artist();