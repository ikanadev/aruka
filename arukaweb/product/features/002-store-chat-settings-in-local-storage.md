### Description

The `/chat/new` screen has a form with a "General settings" section containing two settings fields: **custom prompt** (`general.customPrompt`) and **model** (`general.modelId`). Currently these settings reset to defaults every time the user visits the new chat screen.

We want to persist the last-used chat settings so they are automatically restored when the user returns to `/chat/new`. The settings should only be saved to local storage **when the user actually starts a chat** (i.e., on successful mutation), not when the user merely changes the form fields. This means if a user tweaks settings but navigates away without chatting, nothing is persisted.

The persisted settings are: `modelId` and `customPrompt`.

### Tasks

1. **Create `src/features/chats/store/use-chat-settings-store.ts`** — a new zustand store with `persist` middleware (using `localStorage`).
   - Interface `ChatSettingsStore` with:
     - `modelId: string | null`
     - `customPrompt: string`
     - `saveSettings: (settings: { modelId: string; customPrompt: string }) => void`
   - Use `zustand/middleware` `persist` with a storage key like `"chat-settings"`.
   - Default values: `modelId: null`, `customPrompt: ""`.

2. **Update `useNewChatForm` (`src/features/chats/screen/NewChat/useNewChatForm.ts`)** — read persisted settings from the store and use them as `initialValues` for the form.
   - Import `useChatSettingsStore` and read `modelId` and `customPrompt`.
   - If `modelId` from the store is `null`, fall back to `DEFAULT_SELECT_MODEL`.
   - If `customPrompt` from the store is empty, fall back to `""` (same as current default).

3. **Update `useNewChatMutation` (`src/features/chats/screen/NewChat/useNewChatMutation.ts`)** — save settings to the store on successful chat creation.
   - In the `onSuccess` callback, call `saveSettings` with the `general.modelId` and `general.customPrompt` from the submitted form values.
