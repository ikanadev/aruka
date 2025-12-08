import { chatClient } from "@common/utils/clients";
import { chatQueryKeys } from "@features/chats/data/chat-query-keys";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { NewChatFormValues } from "./useNewChatForm";
import { useFirstMessageStore } from "@features/chats/store/useFirstMessageStore";

export function useNewChatMutation() {
  const queryClient = useQueryClient();
  const setFirstMessage = useFirstMessageStore((state) => state.setFirstMessage);
  return useMutation({
    mutationFn: (data: NewChatFormValues) => {
      return chatClient.newChat({
        modelId: data.general.modelId,
        prompt: data.general.customPrompt,
      });
    },
    onSuccess: (_, { userText }) => {
      queryClient.invalidateQueries({ queryKey: chatQueryKeys.chats() });
      setFirstMessage(userText);
    },
  });
}
