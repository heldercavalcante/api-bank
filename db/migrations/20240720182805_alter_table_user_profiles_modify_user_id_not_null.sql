-- migrate:up
ALTER TABLE user_profiles MODIFY user_id INT NOT NULL;


-- migrate:down
ALTER TABLE user_profiles MODIFY user_id INT;
