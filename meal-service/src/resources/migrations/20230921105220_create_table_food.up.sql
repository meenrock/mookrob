CREATE TABLE IF NOT EXISTS food
(
    id              UUID                NOT NULL DEFAULT uuid_generate_v4 (),
    name            VARCHAR(100)        NOT NULL,
    energy          DECIMAL             NOT NULL,
    protein         DECIMAL             NULL,
    carbohydrate    DECIMAL             NULL,
    fat             DECIMAL             NULL,
    sodium          DECIMAL             NULL,
    cholesterol     DECIMAL             NULL,
    created_at      TIMESTAMP           NOT NULL,
    updated_at      TIMESTAMP           NOT NULL,
    PRIMARY KEY (id)
);