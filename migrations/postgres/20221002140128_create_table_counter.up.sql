CREATE TABLE IF NOT EXISTS "table_row_counter" (
    table_oid Oid PRIMARY KEY,
    table_name VARCHAR,
    count_hard_data int,
    count_soft_data int
);