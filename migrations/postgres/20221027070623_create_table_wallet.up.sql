BEGIN;

CREATE TABLE IF NOT EXISTS "wallets" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "name" VARCHAR NOT NULL,
  "balance" BIGINT NOT NULL CHECK(balance >= 0),
  "duration" BIGINT NOT NULL,
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
CREATE TRIGGER "log_wallet_insert" BEFORE
INSERT
  ON "wallets" FOR EACH ROW EXECUTE PROCEDURE log_insert_master();

CREATE TRIGGER "log_wallet_update" BEFORE
UPDATE
  ON "wallets" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

-- counter
CREATE TRIGGER "wallet_increment_trig"
AFTER
INSERT
  ON "wallets" FOR EACH ROW EXECUTE PROCEDURE count_increment();

CREATE TRIGGER "wallet_decrement_hard_trig"
AFTER
  DELETE ON "wallets" FOR EACH ROW EXECUTE PROCEDURE count_hard_decrement();

CREATE TRIGGER "wallet_decrement_soft_trig"
AFTER
UPDATE
  ON "wallets" FOR EACH ROW EXECUTE PROCEDURE count_soft_decrement();

CREATE INDEX "wallets_id_index" ON "wallets" ("id");

ALTER TABLE
  "wallets"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

INSERT INTO
  "table_row_counter"
VALUES
  ('wallets' :: regclass, 'wallets', 0, 0);

COMMIT;