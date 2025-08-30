CREATE SCHEMA adoption;

CREATE TYPE property_type AS ENUM ('house', 'apartment');

CREATE TABLE adoption.request (
    id UUID PRIMARY KEY,
    pet_id UUID NOT NULL REFERENCES pets.pets(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    approved_at TIMESTAMPTZ NULL,
    age INT NOT NULL,
    house_hold_agreed BOOLEAN NOT NULL,
    alreadyPets INT NOT NULL,
    already_pets_castrated_and_vaccinated BOOLEAN NOT NULL,
    property property_type NOT NULL,
    own_property BOOLEAN NOT NULL,
    address TEXT NOT NULL,
    income INT NOT NULL,
    suitable_location TEXT NOT NULL,
    access_to_the_street BOOLEAN NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
