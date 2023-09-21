CREATE TYPE status AS ENUM('ACTIVE', 'INACTIVE');

CREATE TABLE IF NOT EXISTS users
(
    id                          UUID                NOT NULL DEFAULT uuid_generate_v4 (),
    status                      status              NOT NULL,
    first_name                  VARCHAR(100)        NOT NULL,
    last_name                   VARCHAR(100)        NOT NULL,
    nick_name                   VARCHAR(100)        NOT NULL,
    phone_number                VARCHAR(10)         NULL,
    email                       VARCHAR(100)        NOT NULL,
    gender                      VARCHAR(20)         NOT NULL,
    age                         INT                 NOT NULL,
    height                      DECIMAL             NOT NULL,
    weight                      DECIMAL             NOT NULL,
    expected_bmi                DECIMAL             NULL,
    created_at                  TIMESTAMP           NOT NULL,
    updated_at                  TIMESTAMP           NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (email)
);