BEGIN;

DROP TRIGGER IF EXISTS "log_expense_insert" ON "expenses" CASCADE;

DROP TRIGGER IF EXISTS "log_expense_update" ON "expenses" CASCADE;

DROP TRIGGER IF EXISTS "expense_increment_trig" ON "expenses" CASCADE;

DROP TRIGGER IF EXISTS "expense_decrement_trig" ON "expenses" CASCADE;

DROP INDEX IF EXISTS "expenses_id_index";

DROP TABLE IF EXISTS "expenses";

DROP TABLE IF EXISTS "expenses" CASCADE;

DELETE FROM
    "table_row_counter"
WHERE
    table_name = 'expenses';

COMMIT;