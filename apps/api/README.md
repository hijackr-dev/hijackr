# hijackr-api

The backend API for the Hijackr platform — licence validation, Stripe webhooks, and account management for all Hijackr products.

**Base URL:** `https://api.hijackr.io`

---

## Stack

- **Language:** Go 1.23
- **Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** PostgreSQL (via [pgx](https://github.com/jackc/pgx))
- **Payments:** [Stripe Go SDK](https://github.com/stripe/stripe-go)
- **Deployment:** Fly.io or Railway

---

## API Endpoints

### Public

| Method | Path | Description |
|---|---|---|
| `GET` | `/health` | Health check |
| `POST` | `/v1/licence/validate` | Validate a licence key (called by desktop apps on launch) |
| `POST` | `/v1/webhooks/stripe` | Stripe webhook receiver |
| `POST` | `/v1/auth/magic-link` | Request a magic link login email |
| `GET` | `/v1/auth/verify?token=...` | Verify magic link token, returns session |

### Authenticated (Bearer token required)

| Method | Path | Description |
|---|---|---|
| `GET` | `/v1/me` | Get current customer info |
| `GET` | `/v1/licences` | List all licences for the current customer |
| `GET` | `/v1/licences/:id/machines` | List activated machines for a licence |
| `DELETE` | `/v1/licences/:id/machines/:machine_id` | Deactivate a machine |
| `POST` | `/v1/billing/portal` | Create a Stripe Customer Portal session URL |

---

## Licence Validation (Desktop App Integration)

Desktop apps call `POST /v1/licence/validate` on every launch:

```json
// Request
{
  "licence_key": "ABCD-1234-EFGH-5678",
  "product": "offloadr",
  "machine_id": "<hardware-fingerprint>",
  "machine_name": "Alex's MacBook Pro"
}

// Response (valid)
{
  "valid": true,
  "tier": "pro",
  "expires_at": null,
  "features": ["unlimited_destinations", "cloud_upload", "provr_manifests", "lto_archive"]
}

// Response (invalid)
{
  "valid": false,
  "reason": "machine_limit_reached"
}
```

**Reason codes:**
- `licence_not_found` — key does not exist
- `product_mismatch` — key is for a different product
- `licence_expired` — subscription has lapsed
- `licence_revoked` — manually revoked or subscription cancelled
- `machine_limit_reached` — too many machines activated (customer must deactivate one via app.hijackr.io)

**Offline grace period:** The desktop app caches the last valid response (encrypted, signed) for 24 hours. The app works offline for up to 24 hours without re-validating.

---

## Local Development

### Prerequisites

- Go 1.23+
- PostgreSQL 15+
- A Stripe account (test mode)

### Setup

```bash
# 1. Clone the repo
git clone https://github.com/hijackr-dev/hijackr-api.git
cd hijackr-api

# 2. Copy environment variables
cp .env.example .env
# Edit .env with your values

# 3. Create the database
createdb hijackr_api

# 4. Install dependencies
go mod download

# 5. Run the server (migrations run automatically on startup)
go run ./cmd/server
```

The API will be available at `http://localhost:8080`.

### Stripe Webhooks (local testing)

Install the [Stripe CLI](https://stripe.com/docs/stripe-cli) and forward webhooks to your local server:

```bash
stripe listen --forward-to localhost:8080/v1/webhooks/stripe
```

Copy the webhook signing secret printed by the CLI into your `.env` as `STRIPE_WEBHOOK_SECRET`.

### Test a licence validation

```bash
curl -X POST http://localhost:8080/v1/licence/validate \
  -H "Content-Type: application/json" \
  -d '{
    "licence_key": "TEST-0000-0000-0001",
    "product": "offloadr",
    "machine_id": "test-machine-001",
    "machine_name": "Dev Machine"
  }'
```

---

## Stripe Setup

### Products and Prices

Create a product in Stripe for each Hijackr app. Set the following metadata on each **Checkout Session** (in your payment link or checkout code):

```
metadata.product = "offloadr"   (or "scrollr", "transcodr", etc.)
metadata.tier    = "pro"        (or "studio")
```

The webhook handler reads these values to create the correct licence record.

### Webhook Events to Enable

In the Stripe Dashboard → Webhooks, enable:
- `checkout.session.completed`
- `customer.subscription.updated`
- `customer.subscription.deleted`
- `invoice.payment_failed`

---

## Deployment

### Fly.io (recommended)

```bash
fly launch
fly secrets set DATABASE_URL="..." STRIPE_SECRET_KEY="..." STRIPE_WEBHOOK_SECRET="..."
fly deploy
```

### Environment Variables (production)

| Variable | Description |
|---|---|
| `ENV` | Set to `production` |
| `PORT` | Server port (default: `8080`) |
| `DATABASE_URL` | PostgreSQL connection string |
| `STRIPE_SECRET_KEY` | Stripe secret key (`sk_live_...`) |
| `STRIPE_WEBHOOK_SECRET` | Stripe webhook signing secret (`whsec_...`) |
| `APP_BASE_URL` | Base URL of app.hijackr.io (`https://app.hijackr.io`) |

---

## Adding a New Product

1. Add the product slug and its tier features to `featuresByTier` in `internal/handlers/licence.go`
2. Create the product in Stripe with `metadata.product = "your-product-slug"`
3. That's it — the licence system handles the rest automatically

---

## Project Structure

```
cmd/server/main.go              ← entry point, router setup
internal/
  db/db.go                      ← PostgreSQL connection + schema migrations
  handlers/
    auth.go                     ← magic link auth, session management
    billing.go                  ← Stripe Customer Portal
    licence.go                  ← licence validation, machine management
    webhooks.go                 ← Stripe webhook handler
  middleware/
    auth.go                     ← Bearer token session validation
.env.example                    ← environment variable template
go.mod