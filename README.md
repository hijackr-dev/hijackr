# Hijackr

Monorepo for the Hijackr platform — web, app, and API.

## Structure

```
hijackr/
├── apps/
│   ├── api/    # Go/Gin REST API
│   ├── app/    # SvelteKit user dashboard
│   └── web/    # SvelteKit marketing site
└── packages/
    └── types/  # Shared TypeScript types
```

## Prerequisites

- [pnpm](https://pnpm.io) ≥ 9
- [Go](https://go.dev) ≥ 1.23
- [Node.js](https://nodejs.org) ≥ 20

## Getting Started

```bash
# Install all JS/TS dependencies
pnpm install

# Run the dashboard (apps/app) in dev mode
pnpm dev:app

# Run the marketing site (apps/web) in dev mode
pnpm dev:web

# Run the API
cd apps/api && go run ./cmd/server
```

## Apps

| App | Description | Port |
|-----|-------------|------|
| `apps/api` | Go REST API (Gin + PostgreSQL + Stripe) | 8080 |
| `apps/app` | SvelteKit user dashboard | 5173 |
| `apps/web` | SvelteKit marketing site | 5174 |

## Packages

| Package | Description |
|---------|-------------|
| `packages/types` | Shared TypeScript interfaces (User, Machine, Billing) |