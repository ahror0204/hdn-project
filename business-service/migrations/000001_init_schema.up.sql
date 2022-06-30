CREATE TABLE IF NOT EXISTS business (
    id uuid NOT NULL PRIMARY KEY,
    salon_name VARCHAR(255),
    phone_numbers []TEXT,
    status VARCHAR(255),
    staff []TEXT,
    location TEXT
);

CREATE TABLE IF NOT EXISTS men_services (
    id uuid PRIMARY KEY NOT NULL,
    hair_cut BOOLEAN NOT NULL,
    beard_cut BOOLEAN NOT NULL,
    hair_coloring BOOLEAN NOT NULL,
    special_hair_cut BOOLEAN NOT NULL,
    beard_coloring BOOLEAN NOT NULL,
    beard_trim BOOLEAN NOT NULL,
    beard_shave BOOLEAN NOT NULL,
    face_shave BOOLEAN NOT NULL,
    boy_hair_cut BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS women_services (
    id uuid PRIMARY KEY NOT NULL,
    hair_cut BOOLEAN NOT NULL,
    hair_coloring BOOLEAN NOT NULL,
    special_hair_cut BOOLEAN NOT NULL,
    eyebrow_arching BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS staff (
    id uuid NOT NULL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone_numbers []TEXT,
    cost INTEGER,
    status VARCHAR(255),
    comment_id uuid,
    business-id uuid NOT NULL,
    calendar_id uuid,
    client_id uuid NOT NULL,
    men_service_id uuid,
    women_service_id uuid,
    FOREIGN KEY (business_id) REFERENCES businesss (id) ON DELETE CASCADE,
    FOREIGN KEY (men_service_id) REFERENCES men_services (id) ON DELETE CASCADE,
    FOREIGN KEY (women_service_id) REFERENCES women_services (id) ON DELETE CASCADE
);
