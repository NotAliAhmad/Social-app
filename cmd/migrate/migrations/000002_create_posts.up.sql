CREATE TABLE IF NOT EXISTS posts(
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    user_id uuid NOT NULL,
    content text NOT NULL,
    created_at TIMESTAMP(0) WITH time zone NOT NULL DEFAULT NOW()
);