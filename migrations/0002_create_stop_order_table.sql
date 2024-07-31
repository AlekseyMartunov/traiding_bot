-- +goose Up
CREATE TABLE IF NOT EXISTS kucoin_market_stop_order (
    id serial PRIMARY KEY,
    order_id VARCHAR(50) NOT NULL,
    client_order_id VARCHAR(50) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    side VARCHAR(4) NOT NULL,
    stop_trigger VARCHAR(5) NOT NULL,
    stop_price NUMERIC(20, 10) NOT NULL,
    order_size NUMERIC(20, 10) NOT NULL,
    order_price NUMERIC(20, 10) NOT NULL,
    commission NUMERIC(20, 10) NOT NULL,
    commission_currency VARCHAR(10) NOT NULL,
    created_time TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS market_order;