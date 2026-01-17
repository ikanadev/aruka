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

  const addMessageToCache = (message: Message) => {
    queryClient.setQueryData(
      chatQueryKeys.chatMessages(chatId),
      (prev: ChatMessagesResponse | undefined): ChatMessagesResponse => {
        if (prev === undefined) {
          return {
            $typeName: "chat.v1.ChatMessagesResponse",
            messages: [message],
          };
        }
        return {
          ...prev,
          messages: [...prev.messages, message],
        };
      },
    );
  };

  return {
    messages,
    addMessageToCache,
    loadingMessages: query.isLoading,
    fetchingMessages: query.isFetching,
    isFetchedMessages: query.isFetched,
  };
}
