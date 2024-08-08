-- name: GetUser :one
SELECT * FROM users
WHERE user_id = ?;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :execresult
INSERT INTO users (username, password_hash, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?);

-- name: GetUserProfile :one
SELECT * FROM user_profiles
WHERE profile_id = ?;

-- name: GetUsersProfile :many
SELECT * FROM user_profiles;

-- name: CreateUserProfile :execresult
INSERT INTO user_profiles (user_id, first_name, last_name, phone_number, address_id, date_of_birth, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);


-- name: CreateUserAddress :execresult
INSERT INTO user_address (user_id, street_address, house_number, complement, city, zone, district, postal_code, country) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);