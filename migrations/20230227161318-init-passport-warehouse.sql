-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE goods
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);

-- +migrate Down
DROP TABLE goods;
