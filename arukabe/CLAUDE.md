# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go gRPC backend for a chat application backed by Anthropic LLMs and PostgreSQL. The project dir is `/home/kevv/pys/aruka/arukabe` — all file references are relative to this path. There are no tests in this project.

## Commands

**Run locally:**
```bash
./runlocal.sh   # loads .env, sets env vars, runs go run ./cmd/grpc/main.go
```

**Database migrations** (requires `migrate` CLI):
```bash
./migrate.sh up          # apply all pending migrations
./migrate.sh down        # revert last migration
./migrate.sh goto <v>    # migrate to specific version
./migrate.sh version     # print current version
```

**Regenerate SQLC queries** (after editing `core/*/queries/*.sql`):
```bash
sqlc generate
```

**Docker: (used only for database in development mode)**
```bash
docker compose up        # starts PostgreSQL on port 5433
```

## Environment

Copy `.env.example` to `.env` and fill in:
- `PG_*` — PostgreSQL connection details
- `ANTHROPIC_API_KEY`

## Architecture

**Request flow:** Handler → Service → SQLC queries (DB) / Anthropic SDK

```
core/
└── <domain>/          # chat, provider
    ├── handler/       # thin RPC handler, delegates to service
    ├── service/       # business logic, returns Connect-RPC response types
    └── queries/       # raw SQL (input to sqlc generate)
gen/
├── sqlc/              # generated type-safe DB code (do not edit)
└── connect/           # generated Connect-RPC handlers (do not edit)
```

**Server startup** (`grpc/grpc.go`): loads config → opens PG pool → runs migrations → seeds providers → creates Anthropic client → registers services → starts HTTP/2 server on port 5000.

**Message content** is stored as JSONB and supports multiple block types (text, thinking, tool use). Conversions between DB ↔ Protobuf ↔ Anthropic SDK formats live in `core/common/types/`.

**Title generation** uses `constants.TitleGenerationModel` (Haiku) via a non-streaming Anthropic call in `core/chat/service/`.

## Coding Conventions

- **Pattern**: each handler method has a 1:1 corresponding service method with the same name.
- **Constructors**: `NewXxx()` returning concrete types (interfaces rarely used).
- **IDs**: UUID v7 everywhere (`github.com/google/uuid`).
- **Errors**: return `connect.NewError(connect.CodeXxx, err)` from service layer.
- **DB calls**: always pass `context.Context`; use SQLC-generated functions only.
- **Imports**: stdlib → external → local, separated by blank lines.

## Code Generation

- **SQLC** (`sqlc.yml`): SQL files in `core/*/queries/` → generated Go in `gen/sqlc/`.
- **Connect-RPC**: proto definitions → generated handlers in `gen/connect/`. Regenerate with `buf generate` (or whichever buf command is configured).
