create extension if not exists "uuid-ossp";

create table if not exists role
(
    id      uuid            not null
        constraint pk_role_id primary key,
    name    varchar         not null,
    code    varchar         not null
);

create table if not exists country (
    id uuid not null,
    name varchar(128) not null,
    constraint pk_country primary key (id)
);

create table if not exists state (
    id uuid not null,
    name varchar(128) not null,
    country_id uuid,
    constraint pk_state primary key (id),
    constraint fk_state_country_id foreign key (country_id) references country (id)
);

create table if not exists city (
    id uuid not null,
    name varchar(128) not null,
    state_id uuid,
    constraint pk_city primary key (id),
    constraint fk_city_state_id foreign key (state_id) references state (id)
);

create table if not exists record_company (
    id uuid not null,
    name varchar(128) not null,
    founded_at date not null,
    website_url varchar(256),
    country_id uuid,
    constraint pk_record_company primary key (id),
    constraint fk_record_company_country_id foreign key (country_id) references country (id)
);

create table if not exists account
(
    id          uuid        not null
        constraint pk_account_id primary key
        constraint df_account_id default uuid_generate_v4(),
    role_id     uuid        not null,
    name        varchar(64) not null,
    email       varchar(64) not null,
    hash        varchar(64) not null,
    password    text        not null,
    last_login  timestamp   not null,
    picture_url text        null,
    
    constraint fk_account_role_id
        foreign key (role_id) references role (id)
);

create table if not exists artist 
(
    id                  uuid        not null
        constraint pk_artist_id primary key
        constraint df_artist_id default uuid_generate_v4(),
    super_artist_id uuid        null,
    name                varchar(64) not null,
    description         text        null,
    founded_at          date        not null,
    terminated_at       date        null,
    image_url           text        null,
    record_company_id   uuid      null
        constraint fk_artist_record_company_id
            references record_company (id),
    country_id          uuid        null
        constraint fk_artist_country_id
            references country (id),
    spotify_url         text        null,

    constraint fk_artist_super_artist_id
        foreign key (super_artist_id) references artist (id)
);

create table if not exists album
(
    id              uuid        not null
        constraint pk_album_id primary key
        constraint df_album_id default uuid_generate_v4(),
    name            varchar(64) not null,
    artist_id       uuid        not null,
    release_date    date        not null,
    description     text        null,
    image_url       text        null,

    constraint fk_album_artist_id
        foreign key (artist_id) references artist (id)
);

create table if not exists song
(
    id              uuid        not null
        constraint pk_song_id primary key
        constraint df_song_id default uuid_generate_v4(),
    name            varchar(64) not null,
    album_id        uuid        null,
    release_date    date        not null,
    duration        text        not null,
    lyrics          text        null,
    track_number    integer     null,
    spotify_url     text        null,

    constraint fk_song_album_id
        foreign key (album_id) references album (id)
);

create table if not exists song_artist
(
    artist_id   uuid    not null,
    song_id     uuid    not null,

    constraint pk_song_artist 
        primary key (artist_id, song_id),
    constraint fk_song_artist_artist_id 
        foreign key (artist_id) references artist (id),
    constraint fk_song_artist_song_id 
        foreign key (song_id) references song (id)
);

create table if not exists followed_artist
(
    id uuid not null
        constraint pk_followed_artist_id primary key
        constraint df_followed_artist_id default uuid_generate_v4(),
    artist_id       uuid    not null,
    account_id      uuid    not null,
    followed_at     date    not null,
    unfollowed_at   date    null,

    constraint fk_followed_artist_artist_id
        foreign key (artist_id) references artist (id),
    constraint fk_followed_artist_account_id
        foreign key (account_id) references account (id)
);

create table if not exists genre
(
    id              uuid        not null
        constraint pk_genre_id primary key
        constraint df_genre_id default uuid_generate_v4(),
    name            varchar     not null,
    description     text        null,
    created_at      timestamp   null
);

create table if not exists artist_genre
(
    id              uuid        not null
            constraint pk_artist_genre_id primary key
            constraint df_artist_genre_id default uuid_generate_v4(),
    artist_id       uuid        not null,
    genre_id        uuid        not null,

    constraint fk_artist_genre_artist_id
        foreign key (artist_id) references artist (id),
    constraint fk_artist_genre_genre_id
        foreign key (genre_id) references genre (id)
);

create table if not exists song_genre
(
    id              uuid        not null
            constraint pk_song_genre_id primary key
            constraint df_song_genre_id default uuid_generate_v4(),
    song_id         uuid        not null,
    genre_id        uuid        not null,

    constraint fk_song_genre_song_id
        foreign key (song_id) references song (id),
    constraint fk_song_genre_genre_id
        foreign key (genre_id) references genre (id)
);

create table if not exists album_genre
(
    id              uuid        not null
            constraint pk_album_genre_id primary key
            constraint df_album_genre_id default uuid_generate_v4(),
    album_id        uuid        not null,
    genre_id        uuid        not null,

    constraint fk_album_genre_album_id
        foreign key (album_id) references album (id),
    constraint fk_album_genre_genre_id
        foreign key (genre_id) references genre (id)
);

create table if not exists logged_user (
    id uuid not null
        constraint pk_logged_user_id primary key
        constraint df_logged_user_id default uuid_generate_v4(),
    token_id uuid not null
        constraint fk_logged_user_token_id 
        references token (id),
    name varchar(32)
);

create table if not exists token (
    refresh_token text not null,
    access_token text not null

    constraint pk_token (refresh_token, access_token)
);
