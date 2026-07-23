CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS parts (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name                TEXT           NOT NULL,
    category            TEXT           NOT NULL,
    current_stock       INTEGER        NOT NULL,
    minimum_stock       INTEGER        NOT NULL,
    average_daily_sales INTEGER        NOT NULL,
    lead_time_days      INTEGER        NOT NULL,
    criticality_level   INTEGER        NOT NULL,
    unit_cost           NUMERIC(12, 2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_parts_category ON parts (category);
