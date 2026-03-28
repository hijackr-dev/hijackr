# Hijackr

Monorepo for the Hijackr platform.

## Structure

```
hijackr/
├── apps/
│   ├── api/    # api.hijackr.io — Go/Gin backend
│   ├── app/    # app.hijackr.io — SvelteKit account portal
│   └── web/    # hijackr.io — SvelteKit marketing site
└── packages/
    └── types/  # Shared TypeScript types
```

## Apps

| App | URL | Description |
|-----|-----|-------------|
| `apps/api` | `api.hijackr.io` | Go/Gin backend — licence validation, Stripe webhooks, auth |
| `apps/app` | `app.hijackr.io` | SvelteKit account portal — licences, machines, billing |
| `apps/web` | `hijackr.io` | SvelteKit marketing site |

## Prerequisites

- [pnpm](https://pnpm.io) ≥ 9
- [Go](https://go.dev) ≥ 1.23
- [Node.js](https://nodejs.org) ≥ 20

## Getting Started

```bash
# Install all JS/TS dependencies
pnpm install

# Run the account portal in dev mode
pnpm dev:app

# Run the marketing site in dev mode
pnpm dev:web

# Run the API
cd apps/api && go run ./cmd/server
```

## Packages

| Package | Description |
|---------|-------------|
| `packages/types` | Shared TypeScript interfaces (User, Machine, Subscription, Invoice) |