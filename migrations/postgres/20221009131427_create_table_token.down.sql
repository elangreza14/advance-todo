BEGIN;

DROP TRIGGER IF EXISTS "log_token_insert" ON "tokens" CASCADE;

DROP TRIGGER IF EXISTS "log_token_update" ON "tokens" CASCADE;

DROP TRIGGER IF EXISTS "token_increment_trig" ON "tokens" CASCADE;

DROP TRIGGER IF EXISTS "token_decrement_trig" ON "tokens" CASCADE;

DROP INDEX IF EXISTS "tokens_id_index";

DROP TABLE IF EXISTS "tokens";

DROP TABLE IF EXISTS "tokens" CASCADE;

DELETE FROM
    "table_row_counter"
WHERE
    table_name = 'tokens';

COMMIT;