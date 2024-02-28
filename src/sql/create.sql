create table if not exists company_posts
(
    id serial primary key,
    post text,
    is_admin boolean
);

create table if not exists document_types
(
    id serial primary key,
    type text
);

create table if not exists document_field_types
(
    id serial primary key,
    type text
);

create table if not exists company
(
    id serial primary key,
    name text,
    city text
);

create table if not exists employee
(
    id serial primary key,
    phone_number text,
    full_name text,
    company_id int references company(id) on delete cascade,
    post int references company_posts(id) on delete cascade,
    date_of_birth date
);

create table if not exists credentials
(
    id serial primary key,
    employee_id int references employee(id) on delete cascade unique,
    password text,
    created_date date default now()
);

create table if not exists info_card
(
    id serial primary key,
    created_employee_id int references employee(id) on delete cascade,
    is_confirmed boolean,
    created_date date default now()
);

create table if not exists document
(
    id serial primary key,
    info_card_id int references info_card(id),
    type int references document_types(id) on delete cascade
);

create table if not exists photo
(
    id serial primary key,
    document_id int references document(id) on delete cascade,
    key text
);

create table if not exists field
(
    id serial primary key,
    document_id int references document(id) on delete cascade,
    type int references document_field_types(id) on delete cascade,
    value text
);
