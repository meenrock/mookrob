CREATE TYPE user_food_type AS ENUM('LIKE', 'DISLIKE');

CREATE TABLE IF NOT EXISTS user_food
(
    food_id             UUID                NOT NULL,
    user_id             UUID                NOT NULL,
    user_food_type      user_food_type      NOT NULL,
    PRIMARY KEY (food_id, user_id)
);