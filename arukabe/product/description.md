### Go Backend Style

**Architecture Pattern**: Handler → Service → Repository, where almost always, each handler has a corresponding service method with the same name, and each service has access to the database methods via sqlc or any other LLM provider.

```
core/
└── domain/              # E.g., "chat", "provider"
    ├── handler/         # RPC handlers (Connect-RPC)
    ├── service/         # Business logic (uses the request and response types of handler)
    └── queries/         # SQLC SQL queries
```

**Generated code**: We use connect grpc to generate handlers and SQLC to generate typed SQL queries. Both go in `gen/` directory. Usually the /core/DOMAIN/service use the connect grpc generated methods, and SQLC generated code and returns connect GRPC responses (to the handler)

**Coding Conventions**:

1. **Naming**: 
   - Constructors: `NewXxx()` pattern
   - Interfaces: Not heavily used; prefer concrete types
2. **Imports**: Standard library first, then external packages, then local packages (separated by blank lines)
3. **Error handling**: 
   - Return `connect.Error` with proper codes in services
4. **Database access**:
   - Use SQLC generated functions (type-safe)
   - UUID for all IDs (`github.com/google/uuid`)
   - Pass `context.Context` to all DB operations
5. **Dependencies**: Inject via constructor parameters
6. **IDs**: Use UUID for all primary keys, use v7 variant

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
) (*chatv1.NewChatResponse, error) {
	resp, err := c.service.NewChat(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
```

