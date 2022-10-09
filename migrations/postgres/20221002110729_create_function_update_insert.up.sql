BEGIN;

CREATE OR REPLACE FUNCTION log_update_master()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    NEW.version = OLD.version + 1;
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION log_insert_master()
RETURNS TRIGGER AS $$
BEGIN
    NEW.created_at = NOW();
    NEW.version = 0;
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

COMMIT;
