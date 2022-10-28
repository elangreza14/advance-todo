BEGIN;

DROP TRIGGER IF EXISTS "log_income_insert" ON "incomes" CASCADE;

DROP TRIGGER IF EXISTS "log_income_update" ON "incomes" CASCADE;

DROP TRIGGER IF EXISTS "income_increment_trig" ON "incomes" CASCADE;

DROP TRIGGER IF EXISTS "income_decrement_trig" ON "incomes" CASCADE;

DROP INDEX IF EXISTS "incomes_id_index";

DROP TABLE IF EXISTS "incomes";

DROP TABLE IF EXISTS "incomes" CASCADE;

DELETE FROM
    "table_row_counter"
WHERE
    table_name = 'incomes';

COMMIT;