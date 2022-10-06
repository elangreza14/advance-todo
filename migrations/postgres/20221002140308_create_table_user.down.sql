BEGIN;

DROP TRIGGER IF EXISTS "log_user_insert" ON "users" CASCADE;
DROP TRIGGER IF EXISTS "log_user_update" ON "users" CASCADE;
DROP TRIGGER IF EXISTS "user_increment_trig" ON "users" CASCADE;
DROP TRIGGER IF EXISTS "user_decrement_trig" ON "users" CASCADE;

DROP TABLE IF EXISTS "users";

COMMIT;