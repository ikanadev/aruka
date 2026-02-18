### Description

Add the ability for users to manually trigger a chat title update from the sidebar chat list. Each chat item in the sidebar should display a three-dot menu icon on hover. Clicking the icon opens a dropdown menu with an "Update chat title" option. Selecting this option calls the `autoChatTitleUpdate` connect service with the `chatId`, the backend generates a new title based on the chat content, and the response title is used to update the tanstack query cache for the chat list.

**Key points:**
- The user does NOT provide a title manually; the backend generates it automatically from the chat content.
- The connect service to use is `chatClient.autoChatTitleUpdate({ chatId })` which returns `{ title: string }`.
- The three-dot icon button should only appear on hover over the chat item, keeping the UI clean.
- Left-clicking the chat item still navigates to the chat as before; only the three-dot icon triggers the dropdown.
- After a successful response, update the chat title in the tanstack infinite query cache (`chatQueryKeys.chats()`) without refetching.

### Tasks

1. **Modify `ChatNavbarLink` component** (`src/features/chats/components/ChatNavbarLink/ChatNavbarLink.tsx`)
   - Add a three-dot icon button (use Mantine `ActionIcon` with an ellipsis/dots icon) positioned at the right end of each chat item.
   - The icon should be hidden by default and visible only when hovering the chat item (use CSS modules or Mantine `styles`/`className` with `:hover` to toggle visibility).
   - Wrap the icon in a Mantine `Menu` component as the `Menu.Target`.
   - Add a `Menu.Dropdown` with a single `Menu.Item` labeled "Update chat title".
   - The `Menu.Item` should call an `onUpdateTitle` callback prop passed from the parent.
   - Make sure clicking the three-dot icon does NOT trigger navigation (stop event propagation).
   - Add new props: `onUpdateTitle: () => void` and `updatingTitle: boolean` (to show loading state on the menu item or icon while the request is in flight).

2. **Create `useUpdateChatTitle` data hook** (`src/features/chats/data/use-update-chat-title.ts`)
   - Create a custom hook that wraps a `useMutation` calling `chatClient.autoChatTitleUpdate({ chatId })`.
   - On success, update the infinite query cache for `chatQueryKeys.chats()`:
     - Access the cached `InfiniteData` for the chats list.
     - Find the chat by `chatId` across all pages and update its `title` with the response `title`.
     - Use `queryClient.setQueryData` to write the updated data back.
   - Return `{ updateChatTitle, updatingChatTitle }` (the mutate function and `isPending` state).

3. **Update `ChatNavbar` component** (`src/features/chats/components/ChatNavbar/ChatNavbar.tsx`)
   - Import and use the `useUpdateChatTitle` hook.
   - Pass `onUpdateTitle` and `updatingTitle` props to each `ChatNavbarLink`.
   - Wire `onUpdateTitle` to call `updateChatTitle(chatId)` from the hook.
