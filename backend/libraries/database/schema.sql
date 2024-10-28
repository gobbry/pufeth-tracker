CREATE EXTENSION IF NOT EXISTS timescaledb;

CREATE TABLE
    IF NOT EXISTS puffer_conversion_rates (
        chain VARCHAR(255) NOT NULL,
        block_number BIGINT NOT NULL,
        block_timestamp_ms BIGINT NOT NULL,
        total_asset NUMERIC NOT NULL,
        total_supply NUMERIC NOT NULL,
        conversion_rate NUMERIC NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        PRIMARY KEY (chain, block_number)
    );

SELECT
    create_hypertable (
        'puffer_conversion_rates',
        'block_timestamp_ms',
        chunk_time_interval => 86400000, -- 1 day in milliseconds
        if_not_exists => TRUE
    );

CREATE INDEX IF NOT EXISTS idx_pcr_chain_block_number ON puffer_conversion_rates (chain, block_number);

CREATE INDEX IF NOT EXISTS idx_pcr_chain_block_timestamp ON puffer_conversion_rates (chain, block_timestamp_ms);

CREATE INDEX IF NOT EXISTS idx_pcr_chain ON puffer_conversion_rates (chain);