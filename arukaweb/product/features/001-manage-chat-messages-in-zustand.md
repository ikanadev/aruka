### Description

Currently, chat messages are fetched via TanStack Query (`useChatMessages` hook) and new messages are added by directly mutating the query cache with `queryClient.setQueryData`. This causes conflicts because the query cache can refetch/overwrite data while the UI is also pushing optimistic updates (user messages, streamed AI responses) into the same cache.

The goal is to decouple message fetching from message state management:

1. **Fetch messages once** when the user navigates to `/chat/[ID]` using the existing `chatClient.chatMessages` RPC call.
2. **Store fetched messages in a new Zustand store** (`chat-messages-store.ts`) keyed by `chatId`, instead of relying on TanStack Query cache for message state.
3. **All subsequent message mutations** (adding user messages, appending streamed AI responses, saving the generated response on navigation) should update the Zustand store directly — no more `queryClient.setQueryData` for messages.
4. **TanStack Query remains only for the initial fetch**. The query cache for `chat-messages` keys is no longer used as the source of truth.
5. The `Chat.tsx` component should use a **Zustand selector** scoped to the current `chatId` so it only re-renders when messages for that specific chat change.

### Tasks

1. **Create `src/features/chats/store/chat-messages-store.ts`**
   - Define a Zustand store with state shape: `Record<string, Message[]>` keyed by `chatId`.
   - Actions:
     - `setMessages(chatId: string, messages: Message[])` — replaces all messages for a chat (used after initial fetch).
     - `addMessage(chatId: string, message: Message)` — appends a single message to a chat's message list.
     - `updateLastMessage(chatId: string, message: Message)` — replaces the last message in the list (useful for updating the streaming AI response in-place, if needed).
   - Use the generated `Message` type from `@connect/arukabe/models/v1/message_pb`.
   - Export a selector helper: `selectChatMessages(chatId: string)` that returns `Message[]` for that chat (defaulting to `[]`).

2. **Refactor `src/features/chats/data/use-chat-messages.ts`**
   - Keep the TanStack Query fetch for the initial load (`chatClient.chatMessages`), but on success, call `setMessages(chatId, data.messages)` on the Zustand store instead of relying on `query.data`.
   - Remove `addMessageToCache` (the `queryClient.setQueryData` logic).
   - The hook should return:
     - `messages` — sourced from the Zustand store via selector (`useChatMessagesStore(selectChatMessages(chatId))`).
     - `loadingMessages`, `isFetchedMessages` — still from the query (needed for initial load UX).
   - Remove `chatMessages` from `chat-query-keys.ts` if no longer needed, or keep it if the query is still used for fetching (just not as state source).

3. **Update `src/features/chats/screen/Chat/Chat.tsx`**
   - Replace all `addMessageToCache(...)` calls with the Zustand store's `addMessage(chatId, ...)`.
   - The `generatedResponse` save on `onBeforeNavigate` should also use `addMessage` from the store.
   - `messages` should come from the refactored `useChatMessages` hook (which now reads from Zustand).
   - No other behavioral changes — scroll logic, form handling, streaming, retry, and auto-send on first message should remain the same.

4. **Cleanup**
   - Ensure no other files import `addMessageToCache` or depend on the old query cache pattern for messages.
