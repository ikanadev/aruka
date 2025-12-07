import { useForm } from "@mantine/form";
import { DEFAULT_SELECT_MODEL } from "./utils";

export type NewChatFormValues = ReturnType<typeof useNewChatForm>["values"];
export function useNewChatForm() {
  return useForm({
    mode: "uncontrolled",
    initialValues: {
      userText: "",
      general: {
        modelId: DEFAULT_SELECT_MODEL,
        customPrompt: "",
      },
    },
    validate: {
      userText: (value) => (value ? null : "Please enter a message"),
      general: {
        modelId: (value) =>
          value === DEFAULT_SELECT_MODEL ? "Please select a provider" : null,
      },
    },
  });
}
