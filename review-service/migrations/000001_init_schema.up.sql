CREATE TABLE IF NOT EXISTS reviews (
    id uuid PRIMARY KEY NOT NULL,
    client_id uuid NOT NULL,
    business_id uuid NOT NULL,
    likes BOOLEAN,
    dislike BOOLEAN,
    comment TEXT[],
    created_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS reply_comments (
    id uuid PRIMARY KEY NOT NULL,
    review_id uuid NOT NULL,
    reply_comment TEXT[],
    created_at TIMESTAMP,
    FOREIGN KEY (review_id) REFERENCES reviews (id)

);