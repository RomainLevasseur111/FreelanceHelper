BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS users (
    id TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password BLOB NOT NULL,
    timestamp DATETIME NOT NULL
)
CREATE TABLE IF NOT EXISTS simulations (
    id TEXT NOT NULL,
    user_id TEXT REFERENCES users(id),
    status TEXT NOT NULL,
    daily_rate INTEGER NOT NULL,
    contract_length INTEGER NOT NULL,
    monthly_charges INTEGER
)
COMMIT;