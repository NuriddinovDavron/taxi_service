create table if not exists taxis(
    id uuid primary key,
    first_name text not null,
    last_name text not null,
    email text unique not null,
    password text not null,
    birthday timestamp not null,
    car_id uuid,
    phone_number text unique not null,
    gender text not null,
    profile_photo text not null,
    refresh_token text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table if not exists cars(
    id uuid primary key,
    model text not null,
    image_url text not null,
    colour text not null,
    number_passenger integer not null,
    number_bags integer not null,
    number text not null,
    free_days text not null,
    from_location text not null,
    to_location text not null,
    price integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table if not exists review(
    car_id uuid references taxis(id),
    comment text,
    stars integer,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);