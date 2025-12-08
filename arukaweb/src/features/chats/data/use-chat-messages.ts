import { chatClient } from "@common/utils/clients";
import { useQuery } from "@tanstack/react-query";
import { chatQueryKeys } from "./chat-query-keys";

export function useChatMessages(chatId: string) {
  const query = useQuery({
    queryFn: () => chatClient.chatMessages({ chatId }),
    queryKey: chatQueryKeys.chatMessages(chatId),
  });

  const messages = query.data?.messages || [];

  return {
    messages,
    loadingMessages: query.isLoading,
    fetchingMessages: query.isFetching,
  }
}
