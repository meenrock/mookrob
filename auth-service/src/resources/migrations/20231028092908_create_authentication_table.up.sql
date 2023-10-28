CREATE TYPE status AS ENUM('ACTIVE', 'INACTIVE');
CREATE TYPE role AS ENUM('GENERAL_USER', 'ADMIN');

CREATE TABLE IF NOT EXISTS authentication_data
(
    id                  UUID                NOT NULL DEFAULT uuid_generate_v4 (),
    username            VARCHAR(100)        NOT NULL,
    password            VARCHAR(255)        NOT NULL,
    user_id             UUID                NULL,
    role                role                NOT NULL,
    status              status              NOT NULL,
    created_at          TIMESTAMP           NOT NULL,
    updated_at          TIMESTAMP           NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (username),
    UNIQUE (user_id)
);