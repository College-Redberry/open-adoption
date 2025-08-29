-- name: CreateRequest :exec
INSERT INTO adoption.request (
    id,
    pet_id,
    name,
    email,
    phone,
    age,
    house_hold_agreed,
    alreadyPets,
    already_pets_castrated_and_vaccinated,
    property,
    own_property,
    address,
    income,
    suitable_location,
    access_to_the_street
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10, $11,
    $12, $13, $14, $15
);

-- name: ApproveRequest :exec
UPDATE adoption.request
SET approved_at = NOW()
WHERE id = $1;

-- name: ListRequests :many
SELECT
    id,
    pet_id,
    name,
    email,
    phone,
    approved_at,
    age,
    house_hold_agreed,
    alreadyPets,
    already_pets_castrated_and_vaccinated,
    property,
    own_property,
    address,
    income,
    suitable_location,
    access_to_the_street,
    created_at
FROM adoption.request
ORDER BY created_at DESC;

-- name: GetRequestById :one
SELECT
    id,
    pet_id,
    name,
    email,
    phone,
    approved_at,
    age,
    house_hold_agreed,
    alreadyPets,
    already_pets_castrated_and_vaccinated,
    property,
    own_property,
    address,
    income,
    suitable_location,
    access_to_the_street,
    created_at
FROM adoption.request
WHERE id = $1;
