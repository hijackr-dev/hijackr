# hijackr-web

The official marketing website for [hijackr.io](https://hijackr.io).

## Stack

- **Framework:** [SvelteKit](https://kit.svelte.dev) with `@sveltejs/adapter-static`
- **UI:** [Skeleton UI](https://www.skeleton.dev) + [Tailwind CSS](https://tailwindcss.com)
- **Language:** TypeScript
- **Deployment:** Cloudflare Pages (or Vercel)

## Architecture

All product data lives in a single source of truth:

```
src/lib/products.ts   ← add/edit products here
```

Routes are generated dynamically from that data:

```
/                     → homepage with product grid
/[product]            → individual product page (e.g. /offloadr, /scrollr)
/about                → company page
```

Individual product domains (e.g. `offloadr.io`, `scrollr.io`) redirect here via DNS CNAME.

## Development

```bash
npm install
npm run dev
```

The site will be available at **http://localhost:5173/**

| Route | Page |
|---|---|
| `http://localhost:5173/` | Homepage — hero + all product cards |
| `http://localhost:5173/offloadr` | Offloadr product page |
| `http://localhost:5173/provr` | Provr product page |
| `http://localhost:5173/scrollr` | Scrollr product page |
| `http://localhost:5173/[any-slug]` | Any other product page |

If the dev server has stopped, restart it with:

```bash
cd /Users/alexseery/hijackr/hijackr-web && npm run dev
```

## Build

```bash
npm run build
# Output: ./build (static HTML — deploy to Cloudflare Pages)
```

## Adding a Product

Edit `src/lib/products.ts` and add a new entry to the `products` array. The route and product page are generated automatically.

## Deployment

Connect the `hijackr-dev/hijackr-web` repository to Cloudflare Pages:

- **Build command:** `npm run build`
- **Output directory:** `build`
- **Node version:** 20

Set the custom domain to `hijackr.io` in the Cloudflare Pages dashboard.