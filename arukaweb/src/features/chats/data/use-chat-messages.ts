import { chatClient } from "@common/utils/clients";
import { selectChatMessages, useChatMessagesStore } from "@features/chats/store/chat-messages-store";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import { chatQueryKeys } from "./chat-query-keys";

export function useChatMessages(chatId: string) {
  const setMessages = useChatMessagesStore((s) => s.setMessages);
  const messages = useChatMessagesStore(selectChatMessages(chatId));

  const query = useQuery({
    queryFn: async () => {
      await new Promise((resolve) => setTimeout(resolve, 2000));
      return chatClient.chatMessages({ chatId });
    },
    queryKey: chatQueryKeys.chatMessages(chatId),
  });

  useEffect(() => {
    if (query.data) {
      setMessages(chatId, query.data.messages);
    }
  }, [chatId, query.data, setMessages]);

  return {
    messages,
    loadingMessages: query.isLoading,
    isFetchedMessages: query.isFetched,
  };
}
