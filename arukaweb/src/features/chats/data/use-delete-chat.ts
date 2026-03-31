import { chatClient } from "@common/utils/clients";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { chatQueryKeys } from "./chat-query-keys";

export function useDeleteChat() {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (chatId: string) => chatClient.deleteChat({ chatId }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: chatQueryKeys.chats() });
    },
  });

  return {
    deleteChat: mutation.mutate,
    deletingChat: mutation.isPending,
    deletingChatId: mutation.variables,
  };
}
