CREATE TABLE IF NOT EXISTS likes(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid ,
    staff_id uuid ,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS comments(
    id uuid PRIMARY KEY NOT NULL,
    user_id uuid ,
    staff_id uuid ,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
    deleted boolean,
    comment text
    );

CREATE TABLE IF NOT EXISTS reply_comments(
    id uuid PRIMARY KEY NOT NULL,
    user_id uuid ,
    staff_id uuid ,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
    deleted boolean,
    comment text
    );

