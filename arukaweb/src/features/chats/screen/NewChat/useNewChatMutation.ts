import { chatClient } from "@common/utils/clients";
import { chatQueryKeys } from "@features/chats/data/chat-query-keys";
import { useChatSettingsStore } from "@features/chats/store/use-chat-settings-store";
import { useFirstMessageStore } from "@features/chats/store/useFirstMessageStore";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { NewChatFormValues } from "./useNewChatForm";

export function useNewChatMutation() {
  const queryClient = useQueryClient();
  const setFirstMessage = useFirstMessageStore((state) => state.setFirstMessage);
  const saveSettings = useChatSettingsStore((state) => state.saveSettings);
  return useMutation({
    mutationFn: (data: NewChatFormValues) => {
      return chatClient.newChat({
        modelId: data.general.modelId,
        prompt: data.general.customPrompt,
      });
    },
    onSuccess: (_, { userText, general }) => {
      queryClient.invalidateQueries({ queryKey: chatQueryKeys.chats() });
      setFirstMessage(userText);
      saveSettings({ modelId: general.modelId, customPrompt: general.customPrompt });
    },
  });
}
