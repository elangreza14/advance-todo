BEGIN;

-- id of token will be generated manually by service
-- because it's easier to get token by the same id
CREATE TABLE IF NOT EXISTS "tokens" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "token" VARCHAR NOT NULL,
  "token_type" VARCHAR NOT NULL,
  "ip" VARCHAR NOT NULL,
  "issued_at" TIMESTAMPTZ NOT NULL,
  "expired_at" TIMESTAMPTZ NOT NULL,
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
CREATE TRIGGER "log_token_insert" BEFORE
INSERT
  ON "tokens" FOR EACH ROW EXECUTE PROCEDURE log_insert_master();

CREATE TRIGGER "log_token_update" BEFORE
UPDATE
  ON "tokens" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

-- counter
CREATE TRIGGER "token_increment_trig"
AFTER
INSERT
  ON "tokens" FOR EACH ROW EXECUTE PROCEDURE count_increment();

CREATE TRIGGER "token_decrement_hard_trig"
AFTER
  DELETE ON "tokens" FOR EACH ROW EXECUTE PROCEDURE count_hard_decrement();

CREATE TRIGGER "token_decrement_soft_trig"
AFTER
UPDATE
  ON "tokens" FOR EACH ROW EXECUTE PROCEDURE count_soft_decrement();

CREATE INDEX "tokens_id_index" ON "tokens" ("id");

ALTER TABLE
  "tokens"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

INSERT INTO
  "table_row_counter"
VALUES
  ('tokens' :: regclass, 'tokens', 0, 0);

COMMIT;