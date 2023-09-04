create table if not exists reset_password (
    account_id uuid not null,
    email varchar(64) not null,
    code varchar(16) not null,
    constraint pk_reset_password primary key (account_id),
    constraint fk_reset_password_account_id foreign key (account_id) references account (id)
);

create table if not exists category (
    id uuid not null,
    name varchar(64) not null,
    description text,
    constraint pk_category primary key (id)
);

create table if not exists liked_song (
    song_id uuid not null,
    account_id uuid not null,
    liked_at timestamp not null,
    unliked_at timestamp,
    constraint pk_liked_song primary key (song_id, account_id),
    constraint fk_liked_song_song_id foreign key (song_id) references song (id),
    constraint fk_liked_song_account_id foreign key (account_id) references account (id)
);

create table if not exists news (
    id uuid not null,
    title varchar(128) not null,
    released_at date not null,
    description text,
    updated_at timestamp,
    created_at timestamp not null default current_timestamp,
    constraint pk_news primary key (id)
);

create table if not exists artist_news (
    news_id uuid not null,
    artist_id uuid not null,
    constraint pk_artist_news primary key (news_id, artist_id),
    constraint fk_artist_news_news_id foreign key (news_id) references news (id),
    constraint fk_artist_news_artist_id foreign key (artist_id) references artist (id)
);

create table if not exists news_category (
    news_id uuid not null,
    category_id uuid not null,
    constraint pk_news_category primary key (news_id, category_id),
    constraint fk_news_category_news_id foreign key (news_id) references news (id),
    constraint fk_news_category_category_id foreign key (category_id) references category (id)
);

create table if not exists company_country (
    country_id uuid,
    record_company_id uuid,
    constraint pk_company_country primary key (country_id, record_company_id),
    constraint fk_company_country_country_id foreign key (country_id) references country (id),
    constraint fk_company_country_record_company_id foreign key (record_company_id) references record_company (id)
);

