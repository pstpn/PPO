-- CREATE DATABASE tests
--     ENCODING 'UTF-8'
--     LC_COLLATE 'ru_RU.UTF-8'
--     LC_CTYPE 'ru_RU.UTF-8'
--     TEMPLATE template0;

create table if not exists company
(
    id text primary key,
    name text,
    city text
);

create table if not exists employee
(
    id text primary key,
    phone_number text unique,
    full_name text,
    company_id text references company(id) on delete cascade,
    post text,
    password text,
    refresh_token text,
    token_expired_at timestamp,
    date_of_birth date
);

create table if not exists info_card
(
    id text primary key,
    created_employee_id text references employee(id) on delete cascade,
    is_confirmed boolean,
    created_date date default now()
);

create table if not exists document
(
    id text primary key,
    serial_number text,
    info_card_id text references info_card(id) on delete cascade,
    type text
);

create table if not exists photo
(
    id text primary key,
    document_id text references document(id) on delete cascade,
    key text
);

create table if not exists field
(
    id text primary key,
    document_id text references document(id) on delete cascade,
    type text,
    value text,

    unique (document_id, type)
);

create table if not exists checkpoint
(
    id text primary key,
    phone_number text
);

create table if not exists passage
(
    id text primary key,
    checkpoint_id text references checkpoint(id) on delete cascade,
    document_id text references document(id) on delete cascade,
    type text,
    time timestamp
);