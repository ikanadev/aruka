# AGENTS.md

## Build & Run Commands
- **Backend**: `cd arukabe && go run ./cmd/grpc/main.go` or `./runlocal.sh`
- **Frontend**: `cd arukaweb && pnpm dev` (dev), `pnpm build` (prod)
- **Lint**: `cd arukaweb && pnpm lint` (ESLint + Biome)
- **Database**: `docker-compose up` (PostgreSQL), `./migrate.sh up` (migrations)
- **Codegen**: `sqlc generate` (Go from SQL), `buf generate` (proto)

## Code Style
### Go (arukabe/)
- Architecture: handler → service → repository pattern
- Files: snake_case (`chat_handler.go`), packages: lowercase single words
- Errors: return up chain, use `connect.NewError()` for gRPC errors
- Constructors: `NewXxx()` pattern (e.g., `NewChatService`)

### TypeScript (arukaweb/)
- Formatting: tabs, double quotes (Biome enforced)
- Imports: use path aliases (`@assets/*`, `@common/*`, `@connect/*`, `@features/*`)
- Components: feature-based (`src/features/`), each in own folder with `.tsx`
- Data: TanStack Query hooks, query keys in `*-query-keys.ts`
- Routing: TanStack Router file-based (`src/routes/`)
