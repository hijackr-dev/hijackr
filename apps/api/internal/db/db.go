package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect opens a connection pool to PostgreSQL.
func Connect(databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}
	return pool, nil
}

// Migrate runs the embedded SQL schema against the database.
// Idempotent — safe to run on every startup.
func Migrate(pool *pgxpool.Pool) error {
	_, err := pool.Exec(context.Background(), schema)
	return err
}

const schema = `
-- Customers (one per Stripe customer)
CREATE TABLE IF NOT EXISTS customers (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stripe_customer_id  TEXT UNIQUE NOT NULL,
    email               TEXT NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Sessions (magic link auth tokens)
CREATE TABLE IF NOT EXISTS sessions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    token       TEXT UNIQUE NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Magic link tokens (one-time use)
CREATE TABLE IF NOT EXISTS magic_links (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email       TEXT NOT NULL,
    token       TEXT UNIQUE NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL,
    used        BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Licences (one per active subscription/purchase)
CREATE TABLE IF NOT EXISTS licences (
    id                      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id             UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    product                 TEXT NOT NULL,
    tier                    TEXT NOT NULL DEFAULT 'pro',
    licence_key             TEXT UNIQUE NOT NULL,
    stripe_subscription_id  TEXT,
    stripe_price_id         TEXT,
    status                  TEXT NOT NULL DEFAULT 'active',
    machine_limit           INT NOT NULL DEFAULT 2,
    expires_at              TIMESTAMPTZ,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Machine activations (one per device that has activated a licence)
CREATE TABLE IF NOT EXISTS licence_activations (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    licence_id   UUID NOT NULL REFERENCES licences(id) ON DELETE CASCADE,
    machine_id   TEXT NOT NULL,
    machine_name TEXT,
    activated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_seen_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (licence_id, machine_id)
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_licences_customer_id ON licences(customer_id);
CREATE INDEX IF NOT EXISTS idx_licences_licence_key ON licences(licence_key);
CREATE INDEX IF NOT EXISTS idx_activations_licence_id ON licence_activations(licence_id);
CREATE INDEX IF NOT EXISTS idx_sessions_token ON sessions(token);
CREATE INDEX IF NOT EXISTS idx_magic_links_token ON magic_links(token);
`