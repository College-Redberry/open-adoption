-- name: CreatePet :exec
INSERT INTO pets.pets (
    id, name, breed, age, gender, is_adopted
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: UpdatePet :exec
UPDATE pets.pets
SET name = $2,
    breed = $3,
    age = $4,
    gender = $5
WHERE id = $1;

-- name: AdoptPetById :exec
UPDATE pets.pets
SET is_adopted = TRUE
WHERE id = $1;

-- name: GetPetById :one
SELECT id, name, breed, age, gender, is_adopted
FROM pets.pets
WHERE id = $1;

-- name: ListPets :many
SELECT id, name, breed, age, gender, is_adopted
FROM pets.pets
ORDER BY name;

-- name: ListImagesById :many
SELECT url
FROM pets.pet_images
WHERE pet_id = $1;
