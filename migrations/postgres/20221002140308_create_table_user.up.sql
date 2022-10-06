BEGIN;

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "full_name" VARCHAR(50) NOT NULL,
  "email" VARCHAR(25) NOT NULL,
  "password" VARCHAR NOT NULL,

  "version" INT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL,
  "created_by" VARCHAR(50) NULL,
  "updated_at" TIMESTAMPTZ NOT NULL,
  "updated_by" VARCHAR(50) NULL,
  "deleted_at" TIMESTAMPTZ NULL,
  "deleted_by" VARCHAR(50) NULL, 
  UNIQUE ("id")
);

-- updater
CREATE TRIGGER "log_user_insert" BEFORE INSERT ON "users" FOR EACH ROW EXECUTE PROCEDURE log_insert_master();
CREATE TRIGGER "log_user_update" BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE log_update_master();
-- counter
CREATE TRIGGER "user_increment_trig" AFTER INSERT ON "users" FOR EACH ROW EXECUTE PROCEDURE count_increment();
CREATE TRIGGER "user_decrement_hard_trig" AFTER DELETE ON "users" FOR EACH ROW EXECUTE PROCEDURE count_hard_decrement();
CREATE TRIGGER "user_decrement_soft_trig" AFTER UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE count_soft_decrement();

INSERT INTO "table_row_counter" VALUES ('users'::regclass, 'users', 0, 0);

COMMIT;