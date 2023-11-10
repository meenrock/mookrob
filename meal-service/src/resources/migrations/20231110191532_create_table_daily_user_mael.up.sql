CREATE TYPE meal_time AS ENUM('BREAKFAST', 'LUNCH', 'DINNER');

CREATE TABLE IF NOT EXISTS daily_user_meal
(
    id              UUID                NOT NULL DEFAULT uuid_generate_v4 (),
    meal_id         UUID                NOT NULL,
    user_id         UUID                NOT NULL,
    meal_time       meal_time           NOT NULL,
    date            DATE                NOT NULL,
    created_at      TIMESTAMP           NOT NULL,
    updated_at      TIMESTAMP           NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (meal_time, date)
);