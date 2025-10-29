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
WHERE (sqlc.narg(name)::text IS NULL OR name ILIKE '%' || sqlc.narg(name) || '%')
  AND (sqlc.narg(breed)::text IS NULL OR breed ILIKE '%' || sqlc.narg(breed) || '%')
  AND (sqlc.narg(age)::text IS NULL OR age = sqlc.narg(age))
  AND (sqlc.narg(gender)::text IS NULL OR gender = sqlc.narg(gender)::pet_gender)
  AND (sqlc.narg(is_adopted)::boolean IS NULL OR is_adopted = sqlc.narg(is_adopted))
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: ListImagesById :many
SELECT url
FROM pets.pet_images
WHERE pet_id = $1;

-- name: InsertPetImages :exec
INSERT INTO pets.pet_images (id, pet_id, url)
SELECT gen_random_uuid(), $1, url
FROM unnest($2::text[]) AS url;

-- name: CountPets :one
SELECT COUNT(1)
FROM pets.pets
WHERE (sqlc.narg(name)::text IS NULL OR name ILIKE '%' || sqlc.narg(name) || '%')
  AND (sqlc.narg(breed)::text IS NULL OR breed ILIKE '%' || sqlc.narg(breed) || '%')
  AND (sqlc.narg(age)::text IS NULL OR age = sqlc.narg(age))
  AND (sqlc.narg(gender)::text IS NULL OR gender = sqlc.narg(gender)::pet_gender)
  AND (sqlc.narg(is_adopted)::boolean IS NULL OR is_adopted = sqlc.narg(is_adopted));