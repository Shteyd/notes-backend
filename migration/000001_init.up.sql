-- Extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Updated trigger
CREATE OR REPLACE FUNCTION update_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Tables
CREATE TABLE "Users" (
    "id" uuid DEFAULT uuid_generate_v4 (),
    "email" VARCHAR(320) NOT NULL,
    "password_hash" VARCHAR(32) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

    PRIMARY KEY ("id")
);

CREATE TRIGGER _update_users_timestamp
    BEFORE UPDATE ON "Users"
    FOR EACH ROW EXECUTE PROCEDURE update_set_timestamp();

CREATE TABLE "Notes" (
    "id" uuid DEFAULT uuid_generate_v4 (),
    "user_id" uuid NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "content" VARCHAR,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "Users" ("id")
);

CREATE TRIGGER _update_notes_timestamp
    BEFORE UPDATE ON "Notes"
    FOR EACH ROW EXECUTE PROCEDURE update_set_timestamp();