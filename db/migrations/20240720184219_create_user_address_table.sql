-- migrate:up
CREATE TABLE IF NOT EXISTS user_address (
    user_address_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    street_address VARCHAR(50) NOT NULL,
    house_number int,
    complement VARCHAR(50),
    city VARCHAR(50) NOT NULL,
    zone VARCHAR(50) NOT NULL,
    district VARCHAR(50),
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(50) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- migrate:down
DROP TABLE user_address;
