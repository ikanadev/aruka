import { useChatSettingsStore } from "@features/chats/store/use-chat-settings-store";
import { useForm } from "@mantine/form";
import { DEFAULT_SELECT_MODEL } from "./utils";

export type NewChatFormValues = ReturnType<typeof useNewChatForm>["values"];
export function useNewChatForm() {
  const modelId = useChatSettingsStore((state) => state.modelId);
  const customPrompt = useChatSettingsStore((state) => state.customPrompt);

  return useForm({
    mode: "uncontrolled",
    initialValues: {
      userText: "",
      general: {
        modelId: modelId ?? DEFAULT_SELECT_MODEL,
        customPrompt: customPrompt,
      },
    },
    validate: {
      userText: (value) => (value ? null : "Please enter a message"),
      general: {
        modelId: (value) => (value === DEFAULT_SELECT_MODEL ? "Please select a provider" : null),
      },
    },
  });
}
