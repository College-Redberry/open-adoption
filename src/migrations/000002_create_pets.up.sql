CREATE SCHEMA pets;

CREATE TYPE pet_gender AS ENUM ('male', 'female');

CREATE TABLE pets.pets (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    breed VARCHAR(255) NOT NULL,
    age VARCHAR(255) NOT NULL,
    gender pet_gender NOT NULL,
    is_adopted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE pets.pet_images (
    id UUID PRIMARY KEY,
    pet_id UUID REFERENCES pets.pets(id) ON DELETE CASCADE,
    url TEXT NOT NULL
);