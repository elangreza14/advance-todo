BEGIN;

CREATE OR REPLACE FUNCTION count_increment() 
RETURNS TRIGGER AS $_$
BEGIN
    UPDATE table_row_counter SET count_hard_data = count_hard_data + 1 WHERE table_oid = TG_RELID;
    UPDATE table_row_counter SET count_soft_data = count_soft_data + 1 WHERE table_oid = TG_RELID;
    RETURN NEW;
END; 
$_$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION count_soft_decrement() 
RETURNS TRIGGER AS $_$
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        UPDATE table_row_counter SET count_soft_data = count_soft_data - 1  WHERE table_oid = TG_RELID;
    END IF;

    IF NEW.deleted_at IS NULL AND OLD.deleted_at IS NOT NULL THEN
        UPDATE table_row_counter SET count_soft_data = count_soft_data + 1  WHERE table_oid = TG_RELID;
    END IF;
    
    RETURN NEW;
END; 
$_$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION count_hard_decrement() 
RETURNS TRIGGER AS $_$
BEGIN
    UPDATE table_row_counter SET count_hard_data = count_hard_data - 1  WHERE table_oid = TG_RELID;
    IF OLD.deleted_at IS NULL THEN
        UPDATE table_row_counter SET count_soft_data = count_soft_data - 1  WHERE table_oid = TG_RELID;
    END IF;
    RETURN NEW;
END; 
$_$ LANGUAGE 'plpgsql';

COMMIT;