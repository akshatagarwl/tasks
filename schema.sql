CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE task_status AS ENUM ('TODO', 'IN_PROGRESS', 'COMPLETED');

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    status task_status NOT NULL DEFAULT 'TODO',
    created_at timestamptz NOT NULL DEFAULT now(),
    last_modified_at timestamptz NOT NULL DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_last_modified_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.last_modified_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_task_last_modified_at
BEFORE UPDATE ON tasks
FOR EACH ROW
EXECUTE FUNCTION update_last_modified_at();
