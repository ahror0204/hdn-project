CREATE TABLE IF NOT EXISTS business (
    id uuid NOT NULL PRIMARY KEY,
    salon_name VARCHAR(255),
    phone_numbers TEXT[],
    status VARCHAR(255),
    staff TEXT[],
    location TEXT, 
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS staff (
    id uuid NOT NULL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone_numbers TEXT[],
    cost INTEGER,
    status VARCHAR(255),
    business_id uuid NOT NULL,
    calendar_id uuid,
    user_id uuid NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (business_id) REFERENCES business (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS men_services (
    id uuid PRIMARY KEY NOT NULL,
    hair_cut BOOLEAN NOT NULL DEFAULT false,
    beard_cut BOOLEAN NOT NULL DEFAULT false,
    hair_coloring BOOLEAN NOT NULL DEFAULT false,
    special_hair_cut BOOLEAN NOT NULL DEFAULT false,
    beard_coloring BOOLEAN NOT NULL DEFAULT false,
    beard_trim BOOLEAN NOT NULL DEFAULT false,
    beard_shave BOOLEAN NOT NULL DEFAULT false,
    face_shave BOOLEAN NOT NULL DEFAULT false,
    boy_hair_cut BOOLEAN NOT NULL DEFAULT false,
    user_id uuid NOT NULL
);

CREATE TABLE IF NOT EXISTS women_services (
    id uuid PRIMARY KEY NOT NULL,
    hair_cut BOOLEAN NOT NULL DEFAULT false,
    hair_coloring BOOLEAN NOT NULL DEFAULT false,
    special_hair_cut BOOLEAN NOT NULL DEFAULT false,
    eyebrow_arching BOOLEAN NOT NULL DEFAULT false,
    user_id uuid NOT NULL
);