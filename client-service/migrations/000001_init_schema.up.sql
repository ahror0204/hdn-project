CREATE TABLE IF NOT EXISTS clients(
    id uuid PRIMARY KEY NOT NULL,
    calendar_id uuid,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone_numbers TEXT[],
    email VARCHAR(255),
    status VARCHAR(255),
    payment_card VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (calendar_id) REFERENCES calendars(id) 
);
