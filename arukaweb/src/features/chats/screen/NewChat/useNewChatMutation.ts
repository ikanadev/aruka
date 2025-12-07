import { chatClient } from "@common/utils/clients";
import { chatQueryKeys } from "@features/chats/data/chat-query-keys";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { NewChatFormValues } from "./useNewChatForm";

export function useNewChatMutation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: NewChatFormValues) => {
      return chatClient.newChat({
        modelId: data.general.modelId,
        prompt: data.general.customPrompt,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: chatQueryKeys.chats() });
    },
  });
}
