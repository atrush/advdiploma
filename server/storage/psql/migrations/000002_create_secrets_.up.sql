CREATE TABLE IF NOT EXISTS secrets
(
    id uuid primary key default uuid_generate_v4(),
    device_id      uuid NOT NULL,
    user_id      uuid NOT NULL,
    "data"     TEXT,
    created_at timestamp    not null default now(),
    deleted_at timestamp    not null default now(),
    is_deleted boolean not null default false,

    FOREIGN KEY (user_id) REFERENCES users (id)
);