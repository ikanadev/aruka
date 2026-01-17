# Aruka Development Guide for AI Coding Agents

This guide provides comprehensive instructions for AI coding agents working on the Aruka codebase - an AI chat application built with Go backend, React frontend, and Protocol Buffers for API contracts.

## Repository Structure

```
/home/kevv/pys/aruka/
├── proto/           # Protocol Buffers definitions (buf)
├── arukabe/         # Go backend (Connect-RPC)
└── arukaweb/        # React/TypeScript frontend (Vite)
```

## Build, Lint, and Test Commands

### Protocol Buffers (`/home/kevv/pys/aruka/proto`)

```bash
# Generate code for all targets (Go + TypeScript)
buf generate

# Lint protobuf files
buf lint

# Check for breaking changes
buf breaking --against '.git#branch=main'

# Format proto files
buf format -w
```

### Go Backend (`/home/kevv/pys/aruka/arukabe`)

```bash
# Run development server
./runlocal.sh

# Build
go build ./cmd/grpc/main.go

# Format code
go fmt ./...

# Database migrations
./migrate.sh up              # Apply all pending migrations
./migrate.sh up 1            # Apply next migration
./migrate.sh down            # Revert last migration
./migrate.sh version         # Show current version

# Generate SQLC code (after modifying queries)
sqlc generate

# Start PostgreSQL database
docker-compose up -d
```

### Frontend (`/home/kevv/pys/aruka/arukaweb`)

```bash
# Install dependencies (uses pnpm)
pnpm install

# Development server
pnpm dev

# Build for production
pnpm build

# Lint TypeScript/React
pnpm lint

# Format TypeScript/React
pnpm format

# Preview production build
pnpm preview

```

## Code Style Guidelines

### General Principles

1. **No tests** - No test in any project
2. **Oxlint for lint in frontend**
3. **Type safety is paramount** - Use generated types from protobuf; avoid `any` types
4. **Follow existing patterns** - Match the established architecture in each layer

### Protocol Buffers Style

1. **File naming**: Use snake_case for .proto files (e.g., `chat_message.proto`)
2. **Versioning**: All packages include version (e.g., `chat.v1`, `models.v1`)
3. **Request/Response pairs**: Every RPC has dedicated `XxxRequest` and `XxxResponse` messages
4. **Enums**: Zero value must be `XXX_UNSPECIFIED`
5. **Timestamps**: Use `google.protobuf.Timestamp` consistently
6. **Shared models**: Common types go in `models.v1` package
7. **Follow STANDARD lint rules**: Buf enforces Google's style guide automatically

Example:
```protobuf
syntax = "proto3";

package chat.v1;

import "google/protobuf/timestamp.proto";
import "models/v1/chat.proto";

message NewChatRequest {
  string title = 1;
  string model_id = 2;
}

message NewChatResponse {
  models.v1.Chat chat = 1;
}
```

### Go Backend Style

**Architecture Pattern**: Handler → Service → Repository, where in most of cases, each handler has a corresponding service method with the same name, and each service method can call any repository methods. Only repository methods can call SQLC generated code. Only Service method can call repository methods. Only Handler can call service methods.

```
core/
└── domain/              # E.g., "chat", "provider"
    ├── handler/         # RPC handlers (Connect-RPC)
    ├── service/         # Business logic
    ├── repository/      # Data access
    └── queries/         # SQLC SQL queries
```

**Generated code**: We use connect grpc to generate handlers and SQLC to generate typed SQL queries. Both go in `gen/` directory. Only /core/DOMAIN/handler and /core/DOMAIN/service use the connect grpc generated methods, and only core/DOMAIN/repository can call SQLC generated code, but core/DOMAIN/service receives SQLC models (by calling repository) and return connect GRPC responses (to the handler)

**Coding Conventions**:

1. **File organization**: One method per file in service/repository layers; struct definition in `service.go`/`repository.go`
2. **Naming**: 
   - Constructors: `NewXxx()` pattern
   - Interfaces: Not heavily used; prefer concrete types
3. **Imports**: Standard library first, then external packages, then local packages (separated by blank lines)
4. **Error handling**: 
   - Return `connect.Error` with proper codes from handlers
   - Propagate context throughout call chain
   - Use `fmt.Errorf()` with `%w` for wrapping errors
5. **Database access**:
   - Use SQLC generated functions (type-safe)
   - UUID for all IDs (`github.com/google/uuid`)
   - Pass `context.Context` to all DB operations
6. **Dependencies**: Inject via constructor parameters

Example handler:
```go
package handler

import (
	"arukabe/core/chat/service"
	chatv1 "arukabe/gen/connect/chat/v1"
	"context"

	"connectrpc.com/connect"
)

type ChatHandler struct {
	service service.ChatService
}

func NewChatHandler(service service.ChatService) *ChatHandler {
	return &ChatHandler{service}
}

func (c *ChatHandler) NewChat(
	ctx context.Context,
	req *chatv1.NewChatRequest,
) (*connect.Response[chatv1.NewChatResponse], error) {
	resp, err := c.service.NewChat(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
```

### TypeScript/React Frontend Style

**Architecture Pattern**: Feature-based organization

```
features/
└── feature-name/
    ├── components/      # UI components
    ├── data/            # TanStack Query hooks that deals with server data using connect grpc
    ├── screen/          # Page-level components
    ├── store/           # Zustand stores (if needed)
    └── utils/           # Feature-specific utilities
```

**Coding Conventions**:

1. **File naming**: Use kebab-case for files (e.g., `use-chats.ts`, `chat-query-keys.ts`)
2. **Component naming**: PascalCase for components matching filename
3. **Formatting**: 
   - Use Oxfmt default options
4. **Imports**:
   - Use path aliases: `@assets/*`, `@common/*`, `@connect/*`, `@features/*`
   - Auto-organized by Oxlint, Oxfmt
   - React imports not required (React 19 JSX transform)
5. **Type safety**:
   - Strict mode enabled
   - Use generated protobuf types from `@connect/*`
   - Define props interfaces inline or with `type`
   - No unused variables/parameters allowed
6. **State management**:
   - TanStack Query for server state
   - Zustand for client state
   - Query key factories for cache management
7. **Hooks**:
   - Custom hooks use `use` prefix (e.g., `useChats`, `useChatMessages`)
   - Return objects with descriptive names (e.g., `chats`, `loadingChats`)
8. **Components**:
   - Functional components with hooks (no class components)
   - Use Mantine UI library components
   - CSS Modules for component-specific styles

Example data hook:
```typescript
import { chatClient } from "@common/utils/clients";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { chatQueryKeys } from "./chat-query-keys";
import type { ChatMessagesResponse } from "@connect/chat/v1/chat_messages_pb";
import type { Message } from "@connect/models/v1/message_pb";

export function useChatMessages(chatId: string) {
  const queryClient = useQueryClient();

  const query = useQuery({
    queryFn: () => chatClient.chatMessages({ chatId }),
    queryKey: chatQueryKeys.chatMessages(chatId),
  });

  const messages = query.data?.messages || [];

  return {
    messages,
    addMessageToCache,
    loadingMessages: query.isLoading,
    fetchingMessages: query.isFetching,
    isFetchedMessages: query.isFetched,
  };
}
```

## Database Conventions

1. **IDs**: Use UUID for all primary keys
2. **Timestamps**: Include `created_at`, `updated_at` on all tables
3. **Soft deletes**: Use `deleted_at`, `archived_at` (nullable timestamptz)
4. **Schema changes**: Always create migrations (never modify existing migrations)
5. **Queries**: Write SQL in `queries/*.sql` files, generate Go code with SQLC
6. **Indexes**: Add GIN indexes for JSONB columns, composite indexes for common queries

## Important Notes for Agents

1. **Code generation workflow**:
   - Modify proto files → `buf generate` → generates Go + TypeScript code
   - Modify SQL queries → `sqlc generate` → generates Go repository code
   - **Always regenerate after modifying proto or SQL files**

2. **Path references**:
   - Backend working directory: `/home/kevv/pys/aruka/arukabe`
   - Frontend working directory: `/home/kevv/pys/aruka/arukaweb`
   - Proto working directory: `/home/kevv/pys/aruka/proto`

3. **No existing coding rules**:
   - No `.cursorrules`, `.cursor/rules/`, or `.github/copilot-instructions.md` files exist
   - Follow patterns established in existing code

4. **No testing in this project**: No test framework configured; never add test files

5. **Common pitfalls to avoid**:
   - Don't modify generated code in `gen/`, `connect/` directories
   - Don't skip migrations when updating schema
   - Don't use `any` type in TypeScript
   - Don't forget to run `buf generate` after proto changes

## Quick Reference

| Task | Command | Directory |
|------|---------|-----------|
| Generate proto code | `buf generate` | `/home/kevv/pys/aruka/proto` |
| Run backend | `./runlocal.sh` | `/home/kevv/pys/aruka/arukabe` |
| Run frontend | `pnpm dev` | `/home/kevv/pys/aruka/arukaweb` |
| Generate SQL code | `sqlc generate` | `/home/kevv/pys/aruka/arukabe` |
| Apply migrations | `./migrate.sh up` | `/home/kevv/pys/aruka/arukabe` |
| Lint proto | `buf lint` | `/home/kevv/pys/aruka/proto` |
| Lint frontend | `pnpm lint` | `/home/kevv/pys/aruka/arukaweb` |
| Format frontend | `pnpm format` | `/home/kevv/pys/aruka/arukaweb` |
| Format Go | `go fmt ./...` | `/home/kevv/pys/aruka/arukabe` |
