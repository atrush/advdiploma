CREATE TABLE IF NOT EXISTS secrets
(
    id           serial primary key,
    user_id      uuid NOT NULL,
    client_id   int,
    "data"     TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id)
);