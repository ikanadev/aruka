# Aruka - Unified LLM Interface

Aruka is a multi-platform application that provides a single interface to access multiple LLM providers (OpenAI, Anthropic, Google, etc.) Users store their API keys securely in the backend and can seamlessly switch between different models and providers without managing multiple subscriptions or exposing credentials to the frontend.

### Libraries
```
    "@bufbuild/protobuf": "^2.10.1",
    "@connectrpc/connect": "^2.1.1",
    "@connectrpc/connect-web": "^2.1.1",
    "@mantine/core": "^8.3.9",
    "@mantine/form": "^8.3.9",
    "@mantine/hooks": "^8.3.9",
    "@tanstack/react-query": "^5.90.11",
    "@tanstack/react-router": "^1.139.11",
    "react": "^19.2.0",
    "react-dom": "^19.2.0",
    "remarkable": "^2.0.1",
    "zustand": "^5.0.9"
```

### Available commands

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

### TypeScript/React Frontend Style
**Routing**: TanStack React Router with file-based routing
**Architecture Pattern**: Feature-based organization, the `features/common` holds stuff that is shared between features
```
src/features/
    └── feature-name/
        ├── components/      # UI components
        ├── data/            # TanStack Query hooks that deals with server data using connect grpc
        ├── screen/          # Page-level components
        ├── store/           # Zustand stores (if needed)
        └── utils/           # Feature-specific utilities
```

**Coding Conventions**:
1. **File naming**: Use kebab-case for all files (e.g., `use-chats.ts`, `chat-query-keys.ts`, `chat-screen.tsx`)
2. **Formatting**: 
   - Use Oxfmt default options
4. **Imports**:
   - Use path aliases: `@assets/*`, `@common/*`, `@connect/*`, `@features/*`
   - Auto-organized by Oxlint, Oxfmt
5. **Type safety**:
   - Strict mode enabled
   - Use generated protobuf types from `@connect/*`
   - For app models use interfaces for main data and types for derived data
6. **State management**:
   - TanStack Query for server state
   - Zustand for client state
   - Query key factories for cache management
   - Query cache configured to last forever, manually invalidate when needed
   - Components almost never use tanstack query or mutation directly, use custom hooks located in `feature-name/data` instead
7. **Hooks**:
   - Custom hooks use `use` prefix (e.g., `useChats`, `useChatMessages`)
   - Return objects with descriptive names (e.g., `chats`, `loadingChats`)
8. **Components**:
   - Functional components with hooks, don't use arrow functions for component, use named functions
   - Use Arrow functions for any function inside a component
   - Use Mantine UI library components
   - CSS Modules for component-specific styles

Example data hook:
```typescript
import { chatClient } from "@common/utils/clients";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { chatQueryKeys } from "./chat-query-keys";
import type { ChatMessagesResponse } from "@connect/arukabe/chat/v1/chat_messages_pb";
import type { Message } from "@connect/arukabe/models/v1/message_pb";

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

