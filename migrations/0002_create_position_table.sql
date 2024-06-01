-- +goose Up
CREATE TABLE IF NOT EXISTS kucoin_position (
    id serial PRIMARY KEY,
    bot_name VARCHAR(10) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    fk_open_order_id INTEGER REFERENCES kucoin_market_order(id) NOT NULL,
    fk_close_order_id INTEGER REFERENCES kucoin_market_order(id)
);

-- +goose Down
DROP TABLE IF EXISTS kucoin_position;