CREATE TABLE IF NOT EXISTS reservations
(
    id         UUID PRIMARY KEY,
    room_id    VARCHAR(255) NOT NULL,
    start_time TIMESTAMP    NOT NULL,
    end_time   TIMESTAMP    NOT NULL
);
