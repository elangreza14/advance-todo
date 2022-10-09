BEGIN;

DROP TRIGGER IF EXISTS "log_user_insert" ON "users" CASCADE;
DROP TRIGGER IF EXISTS "log_user_update" ON "users" CASCADE;
DROP TRIGGER IF EXISTS "user_increment_trig" ON "users" CASCADE;
DROP TRIGGER IF EXISTS "user_decrement_trig" ON "users" CASCADE;

DROP INDEX IF EXISTS "users_id_index";
DROP INDEX IF EXISTS "users_email_index";

DROP TABLE IF EXISTS "users";

DROP TABLE IF EXISTS "users" CASCADE;
DELETE FROM "table_row_counter"
WHERE table_name = 'users';

COMMIT;