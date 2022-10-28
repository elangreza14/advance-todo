BEGIN;

CREATE TABLE IF NOT EXISTS "incomes" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "wallet_id" uuid NOT NULL,
  "name" VARCHAR NOT NULL,
  "amount" BIGINT NOT NULL CHECK(amount > 0),
  "month" BIGINT NOT NULL,
  "version" INT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_at" TIMESTAMPTZ NULL,
  "updated_by" uuid NULL,
  "deleted_at" TIMESTAMPTZ NULL,
  "deleted_by" uuid NULL,
  UNIQUE ("id")
);

-- updater
CREATE TRIGGER "log_income_insert" BEFORE
INSERT
  ON "incomes" FOR EACH ROW EXECUTE PROCEDURE log_insert_master();

CREATE TRIGGER "log_income_update" BEFORE
UPDATE
  ON "incomes" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

-- counter
CREATE TRIGGER "income_increment_trig"
AFTER
INSERT
  ON "incomes" FOR EACH ROW EXECUTE PROCEDURE count_increment();

CREATE TRIGGER "income_decrement_hard_trig"
AFTER
  DELETE ON "incomes" FOR EACH ROW EXECUTE PROCEDURE count_hard_decrement();

CREATE TRIGGER "income_decrement_soft_trig"
AFTER
UPDATE
  ON "incomes" FOR EACH ROW EXECUTE PROCEDURE count_soft_decrement();

CREATE INDEX "incomes_id_index" ON "incomes" ("id");

ALTER TABLE
  "incomes"
ADD
  FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id") ON DELETE CASCADE;

INSERT INTO
  "table_row_counter"
VALUES
  ('incomes' :: regclass, 'incomes', 0, 0);

COMMIT;