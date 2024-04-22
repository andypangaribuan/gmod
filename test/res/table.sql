CREATE TABLE IF NOT EXISTS users (
    created_at  TIMESTAMP(3)  WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP(3)  WITH TIME ZONE NOT NULL,
    deleted_at  TIMESTAMP(3)  WITH TIME ZONE,
    uid         VARCHAR(20)   NOT NULL,
    name        VARCHAR(500)  NOT NULL,
    address     VARCHAR(500),
    height      INTEGER,
    gold_amount NUMERIC(24, 14)
);