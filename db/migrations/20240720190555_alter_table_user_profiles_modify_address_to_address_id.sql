-- migrate:up
ALTER TABLE user_profiles 
CHANGE COLUMN address address_id INT NOT NULL,
ADD CONSTRAINT fk_address_id FOREIGN KEY (address_id) REFERENCES user_address(user_address_id);

-- migrate:down
ALTER TABLE user_profiles 
DROP FOREIGN KEY fk_address_id,
CHANGE COLUMN address_id address TEXT;

