-- +goose Up
-- SQL in this section is executed when the migration is applied.

DROP TABLE IF EXISTS srv_events;

CREATE TABLE srv_events (
    uid BIGINT NOT NULL,
    type INT NOT NULL,
    data JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    rev BIGINT NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS srv_events;
