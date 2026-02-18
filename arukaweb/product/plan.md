### 001: Manage chat messages in zustand
In `src/features/chats/screen/Chat/Chat.tsx` we are fetching message from backend and updating the react/query cache for new messages, this caused conflicts, we need to change the strategy and fetch the chat messages when user visits `/chat/[ID]` and store it in a zustand store, and from there only update the zustand store for new messages. Create a `chat-messages-store.ts` to store such messages by chat id and update the needed `features/chats/data` hooks to add, or update stored messages. Use the generated connect grpc `Message` type for zustand, and in the chat screen use a proper selector to get only messages of that chat.

### 002: Store chat settings in local storage
In the `/chat/new` screen, we have a form where the user besides the initial message there is a section of general config, for now we store only the custom prompt and the model, we want to save those in local storage, create a new zustand store (that persist the data via local storage) where we have last chat settings data and restore it when we visit the new chat screen, we should only save the changed settings when we start chatting, I mean if the user updates the settings and didn't start a chat, so we don't save that.


### 003: Manuat chat title update
In the `src/routes/chat/route.tsx` we have a sidebar with the chat list, we need to add the manually update chat title feature. There is a connect service ready to use. In the chat list, for each item we need to add a dropdown menu fired by pressing the chat (no extra buttons, just make the chats pressable) between the options, we need the "Update chat title" option that request a chat title update, after a success response, we need to update the tanstack query cache to update the desired chat title)
