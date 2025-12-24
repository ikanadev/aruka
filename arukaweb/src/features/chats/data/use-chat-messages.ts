import { chatClient } from "@common/utils/clients";
import { useQuery } from "@tanstack/react-query";
import { chatQueryKeys } from "./chat-query-keys";

interface Props {
  chatId: string;
  enabled?: boolean;
}
export function useChatMessages(props: Props) {
  const { chatId, enabled = true } = props;
  const query = useQuery({
    queryFn: () => chatClient.chatMessages({ chatId }),
    queryKey: chatQueryKeys.chatMessages(chatId),
    enabled,
  });

  const messages = query.data?.messages || [];

  return {
    messages,
    loadingMessages: query.isLoading,
    fetchingMessages: query.isFetching,
    isFetchedMessages: query.isFetched,
  }
}
