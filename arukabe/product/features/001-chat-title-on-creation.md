### Description

Currently, when a chat is created via `NewChat`, the title is saved as an empty string (`""`). The goal is to implement the `AutoChatTitleUpdate` RPC endpoint (already defined in the proto/generated code but currently panics with `"unimplemented"` in the handler).

This endpoint receives a `chat_id`, fetches all messages for that chat, extracts the text content from them, and calls the Anthropic API with the small model `claude-3-5-haiku-20241022` to generate a short, descriptive title for the chat. The generated title is then saved to the database via the existing `UpdateChat` query and returned in the response.

**Key behaviors:**
- If the chat has **zero messages**, do NOT call the LLM. Return a successful response with an empty title (`""`).
- The model identifier `claude-3-5-haiku-20241022` must be stored as a constant in `core/common/constants/constants.go`.
- The title should be short and help the user understand what the chat was about.

**Already in place:**
- `AutoChatTitleUpdateRequest` (field: `chat_id` string) and `AutoChatTitleUpdateResponse` (field: `title` string) are generated.
- Handler stub exists at `core/chat/handler/chat_handler.go` (line 97) — currently panics.
- `UpdateChat` SQL query supports partial title update via `COALESCE(sqlc.narg('title'), title)`.
- `GetChatMessages` SQL query fetches all messages for a chat ordered by `created_at ASC`.
- Anthropic client (`*anthropic.Client`) is already injected into `ChatService`.

### Tasks

1. **Add the title generation model constant** in `core/common/constants/constants.go`:
   - Add `TitleGenerationModel = "claude-3-5-haiku-20241022"` as a string constant.

2. **Create the service method** `AutoChatTitleUpdate` in a new file `core/chat/service/auto_chat_title_update.go`:
   - Parse `req.ChatId` as UUID.
   - Fetch messages via `cs.db.GetChatMessages(ctx, chatUUID)`.
   - If no messages, return `&chatv1.AutoChatTitleUpdateResponse{Title: ""}` immediately (success, no LLM call).
   - Extract text content from all messages (iterate message sections, grab text parts) and build a summary string of the conversation to send to the LLM.
   - Call Anthropic API using `cs.antClient.Messages.New(ctx, ...)` (non-streaming, synchronous) with:
     - `Model`: `constants.TitleGenerationModel`
     - `MaxTokens`: small value (e.g., 50)
     - `Temperature`: low (e.g., 0.2)
     - A system prompt instructing the model to generate a short chat title (e.g., "Generate a short, descriptive title (max 6 words) for the following conversation. Reply with only the title, nothing else.")
     - A single user message containing the concatenated conversation text.
   - Extract the text response from the LLM result.
   - Save the title via `cs.db.UpdateChat(ctx, sqlc.UpdateChatParams{ChatID: chatUUID, Title: pgtype.Text{String: title, Valid: true}})`.
   - Return `&chatv1.AutoChatTitleUpdateResponse{Title: title}`.

3. **Wire up the handler** in `core/chat/handler/chat_handler.go`:
   - Replace the `panic("unimplemented")` in `AutoChatTitleUpdate` with a call to `c.service.AutoChatTitleUpdate(ctx, req)` following the same pattern as the other handlers.
