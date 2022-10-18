BEGIN;

DROP FUNCTION IF EXISTS count_increment();

DROP FUNCTION IF EXISTS count_hard_decrement();

DROP FUNCTION IF EXISTS count_soft_decrement();

COMMIT;