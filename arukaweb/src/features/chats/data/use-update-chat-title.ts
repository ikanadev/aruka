import { chatClient } from "@common/utils/clients";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { chatQueryKeys } from "./chat-query-keys";
import type { InfiniteData } from "@tanstack/react-query";
import type { ListChatsResponse } from "@connect/arukabe/chat/v1/list_chats_pb";

export function useUpdateChatTitle() {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (chatId: string) => chatClient.autoChatTitleUpdate({ chatId }),
    onSuccess: (data, chatId) => {
      queryClient.setQueryData<InfiniteData<ListChatsResponse>>(
        chatQueryKeys.chats(),
        (oldData) => {
          if (!oldData) return oldData;

          return {
            ...oldData,
            pages: oldData.pages.map((page) => ({
              ...page,
              chats: page.chats.map((chat) =>
                chat.id === chatId ? { ...chat, title: data.title } : chat,
              ),
            })),
          };
        },
      );
    },
  });

  return {
    updateChatTitle: mutation.mutate,
    updatingChatTitle: mutation.isPending,
    updatingChatId: mutation.variables,
  };
}
