# Dim Kb

Personal knowledge base built with [Astro Starlight](https://starlight.astro.build).

The site supports two languages:

- **English** (source) — https://natchanon9.github.io/dim-kb/
- **Thai** (translation) — https://natchanon9.github.io/dim-kb/th/

Use the language picker in the header to switch. Pages not translated yet fall back to English.

## Run

Requires Node 22+ and [pnpm](https://pnpm.io) (version pinned in `package.json` → `packageManager`; `corepack enable` picks it up).

```bash
pnpm install
pnpm dev          # http://localhost:4321/dim-kb/
pnpm check        # type-check (astro check)
pnpm build        # output -> dist/ (generated HTML, not committed)

cd examples/go && go test ./...   # snippets used in the docs
```

## Structure

```text
src/content/docs/
├─ index.mdx              # home page (English)
├─ programming/           # one folder per language: go/<section>/<page>.md, rust/, typescript/
└─ th/                    # Thai translations, same paths
astro.config.ts           # Starlight config: locales, base path, plugins
src/sidebar.ts            # one sidebar topic per language (English labels + Thai translations)
src/content.config.ts     # extra frontmatter: lastReviewed, status
src/styles/custom.css     # site styling (Inter, Noto Sans Thai, JetBrains Mono; self-hosted)
examples/go/              # Go module holding every code snippet
```

## Writing a page

1. Read up on the topic and take notes in your own words.
2. Write the code in `examples/go/<section>/<page>/` and make `go test` pass.
3. Write the page: a 1–2 sentence intro, one short section per concept with its code, and a Pitfalls callout at the end.
4. Set `status: done` and `lastReviewed: YYYY-MM-DD`.
5. Translate it to `src/content/docs/th/programming/go/<section>/<page>.md`.

Find pages not written yet: `grep -rl "status: draft" src/content/docs/programming`

## Deploy

Every push to `main` builds and deploys to GitHub Pages (`.github/workflows/deploy.yml`). Go snippets are checked by `.github/workflows/examples.yml`.
