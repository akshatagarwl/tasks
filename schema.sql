CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE task_status AS ENUM ('PENDING', 'IN_PROGRESS', 'COMPLETED');

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    status task_status NOT NULL DEFAULT 'PENDING',
    created_at timestamptz NOT NULL DEFAULT now(),
    last_modified_at timestamptz NOT NULL DEFAULT now()
);

-- Function to update last_modified_at
CREATE OR REPLACE FUNCTION update_last_modified_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.last_modified_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to automatically update last_modified_at on UPDATE
CREATE TRIGGER update_task_last_modified_at
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION update_last_modified_at();
