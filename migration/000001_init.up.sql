-- Tables
CREATE TABLE "Users" (
    "id" INTEGER NOT NULL AUTO_INCREMENT,
    "email" VARCHAR(320) NOT NULL,
    "password_hash" VARCHAR(32) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

    PRIMARY KEY ("id")
);

CREATE TABLE "Notes" (
    "id" INTEGER NOT NULL AUTO_INCREMENT,
    "user_id" INTEGER NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "content" VARCHAR,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "Users" ("id")
);

-- Updated trigger
CREATE OR REPLACE FUNCTION update_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger activation loop for tables
DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE update_set_timestamp()', t,t);
    END loop;
END;
$$ language 'plpgsql';
