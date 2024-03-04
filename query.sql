-- name: FindContactsByName :many
SELECT * FROM contacts WHERE name LIKE $1;

-- name: FindContactsByPhoneNumber :many
SELECT * FROM contacts WHERE phone_number LIKE $1;