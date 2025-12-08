export const chatQueryKeys = {
  chats: () => ['chats-list'],
  chatMessages: (chatId: string) => ['chat-messages', chatId],
};
