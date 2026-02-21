CREATE TABLE IF NOT EXISTS blood_bressure_record
(
    id serial primary key,
    user_name varchar(64),
    systolic integer,
    diastolic integer,
    pulse integer,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS tag_record
(
    id serial primary key,
    record_id integer unique not null references blood_bressure_record(id) on delete cascade,
    tag_name varchar(64)
);

CREATE TABLE IF NOT EXISTS telegram_user
(
    id serial primary key,
    chat_id integer,
    created_at timestamp
);