ALTER TABLE food RENAME TO meal;
ALTER TABLE user_food RENAME TO user_meal;

ALTER TABLE user_meal RENAME COLUMN food_id TO meal_id;
ALTER TABLE user_meal RENAME COLUMN user_food_type TO user_meal_type;

CREATE TYPE user_meal_type AS ENUM('LIKE', 'DISLIKE');

ALTER TABLE user_meal ALTER COLUMN user_meal_type TYPE user_meal_type USING user_meal_type::text::user_meal_type;

DROP TYPE IF EXISTS user_food_type;