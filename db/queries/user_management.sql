-- name: GetUser :one
SELECT * FROM users
WHERE user_id = ?;

-- name: UserNameOrEmailExists :one
SELECT COUNT(*) FROM users
WHERE username = ? OR email = ?;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserCompleteDataById :one
SELECT 
    u.user_id,
    u.username, 
    u.email, 
    u.created_at, 
    u.updated_at, 
    up.first_name,  
    up.last_name,
    up.phone_number,
    up.date_of_birth,
    ua.street_address,
    ua.house_number,
    ua.complement,
    ua.city,
    ua.zone,
    ua.district,
    ua.postal_code,
    ua.country
FROM users u
JOIN user_profiles up ON up.user_id = u.user_id
JOIN user_address ua ON up.address_id = ua.user_address_id
WHERE u.user_id = ?;

-- name: GetUsersCompleteData :many
SELECT
    u.user_id,
    u.username, 
    u.email, 
    u.created_at, 
    u.updated_at, 
    up.first_name,  
    up.last_name,
    up.phone_number,
    up.date_of_birth,
    ua.street_address,
    ua.house_number,
    ua.complement,
    ua.city,
    ua.zone,
    ua.district,
    ua.postal_code,
    ua.country
FROM users u
JOIN user_profiles up ON up.user_id = u.user_id
JOIN user_address ua ON up.address_id = ua.user_address_id;

-- name: CreateUser :execresult
INSERT INTO users (username, password_hash, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?);

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = ?;

-- name: UpdateUser :exec
UPDATE users SET username = ?, password_hash = ?, email = ?, updated_at = ?
WHERE user_id = ?;

-- name: GetUserProfile :one
SELECT * FROM user_profiles
WHERE profile_id = ?;

-- name: GetUsersProfile :many
SELECT * FROM user_profiles;

-- name: CreateUserProfile :execresult
INSERT INTO user_profiles (user_id, first_name, last_name, phone_number, address_id, date_of_birth, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteUsersProfileByUserId :exec
DELETE FROM user_profiles WHERE user_id = ?;

-- name: UpdateUserProfile :exec
UPDATE user_profiles SET first_name = ?, last_name = ?, phone_number = ?, date_of_birth = ?
WHERE user_id = ?;


-- name: CreateUserAddress :execresult
INSERT INTO user_address (user_id, street_address, house_number, complement, city, zone, district, postal_code, country) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteUsersAddressByUserId :exec
DELETE FROM user_address WHERE user_id = ?;

-- name: UpdateUserAddress :exec
UPDATE user_address 
SET street_address = ?, house_number = ?, complement = ?, city = ?, zone = ?, district = ?, postal_code = ?, country = ?
WHERE user_id = ?;