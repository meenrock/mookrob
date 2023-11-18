ALTER TABLE daily_user_meal
DROP CONSTRAINT daily_user_meal_meal_time_date_key;

ALTER TABLE daily_user_meal
ADD CONSTRAINT daily_user_meal_user_id_meal_time_date_key
UNIQUE (user_id, meal_time, date);