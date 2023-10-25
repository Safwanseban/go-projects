
-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email=$1;

-- name: CreateUser :one
INSERT INTO users(
    
username,email,password,phone_number
)VALUES(
   
$1,$2,$3,$4
)
RETURNING id;



